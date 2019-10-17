package entity

type User struct {
	Name string
	Password string
	Email string
	Phone string
}

func GetName(a User)string {
	return a.Name
}

func GetPhone(a User) string{
	return a.Phone
}

func GetEmail(a User) string{
	return a.Email
}

func GetPassword(a User) string{
	return a.Password
}