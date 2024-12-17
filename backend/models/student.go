package models

// Student 数据模型
type Student struct {
	Name string `json:"name" bson:"name" binding:"required"`
	Age  int    `json:"age" bson:"age" binding:"required"`
}
