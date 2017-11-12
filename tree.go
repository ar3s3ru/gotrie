package gotrie

type Tree interface {
	Insert(key string, data interface{}) (Tree, bool)
	Update(key string, data interface{}) (Tree, bool)
	Delete(key string) (interface{}, Tree, bool)
	Keys() []string

	// Query stuff
	Get(key string) (interface{}, bool)
}
