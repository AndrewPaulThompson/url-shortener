package router

import (
    "encoding/json"
    "net/http"
    "fmt"
    "time"

    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/gorilla/mux"
    "github.com/speps/go-hashids"
)

func (a *App) createUrlEndpoint(w http.ResponseWriter, req *http.Request) {
    var url Url

    err := json.NewDecoder(req.Body).Decode(&url)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
    }

    result := a.lookupLongUrl(url)
    if len(result.Item) == 0 {
        hd          := hashids.NewData()
        h, _        := hashids.NewWithData(hd)
        now         := time.Now()
        url.ID, _    = h.Encode([]int{int(now.Unix())})
        url.ShortUrl = "http://localhost:8080/" + url.ID

        a.addToDatabase(url)
        w.WriteHeader(http.StatusCreated)
    } else {
        url = Url{}
        err = dynamodbattribute.UnmarshalMap(result.Item, &url)

        if err != nil {
            panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
        }
        w.WriteHeader(http.StatusFound)
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    json.NewEncoder(w).Encode(url)
}

func (a *App) redirectEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    result := a.lookupId(params["id"])

    if len(result.Items) == 0 {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "Not found"})
        return
    }

    url := Url{}
    err := dynamodbattribute.UnmarshalMap(result.Items[0], &url)

    if err != nil {
        fmt.Println("Got error unmarshalling:")
        fmt.Println(err.Error())
    }

    http.Redirect(w, req, url.LongUrl, 301)
}
