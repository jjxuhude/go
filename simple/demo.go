package simple

type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

func NewStudent(id uint, name string, male bool, score float64) *Student {
	return &Student{id, name, male, score}
}

func (s *Student) SetName(name string) {
	s.name = name
}

type Number1 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
}

type Number2 interface {
	Equal(i int) bool
	MoreThan(i int) bool
	LessThan(i int) bool
}

type Number int

func (n Number) Equal(i int) bool {
	return int(n) == i
}

func (n Number) LessThan(i int) bool {
	return int(n) < i
}

func (n Number) MoreThan(i int) bool {
	return int(n) > i
}
