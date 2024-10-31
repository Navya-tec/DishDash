package routes

import (
	"database/sql"
	"dish-dash/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	r.POST("/create-flat", func(c *gin.Context) { CreateFlat(c, db) })
	r.POST("/create-flatmate", func(c *gin.Context) { CreateFlatmate(c, db) })
	r.POST("/join-flat", func(c *gin.Context) { JoinFlat(c, db) })
}

func CreateFlat(context *gin.Context,db *sql.DB) {

	var flat models.Flat
	err := context.ShouldBindJSON(&flat)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request data: CreateFlat"})
		return
	}

	flatId,err := models.CreateFlat(db,flat)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error in saving flat in Database"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Flat created successfully","flatId: ":flatId})
}

func CreateFlatmate(context *gin.Context,db *sql.DB) {

	var flatmate models.Flatmate
	err := context.ShouldBindJSON(&flatmate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request data: Createflatmate"})
		return
	}

	err = models.AddFlatmateToFlat(db,flatmate)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error in adding flatmate to flat"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Flatmate created successfully"})
}

func JoinFlat(context *gin.Context,db *sql.DB) {
    var flatmate models.Flatmate

	err:=context.ShouldBindJSON(&flatmate)
	if err!=nil{
		context.JSON(http.StatusBadRequest,gin.H{"message":"Couldn't parse data: Join Flat"})
		return
	}

	joined, err := models.CheckIfUserJoined(db, flatmate.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error in checking user flat status"})
		return
	}
	if joined {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User already joined a flat"})
		return
	}

	err=models.AddFlatmateToFlat(db,flatmate)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Error in adding flatmate to flat"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Flatmate joined the flat"})
}
