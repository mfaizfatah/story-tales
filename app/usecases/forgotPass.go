package usecases

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mfaizfatah/story-tales/app/helpers/encryption"
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
func ValidateToken(token, kode string, IsExpired bool) (Tokens, error) {
	text, err := encryption.DecryptorTokenForgotPass(token)
	if err != nil {
		return nil, err
	}
	field := strings.Split(text, `~`)

	raw, err := strconv.Atoi(field[2])
	if err != nil {
		return nil, err
	}
	if IsExpired {
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

		user models.User
	)

	err = r.query.FindOne(tableUser, user, "email = ?", "id, email", req.Email)
	if user.Email != "" || err != nil {
		return ctx, nil, ErrNotFound, http.StatusNotFound, repository.ErrRecordNotFound
	}

	url := "url"
	token := TokenForgotPass(user.Email, "kode", 5*time.Minute)
	link := fmt.Sprintf("%v/%v", url, token)

	go r.smtp.ForgotPassword(user.Name, link).SendEmail(user.Email)

	code = http.StatusOK
	res = "email sent successfully"
	return ctx, res, msg, code, nil
}
