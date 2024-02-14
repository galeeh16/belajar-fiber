### Requirements
- mysql >= v8.0.0
- golang >= v1.21.0

### Cara menjalankan aplikasi
- buat file baru `.env` pada root aplikasi, lalu copy isi file `.env.example` ke `.env`, sesuaikan isinya
- run `go mod tidy` untuk mendownload library-library yang dibutuhkan
- run `go run main.go` untuk menjalankan aplikasi

Dokumentasi Fiber [Fiber Docs](https://docs.gofiber.io/).

### Routes
- `[GET]    /` (index route)
- `[POST]   /api/v1/users/login` (Logged in)
- `[POST]   /api/v1/users` (Create User)
- `[PUT]    /api/v1/users/:id` (Update user by id)
- `[DELETE] /api/v1/users/:id` (Delete user by id)