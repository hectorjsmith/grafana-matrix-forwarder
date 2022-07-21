package formatter

import (
	"bytes"
	htmlTemplate "html/template"
	"regexp"
	"strings"
)

var (
	htmlTagRegex       = regexp.MustCompile(`<.*?>`)
	htmlParagraphRegex = regexp.MustCompile(`</?p>`)
)

func formattedMessageToPlainMessage(input string) string {
	return strings.TrimSpace(stripHtmlTagsFromString(input))
}

func stripHtmlTagsFromString(input string) string {
	bodyWithoutParagraphs := htmlParagraphRegex.ReplaceAllString(input, " ")
	plainBody := htmlTagRegex.ReplaceAllString(bodyWithoutParagraphs, "")
	return plainBody
}

func executeHtmlTemplate(template *htmlTemplate.Template, data alertMessageData) (string, error) {
	buffer := new(bytes.Buffer)
	err := template.Execute(buffer, data)
	return buffer.String(), err
}
