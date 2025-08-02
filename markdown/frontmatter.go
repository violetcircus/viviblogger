package markdown

// handle frontmatter
// like upload date, tags etc

import (
	"bufio"
	"github.com/violetcircus/viviblogger/configReader"
	"github.com/violetcircus/viviblogger/htmlWriter"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

// im assuming all frontmatter values are formatted as list items
func handleFrontMatter(lines []string) htmlWriter.FrontMatter {
	field := regexp.MustCompile(`[^-*]:$`)
	value := regexp.MustCompile(`\s*-[^0-9]\s*`)

	var result htmlWriter.FrontMatter
	var currentField string
	for _, line := range lines {
		if field.MatchString(line) {
			currentField = line
			// log.Println("frontmatter field:", currentField)
		} else if value.MatchString(line) {
			switch currentField {
			case "tags:":
				result.Tags = append(result.Tags, value.ReplaceAllString(line, ""))
				// log.Println("frontmatter value:", result.Tags)
			case "created:":
				result.Created = value.ReplaceAllString(line, "")
				// log.Println("frontmatter value:", result.Created)
			case "uploaded:":
				result.Uploaded = value.ReplaceAllString(line, "")
				log.Println("frontmatter value:", result.Uploaded)
			case "updated:":
				// log.Println("frontmatter value:", result.Updated)
				result.Updated = value.ReplaceAllString(line, "")
			}
		}
	}
	// log.Printf("created: '%s'", result.Created)
	if result.Uploaded == "" {
		result.Uploaded = handleTime("uploaded")
	}
	result.Updated = handleTime("updated")
	return result
}

// you can change the date and time format in your notes in the config file.
// it has to use go's specific reference time as a base, though.
// also, only the latest upload date is passed to the struct, despite the markdown file having each
// new upload date plonked at the top of the updated field. this is by design: it's cool for you to know
// every time you edited your blog, but you might not want that information to be as easily accessed by other
// people. this way it doesn't have to be! hooray. if you're reading this, you can easily change that:
// simply change "updated" to an array in the struct definition and have updated behave similarly to the tags field
// in the switch case above. to make it not track update history at all, have it delete everything in that field
// before adding the new datetime. i believe in you. or i'll add it later as a config file option idk but it's not
// important to me right now
func handleTime(field string) string {
	config := configReader.GetConfig()
	file, err := os.OpenFile(os.Args[1], os.O_RDWR, os.ModeAppend)
	if err != nil {
		log.Fatal("error opening markdown file!", err)
	}
	dateTime := time.Now().Format(config.DateTimeFormat)

	// construct array out of file contents with current time and date added
	scanner := bufio.NewScanner(file)
	var fileLines []string
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
		if strings.Contains(scanner.Text(), field) {
			fileLines = append(fileLines, "   - "+dateTime)
		}
	}
	file.Close()

	// rewrite file with new values in it
	editedFile := []byte(strings.Join(fileLines, "\n"))
	err = os.WriteFile(os.Args[1], editedFile, 0644)
	if err != nil {
		log.Fatalf("error writing %s time to markdown file! %s", field, err)
	}

	return dateTime
}
