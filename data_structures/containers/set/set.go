package set

// A Go map type looks like this: map[KeyType]ValueType
// Our Set can take any type
type SetWrapper map[interface{}]struct{}

// Returns a new Set
func Set() SetWrapper {
	return make(SetWrapper)
}

func (S SetWrapper) Add(e interface{}) {
	S[e] = struct{}{}
}

func (S SetWrapper) Remove(e interface{}) {
	delete(S, e)
}

func (S SetWrapper) Contains(e interface{}) bool {
	_, ok := S[e]
	return ok
}

func (S SetWrapper) Discard(e interface{}) {
	if S.Contains(e) {
		S.Remove(e)
	}
}

func (S SetWrapper) Clean() {
	for k := range S {
		delete(S, k)
	}
}

func (S SetWrapper) Equals(s SetWrapper) bool {
	if len(S) != len(s) {
		return false
	}

	for e := range S {
		if !s.Contains(e) {
			return false
		}
	}

	return true
}

func (S SetWrapper) Subset(sup SetWrapper) bool {
	for e := range S {
		if !sup.Contains(e) {
			return false
		}
	}
	return true
}

func (S SetWrapper) Union(sets ...SetWrapper) SetWrapper {
	u := Set()
	for _, ss := range append(sets, S) {
		for e := range ss {
			u.Add(e)
		}
	}
	return u
}

func (S SetWrapper) Intersection(sets ...SetWrapper) SetWrapper {
	i := Set()
	for e := range S {
		a := true
		for _, s := range sets {
			if !s.Contains(e) {
				a = false
			}
		}
		if a {
			i.Add(e)
		}

	}
	return i
}
