package app

import (
    "log"
    "net/http"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/joho/godotenv"
    "github.com/gorilla/mux"
)

// Application specific objects
// Router - Mux router, used for routing
// Client - DynamoDB Client, used for connecting to DynamoDB instance
type App struct {
    Router      *mux.Router
    Client      *dynamodb.DynamoDB
    AwsConfig   AwsConfig
}

// Url structure inside DynamoDB
// ID       - ID of the shortened url
// LongUrl  - The original (long) url
// ShortUrl - The hosturl + ID
type Url struct {
    ID          string  `json:"id"`
    LongUrl     string  `json:"longUrl"`
    ShortUrl    string  `json:"shortUrl"`
}

// AWS config needed for DynamoDB
// Access Key & Secret Key are stored in environment variables
// region       - Region of DynamoDB
// tableName    - Table Name for entries
type AwsConfig struct {
    region          string
    tableName       string
}

// Initialise the App object
func (a *App) Initialise() {
    a.initialiseConfig()
    a.Router = mux.NewRouter()
    a.Client = a.initialiseDatabase()
    a.initialiseRoutes()
}

// Initialise the DynamoDB Client
// Returns DynamoDB client object
func (a *App) initialiseDatabase() *dynamodb.DynamoDB {
    sess, _ := session.NewSession(&aws.Config{
        Region: aws.String(a.AwsConfig.region),
        Endpoint: aws.String("http://localhost:8001")},
    )

    return dynamodb.New(sess)
}

// Initialise Application Routes
func (a *App) initialiseRoutes() {
    a.Router.HandleFunc("/create", a.createUrlEndpoint).Methods("POST")
    a.Router.HandleFunc("/{id}", a.redirectEndpoint).Methods("GET")
}

// Runs the sever and listens for connections
func (a *App) Run(addr string) {
    log.Fatal(http.ListenAndServe(addr, a.Router))
}

// Initialise AWS config
func (a *App) initialiseConfig() {
    // Load .env (for AWS creds)
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    a.AwsConfig = AwsConfig{}
    a.AwsConfig.tableName = os.Getenv("AWS_DYNAMODB_TABLE_NAME")
    a.AwsConfig.region = os.Getenv("AWS_REGION")
}
