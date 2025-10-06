package service

import (
	"flashchat/models"
	"flashchat/utils"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetUsers
// @Description get all users
// @Tags Users
// @Produce json
//
// @Success 200 {} json {code, message}
//
// @Router /users/getUsers [get]
func GetUsers(c *gin.Context) {
	users := make([]*models.UserBasic, 10)
	users = models.GetUsers()
	c.JSON(200, gin.H{
		"message": users,
	})
}

// @BasePath /api/v1

// GetUser
// @Description get user by name and password
// @Tags Users
// @Produce json
//
// @Success 200 {} json {code, message}
//
// @Router /users/getUsers [get]
func GetUser(c *gin.Context) {
	// user := models.UserBasic{}
	// name := c.Query("name")
	// password := c.Query("password")
	// user = models.FindUserByName(name)
	// if user.Identity == "" {
	// 	c.JSON(400, gin.H{
	// 		"message": "user does not exist",
	// 	})
	// 	return
	// }
	// match := utils.ValidPassword(password, user.Salt, user.Password)
	// if !match {
	// 	c.JSON(400, gin.H{
	// 		"message": "login fails",
	// 	})
	// 	return
	// }
	c.JSON(200, gin.H{
		"message": "user",
	})
}

type CreateUserRequest struct {
	Name       string `json:"name" binding:"required,min=2"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}

// @BasePath /api/v1

// CreateUser
// @Description create new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param  user  body  CreateUserRequest  true  "User Info"
// @Produce json
//
// @Success 200 {} json {code, message}
//
// @Router /users/createUser [post]
func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := models.UserBasic{}
	user.Name = req.Name
	user.Email = req.Email
	password := req.Password
	rePassword := req.RePassword
	if rePassword != password {
		c.JSON(400, gin.H{
			"error": "passwords are not matched",
		})
	}

	salt, _ := utils.GenerateSalt(31)
	user.Password = utils.MakePassword(password, salt)
	if result := models.CreateUsers(user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "user is successfully created",
	})
}

// @BasePath /api/v1

// Update User
// @Description update existing user by ID
// @Tags Users
//
// @Produce json
// @Param id formData string false "id"
// @Param name formData string false "name"
// @Param password formData string false "password"
// @Success 200 {} json {code, message}
//
// @Router /users/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	// user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	salt, _ := utils.GenerateSalt(31)
	plainpwd := c.PostForm("password")
	user.Password = utils.MakePassword(plainpwd, salt)

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid phone or email",
		})
	}
	if result := models.UpdateUser(user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "user is successfully updated",
	})
}

// @BasePath /api/v1

// Delete User
// @Description delete existing user by ID
// @Tags Users
// @Param        id        query     string  true  "User ID"
// @Produce json
//
// @Success 200 {} json {code, message}
//
// @Router /users/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)

	if result := models.DeleteUser(user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "user is successfully deleted",
	})
}
