/*
 * The search engine adapter for Antares
 * 
 *   https://golang.org/pkg/net/http/
 *
 */

package main

import (
  "fmt"
  "net/http"
  "log"
  "github.com/schul-cloud/arix-search-adapter/arix"
  "bytes"
  "strconv"
  "net/url"
  "encoding/json"
  "io"
  "strings"
  "os"
  "path"
)


const PORT = 8080
const SEARCH_BASE = "/v1/search"
const URL_BASE = "/v1/url/"
const SERVER = "http://arix.datenbank-bildungsmedien.net/"
const CONTEXT = "HH"
const SEARCH_LIMIT = 10
const SERVER_ID = "ARIX"

const CODE_ENDPOINT = "/code"
const CODE_DIR = "github.com/schul-cloud/arix-search-adapter/"

/*
 * See the example https://github.com/schul-cloud/resources-api-v1/blob/master/schemas/search-response/examples/valid/five-fictive-resources-null-links.json
 */
type SuccessfulSearchResponse struct {
  Jsonapi    Jsonapi         `json:"jsonapi"`
  Links      Links           `json:"links"`
  Data       []ResourceData  `json:"data"`
}

type ErrorSearchResponse struct {
  Jsonapi    Jsonapi         `json:"jsonapi"`
  Errors     []HTTPError     `json:"errors"`
}

type Jsonapi struct {
  Version    string      `json:"version"`
  Meta       JsonapiMeta `json:"meta"`
}

type JsonapiMeta struct {
  Name         string    `json:"name"`
  Source       string    `json:"source"`
  Description  string    `json:"description"`
}

type Links struct {
  Self         SelfLink   `json:"self"`
  First        *NullLink  `json:"first"`
  Last        *NullLink   `json:"last"`
  Prev        *NullLink   `json:"prev"`
  Next        *NullLink   `json:"next"`
}

type SelfLink struct {
  Href         string        `json:"href"`
  Meta         SelfLinkMeta  `json:"meta"`
}

type SelfLinkMeta struct {
  Count        int  `json:"count"`
  Offset       int  `json:"offset"`
  Limit        int  `json:"limit"`
}

type NullLink struct {
}

type ResourceData struct {
  Type         string                 `json:"type"`
  Id           string                 `json:"id"`
  Attributes   arix.LearningResource  `json:"attributes"`
}

type ResourceLinks struct {
  Self         string         `json:"string"`  /* TODO: allow requesting single resources */
}

type HTTPError struct {
  Status       string  `json:"status"`
  Title        string  `json:"title"`
  Detail       string  `json:"detail"`
}

func NewWrongArgumentsResponse(host string) ErrorSearchResponse {
  return ErrorSearchResponse{
    Jsonapi: NewJsonapi(host),
    Errors: []HTTPError{
      HTTPError{
        Title: "Bad Request",
        Status: "400",
        Detail: "Only the query parameter Q is supported.",
      },
    },
  }
}

func NewServerErrorResponse(host string, message string) ErrorSearchResponse {
  return ErrorSearchResponse{
    Jsonapi: NewJsonapi(host),
    Errors: []HTTPError{
      HTTPError{
        Title: "Internal Server Error",
        Status: "500",
        Detail: message,
      },
    },
  }
}



func NewInacceptableContentTypeResponse(host string, accepted_content_type string) ErrorSearchResponse {
  return ErrorSearchResponse{
    Jsonapi: NewJsonapi(host),
    Errors: []HTTPError{
      HTTPError{
        Title: "Not Acceptable",
        Status: "406",
        Detail: fmt.Sprintf("The content type \"application/vnd.api+json\" must be accepted. It is not listed in \"%s\".", accepted_content_type),
      },
    },
  }
}

func NewJsonapi(host string) Jsonapi {
  return Jsonapi{
      Version: "1.0",
      Meta: JsonapiMeta{
        Name: "arix-search-adapter",
        Source: fmt.Sprintf("http://%s%s", host, CODE_ENDPOINT),
        Description: fmt.Sprintf(
          "This is a search adapter for Antares connected to %s.",
          SERVER),
      },
   }
}

func RespondWithError(response ErrorSearchResponse, w http.ResponseWriter, r *http.Request) {
  status, _ := strconv.Atoi(response.Errors[0].Status)
  w.WriteHeader(status)
  result, _ := json.MarshalIndent(response, "", "  ")
  io.WriteString(w, string(result))
  io.WriteString(w, "\r\n")
  fmt.Printf("%d %s: %s?%s -> %s\r\n",
             status,
             response.Errors[0].Title,
             r.URL.Path,
             r.URL.RawQuery,
             response.Errors[0].Detail)
}

