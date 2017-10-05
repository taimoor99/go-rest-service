package task

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const (
	index           = `Task`
	userIndex       = `User`
	invalidTaskData = `error: invalid task data`
)

// User defines user attributes
type Task struct {
	ID          string `json:id`
	Title       string `json:"title"`
	Description string `json:"description"`
	Owner       string `json:"email"`
}

// Create an Task
func Create(c context.Context, tsk *Task) (*Task, error) {
	var output *Task
	if tsk == nil || tsk.Title == `` || tsk.Owner == `` {
		return nil, fmt.Errorf(invalidTaskData)
	}

	output, _ = GetByID(c, tsk.Owner, tsk.ID)

	if output == nil {
		parentKey := datastore.NewKey(c, userIndex, tsk.Owner, 0, nil)
		key := datastore.NewKey(c, index, tsk.ID, 0, parentKey)
		insKey, err := datastore.Put(c, key, tsk)

		if err != nil {
			log.Errorf(c, "ERROR INSERTING TASK: %v", err.Error())
			return nil, err
		}

		output, err = GetByID(c, tsk.Owner, insKey.StringID())
		if err != nil {
			log.Errorf(c, "ERROR GETTING TASK OUTPUT: %v", err.Error())
			return nil, err
		}
		return output, nil
	}
	log.Infof(c, "Task was previously saved: %v", tsk.ID)
	return output, nil
}

// GetByID an Task based on its ID
func GetByID(c context.Context, userEmail string, ID string) (*Task, error) {
	if ID == `` {
		return nil, fmt.Errorf(invalidTaskData)
	}
	parentKey := datastore.NewKey(c, userIndex, userEmail, 0, nil)
	key := datastore.NewKey(c, index, ID, 0, parentKey)
	var tsk Task
	err := datastore.Get(c, key, &tsk)

	if err != nil {
		if strings.HasPrefix(err.Error(), `datastore: no such entity`) {
			err = fmt.Errorf(`task '%v' not found`, ID)
		}
		return nil, err
	}
	return &tsk, nil
}

// GetTasks Fetches all users
func GetTasks(c context.Context, userEmail string) ([]Task, error) {
	var output []Task
	parentKey := datastore.NewKey(c, `User`, userEmail, 0, nil)
	q := datastore.NewQuery(index)
	q.Ancestor(parentKey)
	_, err := q.GetAll(c, &output)

	if err != nil {
		log.Errorf(c, "error fetching all tasks for %v", userEmail)
		return nil, err
	}

	if len(output) <= 0 {
		return nil, fmt.Errorf("no tasks found for %v", userEmail)
	}
	return output, nil
}
