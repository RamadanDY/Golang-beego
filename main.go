package main 

//This is Beego, the web framework.
//It lets us create routes, handle HTTP requests, and send responses.
//strconv is commented out because you’re not converting strings to numbers in this version.
import (
	"github.com/beego/beego/v2/server/web"

	//"strconv"
)

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`	
}

//this is our personal slice that we created
var task = []Task{
	{ID: 1, Name: "Task 1"},
	{ID: 2, Name: "Task 2"},
	{ID: 3, Name: "Task 3"},
}



type TaskController struct{
	web.Controller
}




func (c *TaskController) Get() {
	c.Data["json"] = task
	c.ServeJSON()
}



 func main() {
	web.Router("/task",&TaskController{})
 }

































//5️⃣ Controller definition
//we define the controller and give it an identity ie using the web.Controller
//By embedding web.Controller, it gets all the “Beego tools”
//like sending JSON or reading request data.

//so basically this controler make it posible to  send json ,read data and set http status codes 
type TaskController struct {
	web.Controller
}

//Controller method it defines what happens in the controler
//(c *TaskController)  This part is called a receiver.“This function belongs to the TaskController type.”
// So this function is a method of the controller.
// c is just a variable name representing the controller instance.
// The * means pointer, which allows the controller to modify its data.
// Get() ,This is the name of the method that this controler will be using 
func (c *TaskController) Get() {
	//Data is a map inside the controller used to store data that will be returned to the client 
	// or the response that we will get from the user.
	// c.Data["json"] this ensures that we return a json data 
	// Now we assign the value that we have to the slice above.So we are telling Beego:“Send the task slice as JSON.”
	// this line is to actually prepare the data 
	c.Data["json"] = task

	//This command actually sends the response to the client.
	// and this line will send the data to the client
	c.ServeJSON()


}



//Router registration
////Whenever someone visits /task, send them to this controller.”
//Beego then automatically looks for a method matching the HTTP method:
func main(){
	web.Router("/task", &TaskController{})
	web.Run()
}