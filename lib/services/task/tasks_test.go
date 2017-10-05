package task_test

import (
	"os"
	"testing"

	"context"

	"github.com/pedrocelso/go-rest-service/lib/services/task"
	"github.com/stretchr/testify/assert"

	"google.golang.org/appengine/aetest"
)

const email = `pedro@pedrocelso.com.br`

var mainCtx context.Context

func TestMain(m *testing.M) {
	ctx, done, _ := aetest.NewContext()
	mainCtx = ctx
	//_ = createTasks(mainCtx)
	os.Exit(m.Run())
	done()
}

func TestCreateTask(t *testing.T) {
	output, err := task.Create(mainCtx, &task.Task{
		ID:          `87287612387.1221687.13`,
		Title:       `Cool Test`,
		Description: `kind of a task, but different`,
		Owner:       email,
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, `87287612387.1221687.13`, output.ID)
	assert.Equal(t, `Cool Test`, output.Title)
	assert.Equal(t, `kind of a task, but different`, output.Description)
	assert.Equal(t, email, output.Owner)

	output, err = task.Create(mainCtx, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "error: invalid task data", err.Error())
	assert.Nil(t, output)

	output, err = task.Create(mainCtx, &task.Task{
		ID:          `87287612387.1221687.13`,
		Title:       `Cool Test`,
		Description: `kind of a task, but different`,
		Owner:       ``,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "error: invalid task data", err.Error())
	assert.Nil(t, output)

	output, err = task.Create(mainCtx, &task.Task{
		ID:          `87287612387.1221687.13`,
		Title:       ``,
		Description: `kind of a task, but different`,
		Owner:       email,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "error: invalid task data", err.Error())
	assert.Nil(t, output)
}

//func TestGetByEmail(t *testing.T) {
//	err := createUsers(mainCtx)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	output, err := user.GetByEmail(mainCtx, `pedro@pedrocelso.com.br1`)
//	assert.Nil(t, err)
//	assert.NotNil(t, output)
//	assert.Equal(t, "Pedro 1", output.Name)
//	assert.Equal(t, "pedro@pedrocelso.com.br1", output.Email)
//
//	output, err = user.GetByEmail(mainCtx, `bad_email@gmail.com`)
//	assert.NotNil(t, err)
//	assert.Equal(t, "user 'bad_email@gmail.com' not found", err.Error())
//	assert.Nil(t, output)
//
//	output, err = user.GetByEmail(mainCtx, ``)
//	assert.NotNil(t, err)
//	assert.Equal(t, `error: invalid User data`, err.Error())
//	assert.Nil(t, output)
//}
//
//// This test run ina  different context ot ensure that only
//// the created users will be saved on the datastore
//func TestGetUsers(t *testing.T) {
//	ctx, done, err := aetest.NewContext()
//	defer done()
//	err = createUsers(ctx)
//	if err != nil {
//		t.Fatal(err)
//	}
//	// This sleep is needed because it take some milliseconds for the objects
//	// created on `createUsers` to be indexed and returned on query
//	time.Sleep(time.Millisecond * 5e2)
//	output, err := user.GetUsers(ctx)
//	assert.Nil(t, err)
//	assert.NotNil(t, output)
//	assert.Equal(t, 5, len(output))
//}
//
//func TestUpdateUser(t *testing.T) {
//	err := createUsers(mainCtx)
//
//	output, err := user.Update(mainCtx, &user.User{
//		Name:  `Migeh`,
//		Email: `pedro@pedrocelso.com.br0`,
//	})
//	assert.Nil(t, err)
//	assert.NotNil(t, output)
//	assert.Equal(t, "Migeh", output.Name)
//	assert.Equal(t, "pedro@pedrocelso.com.br0", output.Email)
//
//	usr, err := user.GetByEmail(mainCtx, `pedro@pedrocelso.com.br0`)
//	assert.Nil(t, err)
//	assert.NotNil(t, output)
//	assert.Equal(t, "Migeh", usr.Name)
//	assert.Equal(t, "pedro@pedrocelso.com.br0", usr.Email)
//
//	output, err = user.Update(mainCtx, nil)
//	assert.NotNil(t, err)
//	assert.Equal(t, "error: invalid User data", err.Error())
//	assert.Nil(t, output)
//}
//
//func TestDeleteUser(t *testing.T) {
//	err := createUsers(mainCtx)
//
//	usr, err := user.GetByEmail(mainCtx, `pedro@pedrocelso.com.br0`)
//	assert.Nil(t, err)
//	assert.NotNil(t, usr)
//	assert.Equal(t, "Pedro 0", usr.Name)
//	assert.Equal(t, "pedro@pedrocelso.com.br0", usr.Email)
//
//	err = user.Delete(mainCtx, `pedro@pedrocelso.com.br0`)
//	assert.Nil(t, err)
//
//	usr, err = user.GetByEmail(mainCtx, `pedro@pedrocelso.com.br0`)
//	assert.NotNil(t, err)
//	assert.Equal(t, "user 'pedro@pedrocelso.com.br0' not found", err.Error())
//	assert.Nil(t, usr)
//}
//
//func createTasks(ctx context.Context) error {
//	for i := 0; i < 5; i++ {
//		email := fmt.Sprintf(`%v%v`, email, i)
//		name := fmt.Sprintf(`Pedro %v`, i)
//		key := datastore.NewKey(ctx, `User`, email, 0, nil)
//		if _, err := datastore.Put(ctx, key, &user.User{
//			Name:  name,
//			Email: email,
//		}); err != nil {
//			return err
//		}
//	}
//	return nil
//}
