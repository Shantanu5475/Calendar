package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var HolidayCollection *mongo.Collection

func InitHolidayCollection(collection *mongo.Collection) {
	HolidayCollection = collection
}

func ListHolidays(c *gin.Context) {
	cursor, err := HolidayCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching holidays"})
		return
	}

	var holidays []bson.M
	if err = cursor.All(context.TODO(), &holidays); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding holidays"})
		return
	}

	for i := range holidays {
		holidays[i]["id"] = holidays[i]["_id"]
	}

	c.JSON(http.StatusOK, holidays)
}

func AddHoliday(c *gin.Context) {
	var holiday struct {
		Name    string `json:"name"`
		Date    string `json:"date"`
		Country string `json:"country"`
	}
	if err := c.BindJSON(&holiday); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	result, err := HolidayCollection.InsertOne(context.TODO(), holiday)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add holiday"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Holiday added successfully", "id": result.InsertedID})
}

func DeleteHoliday(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	filter := bson.M{"_id": objectID}

	result, err := HolidayCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete holiday"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No holiday found with given ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Holiday deleted successfully"})
}
