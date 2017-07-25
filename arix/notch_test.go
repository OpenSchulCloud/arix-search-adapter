package arix

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "strings"
)



/* Test that the program can generate a link response from the different notches using a secret.
 *
 *
 */
func TestExtractLinks(t *testing.T) {
	for _, c := range []struct {
		response string
		links map[string]string
	} {
		{
	    "<link size='200 MByte'><a href='http://xplay.datenbank-bildungsmedien.net/928c9eaf7ce80072fbbcca7453092c6e/SF-52373-cnv_mp4_a/Die_Physik_Albert_Einsteins-Die_Kosmologie.mp4'>direct</a><a href='http://xplay.datenbank-bildungsmedien.net/928c9eaf7ce80072fbbcca7453092c6e/SF-52373-download/Die_Physik_Albert_Einsteins-Die_Kosmologie.mp4'>download</a></link>",
	    map[string]string{
	      "download": "http://xplay.datenbank-bildungsmedien.net/928c9eaf7ce80072fbbcca7453092c6e/SF-52373-download/Die_Physik_Albert_Einsteins-Die_Kosmologie.mp4",
	      "direct": "http://xplay.datenbank-bildungsmedien.net/928c9eaf7ce80072fbbcca7453092c6e/SF-52373-cnv_mp4_a/Die_Physik_Albert_Einsteins-Die_Kosmologie.mp4",
	    },
	  }, {
	    "  <link size='427 MByte'> <a href='http://xplay.datenbank-bildungsmedien.net/151d1d77f1126fad9b32fd8b6a218095/SF-56395-cnv_mp4_a/10_Dinge-die_du_wissen_musst-Albert_Einstein-HD.mp4'>direct</a><a href='http://xplay.datenbank-bildungsmedien.net/151d1d77f1126fad9b32fd8b6a218095/SF-56395-download/10_Dinge-die_du_wissen_musst-Albert_Einstein-HD.mp4'>download</a></link>",
	    map[string]string{
	      "download": "http://xplay.datenbank-bildungsmedien.net/151d1d77f1126fad9b32fd8b6a218095/SF-56395-download/10_Dinge-die_du_wissen_musst-Albert_Einstein-HD.mp4",
	      "direct": "http://xplay.datenbank-bildungsmedien.net/151d1d77f1126fad9b32fd8b6a218095/SF-56395-cnv_mp4_a/10_Dinge-die_du_wissen_musst-Albert_Einstein-HD.mp4",
	    },
    },
	} {
		got := GetLinksFromLinkResponse(strings.NewReader(c.response))
		assert.Equal(t, c.links, got)
	}
}
