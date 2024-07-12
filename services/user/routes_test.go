package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MohamedAklamaash/GO_Simple_CRUD_Backend_WIth_SQL/types"
	"github.com/gorilla/mux"
)

func TestUserService(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if user payload is invalid", func(t *testing.T) {
		payload := types.UserPayload{
			FirstName: "aklamaash",
			LastName:  "ehsan",
			Email:     "akla123@",
			Password:  "akla",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.ServeHTTP(rr, req)
		router.HandleFunc("/register", handler.handleLogin)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d,got %d", http.StatusBadRequest, rr.Code)
		}
	})

}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserById(id string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}
