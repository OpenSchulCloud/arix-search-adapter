/* Generate objects from parameters
 *
 * The Resource definition:
 * https://github.com/schul-cloud/resources-api-v1/tree/master/schemas/resource#readme
 * 
 * Hints:
 * - looking for the mime type
 *   https://golang.org/pkg/mime/#TypeByExtension
 */

package arix

type License struct {
}

type LearningResource struct {
  Id               string    `json:"id"`
  Title            string    `json:"title"`
  Description      string    `json:"description"`
  Url              string    `json:"url"`
  Licenses         []License `json:"licenses"`
  MimeType         string    `json:"mimeType"`
  ContentCategory  string    `json:"contentCategory"`
  Languages        []string  `json:"languages"`
  Thumbnail        string    `json:"thumbnail"`
}

