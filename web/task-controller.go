package web

import (
	"github.com/sirupsen/logrus"
	"go-backend-sample/dao"
	"go-backend-sample/model"
	"net/http"
	"time"
)

const (
	prefixTask = "/tasks"
)

// TaskController is a controller for tasks resources
type TaskController struct {
	taskDao dao.TaskDAO
	Routes  []Route
	Prefix  string
}

// NewTaskController creates a new task controller to manage tasks
func NewTaskController(taskDAO dao.TaskDAO) *TaskController {
	controller := TaskController{
		taskDao: taskDAO,
		Prefix:  prefixTask,
	}

	var routes []Route
	// Get
	routes = append(routes, Route{
		Name:        "Get one task",
		Method:      http.MethodGet,
		Pattern:     "/{id}",
		HandlerFunc: controller.GetTask,
	})
	// Create
	routes = append(routes, Route{
		Name:        "Create an task",
		Method:      http.MethodPost,
		Pattern:     "",
		HandlerFunc: controller.CreateTask,
	})
	// Update
	routes = append(routes, Route{
		Name:        "Update an task",
		Method:      http.MethodPut,
		Pattern:     "/{id}",
		HandlerFunc: controller.UpdateTask,
	})
	// Delete
	routes = append(routes, Route{
		Name:        "Delete an task",
		Method:      http.MethodDelete,
		Pattern:     "/{id}",
		HandlerFunc: controller.DeleteTask,
	})

	controller.Routes = routes

	return &controller
}

// Get retrieve a task by its id
func (ctrl *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	taskId := ParamAsString("id", r)
	logrus.Println("task : ", taskId)

	task, err := ctrl.taskDao.Get(taskId)
	if err != nil {
		logrus.Error(err)
		SendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logrus.Println("task : ", task)
	SendJSONOk(w, task)
}

// Create create a task
func (ctrl *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	logrus.Println("create task")
	task := &model.Task{}
	err := GetJSONContent(task, r)
	if err != nil {
		logrus.Error(err)
		SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.CreationDate = time.Now()
	task.Status = 0

	task, err = ctrl.taskDao.Upsert(task)
	if err != nil {
		logrus.Error(err)
		SendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logrus.Println("task : ", task)
	SendJSONWithHTTPCode(w, task, http.StatusCreated)
}

// Update update a task by its id
func (ctrl *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	task := &model.Task{}
	err := GetJSONContent(task, r)
	if err != nil {
		logrus.Error(err)
		SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	logrus.Println("update task : ", task.Id)

	task.ModificationDate = time.Now()

	taskExist, err := ctrl.taskDao.Exist(task.Id)
	if err != nil {
		logrus.Error(err)
		SendJSONError(w, err.Error(), http.StatusNotFound)
		return
	} else if taskExist == false {
		SendJSONError(w, "task not found", http.StatusNotFound)
		return
	}

	task, err = ctrl.taskDao.Upsert(task)
	if err != nil {
		logrus.Error(err)
		SendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logrus.Println("task : ", task)
	SendJSONOk(w, task)
}

// Delete delete a task by its id
func (ctrl *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := ParamAsString("id", r)
	logrus.Println("delete task : ", taskId)

	err := ctrl.taskDao.Delete(taskId)
	if err != nil {
		logrus.Error(err)
		SendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logrus.Println("deleted task : ", taskId)
	SendJSONWithHTTPCode(w, nil, http.StatusNoContent)
}