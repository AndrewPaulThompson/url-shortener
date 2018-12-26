package main

import (
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-1"),
        Endpoint: aws.String("http://localhost:8001")},
    )

    // Create DynamoDB client
    svc := dynamodb.New(sess)

    // Create table Movies
    input := &dynamodb.CreateTableInput{
        AttributeDefinitions: []*dynamodb.AttributeDefinition{
            {
                AttributeName: aws.String("longUrl"),
                AttributeType: aws.String("S"),
            },
        },
        KeySchema: []*dynamodb.KeySchemaElement{
            {
                AttributeName: aws.String("longUrl"),
                KeyType:       aws.String("HASH"),
            },
        },
        ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
            ReadCapacityUnits:  aws.Int64(10),
            WriteCapacityUnits: aws.Int64(10),
        },
        TableName: aws.String("Url-Mappings"),
    }

    _, err = svc.CreateTable(input)

    if err != nil {
        fmt.Println("Got error calling CreateTable:")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Println("Created the table Url-Mappings in us-west-1")
}
