package arix

import "testing"

func TestGenerateNodge(t *testing.T) {
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
		if got != c.want {
			t.Errorf("NotchToLinkRequest(%q, %q) == %q, want %q", c.notch, c.secret, got, c.want)
		}
	}
}
