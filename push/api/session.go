package api


/**
 * session 持久化对象，用于快速重连
 */
type SessionStorage interface {
	SaveSession(PersistentSession)
	GetSession() (PersistentSession)
	ClearSession()
}

type PersistentSession struct {
	SessionId  string
	ExpireTime int64
	Cipher0    Cipher
}