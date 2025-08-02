package htmlWriter

import (
	"fmt"
	"github.com/violetcircus/viviblogger/configReader"
	"html/template"
	"log"
	"os"
	"strings"
)

type FrontMatter struct {
	Tags     []string
	Created  string
	Uploaded string
	Updated  string
}

type Post struct {
	FrontMatter FrontMatter
	Title       string
	Body        string
	Preview     string
}

func Build(post Post) {
	config := configReader.GetConfig()

	// read template from file
	tpl, err := os.ReadFile(config.TemplateFile)
	if err != nil {
		log.Fatal("error reading template file:", err)
	}
	templateString := string(tpl)

	t, err := template.New("webpage").Parse(templateString)
	if err != nil {
		log.Fatal("error building webpage:", err)
	}

	// add post body
	t, err = t.Parse(post.Body)
	if err != nil {
		log.Fatal("error building content:", err)
	}

	// add frontmatter
	t, err = t.Parse(buildFrontMatter(post.FrontMatter))
	if err != nil {
		log.Fatal("error building frontmatter:", err)
	}

	f, err := os.Create(config.PostsDir + "index.html")
	if err != nil {
		log.Fatal("error writing to file", err)
	}
	defer f.Close()

	err = t.Execute(f, post)
}

func buildFrontMatter(f FrontMatter) string {
	var builder strings.Builder
	builder.WriteString(`{{ define "frontmatter" }}`)

	builder.WriteString(`<div id="times">`)
	builder.WriteString(fmt.Sprintf(`<p id="uploaded">uploaded: %s</p><p id="updated">updated: %s</p>`, f.Uploaded, f.Updated))
	builder.WriteString(`</div>`)

	builder.WriteString(`<div id="tagcontainer">tags:`)
	for _, tag := range f.Tags {
		builder.WriteString(fmt.Sprintf(`<div id="tag">%s</div>`, tag))
	}
	builder.WriteString(`</div>`)

	builder.WriteString(`{{ end }}`)

	return builder.String()
}
