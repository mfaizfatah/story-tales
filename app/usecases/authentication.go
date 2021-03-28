package usecases

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

const (
	// TableUser is table for user
	tableUser = "users"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func (r *uc) Registration(ctx context.Context, req *models.User) (context.Context, *models.ResponseLogin, string, int, error) {
	var (
		// sha  = sha1.New()
		res  = new(models.ResponseLogin)
		user = new(models.User)
		msg  string
		err  error
	)

	if req == nil || !isEmailValid(req.Email) {
		return ctx, nil, ErrBadRequest, http.StatusBadRequest, repository.ErrBadRequest
	}

	err = r.query.FindOne(tableUser, user, "email = ?", "id, email, username, telp", req.Email)
	if user.Email != "" || user.Username != "" {
		return ctx, nil, ErrAlreadyEmail, http.StatusConflict, repository.ErrConflict
	}

	user = req
	/*
		sha.Write([]byte(user.Password))
		encrypted := sha.Sum(nil)

	user.Password = fmt.Sprintf("%x", encrypted)*/
	user.Password = req.Password
	user.DateOfBirth = req.DateOfBirth

	err = r.query.Insert(tableUser, user)
	if err != nil {
		return ctx, nil, ErrCreated, http.StatusInternalServerError, err
	}
	/*
		ctx, token, duration, err := r.GenerateToken(ctx, user)
		if err != nil {
			return ctx, nil, ErrCreated, http.StatusInternalServerError, err
		}

		res.Token.Key = "bearer"
		res.Token.Value = token
		res.Token.ExpiredIn = fmt.Sprintf("%v", duration)
	*/
	res.Message = "Silahkan verifikasi email atau cek email anda, Terima Kasih!"

	go r.SendLinkVerification(user.Email)

	return ctx, res, msg, http.StatusCreated, err
}

func (r *uc) Login(ctx context.Context, req *models.User) (context.Context, *models.ResponseLogin, string, int, error) {
	var (
		// sha  = sha1.New()
		res  = new(models.ResponseLogin)
		user = new(models.User)
		msg  string
		err  error
	)

	isEmail := isEmailValid(req.User)
	if isEmail {
		req.Email = req.User
	} else {
		req.Username = req.User
	}

	err = r.query.FindOne(tableUser, user, "email = ? OR username = ?", "id, email, password, email_verify", req.Email, req.Username)
	if err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	if user.EmailVerification != 1 {
		return ctx, nil, "You must verify your email!", http.StatusNotAcceptable, errors.New("email_not_verified")
	}
	/*
		sha.Write([]byte(req.Password))
		encrypted := sha.Sum(nil)

		req.Password = fmt.Sprintf("%x", encrypted)
	*/
	if req.Password != user.Password {
		return ctx, nil, ErrNotMatch, http.StatusUnauthorized, repository.ErrUnouthorized
	}

	ctx = logger.Logf(ctx, "user() => %v", user)
	ctx, token, duration, err := r.GenerateToken(ctx, user)
	if err != nil {
		return ctx, nil, ErrCreateToken, http.StatusInternalServerError, err
	}

	res.Token.Key = "bearer"
	res.Token.Value = token
	res.Token.ExpiredIn = fmt.Sprintf("%v", duration)
	res.Message = "Login Success"

	return ctx, res, msg, http.StatusAccepted, nil
}

func (r *uc) Logout(ctx context.Context, token string) (context.Context, interface{}, string, int, error) {
	result, err := r.query.DeleteRedis(token)
	if err != nil {
		return ctx, nil, "Logout Gagal", http.StatusInternalServerError, err
	}
	return ctx, result, "Logout Berhasil", http.StatusOK, nil
}
func (r *uc) CheckSession(ctx context.Context, req *models.User, token string) (context.Context, interface{}, string, int, error) {
	var (
		res  models.TokenResponse
		msg  string
		code = http.StatusOK
		err  error
	)

	result, err := r.query.FindToken(token)
	if err != nil {
		msg = "token expired or not exist"
		return ctx, nil, msg, http.StatusNotFound, err
	}
	ctx = logger.Logf(ctx, "token value() => %v", result)

	ttl, err := r.query.GetTTLRedis(token)
	if err != nil {
		msg = "token expired or not exist"
		return ctx, nil, msg, http.StatusNotFound, err
	}

	res.Key = token
	res.Value = fmt.Sprintf("idUser = %v", req.ID)
	res.ExpiredIn = fmt.Sprintf("%d", ttl)

	return ctx, res, msg, code, nil
}

// segment verification email
func (r *uc) SendLinkVerification(email string) error {
	url := "https://api-storytales.digisoul.id/verify"
	token := TokenForgotPass(email, "kode", 15*time.Minute)
	link := fmt.Sprintf("%v/%v", url, token)

	err := r.smtp.EmailVerification(link).SendEmail(email)
	if err != nil {
		log.Printf("failed send email to user() => %v :: error() => %v", email, err)
		return err
	}
	return nil
}

func (r *uc) EmailVerification(ctx context.Context, token string) (context.Context, interface{}, int, error) {
	var (
		res  interface{}
		code = http.StatusAccepted
		err  error
		user = new(models.User)
	)
	result, err := ValidateToken(token, "kode", false)
	if err != nil {
		res = `<!DOCTYPE html><html><body><h1>invalid link verification</h1></body></html>`
		return ctx, res, http.StatusUnauthorized, err
	}

	email := result.Value()
	err = r.query.FindOne(tableUser, user, "email = ?", "id, email", email)
	if err != nil {
		res = `<!DOCTYPE html><html><body><h1>email not found</h1></body></html>`
		return ctx, res, http.StatusNotFound, err
	}

	data := make(map[string]interface{})
	data["email_verify"] = 1

	err = r.query.Update(tableUser, user, data)
	if err != nil {
		res = `<!DOCTYPE html><html><body><h1>error while update data</h1></body></html>`
		return ctx, res, http.StatusInternalServerError, err
	}

	res = `<!DOCTYPE html><html><body><h1>Your Email is Verified</h1></body></html>`
	return ctx, res, code, nil
}

// end segment
