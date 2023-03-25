package html

import (
	"regexp"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
)

const DOCTYPE = "<!doctype html>"

// CleanHTML takes a given html string and transforms it
// to a string consisting only of its title and content
func CleanHTML(html string) string {
	space := regexp.MustCompile(`\s+`)
	lower := strings.ToLower(html)
	noDoc := strings.ReplaceAll(lower, DOCTYPE, "")
	noTags := strip.StripTags(noDoc)
	noSpaces := space.ReplaceAllString(noTags, " ")
	noSpaces = strings.ReplaceAll(noSpaces, "\n\t", "")
	cleaned := strings.TrimSpace(noSpaces)
	return cleaned
}
