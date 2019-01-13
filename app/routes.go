package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gorilla/mux"
	hashids "github.com/speps/go-hashids"
)

// createUrlEndpoint is the Handler Function for short url creation
// Takes a long url as postdata (eg. {"longUrl":"https://www.google.com"})
// Checks if the url has already been shortened, if not creates it.
// Returns json object of shortened url
// (eg. {"id":"nYmVZXE","longUrl":"https://www.google.com","shortUrl":"http://localhost:8080/nYmVZXE"})
// Or invalid response
// {"error":"Invalid request payload"}
func (a *App) createUrlEndpoint(w http.ResponseWriter, req *http.Request) {
	var url Url
	var err error

	if req.FormValue("longUrl") != "" {
		url.LongUrl = req.FormValue("longUrl")
	} else {
		err = json.NewDecoder(req.Body).Decode(&url)
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	result := a.lookupLongUrl(url)
	if len(result.Item) == 0 {
		hd := hashids.NewData()
		h, _ := hashids.NewWithData(hd)
		now := time.Now()
		url.ID, _ = h.Encode([]int{int(now.Unix())})

		scheme := "http://"
		if req.TLS != nil {
			scheme = "https://"
		}

		url.ShortUrl = scheme + req.Host + "/" + url.ID

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

// Handler Function for redirecting from short -> long urls
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

func (a *App) rootEndpoint(w http.ResponseWriter, req *http.Request) {
	var S struct{}
	tmpl := template.Must(template.ParseFiles("views/index.tmpl"))
	tmpl.Execute(w, S)
}
