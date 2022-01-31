package users

import (
	"net/http"
	"strconv"

	"github.com/Bookstore-GolangMS/bookstore_oauth-go/oauth"
	"github.com/LordRadamanthys/bookstore_users-api/domain/users"
	"github.com/LordRadamanthys/bookstore_users-api/services"
	"github.com/LordRadamanthys/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.BadRequestError("invalid json body", err)
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.UserService.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	isPublicRequest := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusCreated, result.Marshall(isPublicRequest))
}

func GetUser(c *gin.Context) {

	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	if callerId := oauth.GetCallerId(c.Request); callerId == 0 {
		err := rest_errors.RestErr{
			Status:  http.StatusUnauthorized,
			Message: "resource not available",
		}
		c.JSON(err.Status, err)
		return
	}

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := rest_errors.BadRequestError("invalid user id", userErr)
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.UserService.GetUser(int(userId))

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	if oauth.GetCallerId(c.Request) == int64(user.Id) {
		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}
	isPublicRequest := oauth.IsPublic(c.Request)
	c.JSON(http.StatusOK, user.Marshall(isPublicRequest))
}

func UpdateUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := rest_errors.BadRequestError("invalid user id", userErr)
		c.JSON(err.Status, err)
		return
	}
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.BadRequestError("invalid json body", err)
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = int(userId)

	isPatch := c.Request.Method == http.MethodPatch

	result, err := services.UserService.UpdateUser(isPatch, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	isPublicRequest := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusOK, result.Marshall(isPublicRequest))
}

func DeleteUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := rest_errors.BadRequestError("invalid user id", userErr)
		c.JSON(err.Status, err)
		return
	}

	if deleteErr := services.UserService.DeleteUser(int(userId)); deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UserService.FindByStatus(status)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	isPublicRequest := c.GetHeader("X-Public") == "true"

	c.JSON(http.StatusOK, users.Marshall(isPublicRequest))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.BadRequestError("invalid json body", err)
		c.JSON(restErr.Status, restErr)
		return
	}

	user, err := services.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	isPublicRequest := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusOK, user.Marshall(isPublicRequest))
}
