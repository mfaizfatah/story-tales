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
	HandlerCheckSession(w http.ResponseWriter, r *http.Request)
	HandlerEmailVerification(w http.ResponseWriter, r *http.Request)

	// Story
	HandlerPostStory(w http.ResponseWriter, r *http.Request)
	HandlerGetOneStory(w http.ResponseWriter, r *http.Request)
	HandlerGetAllStory(w http.ResponseWriter, r *http.Request)
	HandlerGetAuthorStory(w http.ResponseWriter, r *http.Request)
	HandlerGetDetailEpisode(w http.ResponseWriter, r *http.Request)
	HandlerGetRekomenStory(w http.ResponseWriter, r *http.Request)
	HandlerGetFavoriteStory(w http.ResponseWriter, r *http.Request)
	HandlerPostFavoriteStory(w http.ResponseWriter, r *http.Request)
	HandlerDeleteFavoriteStory(w http.ResponseWriter, r *http.Request)
	HandlerGetStoryGenre(w http.ResponseWriter, r *http.Request)

	// Likes
	HandlerPostLikes(w http.ResponseWriter, r *http.Request)
	HandlerDeleteLikes(w http.ResponseWriter, r *http.Request)
	HandlerGetMyLike(w http.ResponseWriter, r *http.Request)

	// Comment
	HandlerGetTopComment(w http.ResponseWriter, r *http.Request)
	HandlerGetComment(w http.ResponseWriter, r *http.Request)
	HandlerPostComment(w http.ResponseWriter, r *http.Request)
	HandlerDeleteComment(w http.ResponseWriter, r *http.Request)
	HandlerGetMyComment(w http.ResponseWriter, r *http.Request)
	HandlerPostLikeComment(w http.ResponseWriter, r *http.Request)

	// Rating
	HandlerPostRating(w http.ResponseWriter, r *http.Request)
	HandlerDeleteRating(w http.ResponseWriter, r *http.Request)

	// Forgot-Pass
	HandlerSendEmailForgotPass(w http.ResponseWriter, r *http.Request)
	HandlerValidateTokenForgotPass(w http.ResponseWriter, r *http.Request)
	HandlerChangePassword(w http.ResponseWriter, r *http.Request)

	// Banner
	HandlerCreateBanner(w http.ResponseWriter, r *http.Request)
	HandlerGetBannerDetail(w http.ResponseWriter, r *http.Request)
	HandlerGetListBannerThumb(w http.ResponseWriter, r *http.Request)
	HandlerBannerPostPic(w http.ResponseWriter, r *http.Request)

	// UserFollow
	HandlerGetFollow(w http.ResponseWriter, r *http.Request)
	HandlerPostFollow(w http.ResponseWriter, r *http.Request)
	HandlerGetCountFollower(w http.ResponseWriter, r *http.Request)
	HandlerGetCountFollowing(w http.ResponseWriter, r *http.Request)
	HandlerGetListFollower(w http.ResponseWriter, r *http.Request)
	HandlerGetListFollowing(w http.ResponseWriter, r *http.Request)

	// User
	HandlerGetExistAuthor(w http.ResponseWriter, r *http.Request)
	HandlerGetExistUser(w http.ResponseWriter, r *http.Request)
	HandlerGetAuthorProfile(w http.ResponseWriter, r *http.Request)
	HandlerGetUserProfile(w http.ResponseWriter, r *http.Request)
	HandlerUpdateAuthor(w http.ResponseWriter, r *http.Request)
	HandlerUpdateUser(w http.ResponseWriter, r *http.Request)
	HandlerGetUserInfo(w http.ResponseWriter, r *http.Request)

	// test jenkins
	HandlerTestJenkins(w http.ResponseWriter, r *http.Request)

	// Searching
	HandlerSearching(w http.ResponseWriter, r *http.Request)
	HandlerGenerateDoc(w http.ResponseWriter, r *http.Request)
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
