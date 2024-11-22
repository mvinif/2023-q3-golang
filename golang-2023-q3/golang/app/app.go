package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func setupRouter() {
    router := gin.Default()
    router.GET("/ping", func(context *gin.Context) {
        // context.JSON(http.StatusOK, gin.H{ "message":"pong", })
        context.String(http.StatusOK, "pong")

    })
    return router
}

func main() {
    router.Run()
}
