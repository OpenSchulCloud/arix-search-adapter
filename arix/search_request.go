/* Generate a search request from parameters.
 *
 */

package arix
 
import (
  "fmt"
  "strings"
)


const search_template = "<search limit='%d' fields='text,nr,titel,typ,laenge,url,quelle'><condition field='titel_fields'>%s</condition></search>"

/*
 * <search limit='10' fields='text,nr,titel,typ,laenge,url,quelle'>
 * <condition field='titel_fields'>einstein</condition>
 * </search>
 */
func GetSearchRequest(limit int, words string) string {
  escaped_words := strings.Replace(words,         `&`, "&#38;", -1)
  escaped_words  = strings.Replace(escaped_words, `<`, "&#60;", -1)
  escaped_words  = strings.Replace(escaped_words, `>`, "&#62;", -1)
  escaped_words  = strings.Replace(escaped_words, `'`, "&#39;", -1)
  escaped_words  = strings.Replace(escaped_words, `"`, "&#34;", -1)
  return fmt.Sprintf(search_template, limit, escaped_words)
}
