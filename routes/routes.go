package routes

import (
	"net/http"
	"task-api/db"
	"task-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("gfhdesdjhgkjhgijfrdlkhyj") // You can replace this with an environment variable

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := db.GetDB()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user models.User

	// Bind the incoming JSON to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request"})
		return
	}

	database := db.GetDB()

	// Query the database to find a user by username
	var founduser models.User
	result := database.Where("username=?", user.Username).First(&founduser)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		return
	}

	// Compare the provided password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(founduser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid credentials"})
		return
	}

	// If the passwords match, generate JWT token
	claims := &jwt.RegisteredClaims{
		Subject:   founduser.Username,                                 // Set the username as the subject of the token
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
	}

	// Generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not create token"})
		return
	}

	// Send the token back in the response
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})
}

func CreateTask(c *gin.Context) {
	username := c.GetString("username")
	db := db.GetDB()

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	task.Username = username
	db.Create(&task)
	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	username := c.GetString("username")
	db := db.GetDB()

	var tasks []models.Task
	db.Where("username = ?", username).Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	username := c.GetString("username")
	db := db.GetDB()
	id := c.Param("id")

	var task models.Task
	result := db.Where("id = ? AND username = ?", id, username).First(&task)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	username := c.GetString("username")
	db := db.GetDB()
	id := c.Param("id")

	var task models.Task
	result := db.Where("id = ? AND username = ?", id, username).First(&task)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	db.Save(&task)
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	username := c.GetString("username")
	db := db.GetDB()
	id := c.Param("id")

	var task models.Task
	result := db.Where("id = ? AND username = ?", id, username).Delete(&task)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or you don't have permission"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
