package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/StefanKjartansson/sundcloud/laterpay"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
)

var c laterpay.LaterPayClient
var catalog []Song

func init() {
	merchantID := os.Getenv("LP_ID")
	merchantSecret := os.Getenv("LP_SECRET")

	if merchantID == "" {
		log.Fatalln("LP_ID must be set")
	}

	if merchantSecret == "" {
		log.Fatalln("LP_SECRET must be set")
	}

	c = laterpay.LaterPayClient{
		Id:        merchantID,
		SecretKey: []byte(merchantSecret),
		AddURL:    "/foo",
		AccessURL: "/bar",
		WebRoot:   "/baz",
	}
	catalog = []Song{
		{uuid.NewV4().String(), "Adele", "Hello", "http://lorempixel.com/200/100/", "/mp3/adele.mp3", true},
		{uuid.NewV4().String(), "Foo", "World", "http://lorempixel.com/200/100/", "/mp3/adele.mp3", true},
		{uuid.NewV4().String(), "Rammstein", "Bobo", "http://lorempixel.com/200/100/", "/mp3/adele.mp3", true},
		{uuid.NewV4().String(), "Fungi", "XXXX", "http://lorempixel.com/200/100/", "/mp3/adele.mp3", true},
	}

}

type Song struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Image  string `json:"image"`
	Url    string `json:"url"`
	Access bool   `json:"access"`
}

const tpl = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>SundCloud</title>
    <meta name="robots" content="noindex, nofollow">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <script>
      var WebFontConfig = {
        google: { families: [ 'Open+Sans:400,400italic,600:latin,latin-ext' ] }
      };
      (function () {
        var wf = document.createElement('script');
        wf.src = ('https:' == document.location.protocol ? 'https' : 'http') +
          '://ajax.googleapis.com/ajax/libs/webfont/1/webfont.js';
        wf.type = 'text/javascript';
        wf.async = 'true';
        var s = document.getElementsByTagName('script')[0];
        s.parentNode.insertBefore(wf, s);
      })();
    </script>
  </head>
  <body>
    <div id="container" data-token="{{.Token}}"></div>
    <script src="/static/js/sundcloud.js" ></script>
  </body>
</html>
`

func getIdsFromCatalog(songs []Song) []string {
	out := []string{}
	for _, s := range songs {
		out = append(out, s.Id)
	}
	return out
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Token := r.URL.Query().Get("lptoken")

	if Token == "" {
		u, err := c.GetTokenURL("http://46.101.229.91/")
		if err != nil {
			http.Error(w, fmt.Sprintf("%s", err), 400)
			return
		}
		http.Redirect(w, r, u, 302)
		return
	}

	data := struct {
		Token string
	}{
		Token,
	}
	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		log.Fatalln("Unable to parse index")
	}

	err = t.Execute(w, data)

}

func API(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	enc := json.NewEncoder(w)
	localCatalog := catalog[:]
	token := r.Header.Get("X-LP-Token")

	ids := getIdsFromCatalog(localCatalog)

	accessStats := c.Access(token, ids...)

	for id, access := range accessStats {
		if access {
			continue
		}
		for idx, l := range localCatalog {
			if l.Id == id {
				i := laterpay.ItemDefinition{
					Id:      id,
					Pricing: "EUR23",
					Title:   l.Title,
				}
				url, err := c.Add(i)
				if err != nil {
					log.Println(err)
				}
				localCatalog[idx].Url = url
				localCatalog[idx].Access = false
			}
		}
	}

	enc.Encode(localCatalog)

}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/api/songs/", API)
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	log.Fatal(http.ListenAndServe(":80", router))
}
