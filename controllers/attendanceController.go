package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hndev2404/interview_beearning/dto"
	"github.com/hndev2404/interview_beearning/models"
	"github.com/hndev2404/interview_beearning/response"
	"github.com/hndev2404/interview_beearning/services"
)

func AttendanceOwner(c *gin.Context) {
	result, err := services.GetOwnerAddress()
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"owner": result})

}

func AttendanceList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AttendanceList"})
}

func AttendanceDetail(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "AttendanceDetail"})
}

func AttendanceCheckIn(c *gin.Context) {
	// Validation
	var body dto.CheckInDTO
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		response.ResponseError(c, err)
		return
	}

	// Get current Employee ID
	user, exists := c.Get("user")
	if !exists {
		response.ResponseError(c, errors.New("must authentication before authorization"))
		return
	}

	data, err := services.AttendanceCheckIn(user.(models.User).ID, &body)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseSucceed(c, data)
}

func AttendanceCheckout(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "AttendanceCheckout"})
}