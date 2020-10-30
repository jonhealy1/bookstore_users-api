package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonhealy1/bookstore_/bookstore_users-api/domain/users"
	"github.com/jonhealy1/bookstore_/bookstore_users-api/services"
)

// GetUser -
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// CreateUser -
func CreateUser(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	//handle json error
	// 	return
	// }
	// replacing all of the lines of code above
	if err := c.ShouldBindJSON(&user); err != nil {
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//return error
		return
	}
	fmt.Println(user)
	fmt.Println(err)
	c.JSON(http.StatusCreated, result)
}

// SearchUser -
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
