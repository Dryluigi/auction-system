package firebase

type RealtimeDatabase interface {
	Save(key string, data interface{}) error
	Update(key string, data interface{}) error
}
