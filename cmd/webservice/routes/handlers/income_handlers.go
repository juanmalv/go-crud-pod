package handlers

import (
	"github.com/juanmalv/go-crud-pod/internal/income"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"

	"github.com/juanmalv/go-crud-pod/internal/platform/database"

	"github.com/gin-gonic/gin"
)

var incomeCollection = database.Client.Database(database.NameDBMongo).Collection(database.IncomeCollection)

//NewIncome creates a new income id with details
func NewIncome(c *gin.Context) {
	var req income.Income

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request is invalid. Please check the data and try again"})
		return
	}

	err := req.ValidateIncomeRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	req.Date = time.Now().Format(time.RFC1123)

	insertResult, err := incomeCollection.InsertOne(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error while trying to create Income"})
		return
	}

	log.Println("Income saved with ID: ", insertResult.InsertedID)

	c.JSON(http.StatusOK, gin.H{"Income ID": insertResult.InsertedID, "message": "Income saved successfully!", "created_at": req.Date})
}

//GetIncome endpoint updates an existing income
func GetIncome(c *gin.Context) {
	var req []*income.Income

	filter := bson.D{{
		"amount",
		bson.M{
			"$gt" :0,
		},
	}}
	getResult, err := incomeCollection.Find(c,filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for getResult.Next(c) {
		var income income.Income
		err := getResult.Decode(&income)

		if err != nil {
			log.Fatal(err)
		}

		req = append(req, &income)
	}

	c.JSON(http.StatusOK, gin.H{"message": req})
}

//UpdateIncome endpoint updates an existing income
func UpdateIncome(c *gin.Context) {
	var req income.Income

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request is invalid. Please check the data and try again"})
		return
	}

	filter := bson.D{{
		"amount",
		bson.M{
			"$gt" :0,
		},
	}}

	update := bson.D{
		{"$set", bson.D{
			{"amount", req.Amount},
		}},
	}

	_, err := incomeCollection.UpdateOne(c,filter,update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error while trying to Update Income"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Income updated successfully!"})
}

//UpdateIncome endpoint updates an existing income
func DeleteIncome(c *gin.Context) {
	var req income.Income

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request is invalid. Please check the data and try again"})
		return
	}

	filter := bson.D{{
		"amount",
		bson.M{
			"$eq" :req.Amount,
		},
	}}

	deleteOne, err := incomeCollection.DeleteOne(c,filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error while trying to delete Income"})
		return
	}

	if deleteOne.DeletedCount == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Couldn't find the income requested :("})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Income deleted successfully!"})
}