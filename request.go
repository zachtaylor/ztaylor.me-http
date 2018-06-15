package http

import (
	"net/http"
	"strings"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

type Request struct {
	Quest    string
	Remote   string
	Language string
	Data     js.Object
	*Session
	Agent
}

func NewRequest() *Request {
	return &Request{
		Language: "en-US",
		Data:     js.Object{},
	}
}

func RequestFromNet(r *http.Request, w http.ResponseWriter) *Request {
	req := NewRequest()
	req.Quest = r.RequestURI
	req.Remote = r.RemoteAddr[0:strings.LastIndex(r.RemoteAddr, ":")]
	// req.Language = r.Header.Get("Accept-Language")[0:5]
	req.Agent = AgentFromNetHttp(r, w)
	if acceptLang := r.Header.Get("Accept-Language"); len(acceptLang) < 5 || acceptLang[:5] != "en-US" {
		log.Add("Remote", req.Remote).Add("AcceptLanguage", acceptLang).Add("Agent", req.Agent.Name()).Warn("/api/cards: request language failed")
	}
	for k, v := range r.Form {
		req.Data[k] = v
	}
	if session, _ := ReadRequestCookie(r); session != nil {
		req.Session = session
	}
	return req
}

func RequestFromSocketMessage(msg *SocketMessage, s *Socket) *Request {
	req := NewRequest()
	req.Quest = msg.Uri
	req.Remote = s.Name()
	req.Data = msg.Data
	req.Agent = s
	req.Session = s.Session
	return req
}
