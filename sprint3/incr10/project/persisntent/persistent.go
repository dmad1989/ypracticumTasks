package persistent

import "github.com/dmad1989/ypracticumTasks/sprint3/incr10/project/store"

func Lookup(s store.Store, key string) ([]byte, error) {
	// ...
	return s.Get(key)
}
