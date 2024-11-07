package controllers

import (
	"backend/resource"
	"backend/services"
	validator "backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type reportCon struct {
}

var ReportCon reportCon

func (ct *reportCon) GetReport(c *gin.Context) {
	var req resource.ReportReq
	var err error
	if err = c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, fmt.Errorf("You have to input Company or Patent ID. ").Error())
		return

	} else if _, err := validator.Request.ToValidate(req); err != nil {
		c.JSON(http.StatusOK, err.Message)
	} else if val, err := services.ReportServ.Analyze(&req); err != nil {
		c.JSON(http.StatusOK, err.Error())
	} else {
		c.JSON(http.StatusOK, val)
	}
	return
}

func (ct *reportCon) Search(c *gin.Context) {
	var req resource.SearchReq
	var err error
	if err = c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, fmt.Errorf("You have to input Company or Patent ID. ").Error())
		return

	} else if _, err := validator.Request.ToValidate(req); err != nil {
		c.JSON(http.StatusOK, err.Message)
	} else {
		val := services.ReportServ.SearchKeyword(req)
		c.JSON(http.StatusOK, val)
	}
	return
}
