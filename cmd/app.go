package main

import (
    "fmt"
    "os"
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5"
    "github.com/google/uuid"
)

type readPerson struct {
    ID string `json:"id"`
    Nickname string `json:"apelido"`
    Name string `json:"nome"`
    Birthday string `json:"nascimento"`
    Stack []string `json:"stack"`
}

type createPerson struct {
    Nickname string `json:"apelido"`
    Name string `json:"nome"`
    Birthday string `json:"nascimento"`
    Stack []string `json:"stack"`
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
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
        os.Exit(1)
    }
    defer conn.Close(context.Background())

    r.POST("/pessoas", func(c *gin.Context){
        var newPerson createPerson
        uid := uuid.NewString()

        if err := c.BindJSON(&newPerson); err != nil {
            return
        }

        // validate Pessoa object
        if !validate(newPerson) {
            c.JSON(http.StatusBadRequest, gin.H{ "error": "Invalid person data", })
            return
        }
        query := "insert into person values ($1, $2, $3, $4, $5::varchar[]);"
        _, err := conn.Exec(context.Background(),
			 query,
			 uid,
			 newPerson.Nickname,
			 newPerson.Name,
			 newPerson.Birthday,
			 newPerson.Stack)

        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{ "error": err, })
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
        searchTerm := c.Query("t")
        var searchResult string
        err = conn.QueryRow(context.Background(), "select * from person").Scan(&searchResult)
        c.String(http.StatusOK, fmt.Sprintf("%s %s", searchResult, searchTerm))
    })

    r.GET("/contagem-pessoas", func(c *gin.Context) {
        var total int64
        err = conn.QueryRow(context.Background(), "select count(*) from person").Scan(&total)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{ "error": err, })
        }
        c.String(http.StatusOK, fmt.Sprintf("%d", total))
    })

    r.Run(":80")
}
