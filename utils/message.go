package utils

import "errors"

var ErrEmailInvalid = errors.New("email tidak valid")
var ErrEmailExists = errors.New("email sudah terdaftar")
var ErrPhoneInvalid = errors.New("nomor telepon tidak valid")
var ErrPasswordInvalid = errors.New("password minimal 6 karakter")
var ErrUserNotFound = errors.New("user tidak ditemukan")
var ErrPasswordWrong = errors.New("password salah")
