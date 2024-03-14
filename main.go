package main

import (
	"CampusSyncExport/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Replace yourfrontendport with the port your frontend is running on
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}      // Allowed HTTP methods
	config.AllowHeaders = []string{"*"}                                 // Allowed headers, "*" allows all headers
	server.Use(cors.New(config))
	// fmt.Println(result["attendancePeriodSetMap"])
	routes.RegisterRoutes(server)
	server.Run(":9090")
}
