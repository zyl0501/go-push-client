package session

import "github.com/zyl0501/go-push-client/push/api"

type FileSessionStorage struct {
}

func (storage *FileSessionStorage) SaveSession(session api.PersistentSession) {
}
func (storage *FileSessionStorage) GetSession() (api.PersistentSession) {
	return api.PersistentSession{}
}
func (storage *FileSessionStorage) ClearSession() {
}

func encode(session api.PersistentSession) api.PersistentSession {
	return session
}

func decode(value api.PersistentSession) api.PersistentSession {
	return value
}