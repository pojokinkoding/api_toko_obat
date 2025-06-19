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

### 3. Medicines
- **URL:** `/api/medicines`
- **Method:** `GET`
- **Header:**
  - `Authorization: <ID_TOKEN>`
- **Query Params (optional):**
  - `page_offset` (int)
  - `page_limit` (int)
  - `dir` (string, "asc"/"desc")
  - `field` (string, nama kolom untuk sorting)
- **Response (200):**
  ```json
  {
    "record_total": 100,
    "record_total_filtered": 10,
    "data": [
      {
        "id": 1,
        "medicine_code": "MED001",
        "name": "Paracetamol",
        "category_id": 2,
        "manufacturer_id": 1,
        "type": "Tablet",
        "description": "Obat penurun demam",
        "price": 5000,
        "stock": 100,
        "unit": "strip",
        "expiry_date": "2024-12-31T00:00:00Z",
        "created_by": "admin",
        "updated_by": "admin",
        "created_at": "2023-01-01T00:00:00Z",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
  ```

- **URL:** `/api/medicines/:id`
- **Method:** `GET`
- **Header:**
  - `Authorization: <ID_TOKEN>`
- **Response (200):**
  ```json
  {
    "id": 1,
    "medicine_code": "MED001",
    "name": "Paracetamol",
    "category_id": 2,
    "manufacturer_id": 1,
    "type": "Tablet",
    "description": "Obat penurun demam",
    "price": 5000,
    "stock": 100,
    "unit": "strip",
    "expiry_date": "2024-12-31T00:00:00Z",
    "created_by": "admin",
    "updated_by": "admin",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
  ```
- **Response (404):**
  ```json
  { "error": "Medicine not found" }
  ```

### 4. Categories
- **URL:** `/api/categories`
- **Method:** `GET`
- **Header:**
  - `Authorization: <ID_TOKEN>`
- **Query Params (optional):**
  - `page_offset` (int)
  - `page_limit` (int)
  - `dir` (string, "asc"/"desc")
  - `field` (string, nama kolom untuk sorting)
- **Response (200):**
  ```json
  {
    "record_total": 10,
    "record_total_filtered": 5,
    "data": [
      {
        "id": 1,
        "name": "Analgesik",
        "description": "Obat pereda nyeri",
        "created_by": "admin",
        "updated_by": "admin",
        "created_at": "2023-01-01T00:00:00Z",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
  ```

- **URL:** `/api/categories/:id`
- **Method:** `GET`
- **Header:**
  - `Authorization: <ID_TOKEN>`
- **Response (200):**
  ```json
  {
    "id": 1,
    "name": "Analgesik",
    "description": "Obat pereda nyeri",
    "created_by": "admin",
    "updated_by": "admin",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
  ```
- **Response (404):**
  ```json
  { "error": "Category not found" }
  ```

- **URL:** `/api/categories`
- **Method:** `POST`
- **Header:**
  - `Authorization: <ID_TOKEN>`
- **Body:**
  ```json
  {
    "name": "Analgesik",
    "description": "Obat pereda nyeri"
  }
  ```
- **Response (200):**
  ```json
  {
    "success": true,
    "message": "Category created successfully",
    "data": null
  }
  ```
- **Response (400/500):**
  ```json
  {
    "success": false,
    "message": "Invalid request: ..."
  }
  ```

### 5. Root
- **URL:** `/`
- **Method:** `GET`
- **Response:**
  ```json
  { "message": "Welcome to Firebase Auth API with Go!" }
  ```

## Catatan
- Pastikan API key dan file service account sesuai dengan project Firebase Anda.
- Endpoint login menggunakan Firebase REST API (bukan Admin SDK).
- Endpoint untuk manufacturer sudah tersedia di kode, namun belum di-expose melalui API (belum ada route di main.go).
- Semua endpoint (kecuali `/` dan `/login`) membutuhkan header `Authorization` dengan ID Token dari hasil login.

## Contoh Query Params untuk List Endpoint
- `page_offset=0&page_limit=10&dir=desc&field=name`

FIREBASE_API_KEY = AIzaSyDuRLnhw7xdM7I8ZtACrF07oSHJ17oTUDQ
DATABASE_DSN=postgres://postgres:postgres@127.0.0.1:5432/toko_obat?sslmode=disable
