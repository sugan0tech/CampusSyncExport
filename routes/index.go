package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexGreet(context *gin.Context){
  fmt.Println("Got an entry")
  context.JSON(http.StatusOK, gin.H{"message":"From go service"})
}
