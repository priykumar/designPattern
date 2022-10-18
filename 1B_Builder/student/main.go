package main

import "fmt"

type student struct {
	name     string
	age      int
	univName string
	gradYear int
}

type studentBuilder struct {
	name     string
	age      int
	univName string
	gradYear int
}

func newStudentBuilder() studentBuilder {
	return studentBuilder{}
}

func (s *studentBuilder) setName(name string) {
	s.name = name
}

func (s *studentBuilder) setAge(age int) {
	s.age = age
}

func (s *studentBuilder) setUnivName(uName string) {
	s.univName = uName
}

func (s *studentBuilder) setGradYear(gYear int) {
	s.gradYear = gYear
}

func (s *studentBuilder) getStudent() student {
	return student{
		name:     s.name,
		age:      s.age,
		univName: s.univName,
		gradYear: s.gradYear,
	}
}

func main() {
	// craete a new builder
	newStudentBuilder := newStudentBuilder()

	// set all attributes of the builder
	newStudentBuilder.setName("Priyanshu")
	newStudentBuilder.setAge(26)
	newStudentBuilder.setUnivName("PES University")
	newStudentBuilder.setGradYear(2019)

	// create a new student object thru builder
	newStudent := newStudentBuilder.getStudent()
	fmt.Printf("%#v\n\n", newStudent)
}
