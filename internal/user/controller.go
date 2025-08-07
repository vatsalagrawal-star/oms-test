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
		c.JSON(http.StatusBadRequest, gin.H{"error": "UError while running query"})
		return
	}
	if len(user) == 0 {
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

	receivedUser, err := con.service.createUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, receivedUser)
}

func (con *UserContorller) FetchUser(c *gin.Context) {
	id := c.Param("id")

	user, err := con.service.getUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusFound, user)
}

func (con *UserContorller) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	user, err := con.service.getUser(id)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "unable to find the user id that you want to update"},
		)
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
	c.Status(http.StatusAccepted)
}

func (con *UserContorller) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	con.service.deleteUser(id)
	c.Status(http.StatusAccepted)
}
