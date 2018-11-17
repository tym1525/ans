package session

type SessionMgr interface {
	Init(addr string, options ...string)
	Get(sessionId string) (Session, error)
}
