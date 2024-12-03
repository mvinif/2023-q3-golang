package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)


func CreateServer() *gin.Engine {
    r := gin.Default()
    r.Use(gin.Logger())
    r.Run(":80")
}
