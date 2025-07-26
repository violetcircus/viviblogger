package output

import (
	"html/template"
	"log"
	"os"
)

type Post struct {
	FrontMatter map[string]string
	Title       string
	Body        string
	Preview     string
}

const tpl = `
<html>
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>the violet circus</title>
    <link rel="icon" href="/favicon.png" type="image/x-icon">
    <link href="/style.css" rel="stylesheet" type="text/css" media="all">
    <script src="js/quote.js"></script>
    <script src="js/dayssince.js"></script>
  </head>
  <body>
    <!-- title area -->
    <div class="top">
      <div class="title">
        <h1>the violet circus</h1>
        <div id="quoteDisplay"></div>
        <br>
      </div>
    </div>
    <div>
      <div id="blogpost">
        <div id="posthead">
          <div id="title"><h1>{{.Title}}</h1></div>
        </div>
        <div id="postbody">
          <div id="preview">{{.Preview}}</div>
          <div id="content">{{.Body}}</div>
        </div>
      </div>
    </div>
  </body>
</html> `

func Build(post Post) {
	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		log.Fatal("error building webpage:", err)
	}
	err = t.Execute(os.Stdout, post)
}
