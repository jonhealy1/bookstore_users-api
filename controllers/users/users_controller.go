package users

import (
	"encoding/json"
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
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		//handle json error
		return
	}
	fmt.Println(user)
	fmt.Println(string(bytes))
	fmt.Println(err)
	c.String(http.StatusNotImplemented, "implement me!")
}

// SearchUser -
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
