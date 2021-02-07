package mailer

const (
	VerificationTitle = `Email Verification`

	VerificationForeword = `Selamat datang di Story Tales`

	VerificationContent = `Terima kasih sudah mendaftarkan diri Anda di Story Tales. Sebentar lagi Anda akan dapat menikmati akses mudah belajar dengan Story Tales! Silahkan verifikasi email Anda dengan klik tautan berikut:`

	VerificationFooter = `Penting untuk memiliki akun dengan alamat email yang akurat karena semua keterangan aktivitas belajar dan mengajar akan dikirimkan kesini. Harap abaikan email ini jika Anda tidak pernah mendaftar di Story Tales.`

	VerificationLabel = `Konfirmasi email`
)

//NewVerification ...
func NewVerification(link string) Template {
	return &template{
		title:       VerificationTitle,
		foreword:    VerificationForeword,
		content:     VerificationContent,
		buttonLink:  link,
		buttonLabel: VerificationLabel,
		footer:      VerificationFooter,
	}
}
