package role

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (h *Handler) GetRoleByID(c *gin.Context) {
	roleId := c.Query("roleId")
	id, _ := strconv.Atoi(roleId)

	role, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": role,
	})
}

func (h *Handler) Store(c *gin.Context) {
	var createRole CreateRole
	err := c.ShouldBindJSON(&createRole)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	role, err := h.service.Store(createRole)

	c.JSON(http.StatusOK, gin.H{
		"data": role,
	})

}
