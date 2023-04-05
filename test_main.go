package gorm_mysql_di

// import (
// 	"fmt"
// 	"os"
// 	"strings"

// 	"github.com/althenlimzixuan/gorm_mysql_di/app"
// 	"github.com/joho/godotenv"
// 	"github.com/sirupsen/logrus"
// )

// // This main.go serve to test the library

// func main() {

// 	target_env_file := ".env"
// 	env := os.Getenv("ENV")

// 	if !strings.HasPrefix(env, "prod") {

// 		// Load the .env_{lower(ENV)} file
// 		target_env_file = fmt.Sprintf("%s_%s", target_env_file, strings.ToLower(env))
// 	}

// 	err := godotenv.Load(target_env_file)

// 	if err != nil {
// 		logrus.Fatalf("Error loading environment variable: %s", target_env_file)
// 	}

// 	app_instance := app.ProvideAppService()

// 	go app_instance.StartApp()

// 	select {}

// }
