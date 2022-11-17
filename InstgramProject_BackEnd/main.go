package main

import (
    "fmt"
    "log"
    "net/http"
    //jwtmiddleware "github.com/auth0/go-jwt-middleware"
    //jwt "github.com/form3tech-oss/jwt-go"
    "instgramproject/backend"
    "instgramproject/handler"
    "instgramproject/util"
)

func main() {
    fmt.Println("started-service")

    config, err := util.LoadApplicationConfig("conf", "deploy.yml")
    if err != nil {
        panic(err)
    }

    backend.InitElasticsearchBackend(config.ElasticsearchConfig)
    backend.InitGCSBackend(config.GCSConfig)

    log.Fatal(http.ListenAndServe(":8080", handler.InitRouter(config.TokenConfig)))
}