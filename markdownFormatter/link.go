package markdownFormatter

// handle links and image embeds
import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func handleLinks(buf string) string {
	link := regexp.MustCompile(`[\[][^\\]*[\]]\(.*\)`)
	content := strings.Split(buf, "\n")
	var formatted []string
	for _, line := range content {
		lineResult := link.ReplaceAllStringFunc(line, func(match string) string {
			textStart := strings.IndexRune(match, '[')
			textEnd := strings.IndexRune(match, ']')
			text := match[textStart+1 : textEnd]

			urlStart := strings.IndexRune(match, '(')
			urlEnd := strings.IndexRune(match, ')')
			url := match[urlStart+1 : urlEnd]
			urlBasename := filepath.Base(url)

			extensionStart := strings.Index(urlBasename, ".")
			var urlFileExtension string
			if extensionStart >= 0 {
				urlFileExtension = urlBasename[strings.Index(urlBasename, "."):]
			} else {
				urlFileExtension = ""
			}

			switch urlFileExtension {
			case ".png", ".PNG", ".jpg", ".JPG", ".jpeg", ".JPEG", ".avif", ".AVIF", ".webp", ".WEBP", ".gif", ".GIF":
				newUrl := handleImages(urlBasename)
				result := fmt.Sprintf(`<img src=%s alt=%s></img>`, newUrl, text)
				return result
			default:
				result := fmt.Sprintf(`<a href=%s>%s</a>`, url, text)
				return result
			}
		})
		formatted = append(formatted, lineResult)
	}
	result := strings.Join(formatted, "\n")
	fmt.Print(result)
	return result
}

func handleImages(link string) string {
	imageDir := "/home/violet/projects/viviblogger/img/"
	vaultImageDir := "/home/violet/pictures/"
	originalImage := vaultImageDir + link

	data, err := os.ReadFile(originalImage)
	if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
		log.Fatalf("error reading image file %s:%s", originalImage, err)
	} else if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		log.Printf("image file %s does not exist in vault directory!", link)
	} else {
		err = os.WriteFile(imageDir+link, data, 0644)
		if err != nil {
			log.Println("failed to copy image file!")
		}
	}

	imagePath := "img/" + link
	return imagePath
}
