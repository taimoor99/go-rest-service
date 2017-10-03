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
	if tsk == nil || tsk.Title == `` {
		return nil, fmt.Errorf(invalidTaskData)
	}

	if output == nil {
		// func NewKey(c context.Context, kind, stringID string, intID int64, parent *Key) *Key {
		parentKey := datastore.NewKey(c, userIndex, tsk.Owner, 0, nil)
		key := datastore.NewKey(c, index, tsk.ID, 0, parentKey)
		insKey, err := datastore.Put(c, key, tsk)

		if err != nil {
			log.Errorf(c, "ERROR INSERTING TASK: %v", err.Error())
			return nil, err
		}

		output, err = GetByID(c, insKey.StringID())
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
func GetByID(c context.Context, ID string) (*Task, error) {
	if ID == `` {
		return nil, fmt.Errorf(invalidTaskData)
	}
	parentKey := datastore.NewKey(c, userIndex, `pedrocelsonunes@gmail.com`, 0, nil)
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
