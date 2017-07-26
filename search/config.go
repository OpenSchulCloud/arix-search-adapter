
package search

import (
  "os"
  "strconv"
  "fmt"
)


type Configuration struct {
  // The port the server socket binds to
  Port          int
  
  // The ARIX server this adapter connects to
  Server        string
  
  // The endpoints used for requests
  Endpoints     Endpoints
  
  // The Search context for the ARIX server
  Context       string
  
  // The limit of how many resources to search
  Limit         int
  
  // The id of this server
  ServerId      string
  
  // The directory where the code of this server is in, relative to the GOPATH
  CodeDirectory string
  
  // The secret to use to proof the ownership of the license
  Secret        string
  
  // The type of the link which should be given to the user.
  LinkType      string
}

type Endpoints struct {
  // The search endpoint to search for objects
  Search string
  
  // The url endpoint to resolve urls to resources
  Url    string
  
  // The code endpoint to get the source code of this application from
  Code   string
}

func GetEnv(key string, fallback string) string {
  // inspired by https://stackoverflow.com/a/40326580/1320237
  value, success := os.LookupEnv(key)
  if (success) {
    return value
  } else {
    return fallback
  }
}

func GetEnvInt(key string, fallback int) int {
  value, success := os.LookupEnv(key)
  if (success) {
    number, error := strconv.Atoi(value)
    if (error != nil) {
      fmt.Printf("Invalid value for \"%s\", integer expected, not \"%s\"", key, value)
      return fallback
    }
    return number
  } else {
    return fallback
  }
}

func Config() Configuration {
  return Configuration{
      Port:          GetEnvInt("ARIX_SEARCH_PORT", 8080),
      Server:        GetEnv("ARIX_SEARCH_SERVER", "http://arix.datenbank-bildungsmedien.net/"),
      Context:       GetEnv("ARIX_SEARCH_CONTEXT", "HH"),
      Secret:        GetEnv("ARIX_SEARCH_SECRET", ""),
      Limit:         GetEnvInt("ARIX_SEARCH_LIMIT", 10),
      ServerId:      GetEnv("ARIX_SEARCH_SERVER_ID", "ARIX"),
      LinkType:      GetEnv("ARIX_SEARCH_LINK_TYPE", "direct"),
      CodeDirectory: "github.com/schul-cloud/arix-search-adapter/",
      Endpoints: Endpoints {
        Search: "/v1/search",
        Url:    "/v1/url/",
        Code:   "/code",
      },
    }
}
