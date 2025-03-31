package main

import (
	"fmt"
	"html-integrity-tag/internal"
	"html-integrity-tag/internal/flagparse"
	"html-integrity-tag/internal/htmlintegrity"
)

func main() {
	cliArgs := flagparse.Parse()

	htmlTag, err := htmlintegrity.GenerateHtmlIntegrityTag(cliArgs.Url, cliArgs.ShaSize)
	internal.HandleError(err)

	fmt.Println(htmlTag)
}
