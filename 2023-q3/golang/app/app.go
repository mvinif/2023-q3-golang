package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type pessoa struct {
    ID string `json:id`
    Apelido string `json:apelido`
    Nome string `json:nome`
    Nascimento string `json:nascimento`
    Stack []string `json:stack`
}

func validateData(p pessoa) {
    if l := len(p.Apelido); l > 32 {
        return false
    }

    if l := len(p.Nome); l > 100 {
        return false
    }
}

func main() {
    r := gin.Default()

    r.POST("/pessoas", func(c *gin.Context){
        var novaPessoa pessoa

        if err := c.BindJSON(&novaPessoa); err != nil {
            return
        }

        // validate Pessoa object
        if validateData(novaPessoa) {
            
        }
        c.JSON(http.StatusCreated, novaPessoa)
    })

    r.GET("/pessoas/:id", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "id":"pong", 
            "apelido":"pong", 
            "nome":"pong", 
            "nascimento":"pong", 
            "stack":"pong", 
        })
    })

    r.GET("/pessoas", func(c *gin.Context) {
        // pessoas := []pessoa
        // searchterm := c.Query("t")
        c.JSON(http.StatusOK, gin.H{ 
            "id":"pong", 
            "apelido":"pong", 
            "nome":"pong", 
            "nascimento":"pong", 
            "stack":"pong", 
        })
    })

    r.GET("/contagem-pessoas", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{ "message":"pong", })
    })

    r.Run()
}
