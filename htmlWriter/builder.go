package htmlWriter

import (
	"github.com/violetcircus/viviblogger/configReader"
	"html/template"
	"log"
	"os"
)

type Post struct {
	FrontMatter FrontMatter
	Title       string
	Body        string
	Preview     string
}

type FrontMatter struct {
	Tags     []string
	Created  string
	Uploaded string
	Updated  string
}

func Build(post Post) {
	config := configReader.GetConfig()

	tpl, err := os.ReadFile(config.TemplateFile)
	if err != nil {
		log.Fatal("error reading template file:", err)
	}
	templateString := string(tpl)

	t, err := template.New("webpage").Parse(templateString)
	if err != nil {
		log.Fatal("error building webpage:", err)
	}

	t, err = t.Parse(post.Body)
	if err != nil {
		log.Fatal("error building content:", err)
	}
	f, err := os.Create(config.PostsDir + "index.html")
	if err != nil {
		log.Fatal("error writing to file", err)
	}
	defer f.Close()

	// err = t.Execute(os.Stdout, post)
	err = t.Execute(f, post)
}
