package helpers

import "regexp"

func IsValidURL(url string) bool {
	// kalo string kosong gak usah divalidasi
	if url == "" {
		return true
	}

	urlRegex := regexp.MustCompile(`(?i)\b((?:https?|ftp)://|www\.)\S+\.\w{2,4}\b`)
	return urlRegex.MatchString(url)
}
