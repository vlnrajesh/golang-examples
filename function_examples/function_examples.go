package FunctionExamples

import (
	"fmt"
)

//SwapByValue function returns given strings swapped by value
func SwapByValue(x string, y string) (string, string) {
	temp := x
	x = y
	y = temp
	return x, y
}

//SwapByReference function returns given strings swapped by reference
func SwapByReference(x *string, y *string) {
	temp := *x
	*x = *y
	*y = temp
}

// UserStructure structure which has Firstname and Lastname
type UserStructure struct {
	FirstName, LastName, Gender string
}

//MethodGreeting for greeting user
func (u UserStructure) MethodGreeting() string {
	if u.Gender == "Male" {
		return fmt.Sprintf("Hello Mr. %s, %s", u.FirstName, u.LastName)
	} else if u.Gender == "Female" {
		fmt.Println(u)
		return fmt.Sprintf("Hello Ms. %s, %s", u.FirstName, u.LastName)
	} else {
		return fmt.Sprintf("Hello %s, %s I am not aware of your Gender to greet", u.FirstName, u.LastName)
	}
}

//UpdateName FirstName
func (u *UserStructure) UpdateName(newName string) {
	(*u).FirstName = newName
}
