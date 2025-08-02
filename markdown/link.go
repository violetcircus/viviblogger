package markdown

// handle links and image embeds
import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/violetcircus/viviblogger/configReader"
)

func handleLinks(buf string) string {
	config := configReader.GetConfig()
	link := regexp.MustCompile(`[\[][^\\]*[\]]\(.*\)`)
	wikiLink := regexp.MustCompile(`[\[][\[][a-zA-Z0-9\s]+[\]][\]]`)
	content := strings.Split(buf, "\n")
	var formatted []string
	for _, line := range content {
		// handle all normal links on the line
		linkResult := link.ReplaceAllStringFunc(line, func(match string) string {
			// extract link text from square brackets
			textStart := strings.IndexRune(match, '[')
			textEnd := strings.IndexRune(match, ']')
			text := match[textStart+1 : textEnd]

			// extract URL from brackets
			urlStart := strings.IndexRune(match, '(')
			urlEnd := strings.IndexRune(match, ')')
			url := match[urlStart+1 : urlEnd]

			// check if link is to a resource on the system or if on the web
			isExternal := strings.HasPrefix(url, "http")

			// get url basename and file extension
			urlBasename := filepath.Base(url)
			extensionStart := strings.Index(urlBasename, ".")
			var urlFileExtension string
			if extensionStart >= 0 {
				urlFileExtension = urlBasename[strings.Index(urlBasename, "."):]
			} else {
				urlFileExtension = ""
			}

			// create anchor/image tag string
			switch urlFileExtension {
			case ".png", ".PNG", ".jpg", ".JPG", ".jpeg", ".JPEG", ".avif", ".AVIF", ".webp", ".WEBP", ".gif", ".GIF":
				if isExternal {
					result := fmt.Sprintf(`<img src=%s alt=%s></img>`, url, text)
					return result
				} else {
					newUrl := handleImages(urlBasename)
					result := fmt.Sprintf(`<img src=%s alt=%s></img>`, newUrl, text)
					return result
				}
			default:
				result := fmt.Sprintf(`<a href=%s>%s</a>`, url, text)
				return result
			}
		})
		// hanndle all wikilinks on the line
		wikiLinkResult := wikiLink.ReplaceAllStringFunc(linkResult, func(match string) string {
			// extract link from brackets
			link := strings.TrimSuffix(strings.TrimPrefix(match, "[["), "]]")
			url := strings.TrimPrefix(url.PathEscape(link), "/")
			result := fmt.Sprintf(`<a href=%s%s.html>%s</a>`, config.PostsDir, url, link)

			return result
		})
		// append final string to formatted array
		formatted = append(formatted, wikiLinkResult)
	}
	result := strings.Join(formatted, "\n")
	fmt.Print(result)
	return result
}

func handleImages(link string) string {
	config := configReader.GetConfig()
	originalImage := config.SourceImageDir + link
	imageDir := config.SiteDir + config.ImageDir

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
