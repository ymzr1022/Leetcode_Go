package option

type Student struct {
	name  string
	age   int
	score int
}

type options func(msg *Student)

func WithName(name string) options {
	return func(msg *Student) {
		msg.name = name
	}
}

func NewStudent(age int, opts ...options) *Student {
	s := &Student{
		age: age,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
