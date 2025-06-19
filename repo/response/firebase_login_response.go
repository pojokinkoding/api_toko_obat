package response

type FirebaseLoginResponse struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalID      string `json:"localId"`
	Email        string `json:"email"`
	Error        *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}
