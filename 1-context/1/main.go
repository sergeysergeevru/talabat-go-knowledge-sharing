package main

import (
	"context"
	"fmt"
)

/*
Topic: getting a root context
*/
func main() {
	fmt.Println("Background context cmp", context.Background() == context.Background())
	fmt.Println("TODO context cmp", context.TODO() == context.TODO())
	fmt.Println("Background and TODO context cmp", context.Background() == context.TODO())

	// context.Background for the root goroutine => main()
	// context.TODO for non main goroutine and functions

	executeBusinessLogic()

}

func executeBusinessLogic() {
	getFromMemory()
	getFromDB(context.TODO())
}

// legacy function
func getFromMemory() string {
	return "some memory data"
}

func getFromDB(ctx context.Context) string {
	// some logic over DB has been omitted
	return "db data"
}
