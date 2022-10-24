package internal

import (
	"errors"
	"fmt"
	"mytodo/golangtodo/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TaskDB struct {
	db *gorm.DB
}

func NewTaskDB() (*TaskDB, error) {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	out := &TaskDB{db: db}
	err = out.MaybeInit()
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *TaskDB) dbExists() (bool, error) {
	// Check for table.
	hasList := t.db.Migrator().HasTable(model.List{})
	hasEntries := t.db.Migrator().HasTable(model.ListEntry{})
	if hasList != hasEntries {
		return false, ErrDatabaseNotValid
	}
	return hasList, nil
}

// Initialize DB if it does not already exist.
func (t *TaskDB) MaybeInit() error {
	exists, err := t.dbExists()
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	return t.db.AutoMigrate(&model.List{}, &model.ListEntry{})
}

// List Methods defined in interface.
func (t *TaskDB) GetAllListNames() ([]string, error) {
	users := make([]model.List, 0)
	if err := t.db.Select("list_name").Find(&users).Error; err != nil {
		return nil, err
	}
	out := make([]string, 0, len(users))
	for _, user := range users {
		out = append(out, user.ListName)
	}
	return out, nil
}

func (t *TaskDB) CreateNewList(listName string) error {
	// We do not need to check if the list with this name already exists in the
	// DB since the DB will check for us by the uniqueness constraint on the
	// list name.
	list := model.NewList(listName)
	if err := t.db.Create(&list).Error; err != nil {
		return fmt.Errorf("error when trying to create a new list: %w", err)
	}
	return nil
}

func (t *TaskDB) GetList(listName string) (model.List, error) {
	out := model.List{}
	if err := t.db.Where("list_name = ?", listName).First(&out).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return out, ErrListNotFound
		}
		return out, fmt.Errorf("error while getting list %s: %w", listName, err)
	}
	return out, nil
}

func (t *TaskDB) UpdateList(listName, newName string) error {
	list, err := t.GetList(listName)
	if err != nil {
		return err
	}
	list.ListName = newName
	return t.db.Save(&list).Error
}

func (t *TaskDB) DeleteList(listName string) error {
	list, err := t.GetList(listName)
	if err != nil {
		return err
	}
	if err = t.db.Delete(&list).Error; err != nil {
		return fmt.Errorf("error when trying to delete %s: %w", listName, err)
	}
	return nil
}

func (t *TaskDB) ArchiveList(listName string) error {
	return nil
}

func (t *TaskDB) AddTaskToList(listName, taskSummary string) error {
	list, err := t.GetList(listName)
	if err != nil {
		return fmt.Errorf("error getting existing list: %w", err)
	}
	list.Items = append(list.Items, model.NewListEntry(taskSummary))
	if err = t.db.Save(&list).Error; err != nil {
		return fmt.Errorf("error while updating list: %w", err)
	}
	return nil
}

// Task Methods defined in interface.
func (t *TaskDB) GetTask(listName, string, taskId int) (model.ListEntry, error) {
	return model.ListEntry{}, nil
}

func (t *TaskDB) EditTask(listName, string, taskId int, newSummary string) error {
	return nil
}

func (t *TaskDB) DeleteTask(listName, string, taskId int) error {
	return nil
}
