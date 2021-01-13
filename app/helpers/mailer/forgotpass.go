package mailer

import "fmt"

const (
	// ForgotPasswordTitle ...
	ForgotPasswordTitle = `Forgot Password`
	// ForgotPasswordContent ...
	ForgotPasswordContent = `Anda telah mengajukan permintaan untuk merubah kata sandi pada akun Ajar Uji anda. Dengan alasan keamanan, tombol ubah kata sandi di bawah hanya berlaku selama lima belas menit. Silahkan klik tombol di bawah untuk merubah kata sandi akun Ajar Uji anda.`
	// ForgotPasswordFooter ...
	ForgotPasswordFooter = `Harap abaikan email ini bila anda tidak pernah mengajukan permintaan perubahan kata sandi akun Ajar Uji.`
	// ForgotPasswordLabel ...
	ForgotPasswordLabel = `Ubah Kata Sandi`
)

// ForgotPasswordForeword ...
func ForgotPasswordForeword(nama string) string {
	return fmt.Sprintf("Hallo, %v", nama)
}

// NewforgotPassword ...
func NewforgotPassword(name, link string) Template {
	return &template{
		title:       ForgotPasswordTitle,
		foreword:    ForgotPasswordForeword(name),
		content:     ForgotPasswordContent,
		buttonLink:  link,
		buttonLabel: ForgotPasswordLabel,
		footer:      ForgotPasswordFooter,
	}
}
