package handlers

import (
	"context"
	"fmt"
	"net/http"

	"toko_obat/firebase"

	"bytes"
	"encoding/json"
	"os"

	"toko_obat/repo/request"
	"toko_obat/repo/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken := c.GetHeader("Authorization")
		if idToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			c.Abort()
			return
		}

		authClient, err := firebase.App.Auth(context.Background())
		if err != nil {
			fmt.Println("err", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Firebase Auth"})
			c.Abort()
			return
		}

		// Verifikasi token
		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Simpan UID ke context
		c.Set("uid", token.UID)
		c.Next()
	}
}

func ProtectedRoute(c *gin.Context) {
	uid := c.MustGet("uid").(string)
	c.JSON(http.StatusOK, gin.H{
		"message": "You are authenticated!",
		"uid":     uid,
	})
}

func Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	apiKey := os.Getenv("FIREBASE_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing Firebase API Key"})
		return
	}

	firebaseURL := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=" + apiKey
	payload := map[string]interface{}{
		"email":             req.Email,
		"password":          req.Password,
		"returnSecureToken": true,
	}
	jsonPayload, _ := json.Marshal(payload)

	resp, err := http.Post(firebaseURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Firebase"})
		return
	}
	defer resp.Body.Close()

	var firebaseResp response.FirebaseLoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&firebaseResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse Firebase response"})
		return
	}

	if firebaseResp.IDToken == "" {
		msg := "Login failed"
		if firebaseResp.Error != nil {
			msg = firebaseResp.Error.Message
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"idToken":      firebaseResp.IDToken,
		"refreshToken": firebaseResp.RefreshToken,
		"expiresIn":    firebaseResp.ExpiresIn,
		"localId":      firebaseResp.LocalID,
		"email":        firebaseResp.Email,
	})
}
