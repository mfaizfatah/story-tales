package controllers

import (
	"net/http"

	"github.com/mfaizfatah/story-tales/app/usecases"
)

// ctrl struct with value interface Usecases
type ctrl struct {
	uc usecases.Usecases
}

// Controllers represent the Controllers contract
type Controllers interface {
	// Authentication controllers
	HandlerRegistration(w http.ResponseWriter, r *http.Request)
	HandlerLogin(w http.ResponseWriter, r *http.Request)
	HandlerLogout(w http.ResponseWriter, r *http.Request)
	// Story
	HandlerPostStory(w http.ResponseWriter, r *http.Request)
	HandlerGetOneStory(w http.ResponseWriter, r *http.Request)
	HandlerCheckSession(w http.ResponseWriter, r *http.Request)
	HandlerGetAllStory(w http.ResponseWriter, r *http.Request)
	// Forgot-Pass
	HandlerSendEmailForgotPass(w http.ResponseWriter, r *http.Request)
	HandlerValidateTokenForgotPass(w http.ResponseWriter, r *http.Request)
	HandlerChangePassword(w http.ResponseWriter, r *http.Request)
}

/*NewCtrl will create an object that represent the Controllers interface (Controllers)
 * @parameter
 * r - Repository Interface
 *
 * @represent
 * interface Controllers
 *
 * @return
 * uc struct with value interface Usecases
 */
func NewCtrl(u usecases.Usecases) Controllers {
	return &ctrl{uc: u}
}
