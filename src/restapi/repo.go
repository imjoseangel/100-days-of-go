///bin/true; exec /usr/bin/env go run "" "$@"
package main

var currentID int

var todos Todos

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

// // RepoFindTodo ...
// func RepoFindTodo(id int) Todo {
// 	for _, t := range todos {
// 		if t.ID == id {
// 			return t
// 		}
// 	}
// 	// return empty Todo if not found
// 	return Todo{}
// }

// RepoCreateTodo ...
func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

// // RepoDestroyTodo ...
// func RepoDestroyTodo(id int) error {
// 	for i, t := range todos {
// 		if t.ID == id {
// 			todos = append(todos[:i], todos[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
// }
