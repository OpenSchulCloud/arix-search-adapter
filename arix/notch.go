/* The arix package contains tools to parse arix requests and responses.
 *
 *
 */

package arix

import (
  "fmt"
  "encoding/xml"
  "crypto/md5"
  "encoding/hex"
  "bytes"
  "io"
)


func GetNotchRequest(resource_id string) string {
  return fmt.Sprintf("<notch identifier=\"%s\" />", resource_id)
}

/* A parsed XML notch
 * Example:
 * - "<notch id=\"1111111b8b2f3bb370e0b75405aaaaaa\">11111111111111111111111aaaaaaaaa</notch>"
 */
type Notch struct {
  Id          []byte   `xml:"id,attr"`
  Challenge   []byte   `xml:",chardata"`
}

type LinkRequest struct {
  notch Notch
  secret []byte
}

/* From a notch response and a secret, generate the link request.
 * 
 * Example:
 *   notch_response = "<notch id=\"1111111b8b2f3bb370e0b75405aaaaaa\">11111111111111111111111aaaaaaaaa</notch>"
 *   secret = "SECRET"
 *   result = "<link id=\"1111111b8b2f3bb370e0b75405aaaaaa\">5e032c75e20c7d8acd5460cf8023d6e3</link>"
 * 
 * See also
 * - https://golang.org/pkg/encoding/xml/
 * - https://siongui.github.io/2015/02/17/go-parse-xml-example-1/
 * 
 */
func NotchReaderToLinkRequest(notch_response io.Reader, secret string) LinkRequest {
  notch := Notch{}
  decoder := xml.NewDecoder(notch_response)
  err := decoder.Decode(&notch)
	if err != nil {
		fmt.Printf("error: %v", err)
		return LinkRequest{notch:notch, secret:[]byte(secret)}
	}
  //fmt.Printf("notch: %#v <- %s\n", notch, notch_response)
	return LinkRequest{notch:notch, secret:[]byte(secret)}
}
func NotchToLinkRequest(notch_response []byte, secret string) LinkRequest {
  return NotchReaderToLinkRequest(bytes.NewReader(notch_response), secret)
}

func (link_request LinkRequest) String() string {
  h := md5.New()
  h.Write(link_request.notch.Challenge)
  h.Write([]byte{':'})
  h.Write(link_request.secret)
  var buffer bytes.Buffer // from http://stackoverflow.com/a/1766304/1320237
  buffer.WriteString("<link id=\"")
  buffer.Write(link_request.notch.Id)
  buffer.WriteString("\">")
  buffer.WriteString(hex.EncodeToString(h.Sum(nil)))  
  buffer.WriteString("</link>")
  return buffer.String()
}


type ResponseLink struct {
  Href string `xml:"href,attr"`
  Name string `xml:",innerxml"`
}

type ResponseLinkCollection struct {
  Error  string `xml:",innerxml"`
  Links  []ResponseLink  `xml:"a"`
}

func GetLinksFromLinkResponse(link_response io.Reader) map[string]string {
  links := ResponseLinkCollection{}
  decoder := xml.NewDecoder(link_response)
  err := decoder.Decode(&links)
  mapping := map[string]string{}
	if err != nil {
		fmt.Printf("error: %v", err)
		return mapping
	}
	for _, link := range links.Links {
	  mapping[link.Name] = link.Href
	}
	if (len(links.Links) == 0) {
	  mapping["error"] = links.Error
	}
  return mapping
}