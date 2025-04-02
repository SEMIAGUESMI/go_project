package main
/*
import "errors"
import "fmt"

//import "LEARN_GO/types"

func main() {
	division, err := devide(10, 0)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(division)

	username, err := findUser("guest")
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("username", username)
	}

	item, err := getItem(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(item)

	
// custum error
type myerror struct {
	code    int
	message string
}

func (e *myerror) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}

func riskymessage() error {
	return &myerror{code: 404, message: "source not found"}
}
func devide(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}
func findUser(username string) (string, error) {
	if username != "admin" {
		return "", fmt.Errorf("user %s not found", username)
	}
	return "admin", nil

}

var ErrorNotFound = errors.New("item not found")

func getItem(id int) (string, error) {
	if id != 1 {
		return "", ErrorNotFound
	}
	return "item 1", nil
}
*/