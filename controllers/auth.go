package controllers

import (
	"encoding/json"
	"sample-api/models"
	"sample-api/utils"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type AuthBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthController operations for Auth
type AuthController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create Auth
// @Param	body		body 	models.Auth	true		"body for Auth content"
// @Success 201 {object} models.Auth
// @Failure 403 body is empty
// @router / [post]
func (c *AuthController) Post() {
	var v AuthBody
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	u, err := models.GetUserByUsername(v.Username)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	if u.Password == v.Password {
		if tok, err := utils.CreateToken(u.Id); err == nil {
			c.Data["json"] = tok
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Ctx.Output.SetStatus(400)
	}

	c.ServeJSON()
}

// Get ...
// @Title Validate
// @Description validate token
// @Success 201 {object} jwt.StandardClaims
// @Failure 403 body is empty
// @router / [get]
func (c *AuthController) Get() {
	authHeader := c.Ctx.Input.Header("Authorization")

	if len(authHeader) == 0 {
		c.Ctx.Output.SetStatus(401)
		c.Finish()
	}

	headerCon := strings.Split(authHeader, " ")

	if len(headerCon) != 2 {
		c.Ctx.Output.SetStatus(401)
		c.Finish()
	}

	claims, err := utils.ValidateToken(headerCon[1])

	if err != nil || (err == nil && claims == nil) {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = err
	} else {
		c.Data["json"] = claims
	}

	c.ServeJSON()
}
