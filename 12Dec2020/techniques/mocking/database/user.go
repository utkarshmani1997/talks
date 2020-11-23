package database

const (
	insertQuery = `
INSERT INTO users (age, email, first_name, last_name)
VALUES ($1, $2, $3, $4)
RETURNING id
`
)

type Users interface {
	Insert() (int, error)
}

type User struct {
	Age       int
	Email     string
	FirstName string
	LastName  string
}

func NewUser(age int, email, first, last string) Users {
	return User{age, email, first, last}
}

func (u User) Insert() (int, error) {
	id := 0
	err := db.QueryRow(insertQuery, u.Age, u.Email, u.FirstName, u.LastName).Scan(&id)
	return id, err
}
