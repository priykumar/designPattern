package main

type student struct {
	name             string
	age              int
	branchName       string
	univName         string
	expectedGradYear int
}

func newStudent() student {
	return student{}
}

func (s *student) setName(name string) {
	s.name = name
}

func (s *student) setAge(age int) {
	s.age = age
}

func (s *student) setBranchName(branch string) {
	s.branchName = branch
}

func (s *student) setUnivName(uName string) {
	s.univName = uName
}

func (s *student) setGradYear(gYear int) {
	s.expectedGradYear = gYear
}

func (s student) clone() student {
	return s
}
