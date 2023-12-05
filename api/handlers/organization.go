package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.com/ast3am77/test-go/internal/models"
	"net/http"
	"time"
)

func (h *Handler) CreateOrganization(c *gin.Context) {
	temp := models.Organization{}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = h.services.MakeNewOrganization(ctx, &temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}
	c.String(http.StatusCreated, "Organization %s has been created", temp.Name)
}

func (h *Handler) EditOrganization(c *gin.Context) {
	temp := models.Organization{}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = h.services.EditOrganization(ctx, &temp)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}

	c.String(http.StatusCreated, "Organization %s has been edited", temp.Name)
}

func (h *Handler) GetInfoOrganization(c *gin.Context) {
	orgId := c.Param("org_id")
	userId := c.Param("user_id")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := h.services.GetInfoOrganization(ctx, orgId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}

	c.JSON(http.StatusOK,
		gin.H{"name": result.Name,
			"legal_type":    result.LegalType,
			"legal_address": result.LegalAddress,
			"inn":           result.INN,
		})
}

func (h *Handler) GetOrganizationFilials(c *gin.Context) {
	orgId := c.Param("org_id")
	userId := c.Param("user_id")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := h.services.GetOrganizationsFilials(ctx, orgId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"filials": result})
}
