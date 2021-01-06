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

	repo := repository.NewRepo(db, redis)
	uc := usecases.NewUC(repo)
	ctrl := controllers.NewCtrl(uc)

	router := routes.NewRouter(ctrl)
	router.Router(os.Getenv("SERVER_PORT"))
}
