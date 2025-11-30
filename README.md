# 🔐 Authentication System CLI

Aplikasi ini merupakan implementasi sistem login sederhana yang berjalan di terminal dengan menggunakan bahasa Go. Sistem ini dilengkapi dengan berbagai validasi untuk memastikan data yang diinput oleh user sesuai dengan ketentuan yang telah ditetapkan.

## ✨ Fitur

### 1. **Register**

Saat user memilih menu Register, program akan meminta input:

- **Full Name**: Nama lengkap user
- **Email**: Alamat email dengan validasi format
- **Phone Number**: Nomor telepon dengan validasi digit dan panjang
- **Password**: Password dengan minimal karakter tertentu

### 2. **Login**

Saat user memilih menu Login, program akan:

- Meminta Email dan Password
- Memverifikasi kredensial dengan data yang tersimpan
- Menampilkan pesan sukses jika berhasil login

### 3. **Exit**

Keluar dari aplikasi

## 🔒 Validasi Register

Program melakukan **validasi** pada saat Register dengan ketentuan sebagai berikut:

### ✅ Email

- Wajib berformat benar (mengandung `@` dan `.` atau gunakan pola regex sederhana)
- Tidak boleh sama dengan email yang sudah terdaftar di file JSON

### ✅ Phone Number

- Hanya boleh berisi angka (0-9)
- Panjang minimal **10 digit** dan maksimal **15 digit**

### ✅ Password

- Minimal **6 karakter**

### ⚠️ Error Handling

Jika ada validasi yang gagal, tampilkan pesan error yang sesuai dan **jangan simpan** data ke file.

**Pesan Error yang mungkin muncul:**

- `email tidak valid` - Format email salah
- `email sudah terdaftar` - Email sudah digunakan
- `nomor telepon tidak valid` - Phone number tidak sesuai ketentuan
- `password minimal 6 karakter` - Password terlalu pendek
- `user tidak ditemukan` - Email tidak terdaftar saat login
- `password salah` - Password tidak cocok saat login

## 🚀 Cara Menjalankan

```bash
# Masuk ke direktori project
cd authentication-system-cli

# Jalankan aplikasi
go run main.go
```

## Contoh Response

### **Register Berhasil**

```
=== SIMPLE LOGIN SYSTEM ===
1. Register
2. Login
3. Exit
Pilih menu: 1

--- REGISTER ---
Full Name       : Budi Santoso
Email           : budi@example.com
Phone           : 081234567890
Password        : rahasia123
Registrasi berhasil! Data tersimpan di users.json
```

### **Login Berhasil**

```
=== SIMPLE LOGIN SYSTEM ===
1. Register
2. Login
3. Exit
Pilih menu: 2

--- LOGIN ---
Email           : budi@example.com
Password        : rahasia123
Login berhasil, selamat datang Budi Santoso
```

## 👨‍💻 Pengembang

Dibuat oleh: **[Alvin Rama Saputra](https://github.com/Alvinnn-R/)**

Project ini dibuat sebagai bagian dari Challenge **Bootcamp Golang - Intermediate Daytime Class** di Lumoshive Academy.
