package main
import ( 
"github.com/gin-gonic/gin"
"github.com/trey-rustand/go-crud/initializers"
)

func init(){
	initializers.LoadEnvVariables()
}
func main() {
	r := gin.Default()


	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}