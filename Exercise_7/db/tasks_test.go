package db

import "testing"

func TestAddTask(t *testing.T) {
	db, _ := Init("/home/gs-1708/tasks.db")
	err := AddTask("testing123")
	if err != nil {
		t.Errorf("Expected result No error, But got Error %v", err)
	}
	db.Close()
}

func TestInit(t *testing.T) {
	db, err := Init("/home/gs-1708/tasks.db")
	if err != nil {
		t.Errorf("Expected result No error, But got Error %v", err)
	}
	db.Close()
}

func TestInitNegative(t *testing.T) {
	_, err := Init("/home/gs-1707/task.db")
	if err == nil {
		t.Errorf("Expected result error, But got NO Error")
	}
}

func TestListTasks(t *testing.T) {
	db, _ := Init("/home/gs-1708/tasks.db")
	_, err := ListTasks()
	if err != nil {
		t.Errorf("Expected result No error, But got Error %v", err)
	}
	db.Close()
}

func TestDeleteTask(t *testing.T) {
	db, _ := Init("/home/gs-1708/tasks.db")
	err := DeleteTask(1)
	if err != nil {
		t.Errorf("Expected result No error, But got Error %v", err)
	}
	db.Close()
}
