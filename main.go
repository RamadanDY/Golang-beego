package main 

import (
	"github.com/beego/beego/v2/server/web"
	//"strconv"
)

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`	
}

var task = []Task{
	{ID: 1, Name: "Task 1"},
	{ID: 2, Name: "Task 2"},
	{ID: 3, Name: "Task 3"},
}


type TaskController struct {
	web.Controller
}


func (c *TaskController) Get(){
	c.Data["json"] = task
	c.ServeJSON()
}

func main(){
	web.Router("/task", &TaskController{})
	web.Run()
}