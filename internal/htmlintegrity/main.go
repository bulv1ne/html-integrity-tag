package htmlintegrity

func GenerateHtmlIntegrityTag(url string, shaSize int) (string, error) {
	data, err := DownloadFileToBytes(url)
	if err != nil {
		return "", err
	}

	integrityString := IntegirtyString(data, shaSize)

	return HtmlTag(url, integrityString), nil
}
