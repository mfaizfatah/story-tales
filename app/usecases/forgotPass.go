package usecases

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mfaizfatah/story-tales/app/helpers/encryption"
	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

//response ...
type response struct {
	id   string
	text string
}

//Tokens ...
type Tokens interface {
	ID() string
	Value() string
}

//ID ...
func (r *response) ID() string {
	return r.id
}

//Value ...
func (r *response) Value() string {
	return r.text
}

//TokenForgotPass is generate token for url
func TokenForgotPass(text, kode string, d time.Duration) string {

	expired := time.Now().Add(d).UnixNano()

	id := encryption.RdAlpnum(6)

	compose := fmt.Sprintf("%v~%v~%v~%v", id, text, expired, kode)

	return encryption.EncryptorTokenForgotPass(compose)
}

//ValidateToken is func to parse token
func ValidateToken(token, kode string, CheckExpired bool) (Tokens, error) {
	text, err := encryption.DecryptorTokenForgotPass(token)
	if err != nil {
		return nil, err
	}
	field := strings.Split(text, `~`)

	raw, err := strconv.Atoi(field[2])
	if err != nil {
		return nil, err
	}
	if CheckExpired {
		exp := time.Unix(0, int64(raw))
		if time.Now().After(exp) {
			return nil, errors.New(`expired token`)
		}
	}

	if field[3] != kode {
		return nil, errors.New("invalid type of link")
	}

	return &response{
		id:   field[0],
		text: field[1],
	}, nil
}

func (r *uc) SendLinkForgotPass(ctx context.Context, req *models.User) (context.Context, interface{}, string, int, error) {
	var (
		res  interface{}
		msg  string
		code int
		err  error

		user = new(models.User)
	)

	err = r.query.FindOne(tableUser, user, "email = ?", "id, email, name", req.Email)
	if err != nil {
		return ctx, nil, ErrServer, http.StatusInternalServerError, err
	}
	ctx = logger.Logf(ctx, "user() => %v", user)
	if user.Email == "" {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	randString := uuid.New().String()
	rand := strings.Split(randString, "-")
	newPass := rand[0]

	data := make(map[string]interface{})
	data["password"] = newPass

	err = r.query.Update(tableUser, user, data)
	if err != nil {
		return ctx, nil, ErrServer, http.StatusInternalServerError, err
	}

	go r.smtp.ForgotPassword(user.Name, newPass).SendEmail(user.Email)

	code = http.StatusOK
	res = "email sent successfully"
	return ctx, res, msg, code, nil
}

func (r *uc) ValidateTokenForgotPass(ctx context.Context, tokenForgotPass string) (context.Context, string, int, error) {
	var (
		msg  string
		code = http.StatusAccepted
		err  error
	)

	result, err := ValidateToken(tokenForgotPass, "kode", true)
	if err != nil {
		ctx = logger.Logf(ctx, "error reset pass: %v", err)
		msg = "Expired reset password"
		return ctx, msg, http.StatusUnauthorized, err
	}

	msg = "token accepted with email: " + result.Value()

	return ctx, msg, code, nil
}

func (r *uc) ChangePassword(ctx context.Context, idUser int, req *models.ForgotPass) (context.Context, string, int, error) {
	var (
		msg  string
		code = http.StatusOK
		err  error

		user = new(models.User)
		// sha  = sha1.New()
	)

	err = r.query.FindOne(tableUser, user, "id = ?", "id, email", idUser)
	if err != nil {
		msg = "email not found"
		return ctx, msg, http.StatusNotFound, err
	}

	if req.Password != req.RepeatPassword {
		msg = "password not match"
		return ctx, msg, http.StatusForbidden, errors.New("password not match")
	}

	// sha.Write([]byte(req.Password))
	// encrypted := sha.Sum(nil)

	data := make(map[string]interface{})
	data["password"] = fmt.Sprintf("%s", req.Password)

	go r.query.Update(tableUser, user, data)

	msg = "success change password => email: " + user.Email
	return ctx, msg, code, err
}
