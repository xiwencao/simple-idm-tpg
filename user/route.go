package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
)

type Handle struct {
	userService *UserService
}

func NewHandle(userService *UserService) Handle {
	return Handle{
		userService: userService,
	}
}

// Get a list of users
// (GET /user)
func (h Handle) GetUser(w http.ResponseWriter, r *http.Request) *Response {
	return nil
}

// Create a new user
// (POST /user)
func (h Handle) PostUser(w http.ResponseWriter, r *http.Request) *Response {
	data := PostUserJSONRequestBody{}
	err := render.DecodeJSON(r.Body, &data)
	if err != nil {
		return &Response{
			Code: http.StatusBadRequest,
			body: "unable to parse body",
		}
	}

	slog.Debug("request params:", "data", data)

	params := UserParams{
		Email: *data.Email,
	}
	dbUser, err := h.userService.Create(r.Context(), params)
	if err != nil {
		slog.Error("Failed creating user", params, "err", err)
		return &Response{
			body: "Failed creating user",
			Code: http.StatusInternalServerError,
		}
	}
	return &Response{
		Code: http.StatusCreated,
		body: dbUser,
	}
}

// Delete user by UUID
// (DELETE /user/{uuid})
func (h Handle) DeleteUserUUID(w http.ResponseWriter, r *http.Request, uuid string) *Response {
	return nil
}

// Get user details by UUID
// (GET /user/{uuid})
func (h Handle) GetUserUUID(w http.ResponseWriter, r *http.Request, uuid string) *Response {
	return nil
}

// Update user details by UUID
// (PUT /user/{uuid})
func (h Handle) PutUserUUID(w http.ResponseWriter, r *http.Request, uuid string) *Response {
	return nil
}

func Routes(r *chi.Mux, handle Handle) {

	r.Group(func(r chi.Router) {
		// add auth middleware
		r.Mount("/api/v4", Handler(&handle))
	})
}