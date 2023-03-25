package html

import (
	"regexp"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
)

const DOCTYPE = "<!doctype html>"

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
