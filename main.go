package main

import (
	"fmt"
	"net/http"
)

var todayList = []string{"Put the coffee poster up on the wall", "Buy rocket parts", "Buy glue to stick all parts"}
var tomorrowList = []string{"Buy rocket fuel at a nearby andhra restaurant", "Open rocket's mouth and push all food into it", "Wait for the rocket to take off"}

func main() {
	//fmt.Println("Go World, you got this! Bring me a blue lays packet when you come back home\n")
	fmt.Print("##### Welcome to my Todo list app: #######\n\n")

	http.HandleFunc("/hello", helloGo) //Handle function handles the user request
	http.HandleFunc("/list", showTasks)
	http.HandleFunc("/", listPaths)

	http.ListenAndServe(":8080", nil) // Started http server at 8080

}

func helloGo(writer http.ResponseWriter, reader *http.Request) {
	greeting := "Go World, you got this! Bring me a blue lays packet when you come back home"
	fmt.Fprintln(writer, greeting)
}

func showTasks(writer http.ResponseWriter, reader *http.Request) {

	fmt.Fprintln(writer, "Category: Today")
	for i, task := range todayList {
		fmt.Fprintf(writer, "%d) %s\n", i+1, task)
	}

	fmt.Fprint(writer, "\nCategory: Tomorrow\n")
	for i, task := range tomorrowList {
		fmt.Fprintf(writer, "%d) %s\n", i+1, task)
	}
}

func listPaths(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintln(writer, "List of paths you can go to:")
	fmt.Fprintln(writer, "1. /hello")
	fmt.Fprintln(writer, "2. /list")
}

func printTasks(taskItems []string) {
	fmt.Println("To-do list items:")
	for index, task := range taskItems {
		fmt.Printf("%d) %s", index, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	newTaskItems := append(taskItems, newTask)
	return newTaskItems
}
