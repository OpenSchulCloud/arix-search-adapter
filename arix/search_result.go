/* Parse search results.
 *
 */
package arix

import (
  "io"
  "encoding/xml"
  "fmt"
)

type Field struct {
  Name  string `xml:"n,attr"`
  Value string `xml:",innerxml"`
}

type Resource struct {
  Id     string   `xml:"identifier,attr"`
  Fields []Field  `xml:"f"`
  Title  string   `xml:"-"`
}


type SearchResult struct {
  Resources []Resource `xml:"r"`
}


func ParseSearchResult(source io.Reader) SearchResult {
  decoder := xml.NewDecoder(source)
  search_result := SearchResult{}
  err := decoder.Decode(&search_result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return SearchResult{}
	}
  for i, _ := range search_result.Resources {
    resource := &search_result.Resources[i] // http://stackoverflow.com/questions/20185511/golang-range-references-instead-values#comment47156406_29498133
    for _, field := range resource.Fields {
      switch field.Name {
        case "titel":
          resource.Title = field.Value
      }
    }
    resource.Fields = nil // delete unused fields
  }
	return search_result
}