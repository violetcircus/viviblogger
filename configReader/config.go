package configReader

type Config struct {
	SiteDir        string
	PostsDir       string
	NotesDir       string
	ImageDir       string
	SourceImageDir string
	TemplateFile   string
	DateTimeFormat string
}

func GetConfig() Config {
	config := Config{
		SiteDir:        "/home/violet/projects/viviblogger/",
		PostsDir:       "./",
		NotesDir:       "/home/violet/projects/viviblogger/",
		ImageDir:       "img/",
		SourceImageDir: "/home/violet/pictures/",
		TemplateFile:   "blogtemplate.html",
		DateTimeFormat: "2006-01-021504",
	}
	return config
}
