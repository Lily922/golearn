package person

type Student struct {
	Name string
	Job  string
	Age  int
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetJob() string {
	return s.Job
}

func (s *Student) GetAge() int {
	return s.Age
}
