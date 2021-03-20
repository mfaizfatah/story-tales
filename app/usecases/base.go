package usecases

import (
	"context"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/adapter"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

// all variable const
const (
	// all variable for error
	ErrServer        = "Something wrong with our Server. Please try again later. Thank you"
	ErrNotFound      = "User not found"
	ErrCreated       = "Error when create a new User. Please try again later. Thank you"
	ErrAlreadyEmail  = "Email already created. Please use another e-mail. Thank you"
	ErrAlreadyPhone  = "Phone number already created. Please use another Phone Number. Thank you"
	ErrNotVerified   = "Your e-mail is not Verified"
	ErrBadRequest    = "Your Request is Invalid. Please check the payload"
	ErrNotMatch      = "Email or Password not match"
	ErrInvalidHeader = "Invalid Header"
	ErrTimezones     = "Timezone for Asia/Jakarta not found in our Server. Please try again later. Thank you"
	ErrEncryption    = "Encryption failed"
	ErrCreateToken   = "Error when create a token"

	// layout date
	LayoutDate = "2006-01-02 15:04:05"
)

// uc struct with value interface Repository
type uc struct {
	query repository.Repo
	smtp  adapter.MailClient
}

// Usecases represent the Usecases contract
type Usecases interface {
	// Authentication for logic auth
	Registration(ctx context.Context, req *models.User) (context.Context, *models.ResponseLogin, string, int, error)
	Login(ctx context.Context, req *models.User) (context.Context, *models.ResponseLogin, string, int, error)
	Logout(ctx context.Context, token string) (context.Context, interface{}, string, int, error)
	CheckSession(ctx context.Context, req *models.User, token string) (context.Context, interface{}, string, int, error)
	EmailVerification(ctx context.Context, token string) (context.Context, interface{}, int, error)

	//Story
	GetOneStory(ctx context.Context, storyID int) (context.Context, *models.ResponseOneStory, string, int, error)
	GetAllStory(ctx context.Context) (context.Context, []models.ResponseAllStory, string, int, error)
	GetDetailEpisode(ctx context.Context, storyID, episodeID int) (context.Context, *models.ResponseDetailEpisode, string, int, error)
	PostStory(ctx context.Context, req *models.Story, userid int) (context.Context, string, int, error)
	GetRekomendasiStory(ctx context.Context) (context.Context, []models.ResponseRekomenStory, string, int, error)
	GetFavoriteStory(ctx context.Context, limit, userid int) (context.Context, []models.ResponseFavoriteStory, string, int, error)
	GetCheckFavoriteStory(ctx context.Context, storyid, userid int) (context.Context, *models.ResponseCheckFavorite, string, int, error)
	GetLoadFavoriteStory(ctx context.Context, limit, storyid, userid int) (context.Context, []models.ResponseFavoriteStory, string, int, error)
	PostFavoriteStory(ctx context.Context, req *models.PostFavoriteStory, userid int) (context.Context, string, int, error)
	DeleteFavoriteStory(ctx context.Context, storyid, userid int) (context.Context, string, int, error)

	//Likes
	PostLikes(ctx context.Context, req *models.Likes, userid int) (context.Context, string, int, error)
	DeleteLikes(ctx context.Context, storyid, episodeid, userid int) (context.Context, string, int, error)

	//Rating
	PostRating(ctx context.Context, req *models.Rating, userid int) (context.Context, string, int, error)
	DeleteRating(ctx context.Context, storyid, userid int) (context.Context, string, int, error)

	//Banner
	CreateBanner(ctx context.Context, req *models.BannerReq) (context.Context, string, int, error)
	GetBannerDtl(ctx context.Context, id int) (context.Context, *models.BannerDetailRs, string, int, error)
	GetListBannerThumb(ctx context.Context) (context.Context, []models.ListBannerThumbRs, string, int, error)

	//UserFollow
	PostFollow(ctx context.Context, id int) (context.Context, string, int, error)
	GetCountFollowing(ctx context.Context, id int) (context.Context, *models.UserCountFollowing, string, int, error)
	GetCountFollower(ctx context.Context, id int) (context.Context, *models.UserCountFollower, string, int, error)
	GetListFollower(ctx context.Context, id int) (context.Context, []models.ListFollower, string, int, error)
	GetListFollowing(ctx context.Context, id int) (context.Context, []models.ListFollowing, string, int, error)

	// forgot pass
	SendLinkForgotPass(ctx context.Context, req *models.User) (context.Context, interface{}, string, int, error)
	ValidateTokenForgotPass(ctx context.Context, tokenForgotPass string) (context.Context, string, int, error)
	ChangePassword(ctx context.Context, idUser int, req *models.ForgotPass) (context.Context, string, int, error)

	// Process token
	GetUserFromToken(req *http.Request) (*models.User, string, int, error)

	// Searching
	Searching(ctx context.Context, query string) (context.Context, interface{}, string, int, error)
}

/*NewUC will create an object that represent the Usecases interface (Usecases)
 * @parameter
 * r - Repository Interface
 *
 * @represent
 * interface Usecases
 *
 * @return
 * uc struct with value interface Repository
 */
func NewUC(r repository.Repo, m adapter.MailClient) Usecases {
	return &uc{query: r, smtp: m}
}
