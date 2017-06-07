package arix

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
  "os"
)


/* Test that the program can generate a link response from the different notches using a secret.
 *
 *
 */
func TestGenerateNotchLinkResponse(t *testing.T) {
	for _, c := range []struct {
		notch, secret, want string
	}{
		{
      `<notch id="f3daeeeb8b2f3bb370e0b75405c1dd60">229160641805b905d3b2602651e6b840</notch>`,
      `<Secret>`,
      `<link id="f3daeeeb8b2f3bb370e0b75405c1dd60">e4679cdf57cd38c93f31eb8119c9adab</link>`,
    }, {
      "<notch id=\"f3daeeeb8b2f3bb370e0b75405aaaaaa\">229160641805b905d3b2602aaaaaaaaa</notch>",
      "<Secret2>",
      "<link id=\"f3daeeeb8b2f3bb370e0b75405aaaaaa\">9629e6080f2d0b8ec3c56b3010ecf469</link>",
    }, {
      "<notch id=\"1111111b8b2f3bb370e0b75405aaaaaa\">11111111111111111111111aaaaaaaaa</notch>",
      "SECRET",
      "<link id=\"1111111b8b2f3bb370e0b75405aaaaaa\">5e032c75e20c7d8acd5460cf8023d6e3</link>",
    }, 
	} {
		got := NotchToLinkRequest([]byte(c.notch), c.secret).String()
		assert.Equal(t, got, c.want)
	}
}

func GetResultFromFile(file_name string) []LearningResource {
  fi, err := os.Open(file_name)
  if err != nil {
      panic(err)
  }
  // close fi on exit and check for its returned error
  defer func() {
      if err := fi.Close(); err != nil {
          panic(err)
      }
  }()
  return ParseSearchResult(fi)
}


func TestParseResults(t *testing.T) {
  assert := assert.New(t)
  require := require.New(t)
  resources := GetResultFromFile("parse_test_result.txt")
  require.Equal(10, len(resources))
  assert.Equal("SF-56395", resources[0].Id)
  assert.Equal("SF-52309", resources[3].Id)
  assert.Equal("SF-52146", resources[4].Id)
  assert.Equal("SF-52373", resources[9].Id)
  
  assert.Equal("Albert Einstein  (HD)", resources[0].Title)
}
