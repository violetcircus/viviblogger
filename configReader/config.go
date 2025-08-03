package configReader

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Config struct {
	SiteDir        string
	PostsDir       string
	ImageDir       string
	SourceImageDir string
	TemplateFile   string
	DateTimeFormat string
}

// get config directory
func getConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	return configDir + "/vvblogger/"
}

func writeConfig() {
	configFile := getConfigDir() + "config"
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("error getting user home directory!")
	}
	var builder strings.Builder
	builder.WriteString("SiteDir=" + home + "\n")
	builder.WriteString("PostsDir=posts/\n")
	builder.WriteString("ImageDir=img/\n")
	builder.WriteString("SourceImageDir=\n")
	builder.WriteString("TemplateFile=\n")
	builder.WriteString("DateTimeFormat=2006-01-021504")

	config := []byte(builder.String())
	os.WriteFile(configFile, config, 0644)
}

func MakeConfig() {
	configDir := getConfigDir()
	if _, err := os.Stat(configDir + "config"); err == nil {
		log.Println("config already exists..")
	} else {
		os.Mkdir(configDir, 0777)
		configFile, err := os.Create(configDir + "config")
		if err != nil {
			log.Fatal("failed to create config file!", err)
		}
		defer configFile.Close()
		writeConfig()
	}
}

func GetConfig() Config {
	configDir := getConfigDir()
	f, err := os.Open(configDir + "config")
	if err != nil {
		log.Fatal("error reading config file!", err)
	}

	var config Config

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		field, value, _ := strings.Cut(scanner.Text(), "=")
		switch field {
		case "SiteDir":
			config.SiteDir = strings.TrimSpace(value)
		case "PostsDir":
			config.PostsDir = strings.TrimSpace(value)
		case "ImageDir":
			config.ImageDir = strings.TrimSpace(value)
		case "SourceImageDir":
			config.SourceImageDir = strings.TrimSpace(value)
		case "TemplateFile":
			config.TemplateFile = strings.TrimSpace(value)
		case "DateTimeFormat":
			config.DateTimeFormat = strings.TrimSpace(value)
		}
	}

	return config
}
