package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.com/ast3am77/test-go/internal/models"
	"net/http"
	"time"
)

func (h *Handler) CreateFilial(c *gin.Context) {
	temp := models.FilialDTO{}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = h.services.MakeNewFilial(ctx, &temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}
	c.String(http.StatusCreated, "Filial %s has been created", temp.FilialName)
}

func (h *Handler) EditFilial(c *gin.Context) {
	temp := models.FilialDTO{}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = h.services.EditFilial(ctx, &temp)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}

	c.String(http.StatusOK, "Filial %s has been edited", temp.FilialName)
}

func (h *Handler) GetInfoFilial(c *gin.Context) {
	filialId := c.Param("filial_id")
	userId := c.Param("user_id")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := h.services.GetInfoFilial(ctx, filialId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}

	c.JSON(http.StatusOK,
		gin.H{"filial_name": result.FilialName,
			"country":         result.Country,
			"city":            result.City,
			"filial_type":     result.FilialType,
			"phone_list":      result.PhoneList,
			"email_list":      result.EmailList,
			"photo_id_list":   result.PhotoIDList,
			"organization_id": result.OrganizationID,
			"user_id":         userId,
		})
}

func (h *Handler) AddWorkerFilial(c *gin.Context) {
	temp := models.AddWorkersDTO{}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = h.services.AddWorkerFilial(ctx, &temp, c.Request.Host)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}

	c.String(http.StatusCreated, "New worker has been created, you must confirm the invitation using the link")
}