func NewSuccessfulSearchResponse(host string, request_url string, limit int, offset int, resources []arix.LearningResource) SuccessfulSearchResponse {
  self_url := strings.Join([]string{
    "http://",
    host,
    request_url,
    }, "") 
  var data = []ResourceData{}
  for _, resource := range resources {
    data_url := strings.Join([]string{
      "http://",
      host,
      URL_BASE,
      resource.Id,
      }, "")
    resource.Url = data_url
    data = append(data, ResourceData{
      Type: "resource",
      Id: fmt.Sprintf("%s-%s", SERVER_ID, resource.Id),
      Attributes: resource,
    })
  }
  return SuccessfulSearchResponse{
    Jsonapi: NewJsonapi(host),
    Links: Links{
      Self: SelfLink{
        Href: self_url,
        Meta: SelfLinkMeta{
          Count: len(resources),
          Offset: offset,
          Limit: SEARCH_LIMIT,
        },
      },
    },
    Data: data,
  }
}


/*
 * Return whether the content type accepted by the client is
 * application/vnd.api+json
 *
 * TODO: write this shorter with split and contains
 */
func RequestIsAcceptable(accepted string) bool {
  accepted_list := strings.Split(accepted, ",")
  for _, content_type := range accepted_list {
    if (content_type == "*/*" || content_type == "application/*" ||
        content_type == "application/vnd.api+json") {
      return true
    }
  }
  return false
}


func main() {
  fmt.Printf("Server is starting on port http://localhost:%d%s\n", PORT, SEARCH_BASE)
  http.HandleFunc(SEARCH_BASE, func(w http.ResponseWriter, r *http.Request) {
    /* parse the query */
    w.Header().Set("Content-Type", "application/vnd.api+json") // from https://gist.github.com/tristanwietsma/8444cf3cb5a1ac496203#file-routes-go-L26
    query := r.FormValue("Q")  /* https://godoc.org/net/http#Request.FormValue */
    accpepted_content_types := r.Header.Get("Accept")
    if (query == "" || strings.Count(r.URL.RawQuery, "=") != 1) {
      /* The request is invalid. */
      RespondWithError(NewWrongArgumentsResponse(r.Host), w, r)
    } else if (!RequestIsAcceptable(accpepted_content_types)) {
      /* The request can not be fulfilled with this encoding. */
      RespondWithError(NewInacceptableContentTypeResponse(r.Host, accpepted_content_types), w, r)
    } else {
      /* request content from anatares 
       *
       *  https://stackoverflow.com/a/19253970/1320237
       */
      //status_code := 999 /*
      data := url.Values{}
      data.Set("context", CONTEXT)
      data.Set("xmlstatement", arix.GetSearchRequest(SEARCH_LIMIT, query))
      encoded_data := data.Encode()

      client := &http.Client{}
      arix_search_request, _ := http.NewRequest("POST", SERVER, bytes.NewBufferString(encoded_data))
      arix_search_request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
      arix_search_request.Header.Add("Content-Length", strconv.Itoa(len(encoded_data)))
      
      arix_response, error := client.Do(arix_search_request)
      if (error != nil) {
        RespondWithError(NewServerErrorResponse(r.Host, fmt.Sprintf("%s", error)), w, r)
      } else {
        body := arix_response.Body
        found_resources := arix.ParseSearchResult(body)
        status_code := arix_response.StatusCode
          /**/
        
        /*
         * Create the converted search result.
         */

        search_response := NewSuccessfulSearchResponse(
          r.Host, r.URL.String(), SEARCH_LIMIT, 0, found_resources)

        result, _ := json.MarshalIndent(search_response, "", "  ")
        io.WriteString(w, string(result))
        io.WriteString(w, "\r\n")
        fmt.Printf("Searching %s?%s -> Arix (%d)\r\n",
                   r.URL.Path,
                   r.URL.RawQuery,
                   status_code)
      }
    }
  })
  
  /*
   * Serve the code at CODE_ENDPOINT.
   * see 
   * - https://stackoverflow.com/a/26563418/1320237
   * - https://gobyexample.com/environment-variables
   */
  gopath := os.Getenv("GOPATH")
  if (gopath != "") {
    code := path.Join(gopath, "src", CODE_DIR)
    http.Handle("/code/", http.StripPrefix("/code/", http.FileServer(http.Dir(code))))
    fmt.Printf("Serving code from %s at /code/\n", code)
  }

  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
