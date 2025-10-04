package service

import (
	"flashchat/models"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetUsers
// @Description get all users
// @Tags Users
// @Produce json
//
//	@Success 200 {
//		  "message": [
//		    {
//		      "ID": 1,
//		      "CreatedAt": "2025-10-04T00:42:03.878-04:00",
//		      "UpdatedAt": "2025-10-04T00:42:03.878-04:00",
//		      "DeletedAt": null,
//		      "Name": "admin",
//		      "Password": "",
//		      "Email": "",
//		      "Identity": "",
//		      "ClientIP": "",
//		      "ClientPort": "",
//		      "LastLogin": "0001-01-01T00:00:00Z",
//		      "HeartBeatTime": "0001-01-01T00:00:00Z",
//		      "LogOutTime": "0001-01-01T00:00:00Z",
//		      "IsLogout": false,
//		      "DeviceInfo": ""
//		    }
//		  ]
//		} json {code, message}
//
// @Router /users/getUsers [get]
func GetUsers(c *gin.Context) {
	users := make([]*models.UserBasic, 10)
	users = models.GetUsers()
	c.JSON(200, gin.H{
		"message": users,
	})
}
