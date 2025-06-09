package task

import (
	"IOBound/logging"
	"IOBound/pkg/handlers"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type taskHandler struct {
	service TaskService
	logger *logging.Logger
}

func NewHandler(serv TaskService, logger *logging.Logger) handlers.Handler {
	return &taskHandler{service: serv, logger: logger}
}

func (h *taskHandler) Register(router *mux.Router) {
	router.HandleFunc("/tasks/{id}", h.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks/{id}", h.DeleteTaskByID).Methods("DELETE")
	router.HandleFunc("/tasks", h.CreateTask).Methods("POST")
}

func (h *taskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	task.Status = StatusPending
	task.CreatedAt = time.Now().UTC()

	h.logger.Info("Start a create of task")
	taskID, err := h.service.CreateTask(r.Context(), task)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h.logger.Info("Creation completed")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(taskID))
}

func (h *taskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]
	h.logger.Info("Get task by ID")
	task, err := h.service.GetTaskByID(r.Context(), taskID)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	h.logger.Info("Receipt completed")

	h.logger.Info("Transform task from struct to json")
	jsonTask, err := json.Marshal(task)
	if err != nil {
		h.logger.Error(err)
		http.Error(w, "Could not transform from json to bytes", http.StatusInternalServerError)
		return
	}
	h.logger.Info("Transform complited")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonTask)
}

func (h *taskHandler) DeleteTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	h.logger.Info("Start delete task by ID")
	err := h.service.DeleteTaskByID(r.Context(), taskID)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	h.logger.Info("Delete complited")

	w.WriteHeader(http.StatusOK)

}