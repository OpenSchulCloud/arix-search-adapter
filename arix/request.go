package arix

import (
  "strconv"
  "net/url"
  "net/http"
  "bytes"
)


/* Request a response from the ARIX server.
 * 
 * Returns response and error.
 * https://stackoverflow.com/a/19253970/1320237
 */
func Request(server string, context string, request string) (*http.Response, error) {
  data := url.Values{}
  data.Set("context", context)
  data.Set("xmlstatement", request)
  encoded_data := data.Encode()

  client := &http.Client{}
  arix_search_request, _ := http.NewRequest("POST", server, bytes.NewBufferString(encoded_data))
  arix_search_request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  arix_search_request.Header.Add("Content-Length", strconv.Itoa(len(encoded_data)))
  
  return client.Do(arix_search_request)
}
