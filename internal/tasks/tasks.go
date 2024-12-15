package tasks

import (
	"fmt"
	"sync"

	"github.com/DhruvikDonga/task-management-desgin-pattern/internal/notify"
	"github.com/DhruvikDonga/task-management-desgin-pattern/internal/users"
)

const (
	BACKLOG  = "Backlog"
	TODO     = "Todo"
	PROGRESS = "Progress"
	DONE     = "Done"
)

const (
	LOW    = "Low"
	HIGH   = "High"
	URGENT = "Urgent"
)

type TaskManager interface {
	GetTaskDetails(taskid int) *Task
	GetAllTasks() []*Task
	SetTaskStatus(taskid int, status string) string
	UpdateTaskStatus(taskid, userid int) string
}

type TaskRepo struct {
	Tasks    []*Task
	Users    users.UserManager
	Strategy TaskSortStrategy
}

var taskrepo *TaskRepo
var once sync.Once

type TaskData interface {
	DescribeTask() string
	SetTaskStatus(status string, userid int) string
}
type Task struct {
	taskid      int
	name        string
	description string
	assignee    int
	reportedby  int
	priority    string
	status      string
	users       users.UserManager
	subscribers []notify.Notify
}

type TaskDecorator struct {
	task *Task
}

type PrioritizeTask struct {
	TaskDecorator
	priorityLevel string
}

func NewPriorityDecorator(task *Task, priority string) *PrioritizeTask {
	return &PrioritizeTask{
		TaskDecorator: TaskDecorator{task: task},
		priorityLevel: priority,
	}
}

func (pd *PrioritizeTask) SetPriority(priority string) {
	pd.priorityLevel = priority
	pd.task.priority = priority
}

// Singelton to create a task repo
func CreateTaskRepo(users users.UserManager) *TaskRepo {
	once.Do(func() {
		taskrepo = &TaskRepo{}
		taskrepo.Users = users
	})
	return taskrepo
}

func NewTask(repo *TaskRepo, name, description, priority, status string, id, assignee, reportedby int) {
	assigned := repo.Users.GetUserDetails(assignee)
	if assigned == nil {
		fmt.Println(fmt.Errorf("ERROR: Assignee %d is not valid for task %s", assignee, name).Error())
		return
	}
	reporter := repo.Users.GetUserDetails(reportedby)
	if reporter == nil {
		fmt.Println(fmt.Errorf("ERROR: Reporter %d is not valid for task %s", reportedby, name).Error())
		return
	}
	t := &Task{
		taskid:      id,
		assignee:    assignee,
		reportedby:  reportedby,
		name:        name,
		description: description,
		priority:    priority,
		status:      status,
		users:       repo.Users,
	}
	t.RegisterReceiver(assigned)
	t.RegisterReceiver(reporter)
	repo.Tasks = append(repo.Tasks, t)
	fmt.Println("Task ", name, " added")
}

func (t *Task) DescribeTask() string {
	return fmt.Sprintf("Task ID :- %d \n Task name :- %s \n Task description:- %s \n Assigned To:- %s \n Reported by:- %s \n Priority:- %s \n Status:- %s \n",
		t.taskid,
		t.name,
		t.description,
		t.users.GetUserDetails(t.assignee).GetUserName(),
		t.users.GetUserDetails(t.reportedby).GetUserName(),
		t.priority,
		t.status)
}

func (t *Task) RegisterReceiver(user notify.Notify) {
	t.subscribers = append(t.subscribers, user)
}

func (t *Task) NotifcationReceivers(message string) {
	var wg sync.WaitGroup

	for _, ta := range t.subscribers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ta.Update(message)
		}()
	}
	wg.Wait()
}

func (t *TaskRepo) GetAllTasks() []*Task {
	return t.Tasks
}

func (t *TaskRepo) GetTaskDetails(id int) *Task {
	for _, ta := range t.Tasks {
		if ta.taskid == id {
			return ta
		}
	}
	return nil

}

func (t *Task) SetTaskStatus(status string, userid int) string {

	currentStatus := t.status
	t.status = status

	return "Status changed for task " + t.name + " from " + currentStatus + " to " + status + " by " + t.users.GetUserDetails(userid).GetUserName()
}

func (t *TaskRepo) FilterTasksByAssignee(username string) []*Task {
	var tasks []*Task

	for _, task := range t.GetAllTasks() {
		if t.Users.GetUserDetails(task.assignee).GetUserName() == username {
			tasks = append(tasks, task)
		}
	}

	return tasks
}

func (t *TaskRepo) UpdateTaskStatus(taskid, userid int, status string) string {
	res := "failed to update the task"

	task := t.GetTaskDetails(taskid)
	res = task.SetTaskStatus(status, userid)
	task.NotifcationReceivers(res)
	return res
}
