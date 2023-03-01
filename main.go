package main
import ( 
"github.com/gin-gonic/gin"
"log"
"github.com/joho/godotenv"
)
func init(){
	err := godotenv.Load()
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}}
func main() {
	r := gin.Default()


	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}