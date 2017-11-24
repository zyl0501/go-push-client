package session

import "github.com/zyl0501/go-push-client/push/api"

type MemorySessionStorage struct {
	session api.PersistentSession
}

func (storage *MemorySessionStorage) SaveSession(session api.PersistentSession) {
	storage.session = session
}
func (storage *MemorySessionStorage) GetSession() (api.PersistentSession) {
	return storage.session
}
func (storage *MemorySessionStorage) ClearSession() {
	storage.session = api.PersistentSession{}
}
