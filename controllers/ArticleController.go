package controllers

import (
	"fmt"
	"go-api/connection"
	"go-api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetAllArticles(c *gin.Context) {
	articles := []models.Article{}
	err := connection.DB.Find(&articles).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Internal server error"})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Get all articles successfully", "data": articles})
}

func PaginationPostedArticle(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, _ :=  strconv.Atoi(c.Param("offset"))

	articles := []models.Article{}
	err := connection.DB.Find(&articles).Limit(limit).Offset(offset).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "errors": "Failed delete   article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Get article page successfully",  "data": articles})
}

func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	article := models.Article{}
	err := connection.DB.Where("id = ?", id).Find(&article).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message" : "Internal server error"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Get article by id successfully", "data" : article})
}

func CreateNewArticle(c *gin.Context) {

	article := models.Article{
		Created_date: time.Now(),
		Updated_date: time.Now(),
	}
	
	if err := c.ShouldBindJSON(&article); err != nil {

	errorMessages :=  []string{}
	
	for _, e :=  range err.(validator.ValidationErrors) {
		errorMessage := fmt.Sprintf("Error on Filled %s, condition: %s", e.Field(), e.ActualTag())
		errorMessages = append(errorMessages,  errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad request", "errors": errorMessages})
		return
	}

	if errDB := connection.DB.Create(&article).Error; errDB != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "errors": "Failed create   article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Create article  successfully",  "data" : article})
}

func UpdateArticle(c *gin.Context) {

	id := c.Param("id")

	article := models.Article{
		Updated_date: time.Now(),
	}

	if err := c.ShouldBindJSON(&article); err != nil {
	errorMessages :=  []string{}
	for _, e :=  range err.(validator.ValidationErrors) {
		errorMessage := fmt.Sprintf("Error on Filled %s, condition: %s", e.Field(), e.ActualTag())
		errorMessages = append(errorMessages,  errorMessage)
		}
		c.JSON(http.StatusOK, gin.H{"status": "bad request", "errors": errorMessages})
		return
	}

	if errDB := connection.DB.Model(&models.Article{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title": article.Title,
		"content": article.Content,
		"category": article.Category,
		"updated_date": article.Updated_date,
		"status": article.Status,
	}).Error; errDB != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "errors": "Failed update   article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Create article  successfully",  "data" : article})
}

func DeleteArticleById(c *gin.Context) {
	id := c.Param("id")
	err := connection.DB.Delete(&models.Article{}, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "errors": "Failed delete   article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Article deleted successfully"})
}