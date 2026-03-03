package main

import (
	"net/url"
)

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	var parsedURLPassword string
	// make sure we don't  return a User without a Password
	if password, isSet := parsedUrl.User.Password(); isSet {
		parsedURLPassword = password
	}
	return ParsedURL{
		protocol: parsedUrl.Scheme,
		username: parsedUrl.User.Username(),
		password: parsedURLPassword,
		hostname: parsedUrl.Hostname(),
		port:     parsedUrl.Port(),
		pathname: parsedUrl.Path,
		search:   parsedUrl.RawQuery,
		hash:     parsedUrl.Fragment,
	}
}
