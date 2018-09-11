package url

import "net/url"

type URlFactory struct{}

func (URlFactory) Parse(rawurl string) (*url.URL, error) {
	return url.Parse(rawurl)
}

func (URlFactory) ParseRequestURI(rawurl string) (*url.URL, error) {
	return url.ParseRequestURI(rawurl)
}

func (URlFactory) User(username string) *url.Userinfo {
	return url.User(username)
}

func (URlFactory) UserPassword(username, password string) *url.Userinfo {
	return url.UserPassword(username, password)
}

func (URlFactory) ParseQuery(query string) (url.Values, error) {
	return url.ParseQuery(query)
}

func (URlFactory) PathEscape(s string) string {
	return url.PathEscape(s)
}

func (URlFactory) QueryUnescape(s string) (string, error) {
	return url.QueryUnescape(s)
}

func (URlFactory) QueryEscape(s string) string {
	return url.QueryEscape(s)
}
