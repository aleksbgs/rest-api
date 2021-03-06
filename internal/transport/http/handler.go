package http

import (
	"fmt"
	"github.com/aleksbgs/rest-api/internal/comment"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/comment/{id}", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		fmt.Fprintf(w, "I am alive!")
	})
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		fmt.Fprintf(w, "Unable to pars UINT from ID")
	}

	comment, err := h.Service.GetComment(uint(i))

	if err != nil {
		fmt.Fprintf(w, "Error Retrieving Comment By ID")
	}

	fmt.Fprintf(w, "%+v", comment)

}
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()

	if err != nil {
		fmt.Fprintf(w, "Failed to retreive all comments")
	}
	fmt.Fprintf(w, "%+v", comments)
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to post new comment")
	}
	fmt.Fprintf(w, "%+v", comment)

}
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/new",
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to update comment")
	}
	fmt.Fprintf(w, "%+v", comment)
}
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 1, 64)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse uint")
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		fmt.Fprintf(w, "Failed to delete comment by comment ID")
	}

	fmt.Fprintf(w, "Successfully deleted comment")

}
