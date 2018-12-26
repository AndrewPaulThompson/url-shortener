package router

import (
    "log"
    "net/http"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/gorilla/mux"
)

type App struct {
    Router *mux.Router
    Client *dynamodb.DynamoDB
}

type Url struct {
    ID string       `json:"id"`
    LongUrl string  `json:"longUrl"`
    ShortUrl string `json:"shortUrl"`
}

func (a *App) Initialize() {
    a.Router = mux.NewRouter()
    a.Client = a.initializeDatabase()
    a.initializeRoutes()
}

func (a *App) initializeDatabase() *dynamodb.DynamoDB {
    sess, _ := session.NewSession(&aws.Config{
        Region: aws.String("us-west-1"),
        Endpoint: aws.String("http://localhost:8001")},
    )

    return dynamodb.New(sess)
}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/create", a.createUrlEndpoint).Methods("POST")
    a.Router.HandleFunc("/{id}", a.redirectEndpoint).Methods("GET")
}

func (a *App) Run(addr string) {
    log.Fatal(http.ListenAndServe(addr, a.Router))
}
