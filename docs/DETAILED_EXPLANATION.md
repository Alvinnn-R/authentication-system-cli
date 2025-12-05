# Authentication System CLI — Penjelasan Teknis Lengkap

Dokumen ini dibuat khusus untuk membantu mahasiswa informatika semester 5 memahami cara kerja project **Authentication System CLI** secara menyeluruh. Penjelasan dibagi ke beberapa bagian: struktur folder, alur program, detail kode setiap layer, serta bagaimana data diproses hingga tersimpan di file `users.json`.

---

## 1. Struktur Folder

```
authentication-system-cli/
├── data/
│   └── users.json          // Tempat menyimpan data user
├── handler/
│   └── user.go             // Interaksi antar user dengan CLI
├── model/
│   └── user.go             // Definisi struktur User
├── repository/
│   └── user.go             // Akses dan manipulasi data di file JSON
├── service/
│   └── user.go             // Business logic dan validasi
├── utils/
│   ├── file.go             // Helper baca/tulis JSON
│   └── message.go          // Kumpulan pesan error
├── main.go                 // Titik masuk program
└── go.mod                  // Konfigurasi module Go
```

Struktur ini menerapkan **layered architecture**: setiap folder memiliki tanggung jawab jelas sehingga kode mudah dibaca, diuji, dan dikembangkan.

---

## 2. Alur Besar Program

1. **Program dimulai di `main.go`**. Di sini kita membuat object repository → service → handler (dependency injection).
2. **Loop menu** menampilkan pilihan Register, Login, Exit.
3. **Input user** dibaca oleh handler menggunakan `bufio.Scanner` agar bisa menangkap kalimat dengan spasi.
4. **Handler memanggil service** untuk menjalankan logika bisnis (validasi, cek data, simpan ke file).
5. **Service berkomunikasi dengan repository** untuk membaca/menulis `users.json`.
6. **Repository menggunakan utils** untuk operasi file.

Diagram singkat:

```
main.go → handler → service → repository → utils/file → users.json
```

---

## 3. Penjelasan Setiap Layer

### 3.1 Model Layer (`model/user.go`)

```go
type User struct {
    FullName    string `json:"full_name"`
    Email       string `json:"email"`
    PhoneNumber string `json:"phone_number"`
    Password    string `json:"password"`
}
```

- **Tujuan**: mendeskripsikan data user. Tag JSON memastikan saat diserialisasi ke file, nama field mengikuti format `snake_case`.
- **Analoginya**: seperti formulir kosong. Setiap user baru akan mengisi kolom yang sama.

### 3.2 Utils Layer (`utils/file.go` dan `utils/message.go`)

#### `file.go`

- `ReadJSON(path, v)` membuka file, mendecode JSON ke variabel tujuan, dan aman menghadapi kasus file belum ada.
- `WriteJSON(path, v)` menulis slice struct ke file JSON dengan indentasi agar mudah dibaca.

#### `message.go`

- Menyimpan `errors.New(...)` untuk semua pesan error. Dengan begini, pesan konsisten dan mudah diubah.

### 3.3 Repository Layer (`repository/user.go`)

- `NewUserRepository(path)` menyimpan path file JSON.
- `GetAll()` mengambil semua user dengan `utils.ReadJSON`.
- `SaveAll(users)` menulis seluruh slice user ke file.
- `FindByEmail(email)` mencari user tertentu. Mengembalikan pointer ke user jika ketemu atau `nil` bila tidak.

Layer ini fokus pada **akses data** saja tanpa validasi.

### 3.4 Service Layer (`service/user.go`)

- Memegang logika bisnis.
- `Register(...)` menjalankan urutan validasi → cek duplikat → simpan user baru.
- `Login(...)` memastikan email ada dan password sesuai.
- Tiga fungsi validasi:
  - `validateEmail` memakai regex agar format email benar.
  - `validatePhoneNumber` memastikan hanya angka dan panjang 10-15 digit.
  - `validatePassword` mewajibkan minimal 6 karakter.

Jika ada pelanggaran, service mengembalikan error spesifik (contoh: `utils.ErrEmailInvalid`). Dengan begitu handler tinggal menampilkan pesan ke user.

### 3.5 Handler Layer (`handler/user.go`)

- Bertugas membaca input dari terminal dan menampilkan output ke user.
- Menggunakan `bufio.Scanner` agar bisa menerima input dengan spasi (misal: "Budi Santoso").
- Dua fungsi penting:
  - `HandleRegister()` meminta nama, email, phone, password → kirim ke `service.Register` → tampilkan hasil.
  - `HandleLogin()` meminta email, password → kirim ke `service.Login` → tampilkan sukses/gagal.
- Handler tidak melakukan validasi rumit: dia hanya mengantarkan data dan menampilkan pesan.

### 3.6 Main (`main.go`)

