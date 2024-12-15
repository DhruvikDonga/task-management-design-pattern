package tasks

import (
	"fmt"
	"sort"
)

// Strategy pattern
type TaskSortStrategy interface {
	SortTasks(tasks []*Task)
}

type PriorityAscStrategy struct{}

func (p *PriorityAscStrategy) SortTasks(tasks []Task) {
	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].priority < tasks[j].priority
	})
}

type PriorityDescStrategy struct{}

func (p *PriorityDescStrategy) SortTasks(tasks []*Task) {
	fmt.Println("SORTING IN DEC")
	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].priority > tasks[j].priority
	})
}

func (t *TaskRepo) SetStrategy(strategy TaskSortStrategy) {
	t.Strategy = strategy
}

// Sort uses the current strategy to sort the tasks
func (t *TaskRepo) Sort() {
	t.Strategy.SortTasks(t.Tasks)
}
