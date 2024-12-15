package main

import (
	"fmt"

	"github.com/DhruvikDonga/task-management-desgin-pattern/internal/tasks"
	"github.com/DhruvikDonga/task-management-desgin-pattern/internal/users"
)

func main() {
	fmt.Println("TASK management")

	//Create userRepo
	u := users.CreateUserRepo()
	//Create taskRepo
	t := tasks.CreateTaskRepo(u)
	//create users
	users.NewUser(u, "John", "Software Engineer", 1)
	users.NewUser(u, "Jake", "Software Engineer", 2)
	users.NewUser(u, "Alice", "Program Manager", 3)
	users.NewUser(u, "Alice", "Senior Engineer", 4)

	//create tasks
	tasks.NewTask(t, "Backend Revamp", "Go Backend Revamp changes", tasks.HIGH, tasks.BACKLOG, 1, 1, 4)
	tasks.NewTask(t, "Frontend Revamp", "UI Revamp changes", tasks.LOW, tasks.BACKLOG, 2, 2, 3)
	tasks.NewTask(t, "Infrastrucutre optimization", "K8s config changes changes", tasks.URGENT, tasks.BACKLOG, 3, 1, 3)
	tasks.NewTask(t, "Database Revamp", "Database Revamp changes", tasks.HIGH, tasks.BACKLOG, 4, 4, 3)

	//update with decorator
	taskWithPriority := tasks.NewPriorityDecorator(t.GetTaskDetails(4), tasks.HIGH)
	fmt.Println("TASK uppdated:-")
	taskWithPriority.SetPriority(tasks.LOW)

	fmt.Println(taskWithPriority)
	tasks.NewTask(t, "Lunch", "lunch changes", tasks.URGENT, tasks.BACKLOG, 5, 10, 3)

	tasklist := t.GetAllTasks()
	for _, v := range tasklist {
		fmt.Println(v.DescribeTask())
	}

	filtertasklist := t.FilterTasksByAssignee("Alice")
	fmt.Println("Filtered by username :- Alice")
	for _, v := range filtertasklist {
		fmt.Println(v.DescribeTask())
	}

	//update task backend revamp
	fmt.Println("Update task 1 by user 1:-")
	fmt.Println(t.UpdateTaskStatus(1, 1, tasks.PROGRESS))

	fmt.Println("SORTING tasks")
	t.SetStrategy(&tasks.PriorityDescStrategy{})
	t.Sort()
	for _, v := range t.GetAllTasks() {
		fmt.Println(v.DescribeTask())
	}
}
