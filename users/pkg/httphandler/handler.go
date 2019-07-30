package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/rahul-golang/ecommerce/log"

	"github.com/rahul-golang/ecommerce/users/pkg/models"

	"github.com/rahul-golang/ecommerce/users/pkg/service"

	"github.com/gorilla/mux"
)

type UserHttpHandlers struct {
	userService service.UsersService
}

func NewUserRestHandler(userService service.UsersService) *UserHttpHandlers {
	return &UserHttpHandlers{userService: userService}
}

//CreateUser d
func (userHttpHandlers UserHttpHandlers) CreateUser(w http.ResponseWriter, req *http.Request) {

	createUserReq := models.CreateUserReq{}
	createUserResp := &models.CreateUserResp{}
	err := json.NewDecoder(req.Body).Decode(&createUserReq)
	if err != nil {
		createUserResp.Message = err.Error()
		createUserResp.User = nil
		json.NewEncoder(w).Encode(createUserResp)
		writeResponse(w, http.StatusInternalServerError)
		return
	}
	createUserResp, err = userHttpHandlers.userService.CreateUser(req.Context(), createUserReq)

	endpointResp, err := json.Marshal(createUserResp)
	if err != nil {
		createUserResp.Message = err.Error()
		createUserResp.User = nil
		json.NewEncoder(w).Encode(createUserResp)
		writeResponse(w, http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(endpointResp)
}
func (userHttpHandlers UserHttpHandlers) GetUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	getUserResp := &models.GetUserResp{}

	resp, err := userHttpHandlers.userService.GetUser(req.Context(), id)

	if err != nil {
		getUserResp.Message = err.Error()
		getUserResp.User = nil
		json.NewEncoder(w).Encode(getUserResp)
		writeResponse(w, http.StatusInternalServerError)
		return
	}
	getUserResp.Message = "Sucess"
	getUserResp.User = resp
	//log.Println(getUserResp)
	err = json.NewEncoder(w).Encode(getUserResp)
	if err != nil {
		getUserResp.Message = err.Error()
		getUserResp.User = nil
		json.NewEncoder(w).Encode(getUserResp)
		writeResponse(w, http.StatusInternalServerError)
		return
	}
	writeResponse(w, http.StatusOK)

}
func (userHttpHandlers UserHttpHandlers) UpdateUser(w http.ResponseWriter, req *http.Request) {
	updateUserReq := models.UpdateUserReq{}
	updateUserResp := models.UpdateUserResp{}
	err := json.NewDecoder(req.Body).Decode(&updateUserReq)
	if err != nil {
		updateUserResp.Message = err.Error()
		updateUserResp.User = nil
		json.NewEncoder(w).Encode(updateUserResp)
		writeResponse(w, http.StatusInternalServerError)
	}
	user, err := userHttpHandlers.userService.UpdateUser(req.Context(), updateUserReq.User)
	if err != nil {
		updateUserResp.Message = err.Error()
		updateUserResp.User = nil
		json.NewEncoder(w).Encode(updateUserResp)
		writeResponse(w, http.StatusInternalServerError)
	}
	updateUserResp.User = user
	endpointResp, err := json.Marshal(updateUserResp)
	if err != nil {
		updateUserResp.Message = err.Error()
		updateUserResp.User = nil
		json.NewEncoder(w).Encode(updateUserResp)
		writeResponse(w, http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(endpointResp)

}
func (userHttpHandlers UserHttpHandlers) DeleteUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	deleteUserResp := &models.DeleteUserResp{}

	deleteUserResp, err := userHttpHandlers.userService.DeleteUser(req.Context(), id)
	if err != nil {
		json.NewEncoder(w).Encode(deleteUserResp)
		writeResponse(w, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(deleteUserResp)
	if err != nil {
		deleteUserResp.Message = err.Error()
		deleteUserResp.ID = ""
		json.NewEncoder(w).Encode(deleteUserResp)
		writeResponse(w, http.StatusInternalServerError)
		return
	}
	writeResponse(w, http.StatusOK)

}
func (userHttpHandlers UserHttpHandlers) GetAllUser(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	getAllUserResp := &models.GetAllUserResp{}

	users, err := userHttpHandlers.userService.GetAllUser(ctx)
	if err != nil {
		getAllUserResp.Message = err.Error()
		getAllUserResp.User = users
		json.NewEncoder(w).Encode(getAllUserResp)
		writeResponse(w, http.StatusInternalServerError)
		return
	}
	getAllUserResp.User = users
	log.Logger(ctx).Info("sucess")
	err = json.NewEncoder(w).Encode(getAllUserResp)
	if err != nil {
		getAllUserResp.Message = err.Error()
		getAllUserResp.User = users
		json.NewEncoder(w).Encode(getAllUserResp)
		writeResponse(w, http.StatusInternalServerError)
		return
	}
	writeResponse(w, http.StatusOK)
}

func writeResponse(w http.ResponseWriter, errorCode int) {
	w.WriteHeader(errorCode)
}
