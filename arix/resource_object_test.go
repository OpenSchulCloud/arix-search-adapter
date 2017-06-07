
package arix


import (
  "testing"
  "encoding/json"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
  "os"
)





func TestJSONToResourceSerialization(t *testing.T) {
	for _, resource := range []LearningResource {
		{
      Title:            "Albert Einstein  (HD)",
      Description:      "Auf seinem bekanntesten Foto streckt Albert Einstein dem Fotografen die Zunge raus und ist damit zur Ikone der Pop-Kultur geworden, später kopiert von den Rolling-Stones und Kiss-Bassist Gene Simmons. Aber was für ein Mensch war Einstein? Und warum ist er so berühmt geworden? ä10 Dinge, die du wissen musst! Albert Einstein unternimmt eine Reise durch Zeit und Raum und klärt unter anderem folgende Fragen:~br&gt;~br&gt;Hat Einstein sieben Jahre lang als Schuster gearbeitet?~br&gt;~br&gt;Als er im schweizerischen Bern lebte, sprach Albert Einstein von seiner Tagesbeschäftigung als äSchusterarbeit. Aber hat der später weltberühmte Physiker tatsächlich sieben Jahre lang Schuhe zusammengenagelt? Die Geschichte klingt plausibel durch die oft zitierte Tatsache, dass Einstein auf dem Gymnasium eine sechs in Mathe und Physik hat. Erst im zweiten Anlauf wird er an der Universität aufgenommen, und noch drei Jahre vor der Veröffentlichung seiner Relativitätstheorie sucht er mit Gratis-Probestunden verzweifelt Nachhilfeschüler, um seine Miete zahlen zu können. Aber war Einstein wirklich ein schlechter Schüler oder einfach seiner Zeit voraus? ~br&gt;~br&gt;Einstein suchte vergeblich die Weltformel!~br&gt;~br&gt;So unglaublich es klingt, aber Albert Einstein hat für seine Relativitätstheorie nicht den Nobelpreis bekommen. 1905 sorgte sie für Aufsehen in der Fachwelt, aber das Nobelpreiskomitee war sich 1921 immer noch zu unsicher und gab Einstein die Auszeichnung lieber für eine andere wissenschaftliche Arbeit. Dabei hatten die Juroren die Qual der Wahl. Schließlich entschieden sie sich für eine Arbeit Einsteins über Licht-Teilchen. Mit ihr hatte Einstein die zweite große Physik-Revolution zu Beginn des 20. Jahrhunderts angestoßen: die Entwicklung der Quantentheorie. Sie versucht die Vorgänge auf der Ebene von Atomen, Elektronen und noch kleineren Teilchen zu verstehen. Sie ist die Grundlage für unsere Digitalkameras, Laser oder Computer. Aber sie widerspricht der Relativitätstheorie. Einstein arbeitet   e sich die letzten 30 Jahre seines Lebens daran ab, beide Theorien zu einer alles erklärenden äWeltformel zu vereinen. Er scheiterte daran und bis heute suchen die Physiker vergeblich nach ihr.~br&gt;~br&gt;~br&gt;Hat Einstein die Atombombe gebaut?~br&gt;~br&gt;Mit der berühmtesten Formel der Welt äe=mc2 legte Einstein eine der Grundlagen für die Gewinnung von Energie aus der Atomspaltung. Dabei verschwindet Masse und wird in gewaltig viel Energie umgewandelt. Gewaltig viel Energie, denn die Formel lautet ausgesprochen: äEnergie gleich Masse mal Lichtgeschwindigkeit. Wenn man weiß, dass Licht mit unglaublichen 1 Milliarde Kilometer pro Stunde unterwegs ist, wird einem klar, welche Mengen an Energie bei der Umwandlung von Materie, wie etwa in einer Atombombenexplosion, freigesetzt werden. Aber stimmt es, dass ohne Einstein die Atombombe nicht möglich gewesen wäre? Und was genau schrieb Einstein an Roosevelt, den amerikanischen Präsidenten, der Hitler-Deutschland den Krieg erklärte? Rief der überzeugte Pazifist Einstein wirklich zum Bau der Atombombe auf?",
      Url:              "http://xplay.datenbank-bildungsmedien.net/151d1d77f1126fad9b32fd8b6a218095/SF-56395-download/10_Dinge-die_du_wissen_musst-Albert_Einstein-HD.mp4",
      Licenses:         []License{},
      MimeType:         "video/mp4",
      ContentCategory:  "a",
      Languages:        []string{"de-de"},
    }, {
      Title:            "Schul-Cloud",
      Description:      "Eine Umgebung für alle Schulen.",
      Url:              "https://schul-cloud.org",
      Licenses:         []License{},
      MimeType:         "text/html",
      ContentCategory:  "l",
      Languages:        []string{"de-de", "en-en"},
    },
	} {
    // https://golang.org/pkg/encoding/json/#Unmarshal
		encoded_resource, encode_error := json.Marshal(resource)
    require.Nil(t, encode_error, "Encoding works without error.")
    var decoded_resource LearningResource
    decode_error := json.Unmarshal(encoded_resource, &decoded_resource)
    require.Nil(t, decode_error, "Decoding works without error.")
		assert.Equal(t, decoded_resource, resource)
	}
}

func GetResourceFromFile(file_name string) LearningResource {
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
  var resource LearningResource
  decoder := json.NewDecoder(fi)
  error := decoder.Decode(&resource)
  if (error != nil) {
    panic(error)
  }
  return resource
}

func TestParsingResourceFromJSON(t *testing.T) {
  resource := GetResourceFromFile("resource_object_test_example_1.json")
  require.NotNil(t, resource)
  assert.Equal(t, resource.Title, "Example Website")
  assert.Equal(t, resource.Url, "https://example.org")
  assert.Equal(t, resource.Licenses, []License{})
  assert.Equal(t, resource.ContentCategory, "l")
  assert.Equal(t, resource.Languages, []string{"en-en"})
  assert.Equal(t, resource.Thumbnail, "http://cache.schul-cloud.org/thumbs/k32164876328764872384.jpg")
}




