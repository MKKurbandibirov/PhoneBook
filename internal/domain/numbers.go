package domain

type Number struct {
	Id        uint
	FirstName string
	LastName  string
	Country   string
	Number    int64
}

type NumbersFilter struct {
	Limit int
}

type Operation string

const (
	Select Operation = "get"
	Update Operation = "update"
	Insert Operation = "insert"
	Delete Operation = "delete"
	Exit   Operation = "exit"
)
