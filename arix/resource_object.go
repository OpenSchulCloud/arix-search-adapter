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
  Id               string
  Title            string
  Description      string
  Url              string
  Licenses         []License
  MimeType         string
  ContentCategory  string
  Languages        []string
  Thumbnail        string
}

