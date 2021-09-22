package controllers

import (
	"encoding/json"
	"sample-api/models"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

//  TodoController operations for Todo
type TodoController struct {
	beego.Controller
}

// URLMapping ...
func (c *TodoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Todo
// @Param	body		body 	models.Todo	true		"body for Todo content"
// @Success 201 {int} models.Todo
// @Failure 403 body is empty
// @router / [post]
func (c *TodoController) Post() {
	var v models.Todo
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddTodo(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Todo by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Todo
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TodoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	v, err := models.GetTodoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Todo
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	page	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Todo
// @Failure 403
// @router / [get]
func (c *TodoController) GetAll() {
	var limit int = 10
	var page int

	// limit: 10 (default is 10)
	if v, err := c.GetInt("limit"); err == nil {
		limit = v
	}
	// page: 0 (default is 0)
	if v, err := c.GetInt("page"); err == nil {
		page = v
	}

	l, err := models.GetAllTodo(page, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Todo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Todo	true		"body for Todo content"
// @Success 200 {object} models.Todo
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TodoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Todo{Base: models.Base{Id: id}}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateTodoById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Todo
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TodoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteTodo(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