```go
func main() {
    repo := repository.NewUserRepository("data/users.json")
    svc := service.NewUserService(repo)
    h := handler.NewUserHandler(svc)

    for {
        h.ShowMenu()
        var choice string
        fmt.Scanln(&choice)

        switch choice {
        case "1":
            h.HandleRegister()
        case "2":
            h.HandleLogin()
        case "3":
            fmt.Println("Exit")
            return
        default:
            fmt.Println("Pilihan tidak valid")
        }
    }
}
```

- **Dependency injection** terjadi di tiga baris pertama.
- Loop `for { ... }` membuat program terus berjalan sampai user memilih menu 3.
- `switch` menentukan handler mana yang dipanggil berdasarkan input.

---

## 4. Alur Detail Register

1. User memilih menu 1 → `h.HandleRegister()` dipanggil.
2. Handler meminta 4 input: Full Name, Email, Phone, Password.
3. Handler mengoper data ke `service.Register`.
4. Service menjalankan validasi berurutan:
   - Email kosong atau format salah? → error `email tidak valid`.
   - Email sudah ada di file? → error `email sudah terdaftar`.
   - Phone bukan angka atau panjang di luar 10-15? → error `nomor telepon tidak valid`.
   - Password < 6 karakter? → error `password minimal 6 karakter`.
5. Jika semua valid, service mengambil semua user lewat `repo.GetAll()`.
6. Buat struct `model.User` dari input, append ke slice users.
7. Panggil `repo.SaveAll(users)` → `utils.WriteJSON` → file `users.json` diperbarui.
8. Handler menampilkan pesan sukses.

Visualisasi ringkas:

```
Input CLI → Handler → Service (validasi) → Repository → utils.WriteJSON → users.json
```

---

## 5. Alur Detail Login

1. User memilih menu 2 → `h.HandleLogin()` dipanggil.
2. Handler meminta email & password.
3. Data dikirim ke `service.Login`.
4. Service memanggil `repo.FindByEmail(email)`:
   - Jika `nil` → error `user tidak ditemukan`.
5. Jika user ditemukan, bandingkan password:
   - Tidak cocok → error `password salah`.
   - Cocok → return user dan handler menampilkan "Login berhasil...".

---

## 6. Penyimpanan Data (`data/users.json`)

Contoh isi file setelah dua user berhasil register:

```json
[
  {
    "full_name": "Budi Santoso",
    "email": "budi@example.com",
    "phone_number": "081234567890",
    "password": "rahasia123"
  },
  {
    "full_name": "Siti Nurbaya",
    "email": "siti@example.com",
    "phone_number": "081111222233",
    "password": "passwordku"
  }
]
```

Saat register baru:

1. Repository baca file → slice user.
2. Service append user baru ke slice.
3. Repository tulis ulang seluruh slice ke file (overwrite).

**Catatan keamanan**: Password masih plaintext karena project ini bersifat pembelajaran. Di produksi harus memakai hashing (misal bcrypt).

---

## 7. Tips Pemahaman & Modifikasi

1. **Tambah fitur** (misal: hapus user, list user) cukup membuat fungsi baru di handler/service/repository.
2. **Ubah validasi** (misal: password harus campuran huruf & angka) diubah di layer Service saja.
3. **Ganti media penyimpanan** (misal: database) cukup modifikasi layer Repository.
4. **Testing manual**: coba masukkan data salah (email tanpa `@`, phone berhuruf) untuk melihat pesan error.

---

## 8. Ringkasan Tujuan Setiap Bagian

| Bagian       | Fungsi Utama                                     |
| ------------ | ------------------------------------------------ |
| `main.go`    | Titik masuk, membuat dependency, dan loop menu   |
| Handler      | Interaksi dengan user via CLI, menampilkan pesan |
| Service      | Validasi dan aturan bisnis Register/Login        |
| Repository   | Baca dan tulis data ke file JSON                 |
| Utils        | Fungsi helper (IO file, pesan error)             |
| Model        | Struktur data User                               |
| `users.json` | Lokasi penyimpanan semua user yang terdaftar     |

---

## 9. Checklist Pemahaman

Tanya diri sendiri setelah membaca dokumen ini:

- [ ] Bisa menjelaskan kenapa ada folder model/service/repository?
- [ ] Paham kenapa validasi dilakukan di service, bukan di handler?
- [ ] Tahu alur data dari input user sampai tersimpan di file?
- [ ] Bisa jelaskan apa itu dependency injection di `main.go`?
- [ ] Mengerti cara kerja `utils.ReadJSON` dan `utils.WriteJSON`?

Jika semua terjawab, berarti pemahamanmu sudah solid. Tinggal latihan menambah fitur baru supaya makin mahir.

---

Selamat belajar! Bila ada bagian yang masih membingungkan, tinggal baca ulang per layer atau coba debug dengan menambahkan `fmt.Println` untuk melihat alur data secara real-time.
