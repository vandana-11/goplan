package main

import "fmt"

type User struct {
	Fname string
	Lname string
	Age   int64
}

// GetFullName returns the first and lastname
func (u *User) GetFullName() (string, error) {
	return fmt.Sprintf("%s %s", u.Fname, u.Lname), nil
}

// GetAge returns the current user Age
func (u *User) GetAge() (int64, error) {
	return u.Age, nil
}

// GetFullDetails returns complete fname, lname and age
func (u *User) GetFullDetails() (string, error) {
	return fmt.Sprintf("%s %s is of Age %v", u.Fname, u.Lname, u.Age), nil
}

func main() {

	user := User{
		Fname: "Vandana",
		Lname: "Bhagat",
		Age:   36,
	}

	fmt.Println(user.GetFullName())
	fmt.Println(user.GetAge())
	fmt.Println(user.GetFullDetails())

}
