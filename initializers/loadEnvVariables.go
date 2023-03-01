package initializers


import (
	"log"
	"github.com/joho/godotenv"
)
func LoadEnvVariables() {

	func init(){
		err := godotenv.Load()
	
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}