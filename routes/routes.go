package routes

import (
	"Hands_On/auth"
	"Hands_On/database"
	"Hands_On/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterStudent(c *gin.Context) {
	var user model.Student
	err := c.Bind(&user)
	if err != nil {
		c.JSON(400, "inavalid data")
	}
	err = database.CreateStudent(user)
	if err != nil {
		c.JSON(401, "user not created")
	} else {
		c.JSON(200, "user created")
	}
}

func GetBooks(c *gin.Context) {
	user := database.GetBooks()
	c.JSON(200, user)
}

func UpdateBook(c *gin.Context) {
	var update model.Books
	err := c.Bind(&update)
	if err != nil {
		c.JSON(400, "invalid datatype")
	}
	err = database.UpdateBook(update)
	if err != nil {
		c.String(200, "not updated")
	} else {
		c.String(200, "user updated")
	}
}
func Login(c *gin.Context) {
	var user model.Librarian
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong Format")
	}
	var dataUser model.Librarian
	dataUser, err = database.FindLibrarian(user.Librarian_No)
	if err != nil {
		c.String(200, "user Not Found")
	} else if dataUser.Password == user.Password {
		//TODO generate token
		var token *model.TokenDetails
		token, err = auth.CreateToken(user.Librarian_No)
		model.SaveToken[token.AccessToken] = user.Librarian_No
		c.Header("auth", token.AccessToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, "error generating Token")
		}
	} else {

		c.String(200, "Password invalid")
	}
}
func StudentLogin(c *gin.Context) {
	var user model.Student
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Wrong Format")
	}
	var dataUser model.Student
	dataUser, err = database.FindStudent(user.Student_Name)
	if err != nil {
		c.String(200, "user Not Found")
	} else if dataUser.Password == user.Password {
		//TODO generate token
		var token *model.TokenDetails
		token, err = auth.CreateToken(user.Student_Name)
		model.SaveToken[token.AccessToken] = user.Student_Name
		c.Header("auth", token.AccessToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, "error generating Token")
		}
	} else {

		c.String(200, "Password invalid")
	}
}

// func Delete(c *gin.Context) {
// 	var user model.User
// 	err := c.Bind(&user)
// 	log.Println(user)
// 	database.CheckError(err)
// 	database.DeleteUser(user.Name)
// }

func AddBook(c *gin.Context) {
	var user model.Books
	err := c.Bind(&user)
	if err != nil {
		c.JSON(400, "inavalid data")
	}
	err = database.AddBook(user)
	if err != nil {
		c.JSON(401, "book not added")
	} else {
		c.JSON(200, "book added")
	}
}

func Borrow(c *gin.Context) {
	var update model.Borrow
	err := c.Bind(&update)
	if err != nil {
		c.JSON(400, "invalid datatype")
	}
	err = database.Borrow(update)
	if err != nil {
		c.String(200, "not updated")
	} else {
		c.String(200, "user updated")
	}
}
