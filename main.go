package main

import (
  "bytes"
  "html/template"
  "log"
	"net/http"
  "net/url"
	"os"
  "strconv"
  "time"
  "math"
  "fmt"

  "github.com/freshman-tech/news-demo-starter-files/news"
  "github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

type Search struct {
	Query      string
  AccessToken string
  ShopifyStoreName string
  AlertLevel string
  EmailAddress string
  PhoneNumber string
  VariantID string
	NextPage   int
	TotalPages int
	Results    *news.Results
}

// Show screen where you can input your Shopify parameters
func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

//This function will make the template show info on the inventory
func searchHandler(newsapi *news.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
  	u, err := url.Parse(r.URL.String())
  	if err != nil {
  		http.Error(w, err.Error(), http.StatusInternalServerError)
  		return
  	}

  	params := u.Query()
  	searchQuery := params.Get("q")
    AccessToken := params.Get("AccessToken")
    ShopifyStoreName := params.Get("ShopifyStoreName")
    AlertLevel := params.Get("AlertLevel")
    EmailAddress := params.Get("EmailAddress")
    PhoneNumber := params.Get("PhoneNumber")
    VariantID := params.Get("VariantID")
  	page := params.Get("page")
  	if page == "" {
  		page = "1"
  	}

    results, err := newsapi.FetchEverything(searchQuery, AccessToken, ShopifyStoreName, VariantID, AlertLevel, EmailAddress, PhoneNumber, page)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    nextPage, err := strconv.Atoi(page)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    search := &Search{
      Query:      searchQuery,
      AccessToken: AccessToken,
      ShopifyStoreName: ShopifyStoreName,
      AlertLevel: AlertLevel,
      EmailAddress: EmailAddress,
      PhoneNumber: PhoneNumber,
      VariantID: VariantID,
      NextPage:   nextPage,
      TotalPages: int(math.Ceil(float64(results.TotalResults) / float64(newsapi.PageSize))),
      Results:    results,
    }

    buf := &bytes.Buffer{}
		err = tpl.Execute(buf, search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(w)

  }
}

func main() {
  fmt.Println("Started")


  err := godotenv.Load()
  if err != nil {
    log.Println("Error loading .env file")
  }

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}


  apiKey := "N/A"

	myClient := &http.Client{Timeout: 10 * time.Second}
	newsapi := news.NewClient(myClient, apiKey, 20)

  fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()


  mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
  mux.HandleFunc("/monitor", searchHandler(newsapi))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
