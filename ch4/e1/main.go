package main

import (
	"net/url"
)

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	return ParsedURL{
		protocol: "",
		username: "",
		password: "",
		hostname: "",
		port:     "",
		pathname: "",
		search:   "",
		hash:     "",
	}
}
