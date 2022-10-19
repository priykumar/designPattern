package main

type studentCriteria struct {
	registry map[string]student
}

func (sc *studentCriteria) addNewCategory(key string, s student) {
	sc.registry[key] = s
}

func (sc *studentCriteria) get(key string) student {
	return sc.registry[key]
}
