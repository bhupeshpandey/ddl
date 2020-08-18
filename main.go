package main
import (
	"com.webfrm.app/config"
	"com.webfrm.app/models"
)

var configuration *models.Configurations
func main() {
	configuration = config.ReadConfig()


}
