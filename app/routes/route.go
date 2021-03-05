package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mfaizfatah/story-tales/app/controllers"
	"github.com/mfaizfatah/story-tales/app/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/zephinzer/ezpromhttp"
)

// route struct with value Controllers Interface
type route struct {
	ctrl controllers.Controllers
}

// Router represent the Router contract
type Router interface {
	Router(port string)
}

/*NewRouter will create an object that represent the Router interface (Router)
 * @parameter
 * c - controllers Interface
 *
 * @represent
 * interface Router
 *
 * @return
 * struct route with value Controllers Interface
 */
func NewRouter(c controllers.Controllers) Router {
	return &route{ctrl: c}
}

func (c *route) Router(port string) {
	router := chi.NewRouter()

	router.Use(middleware.RecordMiddleware)

	router.Group(func(r chi.Router) {
		r.Use(ezpromhttp.InstrumentHandler)
		r.Get("/verify/{token}", c.ctrl.HandlerEmailVerification)
	})

	router.Group(func(r chi.Router) {
		r.Use(ezpromhttp.InstrumentHandler, middleware.CheckAPIKey)
		r.Post("/user/signup", c.ctrl.HandlerRegistration)
		r.Post("/user/login", c.ctrl.HandlerLogin)

		r.Get("/story", c.ctrl.HandlerGetAllStory)
		r.Get("/story/rekomendasi", c.ctrl.HandlerGetRekomenStory)
		r.Get("/story/{storyID}", c.ctrl.HandlerGetOneStory)
		r.Get("/story/{storyID}/{episodeID}", c.ctrl.HandlerGetDetailEpisode)

		r.Post("/banner", c.ctrl.HandlerCreateBanner)
		r.Get("/banner/{id}", c.ctrl.HandlerGetBannerDetail)
		r.Get("/banner", c.ctrl.HandlerGetListBannerThumb)
		r.Get("/author/story", c.ctrl.HandlerGetListBannerThumb)
		r.Post("/follow/{id}", c.ctrl.HandlerGetCountFollower)
		r.Get("/follower/count/{id}", c.ctrl.HandlerGetCountFollower)
		r.Get("/following/count/{id}", c.ctrl.HandlerGetCountFollowing)
		r.Get("/following/list{id}", c.ctrl.HandlerGetListFollowing)
		r.Get("/follower/list{id}", c.ctrl.HandlerGetListFollower)

		r.Get("/logout", c.ctrl.HandlerLogout)

		r.Post("/forgot-pass", c.ctrl.HandlerSendEmailForgotPass)
		r.Get("/forgot-pass/{token}", c.ctrl.HandlerValidateTokenForgotPass)
	})

	// group router if need to check session
	router.Group(func(r chi.Router) {
		r.Use(ezpromhttp.InstrumentHandler, middleware.CheckSession)
		r.Get("/user/check", c.ctrl.HandlerCheckSession)
		r.Patch("/forgot-pass", c.ctrl.HandlerChangePassword)

		r.Post("/story", c.ctrl.HandlerPostStory)
		r.Get("/story/favorite", c.ctrl.HandlerGetFavoriteStory)
		r.Post("/story/favorite", c.ctrl.HandlerPostFavoriteStory)
		r.Delete("/story/favorite/{storyID}", c.ctrl.HandlerDeleteFavoriteStory)
		r.Post("/story/rating", c.ctrl.HandlerPostRating)
		r.Delete("/story/rating/{storyID}", c.ctrl.HandlerDeleteRating)
		r.Post("/story/likes", c.ctrl.HandlerPostLikes)
		r.Delete("/story/likes/{storyID}/{episodeID}", c.ctrl.HandlerDeleteLikes)

	})

	router.MethodNotAllowed(middleware.NotAllowed)
	router.NotFound(middleware.NotFound)

	optsDoc := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	docs := middleware.Redoc(optsDoc, nil)
	router.Handle("/docs", docs)

	/* 	optsUI := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	   	ui := middleware.SwaggerUI(optsUI, nil)
	   	router.Handle("/docs", ui) */

	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	router.Handle("/metrics", promhttp.Handler())

	logrus.Infof("Server running on port : %s", port)
	logrus.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), cors.AllowAll().Handler(router)))
}
