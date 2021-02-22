package mailer

import "fmt"

const (
	// ForgotPasswordTitle ...
	ForgotPasswordTitle = `Forgot Password`
	// ForgotPasswordContent ...
	ForgotPasswordContent = `Anda telah mengajukan permintaan untuk merubah kata sandi pada akun Story Tales anda. Dengan alasan keamanan, tombol ubah kata sandi di bawah hanya berlaku selama lima belas menit. Silahkan klik tombol di bawah untuk merubah kata sandi akun Story Tales anda.`
	// ForgotPasswordFooter ...
	ForgotPasswordFooter = `Harap segera ganti password kamu atau abaikan email ini bila anda tidak pernah mengajukan permintaan perubahan kata sandi akun Story Tales.`
	// ForgotPasswordLabel ...
	ForgotPasswordLabel = `Ubah Kata Sandi`
)

// ForgotPasswordForeword ...
func ForgotPasswordForeword(nama string) string {
	return fmt.Sprintf("Hallo, %v", nama)
}

// NewforgotPassword ...
func NewforgotPassword(name, newPass string) Template {
	return &template{
		title:       ForgotPasswordTitle,
		foreword:    ForgotPasswordForeword(name),
		content:     ForgotPasswordContent,
		message:     newPass,
		buttonLabel: ForgotPasswordLabel,
		footer:      ForgotPasswordFooter,
	}
}
