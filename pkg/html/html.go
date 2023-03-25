package html

import (
	"regexp"
	"strings"
)

const DOCTYPE = "<!doctype html>"

// CleanHTML takes a given html string and transforms it
// to a string consisting only of its title and content
func CleanHTML(html string) string {
	space := regexp.MustCompile(`\s+`)
	lower := strings.ToLower(html)
	noDoc := strings.ReplaceAll(lower, DOCTYPE, "")
	noTags := removeTags(noDoc)
	noSpaces := space.ReplaceAllString(noTags, " ")
	noSpaces = strings.ReplaceAll(noSpaces, "\n\t", "")
	cleaned := strings.TrimSpace(noSpaces)
	return cleaned
}

func removeTags(s string) string {
	re := regexp.MustCompile(`(?s)<script.*?>.*?</script>|<style.*?>.*?</style>|<[^>]*>`)
	return re.ReplaceAllString(s, "")
}
