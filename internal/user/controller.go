package user

import (
	"net/http"
	"oms-test/models"

	"github.com/gin-gonic/gin"
)

type UserContorller struct {
	service UserService
}

func (con *UserContorller) SearchUser(c *gin.Context) {
	usernameOrEmail := c.Query("query")
	user, err := con.service.searchUsers(&usernameOrEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and Email does not exist!"})
		return
	}
	c.JSON(http.StatusAccepted, user)
}

func (con *UserContorller) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	received_user, err := con.service.createUser(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, received_user)
}

func (con *UserContorller) FetchUser(c *gin.Context) {
	id := c.Param("id")

	user, err := con.service.getUser(id)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, user)
}

func (con *UserContorller) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	user, err := con.service.getUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to find the user id that you want to update"})
		return
	}

	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = con.service.updateUser(user, &newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, newUser)
}

func (con *UserContorller) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	con.service.deleteUser(id)
	c.JSON(http.StatusAccepted, &models.User{})
}

