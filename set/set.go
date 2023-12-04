package set

type Set[T comparable] map[T]struct{}

var SetElement struct{}

func New[T comparable]() Set[T] {
	return Set[T]{}
}

func NewFromMap[T comparable, V any](old map[T]V) Set[T] {
	new := New[T]()
	for k := range old {
		new[k] = SetElement
	}
	return new
}

func (universe Set[T]) Contains(query Set[T]) bool {
	for key := range query {
		if _, ok := universe[key]; !ok {
			return false
		}
	}
	return true
}

func (s Set[T]) CountMatches(query Set[T]) int {
	count := 0
	for key := range query {
		if _, ok := s[key]; ok {
			count++
		}
	}
	return count
}

func (a Set[T]) Equiv(b Set[T]) bool {
	return a.Contains(b) && b.Contains(a)
}
