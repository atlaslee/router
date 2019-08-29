package router

type Context interface {
	Sessions() []Session
	Session(string) Session
	DeleteSession(string)
	SetSession(Session)
}
