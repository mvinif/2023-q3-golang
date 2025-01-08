package main

import (
    "fmt"
    "os"
    "time"
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
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

    config, err := pgxpool.ParseConfig(connectionString)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to parse pool config: %v\n", err)
        os.Exit(1)
    }

    config.MaxConns = 20
    config.MinConns = 0
    config.MaxConnLifetime = 30 * time.Minute
    config.MaxConnIdleTime = 5 * time.Minute
    config.HealthCheckPeriod = time.Minute

    conn, err := pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
        os.Exit(1)
    }
    defer conn.Close()

    r.POST("/pessoas", func(c *gin.Context){
        var newPerson createPerson
        uid := uuid.NewString()

        if err := c.BindJSON(&newPerson); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{ "error": "Invalid json", })
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
        id := c.Param("id")
        if id == "" {
            c.JSON(400, gin.H{"error":"Id should not be null"})
            return
        }

        uid, err := uuid.Parse(id)
        if err != nil {
            c.JSON(400, gin.H{"error":err.Error()})
            return
        }

        query := "select id, nickname, name, birthday, stacks from person where id = $1"

        var person readPerson
        err = conn.QueryRow(context.Background(), query, uid).Scan(
            &person.ID,
            &person.Nickname,
            &person.Name,
            &person.Birthday,
            &person.Stack)
        if err != nil {
            c.JSON(400, gin.H{"error":err.Error()})   
            return
        }

        c.JSON(http.StatusOK, person)
    })

    r.GET("/pessoas", func(c *gin.Context) {
        // pessoas := []pessoa
        searchTerm := c.Query("t")
        query := `select * from person where name ilike '%'||$1||'%' OR nickname ilike '%'||$1||'%' OR $1 = ANY(stacks)`
        rows, err := conn.Query(context.Background(), query, searchTerm)
        if err != nil {
            c.JSON(400, gin.H{"Error":err.Error()})   
            return
        }
        defer rows.Close()

        persons, err := pgx.CollectRows(rows, pgx.RowToStructByName[readPerson])
        if err != nil {
            c.JSON(400, gin.H{"Error":err.Error()})   
            return
        }

        c.JSON(http.StatusOK, persons)
    })

    r.GET("/contagem-pessoas", func(c *gin.Context) {
        var total int64
        err := conn.QueryRow(context.Background(), "select count(*) from person").Scan(&total)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
        } else {
            c.String(http.StatusOK, fmt.Sprintf("%d", total))
        }
    })

    r.Run(":80")
}
