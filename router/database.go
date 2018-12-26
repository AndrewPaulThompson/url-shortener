package router

import(
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

// Puts an Item into DynamoDb
func (a *App) addToDatabase(url Url) {
    av, err := dynamodbattribute.MarshalMap(url)

    if err != nil {
        fmt.Println(err.Error())
    }

    input := &dynamodb.PutItemInput{
        Item: av,
        TableName: aws.String("Url-Mappings"),
    }

    _, err = a.Client.PutItem(input)
}

// Gets Item from DynamoDB using it's longUrl
func (a *App) lookupLongUrl(url Url) *dynamodb.GetItemOutput {
    result, err := a.Client.GetItem(&dynamodb.GetItemInput{
        TableName: aws.String("Url-Mappings"),
        Key: map[string]*dynamodb.AttributeValue{
            "longUrl": {
                S: aws.String(url.LongUrl),
            },
        },
    })

    if err != nil {
        fmt.Println(err.Error())
    }

    return result
}

// Scans DynamoDB for existing routes using its ID
func (a *App) lookupId(id string) *dynamodb.ScanOutput {
    filt := expression.Name("id").Equal(expression.Value(id))
    proj := expression.NamesList(expression.Name("id"), expression.Name("longUrl"), expression.Name("shortUrl"))
    expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

    if err != nil {
        fmt.Println("Got error building expression:")
        fmt.Println(err.Error())
    }

    params := &dynamodb.ScanInput{
        ExpressionAttributeNames:  expr.Names(),
        ExpressionAttributeValues: expr.Values(),
        FilterExpression:          expr.Filter(),
        ProjectionExpression:      expr.Projection(),
        TableName:                 aws.String("Url-Mappings"),
    }

    result, err := a.Client.Scan(params)

    if err != nil {
        fmt.Println("Query API call failed:")
        fmt.Println((err.Error()))
    }

    return result
}
