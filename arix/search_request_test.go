package arix

import (
  "testing"
  "github.com/stretchr/testify/assert"
)


/* Test that the program can generate search request from the given parameters.
 *
 *
 */
func TestGenerateLinkRequest(t *testing.T) {
  for _, c := range []struct {
    limit int
    words, want string
  }{
    {
      10,
      `einstein`,
      `<search limit='10' fields='text,nr,titel,typ,laenge,url,quelle'><condition field='titel_fields'>einstein</condition></search>`,
    }, {
      0,
      `<>"&'`,  // http://stackoverflow.com/a/17448222/1320237
      `<search limit='0' fields='text,nr,titel,typ,laenge,url,quelle'><condition field='titel_fields'>&#60;&#62;&#34;&#38;&#39;</condition></search>`,
    }, {
      33,
      `einstein und ein stein`,
      `<search limit='33' fields='text,nr,titel,typ,laenge,url,quelle'><condition field='titel_fields'>einstein und ein stein</condition></search>`,
    },
  } {
    got := GetSearchRequest(c.limit, c.words)
    assert.Equal(t, got, c.want)
  }
}

