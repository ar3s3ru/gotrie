package gotrie

type Tree interface {
	Insert(key string, data interface{}) (Tree, bool)
	Update(key string, data interface{}) (Tree, bool)
	Delete(key string) (interface{}, Tree, bool)
	Query(key string) (interface{}, bool)
}
