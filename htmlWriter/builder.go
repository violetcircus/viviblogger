package htmlWriter

import (
	"fmt"
	"github.com/violetcircus/viviblogger/configReader"
	"html/template"
	"log"
	"os"
	"strings"
	"time"
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
	t, err = t.Parse(buildFrontMatter(post.FrontMatter, config))
	if err != nil {
		log.Fatal("error building frontmatter:", err)
	}

	log.Println("page path:", config.SiteDir+config.PostsDir+post.Title+".html")
	f, err := os.Create(config.SiteDir + config.PostsDir + strings.TrimSpace(post.Title) + ".html")
	if err != nil {
		log.Fatal("error writing to file", err)
	}
	defer f.Close()

	err = t.Execute(f, post)
}

func changeTimeFormat(input string, config configReader.Config) string {
	t, err := time.Parse(config.DateTimeFormat, input)
	if err != nil {
		log.Fatal("error parsing time", err)
	}
	result := t.Format("02 Jan, 2006 at 15:04")

	return result
}

func buildFrontMatter(f FrontMatter, config configReader.Config) string {
	var builder strings.Builder
	builder.WriteString(`{{ define "frontmatter" }}`)

	uploaded := changeTimeFormat(f.Uploaded, config)
	updated := changeTimeFormat(f.Updated, config)

	builder.WriteString(`<div id="times">`)
	builder.WriteString(fmt.Sprintf(`<p id="uploaded">uploaded: %s</p><p id="updated">updated: %s</p>`, uploaded, updated))
	builder.WriteString(`</div>`)

	builder.WriteString(`<div id="tagcontainer">tags:`)
	for _, tag := range f.Tags {
		builder.WriteString(fmt.Sprintf(`<div id="tag">%s</div>`, tag))
	}
	builder.WriteString(`</div>`)

	builder.WriteString(`{{ end }}`)

	return builder.String()
}
