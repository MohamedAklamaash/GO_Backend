package user

import (
	"log"
	"net/http"

	"github.com/MohamedAklamaash/GO_Simple_CRUD_Backend_WIth_SQL/services/auth"
	"github.com/MohamedAklamaash/GO_Simple_CRUD_Backend_WIth_SQL/types"
	"github.com/MohamedAklamaash/GO_Simple_CRUD_Backend_WIth_SQL/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(routes *mux.Router) {
	routes.HandleFunc("GET /login", h.handleLogin)
	routes.HandleFunc("POST /register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get json payload
	var payload types.UserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, 404, err)
	}
	// check if the user exists already? not! we create the new user
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, 401, err)
		return
	}
	hashedPassword, _ := auth.HashPassword(payload.Password)
	err = h.store.CreateUser(
		types.User{
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Email:     payload.Email,
			Password:  hashedPassword,
		})
	if err != nil {
		log.Fatal(err)
	}
}
