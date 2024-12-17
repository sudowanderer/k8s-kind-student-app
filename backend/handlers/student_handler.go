package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"net/http"
	"student-management-system/db"
	"student-management-system/models"
)

// 获取学生列表
func GetStudentsHandler(c *gin.Context) {
	collection := db.GetCollection("students")

	// 查询所有学生
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(c)

	var students []models.Student
	if err := cursor.All(c, &students); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

// 添加学生
func AddStudentHandler(c *gin.Context) {
	collection := db.GetCollection("students")

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 插入学生到数据库
	_, err := collection.InsertOne(c, student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Student added successfully"})
}
