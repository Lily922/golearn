package person

type Teacher struct {
	Name string
	Job  string
	Age  int
}

func (t *Teacher) GetName() string {
	return t.Name
}

func (t *Teacher) GetJob() string {
	return t.Job
}

func (t *Teacher) GetAge() int {
	return t.Age
}
