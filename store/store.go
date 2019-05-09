package store

// Writer is used to write date to the store
type Writer interface {
	Write(key string, value interface{}) error
}

// Reader is used to read data from the store
type Reader interface {
	Read(key string) (interface{}, error)
}

// NewCacheWriter creates an instance of Writer
func NewCacheWriter() Writer {
	return &Cache{}
}

// NewCacheReader creates an instance of Writer
func NewCacheReader() Reader {
	return &Cache{}
}
