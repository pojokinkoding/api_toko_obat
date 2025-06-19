# api_toko_obat

API sederhana untuk autentikasi menggunakan Firebase (email & password) dengan Gin (Golang).

## Setup

1. **Clone repo & masuk ke folder project**
2. **Buat file `.env` di root project:**
   ```
   FIREBASE_API_KEY=YOUR_FIREBASE_API_KEY
   DATABASE_DSN=postgres://user:password@localhost:5432/dbname?sslmode=disable
   ```
   Ganti `YOUR_FIREBASE_API_KEY` dengan API key dari Firebase Console (Project Settings > Web API Key).
   Ganti `DATABASE_DSN` dengan DSN PostgreSQL Anda.
3. **Letakkan file service account Firebase** (misal: `firebase-adminsdk.json`) di root project.
4. **Jalankan aplikasi:**
   ```
   go run main.go
   ```

## Endpoint

### 1. Login
- **URL:** `/login`
- **Method:** `POST`
- **Body:**
  ```json
  {
    "email": "user@example.com",
    "password": "userpassword"
  }
  ```
- **Response (200):**
  ```json
  {
    "idToken": "...",
    "refreshToken": "...",
    "expiresIn": "3600",
    "localId": "...",
    "email": "user@example.com"
  }
  ```
- **Response (401):**
  ```json
  { "error": "INVALID_PASSWORD" }
  ```

### 2. Protected Route
- **URL:** `/api/protected`
- **Method:** `GET`
- **Header:**
  - `Authorization: <ID_TOKEN>` (dari hasil login)
- **Response (200):**
  ```json
  {
    "message": "You are authenticated!",
    "uid": "..."
  }
  ```
- **Response (401):**
  ```json
  { "error": "Invalid or expired token" }
  ```

### 3. Root
- **URL:** `/`
- **Method:** `GET`
- **Response:**
  ```json
  { "message": "Welcome to Firebase Auth API with Go!" }
  ```

## Catatan
- Pastikan API key dan file service account sesuai dengan project Firebase Anda.
- Endpoint login menggunakan Firebase REST API (bukan Admin SDK). 

FIREBASE_API_KEY = AIzaSyDuRLnhw7xdM7I8ZtACrF07oSHJ17oTUDQ
DATABASE_DSN=postgres://postgres:postgres@127.0.0.1:5432/toko_obat?sslmode=disable
