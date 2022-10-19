// Statement: Suppose you run a university and you need to maintain record of each student (name, age, branchName, expectedGradYear, univName)

package main

import "fmt"

type Protptype interface {
	clone()
}

type student struct {
	name             string
	age              int
	branchName       string
	univName         string
	expectedGradYear int
}

type studentCriteria struct {
	registry map[string]student
}

func (sc *studentCriteria) addNewCategory(key string, s student) {
	sc.registry[key] = s
}

func (sc *studentCriteria) get(key string) student {
	return sc.registry[key]
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

func main() {
	studentCriteria := &studentCriteria{}
	studentCriteria.registry = make(map[string]student)

	// setting up the base criteria for each batch
	CSE_2019 := newStudent()
	CSE_2019.branchName, CSE_2019.expectedGradYear, CSE_2019.univName = "CSE", 2019, "PES"
	studentCriteria.addNewCategory("CSE_2019", CSE_2019)

	ECE_2019 := newStudent()
	ECE_2019.branchName, ECE_2019.expectedGradYear, ECE_2019.univName = "ECE", 2019, "PES"
	studentCriteria.addNewCategory("ECE_2019", ECE_2019)

	CSE_2020 := newStudent()
	CSE_2020.branchName, CSE_2020.expectedGradYear, CSE_2020.univName = "CSE", 2020, "PES"
	studentCriteria.addNewCategory("CSE_2020", CSE_2020)

	ECE_2020 := newStudent()
	ECE_2020.branchName, ECE_2020.expectedGradYear, ECE_2020.univName = "ECE", 2020, "PES"
	studentCriteria.addNewCategory("ECE_2020", ECE_2020)

	// Need to enroll student for CSE_2019 batch
	// here we can create copy of object CSE_2019, which already contains the all the necessary details to quality for CSE_2019 batch.
	// and then we can add additional details to newly clone object
	s1 := studentCriteria.get("CSE_2019").clone()
	s1.setName("S1")
	s1.setAge(20)

	s2 := studentCriteria.get("CSE_2019").clone()
	s2.setName("S2")
	s2.setAge(19)

	// Need to enroll student for ECE_2019 batch
	s3 := studentCriteria.get("ECE_2019").clone()
	s3.setName("S3")
	s3.setAge(21)

	s4 := studentCriteria.get("ECE_2019").clone()
	s4.setName("S4")
	s4.setAge(20)

	fmt.Printf("\n%#v\n\n", studentCriteria.registry)
	fmt.Printf("\n%#v\n\n", s1)
	fmt.Printf("\n%#v\n\n", s2)
	fmt.Printf("\n%#v\n\n", s3)
	fmt.Printf("\n%#v\n\n", s4)
}
