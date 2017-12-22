package dao

import (
	"errors"
	"github.com/satori/go.uuid"
	"go-backend-sample/model"
)

var _ TaskDAO = (*TaskDAOMock)(nil)

// MockedTask is the task returned by this mocked interface
var MockedTask = model.Task{
	Id:          uuid.NewV4().String(),
	Title:       "TestMock",
	Description: "TestMock",
}

// TaskDAOMock is the mocked implementation of the TaskDAO
type TaskDAOMock struct {
	storage map[string]*model.Task
}

// NewTaskDAOMock creates a new TaskDAO with a mocked implementation
func NewTaskDAOMock() TaskDAO {
	daoMock := &TaskDAOMock{
		storage: make(map[string]*model.Task),
	}

	// Adds some fake data
	daoMock.Upsert(&MockedTask)

	return daoMock
}

// Get return a task by its id
func (s *TaskDAOMock) Get(id string) (*model.Task, error) {
	task, ok := s.storage[id]
	if !ok {
		return nil, errors.New("task not found with id " + id)
	}
	return task, nil
}

// GetAll return all tasks
func (s *TaskDAOMock) GetAll() ([]model.Task, error) {
	var tasks []model.Task
	for taskId := range s.storage {
		task := s.storage[taskId]
		tasks = append(tasks, *task)
	}
	return tasks, nil
}

// Upsert update or create a task
func (s *TaskDAOMock) Upsert(task *model.Task) (*model.Task, error) {
	if task.Id == "" {
		task.Id = uuid.NewV4().String()
	}
	s.save(task)
	return task, nil
}

// Delete delete a task by its id
func (s *TaskDAOMock) Delete(id string) error {
	delete(s.storage, id)
	return nil
}

// DeleteAll deletes all tasks
func (s *TaskDAOMock) DeleteAll() error {
	for taskId := range s.storage {
		delete(s.storage, taskId)
	}
	return nil
}

// Exist check if the task exist
func (s *TaskDAOMock) Exist(id string) (bool, error) {
	if s.storage[id] != nil {
		return true, nil
	}
	return false, errors.New("task not found with id " + id)
}

// save the task
func (s *TaskDAOMock) save(task *model.Task) *model.Task {
	s.storage[task.Id] = task
	return task
}

// get a task by its id
func (s *TaskDAOMock) get(id string) *model.Task {
	task := s.storage[id]
	return task
}