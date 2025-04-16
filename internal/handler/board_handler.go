package handler

import (
	"net/http"
	"strconv"
	"task-manager/internal/model"
	"task-manager/internal/usecase"

	"github.com/gin-gonic/gin"
)

type BoardHandler struct {
	boardUsecase usecase.BoardRepository
}

func NewBoardHandler(uc usecase.BoardRepository) *BoardHandler {
	return &BoardHandler{boardUsecase: uc}
}

func (h *BoardHandler) GetAllBoards(c *gin.Context) {
	boards, err := h.boardUsecase.GetAllBoards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, boards)
}

func (h *BoardHandler) GetBoard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid board ID"})
		return
	}

	board, err := h.boardUsecase.GetBoardByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, board)
}

func (h *BoardHandler) CreateBoard(c *gin.Context) {
	var board model.Board
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.boardUsecase.CreateBoard(&board); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, board)
}

func (h *BoardHandler) UpdateBoard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid board ID"})
		return
	}

	var board model.Board
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	board.ID = uint(id)

	if err := h.boardUsecase.UpdateBoard(&board); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, board)
}

func (h *BoardHandler) DeleteBoard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid board ID"})
		return
	}

	if err := h.boardUsecase.DeleteBoard(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *BoardHandler) BlockBoard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid board ID"})
		return
	}

	adminID, err := strconv.Atoi(c.GetHeader("X-Admin-ID"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "admin ID required"})
		return
	}

	if err := h.boardUsecase.BlockBoard(uint(id), uint(adminID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
