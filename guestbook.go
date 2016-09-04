package main

import (
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
)

var tmpl = `
<html>
  <head>
    <title>Slothful Soda Guestbook</title>
  </head>
  <body>
    <h1>Slothful Soda guestbook</h1>
    <div>
       {{ range $i, $msg := .Messages }}
       <p>{{ $msg }}</p>
       {{ end }}
    </div>
    <form method="POST" action="/guestbook">
      <textarea name="message" cols="80" rows="10"></textarea><br />
      <input type="submit" name="Submit" value="Submit" />
    </form>
  </body>
</html>`

var htmlTemplate *template.Template

// Prepare our HTML template so we can use it in this web app;
// panic if there is an error
func init() {
	var err error
	htmlTemplate, err = template.New("guestbook").Parse(tmpl)

	if err != nil {
		panic(err.Error())
	}
}

// In a real server, we would use a database so we don't lose our posts if
// the server goes down an for not taking up too much memory; we're just
// using a slice here to keep the Go script short
var messages = make([]string, 0, 100)

func main() {
	http.HandleFunc("/guestbook", func(w http.ResponseWriter, r *http.Request) {
		reqtxt, _ := httputil.DumpRequest(r, true)
		log.Println(string(reqtxt))

		// If we get a POST request that includes a message, add the message
		// to our "database"
		if r.Method == "POST" {
			// [WARNING!] In a real web app taking in user input, this would
			// be bad design to accept the user input as it was typed; this
			// opens up the possibility for cross-site scripting attacks,
			// a very common security hole on a lot of websites
			if message := r.FormValue("message"); message != "" {
				messages = append(messages, message)
			}
		}
		htmlTemplate.Execute(w, struct{ Messages []string }{messages})
	})
	http.ListenAndServe(":1123", nil)
}
