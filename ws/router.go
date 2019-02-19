package ws

import "regexp"

// RouterSet is used to combine Routers into a single Router
type RouterSet []Router

// Match checks that all included Routers return true
func (set RouterSet) Route(m *Message) bool {
	for _, router := range set {
		if router == nil || !router.Route(m) {
			return false
		}
	}
	return true
}

type routerFunc func(*Message) bool

func (f routerFunc) Route(m *Message) bool {
	return f(m)
}

// RouterFunc turns a func into a Router
func RouterFunc(f func(*Message) bool) Router {
	return routerFunc(f)
}

type routerRegex struct {
	*regexp.Regexp
}

func (rgx *routerRegex) Route(m *Message) bool {
	return rgx.MatchString(m.URI)
}

// RouterRegex creates a regexp match check against Message.Name
func RouterRegex(s string) Router {
	return &routerRegex{regexp.MustCompile(s)}
}

// RouterLit creates a literal match check against Message.Name
type RouterLit string

func (s RouterLit) Route(m *Message) bool {
	return string(s) == m.URI
}
