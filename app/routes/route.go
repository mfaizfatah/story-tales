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
		r.Use(ezpromhttp.InstrumentHandler, middleware.CheckAPIKey)
		r.Post("/user/signup", c.ctrl.HandlerRegistration)
		r.Post("/user/login", c.ctrl.HandlerLogin)
		r.Post("/story", c.ctrl.HandlerPostStory)
		r.Get("/story", c.ctrl.HandlerGetAllStory)
		r.Get("/story/{storyID}", c.ctrl.HandlerGetOneStory)
		r.Get("/logout", c.ctrl.HandlerLogout)
	})

	// group router if need to check session
	router.Group(func(r chi.Router) {
		r.Use(ezpromhttp.InstrumentHandler, middleware.CheckSession)

	})

	router.MethodNotAllowed(middleware.NotAllowed)
	router.NotFound(middleware.NotFound)

	router.Handle("/metrics", promhttp.Handler())

	logrus.Infof("Server running on port : %s", port)
	logrus.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), cors.AllowAll().Handler(router)))
}
