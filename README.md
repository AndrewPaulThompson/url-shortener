# Url Shortener
A Url Shortener API built in Go, using DynamoDB for storage.
### Prerequisites
This uses go modules for managing dependencies, making [Go 1.11](https://github.com/golang/go/wiki/Modules) a requirment.
### Getting Started
First copy the environment template `.env.dist` into `.env`:
```
cp .env.dist .env
```
Then replace the values with your AWS credentials
### Setting up DynamoDB
If you want to run DynamoDB locally:
[Download](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html)\
Change directory into the unzipped folder and run:
```
java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb
```
If the default port is already in use:
```
java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb -port 8001
```
The `database/helper/main.go` file sets up the DynamoDB table with the primary key.
### Built With
* [Go](https://golang.org/) - Go Programming Language
* [DynamoDB](https://aws.amazon.com/dynamodb/) - No SQL Database
