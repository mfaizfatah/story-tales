package config

import (
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mfaizfatah/story-tales/app/adapter"
)

/*LoadConfig is for load all configuration, connection etc
 * @parameters
 * service is for service name
 */
func LoadConfig(service string) {
	dbdriver := strings.ToLower(os.Getenv("DB_DRIVER"))
	if dbdriver == "mysql" {
		sql := os.Getenv("DB_URI")
		adapter.LoadMySQL(sql)
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	adapter.LoadRedis(redisHost, redisPort)

	adapter.LoadMongoDB(os.Getenv("MONGO_URI"))

	// utils.Newprometheus(service)
}
