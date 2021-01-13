package mailer

const (
	VerificationTitle = `Email Verification`

	VerificationForeword = `Selamat datang di Ajar Uji`

	VerificationContent = `Terima kasih sudah mendaftarkan diri Anda di Ajar Uji. Sebentar lagi Anda akan dapat menikmati akses mudah belajar dengan Ajar Uji! Silahkan verifikasi email Anda dengan klik tautan berikut:`

	VerificationFooter = `Penting untuk memiliki akun dengan alamat email yang akurat karena semua keterangan aktivitas belajar dan mengajar akan dikirimkan kesini. Harap abaikan email ini jika Anda tidak pernah mendaftar di Ajar Uji.`

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
