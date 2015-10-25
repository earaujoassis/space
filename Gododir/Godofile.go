package main

import (
    . "gopkg.in/godo.v1"
)

func tasks(p *Project) {
    p.Task("default", D{"serve"})

    p.Task("serve", func() {
        Bash(`go run *.go`)
    })
}

func main() {
    Godo(tasks)
}
