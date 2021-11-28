package repository

type user struct {
	ID      string
	Name    string
	Address string
}

type circle struct {
	ID    *string
	Name  *string
	Owner *string
}

type circleMember struct {
	ID       *string
	CircleID *string
	UserID   *string
}
