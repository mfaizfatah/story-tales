// Package classification Story Tales API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
//     Schemes: http, https
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: x-api-key
//          in: header
//
//
// swagger:meta
package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mfaizfatah/story-tales/app/adapter"
	"github.com/mfaizfatah/story-tales/app/config"
	"github.com/mfaizfatah/story-tales/app/controllers"
	"github.com/mfaizfatah/story-tales/app/repository"
	"github.com/mfaizfatah/story-tales/app/routes"
	"github.com/mfaizfatah/story-tales/app/usecases"
)

func init() {
	service := "story-tales-api"

	config.LoadConfig(service)
}

func main() {
	db := adapter.DBSQL()
	redis := adapter.UseRedis()
	mail := adapter.NewSMTPClient()
	mongodb := adapter.MongoDatabase()

	repo := repository.NewRepo(db, redis, mongodb)
	uc := usecases.NewUC(repo, mail)
	ctrl := controllers.NewCtrl(uc)

	router := routes.NewRouter(ctrl)
	router.Router(os.Getenv("SERVER_PORT"))
}
