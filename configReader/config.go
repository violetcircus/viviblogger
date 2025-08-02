package configReader

type Config struct {
	SiteDir        string
	PostsDir       string
	NotesDir       string
	ImageDir       string
	SourceImageDir string
}

func GetConfig() Config {
	config := Config{
		SiteDir:        "/home/violet/projects/viviblogger/",
		PostsDir:       "./",
		NotesDir:       "/home/violet/projects/viviblogger/",
		ImageDir:       "img/",
		SourceImageDir: "/home/violet/pictures/",
	}
	return config
}
