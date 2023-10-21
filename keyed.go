package gosds

type Keyed interface {
	NodeForKey(key string) Node
	ValueForKey(key string) (any, bool)
	SetValueForKey(key string, value any)
	RemoveValueForKey(key string)
	Keys() []string
	PrimitiveObject() map[string]any
}
