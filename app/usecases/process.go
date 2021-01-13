package usecases

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/mfaizfatah/story-tales/app/models"
)

func (r *uc) GetUserFromToken(req *http.Request) (*models.User, string, int, error) {
	var (
		msg  string
		code = http.StatusOK
		err  error
		user models.User
	)

	token := req.Header.Get("x-app-token")
	if !reflect.ValueOf(token).IsZero() {
		err = json.Unmarshal([]byte(token), &user)
		if err != nil {
			msg = ErrInvalidHeader
			code = http.StatusBadRequest
			return nil, msg, code, err
		}
	}

	return &user, msg, code, nil
}
