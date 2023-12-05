package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) UpdateConfirmStatus(c *gin.Context) {
	token := c.Param("token")
	userId := c.Param("user_id")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := h.services.UpdateConfirmStatus(ctx, userId, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.log.HandlerErrorLog(c.Request, http.StatusBadRequest, "", err)
		return
	}

	c.String(http.StatusOK, "worker has been confirmed")
}
