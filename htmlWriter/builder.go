package htmlWriter

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
  <body>
      <div id="blogpost">
        <div id="posthead">
          <div id="title"><h1>{{.Title}}</h1></div>
        </div>
        <div id="postbody">
          <div id="preview">{{.Preview}}</div>
          <div id="content">{{ template "content" . }}</div>
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

	t, err = t.Parse(post.Body)
	if err != nil {
		log.Fatal("error building content:", err)
	}
	f, err := os.Create("./index.html")
	if err != nil {
		log.Fatal("eror writing to file", err)
	}
	defer f.Close()

	// err = t.Execute(os.Stdout, post)
	err = t.Execute(f, post)
}
