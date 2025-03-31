package htmlintegrity

import (
	"fmt"
	"strings"
)

const jsHtmlFormat = `<script src="%s" integrity="%s" crossorigin="anonymous"></script>`
const cssHtmlFormat = `<link href="%s" rel="stylesheet" integrity="%s" crossorigin="anonymous">`

func HtmlTag(url string, integrity string) string {
	format := jsHtmlFormat
	if strings.HasSuffix(url, ".css") {
		format = cssHtmlFormat
	}
	return fmt.Sprintf(format, url, integrity)
}
