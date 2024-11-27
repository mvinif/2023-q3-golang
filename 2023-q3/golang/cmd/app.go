package main

import (
    "fmt"
    "os"
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5"
)

type readPerson struct {
    ID string `json:id`
    Nickname string `json:apelido`
    Name string `json:nome`
    Birthday string `json:nascimento`
    Stack []string `json:stack`
}

type createPerson struct {
    Nickname string `json:apelido`
    Name string `json:nome`
    Birthday string `json:nascimento`
    Stack []string `json:stack`
}

func validate(p createPerson) bool {

    if len(p.Nickname) > 32 {
        return false
	}

    if len(p.Nickname) == 0 {
        return false
	}

    if len(p.Name) > 100 {
        return false
	}

    if len(p.Name) == 0 {
        return false
	}

    for _, s := range p.Stack {
        if len(s) > 32 {
            return false
        }
    }

    return true
}

func main() {
    r := gin.Default()
    connectionString := os.Getenv("DATABASE_URL")

    conn, err := pgx.Connect(context.Background(), connectionString)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect")
    }
    defer conn.Close(context.Background())

    r.POST("/pessoas", func(c *gin.Context){
        var newPerson createPerson

        if err := c.BindJSON(&newPerson); err != nil {
            return
        }

        // validate Pessoa object
        if !validate(newPerson) {
        }

        c.JSON(http.StatusCreated, newPerson)
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

    r.Run(":80")
}
