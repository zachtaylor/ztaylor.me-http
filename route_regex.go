package http

import (
	"regexp"
)

type RegexRoute struct {
	route
	*regexp.Regexp
}

func NewRouteRegex(s string, r ResponderFunc) Route {
	return &RegexRoute{
		route:  route{r},
		Regexp: regexp.MustCompile(s),
	}
}

func (route *RegexRoute) Match(s string) bool {
	return route.MatchString(s)
}