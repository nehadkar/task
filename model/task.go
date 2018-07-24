package model

import "gopkg.in/mgo.v2/bson"

//Task defines the task information stored in the database
type Task struct {
	Tsk string `bson: "task", json: "task"`
	ID bson.ObjectId  `bson:"_id,omitempty", json:"_id"` 
}
