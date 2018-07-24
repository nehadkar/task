package database

import "testing"
import "task/model"

var initDBTests = [] struct {
	host string
	err bool 
} {
	{"localhost", false},
	{"192.168.123.123", true},
	{"localhost:27017", false},
	{"localhost:3200", true},
}

var AddTaskTests = [] model.Task {
		{Tsk:"Test Me"},
		{Tsk: ""}, 
		{Tsk: ""}, 
}

func TestInitDB (t *testing.T) {
	for _, tt := range initDBTests {
		db := DBSession{}
		err := db.Init(tt.host)
		if err != nil {
			if tt.err == false {
				t.Errorf("FAIL: Init(%s): Expected no error, Received %v", tt.host, err)
			}else {
				t.Logf("SUCCESS: Init(%s) worked as expected", tt.host)
			}
		} else {
			if tt.err == true {
				t.Errorf("FAIL: Init(%s): Expected error, Receiver nil", tt.host)
			}else {
				t.Logf("SUCCESS: Init(%s) worked as expected", tt.host)
			}
			db.Close()
		}
	}
}

func TestAddDB (t *testing.T) {
	db := DBSession{}
	err := db.Init("localhost")
	if err != nil {
		t.Logf("FAIL: Unable to connect to DB, err: %v", err)
		t.FailNow()
	}
	defer db.Close()
	for _, tt := range AddTaskTests {
		err = db.AddTask(tt)
		if (err != nil) {
			t.Errorf("FAIL: AddTask(%v) returned error %v", tt, err)
		}
		tasks, err := db.GetTasks()
		if (err != nil) {
			t.Errorf("FAIL: GetTasks() returned error %v while testing AddTask(%v)", err, tt)
		}
		if len(tasks) == 0 {
			t.Errorf("FAIL: Unable to retrieve recently added task %v", tt)
		}else {
			var found bool
			for _, tk := range(tasks) {
				if tk.Tsk == tt.Tsk {
					found = true
					break
				}
			}
			if found == false{
				t.Errorf("FAIL: Unable to retrieve recently added task %v", tt)
			} else {
				t.Logf("SUCCESS: Could retrieve recently added task %v", tt)
			}
		}
	}
	// Not cleaning up.. will use same set to test DoTask
	tasks, err := db.GetTasks()
	if (err != nil) {
		t.Errorf("FAIL: GetTasks() returned error %v ", err)
	}
	lenT  := len(tasks)
	err = db.DoTask(tasks[0])
	if (err != nil) {
		t.Errorf("FAIL: DoTask(%v) returned error %v ", tasks[0], err)
	}
	tasks, err = db.GetTasks()
	if (err != nil) {
		t.Errorf("FAIL: GetTasks() returned error %v ", err )
	}
	lenN := len(tasks)
	if lenT != lenN + 1{
		t.Errorf("FAIL: Task array length is incorrect, expected %d, got %d", lenT - 1, lenN)
	}
}
