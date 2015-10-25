package main

import (
    "database/sql"
    "log"
    "fmt"

    _ "github.com/lib/pq"
    . "gopkg.in/godo.v1"
)

func databaseUrl() string {
    return fmt.Sprintf(
        "host=%s port=5432 user=%s dbname=%s sslmode=disable",
        "192.168.44.88",
        "postgres",
        "postgres",
    )
}

func dbTasks(p *Project) {
    p.Task("create?", func() {
        var number_of_tables int
        db, err := sql.Open("postgres", databaseUrl())
        if err != nil {
            log.Fatal("Could not connect to database")
        }
        defer db.Close()

        err = db.
            QueryRow(`SELECT COUNT(*) AS n FROM pg_database WHERE datname='space_development'`).
            Scan(&number_of_tables)
        if err != nil {
            log.Fatal("Error while executing query")
        }

        if number_of_tables == 0 {
            db.Query(`CREATE DATABASE space_development OWNER postgres`)
        } else {
            fmt.Println("Database `space_development` already exists; skipping")
        }
    })

    p.Task("destroy?", func() {
        db, err := sql.Open("postgres", databaseUrl())
        if err != nil {
            log.Fatal("Could not connect to database")
        }
        defer db.Close()

        db.Query(`DROP DATABASE IF EXISTS space_development`)
    })

    p.Task("migrate", func() {

    })
}

func tasks(p *Project) {
    p.Task("default", D{"serve"})

    p.Task("server", func() {
        Bash(`go build && ./space`)
    })

    p.Use("db", dbTasks)

    p.Task("setup", D{"db:create", "db:migrate"})
}

func main() {
    Godo(tasks)
}
