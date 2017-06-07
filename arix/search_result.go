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

type ArixResource struct {
  Id     string   `xml:"identifier,attr"`
  Fields []Field  `xml:"f"`
  Title  string   `xml:"-"`
}


type SearchResult struct {
  Resources []ArixResource `xml:"r"`
}


func ParseSearchResult(source io.Reader) []LearningResource {
  decoder := xml.NewDecoder(source)
  var resources = []LearningResource{}
  search_result := SearchResult{}
  err := decoder.Decode(&search_result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return []LearningResource{}
	}
  for i, _ := range search_result.Resources {
    res1 := &search_result.Resources[i] // http://stackoverflow.com/questions/20185511/golang-range-references-instead-values#comment47156406_29498133
    var res2 LearningResource
    res2.Id = res1.Id
    for _, field := range res1.Fields {
      switch field.Name {
        case "titel":
          res2.Title = field.Value
      }
    }
    resources = append(resources, res2)
  }
	return resources
}