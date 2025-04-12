package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"task-manager/internal/model"
	"task-manager/internal/usecase"
)

type BoardHandler struct {
	uc usecase.BoardUsecase
}

func NewBoardHandler(uc usecase.BoardUsecase) *BoardHandler {
	return &BoardHandler{uc}
}

func (h *BoardHandler) GetAllBoards(w http.ResponseWriter, r *http.Request) {
	boards, err := h.uc.GetAllBoards()
	if err != nil {
		http.Error(w, "Failed to fetch boards", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(boards)
}

func (h *BoardHandler) GetBoardByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid board ID", http.StatusBadRequest)
		return
	}
	board, err := h.uc.GetBoardByID(id)
	if err != nil {
		http.Error(w, "Board not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(board)
}

func (h *BoardHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {
	var board model.Board
	if err := json.NewDecoder(r.Body).Decode(&board); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := h.uc.CreateBoard(&board)
	if err != nil {
		http.Error(w, "Failed to create board", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(board)
}

func (h *BoardHandler) UpdateBoard(w http.ResponseWriter, r *http.Request) {
	var board model.Board
	if err := json.NewDecoder(r.Body).Decode(&board); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	h.uc.UpdateBoard(&board)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Board updated"})
}

func (h *BoardHandler) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid board ID", http.StatusBadRequest)
		return
	}
	h.uc.DeleteBoard(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Board deleted"})
}
