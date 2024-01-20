package tbmp

import (
	"net/url"
	"strconv"
)

// Location ...
type Location struct {
	Scheme   string
	User     string
	Host     string
	Port     int
	Path     string
	Query    map[string]string
	Fragment string
}

// NewLocation ...
func NewLocation(u *url.URL) *Location {

	port, _ := strconv.Atoi(u.Port())

	q1 := u.Query()
	q2 := make(map[string]string)
	for name, values := range q1 {
		for _, value := range values {
			q2[name] = value
		}
	}

	l := &Location{}
	l.Scheme = u.Scheme
	l.User = u.User.Username()
	l.Host = u.Hostname()
	l.Port = port
	l.Path = u.Path
	l.Query = q2
	l.Fragment = u.Fragment

	return l
}
