package users

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonhealy1/bookstore_/bookstore_users-api/domain/users"
)

// GetUser -
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// CreateUser -
func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(err)
	fmt.Println(string(bytes))
	c.String(http.StatusNotImplemented, "implement me!")
}

// SearchUser -
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
