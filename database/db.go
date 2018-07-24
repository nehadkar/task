//Package for database communication
package database

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"task/model"
)

//DBSession is a wrapper on the DB Handle 
type DBSession struct {
	sess *mgo.Session 
}

//Initialize connection to the database listening on 'host'
func (db *DBSession) Init(host string) (error) {
        sess, err := mgo.Dial(host)
        if err != nil {
		return err
        }
	db.sess = sess
        db.sess.SetMode(mgo.Monotonic, true)
        db.sess.DB("admin").Login("root","secret")
	index := mgo.Index{
		Key:        []string{"task"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	c :=  db.sess.DB("gopher").C("Tasks")
	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	} 
	return nil
}

//Close the connection to the database
func (db *DBSession) Close(){
	db.sess.Close()
}

//Retrieve the current tasks in the database
func (db *DBSession)GetTasks() ([]model.Task, error){

	c := db.sess.DB("gopher").C("Tasks")

	var tasks []model.Task
	err := c.Find(bson.M{}).All(&tasks)

	if( err != nil) {
		return nil, err 
	}

	return tasks, nil 
}

//Add a task to the database
func (db *DBSession) AddTask(task model.Task) (error) {
	c := db.sess.DB("gopher").C("Tasks")
	err := c.Insert(&task)
	if(err != nil){
		return err
	}

	return nil
}

//Do (Remove)  the task in the database
func (db *DBSession) DoTask( task model.Task) (error) {
	c := db.sess.DB("gopher").C("Tasks")
	//err := c.RemoveId(bson.M{"task":task.Tsk})
	err := c.RemoveId(task.ID)
	if(err != nil){
		return err
	}

	return nil
}

