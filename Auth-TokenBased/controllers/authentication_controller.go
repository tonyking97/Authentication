package controllers

import (
	"../core"
	"../models"
	"../validation/govalidator"
	"encoding/json"
	//"github.com/asaskevich/govalidator"
	"log"
	"net/http"
)

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Contains methods to handle Authentication Functions and returns the response
 */

//Handles Login of a User
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
_:
	decoder.Decode(&requestUser)

	//Validating params. Refer models/users.go
	if _, err := govalidator.ValidateStruct(requestUser); err != nil {
		res := &models.ErrorMessage{err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(res)
		models.CheckError(err, w)
	} else {
		responseStatus, token := core.Login(requestUser, r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		_, err := w.Write(token)
		models.CheckError(err, w)
	}
}

//Check Username is Present in database - used for login
func CheckUsername(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.Username)
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&requestUser)

	if _, err := govalidator.ValidateStruct(requestUser); err != nil {
		res := &models.ErrorMessage{err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(res)
		models.CheckError(err, w)
	} else {
		if ack, username, avatar := core.CheckUsername(requestUser); !ack {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			res := &models.ErrorMessage{"No such Username exist"}
			err := json.NewEncoder(w).Encode(res)
			models.CheckError(err, w)
		} else {
			w.Header().Set("Content-Type", "application/json")
			res := &models.CheckUsername{Username: username, Avatar: avatar}
			err := json.NewEncoder(w).Encode(res)
			models.CheckError(err, w)
		}
	}
}

//Get Username and Avatar
func GetNameDetails(w http.ResponseWriter, r *http.Request) {

	requestUser := new(models.TokenAuthentication)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	if err != nil {
		log.Println(err)
	}

	ff_id, _, err := core.CheckToken(requestUser)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		res := &models.ErrorMessage{err.Error()}
		error := json.NewEncoder(w).Encode(res)
		models.CheckError(error, w)
	} else {

		if ack, result := core.GetNameDetails(&models.Ff_id_user{Ff_id: ff_id}); !ack {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			res := &models.ErrorMessage{"No such Username exist"}
			err := json.NewEncoder(w).Encode(res)
			models.CheckError(err, w)
		} else {
			w.Header().Set("Content-Type", "application/json")
			//res := &models.UserNameDetails{Username:username, Firstname:firstname, Lastname:lastname, Avatar:avatar}
			err := json.NewEncoder(w).Encode(result)
			models.CheckError(err, w)
		}
	}
}

//Check Username is Present in database - used for signup
func CheckUsernameAvailability(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.Username)
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&requestUser)

	if _, err := govalidator.ValidateStruct(requestUser); err != nil {
		res := &models.ErrorMessage{err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(res)
		models.CheckError(err, w)
	} else {
		if ack, _, _ := core.CheckUsername(requestUser); !ack {
			w.Header().Set("Content-Type", "application/json")
			res := &models.SuccessMessage{"Username Available"}
			err := json.NewEncoder(w).Encode(res)
			models.CheckError(err, w)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			res := &models.ErrorMessage{"Username already exist. Try another username."}
			err := json.NewEncoder(w).Encode(res)
			models.CheckError(err, w)
		}
	}
}

//Check email is Present in database - used for signup
func CheckEmailAvailability(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.Email)
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&requestUser)

	if _, err := govalidator.ValidateStruct(requestUser); err != nil {
		res := &models.ErrorMessage{err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(res)
		models.CheckError(err, w)
	} else {
		if ack := core.CheckEmail(requestUser); !ack {
			w.Header().Set("Content-Type", "application/json")
			res := &models.SuccessMessage{""}
			err := json.NewEncoder(w).Encode(res)
			models.CheckError(err, w)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			res := &models.ErrorMessage{"EmailID already exists"}
			err := json.NewEncoder(w).Encode(res)
			models.CheckError(err, w)
		}
	}
}

//Get Active session details
func GetActiveSessionDetails(w http.ResponseWriter, r *http.Request) {
	ffId := "7f338308-c0fd-44e1-589f-e61a726f91ab"
	results := core.GetActiveSessionList(&models.Ff_id{Ff_id: ffId})
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(results)
	models.CheckError(err, w)
}

//Get inactive session details
func GetInActiveSessionDetails(w http.ResponseWriter, r *http.Request) {
	ffId := "7f338308-c0fd-44e1-589f-e61a726f91ab"
	results := core.GetInActiveSessionList(&models.Ff_id{Ff_id: ffId})
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(results)
	models.CheckError(err, w)
}

//Get current session details
func GetSessionDetails(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.TokenAuthentication)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	if err != nil {
		log.Println(err)
	}

	ff_id, _, err := core.CheckToken(requestUser)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		res := &models.ErrorMessage{err.Error()}
		error := json.NewEncoder(w).Encode(res)
		models.CheckError(error, w)
	} else {
		results := core.GetSessionList(&models.Ff_id{Ff_id: ff_id})

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(results)
		models.CheckError(err, w)
	}
}

//Get current session details
func GetCurrentSessionDetails(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.TokenAuthentication)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	if err != nil {
		log.Println(err)
	}

	ff_id, fs_id, err := core.CheckToken(requestUser)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		res := &models.ErrorMessage{err.Error()}
		error := json.NewEncoder(w).Encode(res)
		models.CheckError(error, w)
	} else {
		result := core.GetCurrentSessionList(&models.Fs_id{Fs_id: fs_id})
		tokenCount := core.GetActiveTokensCount(&models.Ff_id{Ff_id: ff_id})
		sessionCount := core.GetActiveTokensCount(&models.Ff_id{Ff_id: ff_id})

		results := &models.CurrentSessionDetail_limited{
			Active_Sessions: sessionCount,
			Active_Tokens:   tokenCount,
			Ip:              result.Ip_Details.Ip,
			City:            result.Ip_Details.City,
			Latitude:        result.Ip_Details.Latitude,
			Longitude:       result.Ip_Details.Longitude,
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(results)
		models.CheckError(err, w)
	}
}

//Refresh token -> Generate new token with old one
func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.TokenAuthentication)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(core.RefreshToken(requestUser))
	models.CheckError(err, w)
}

//Handles Sign UP of a new user
/*TODO make govalidator to returns error as type string instead of error, because json encoder can't encode type error.
TODO Solution: create a custom marshaller logic to convert error to string
*/
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.SignUpDetails)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	models.CheckError(err, w)
	//Validating params. Refer models/signupUsers.go
	if _, err := govalidator.CustomValidateStruct(requestUser); err != nil {
		log.Println(err)
		//errs := err.(govalidator.Errors)
		//log.Print(errs)
		//res := &models.ErrorMessage{err.Error()}
		//for _, e = range errs {
		//	log.Println(e.Name)
		//}
		//log.Println(arrErr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)

	} else {

		if !core.Signup(requestUser) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := &models.ErrorMessage{"Account creation failed"}
			err := json.NewEncoder(w).Encode(res)
			models.CheckError(err, w)
		} else {
			w.Header().Set("Content-Type", "application/json")
			res := &models.SuccessMessage{"Account created successfully"}
			err := json.NewEncoder(w).Encode(res)
			models.CheckError(err, w)
		}
	}
}

//Check token is valid or not
func CheckTokenHandler(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.TokenAuthentication)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	if err != nil {
		log.Println(err)
	}

	_, _, err = core.CheckToken(requestUser)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		res := &models.ErrorMessage{err.Error()}
		error := json.NewEncoder(w).Encode(res)
		models.CheckError(error, w)
	} else {
		w.Header().Set("Content-Type", "application/json")
		res := &models.SuccessMessage{"token valid"}
		err := json.NewEncoder(w).Encode(res)
		models.CheckError(err, w)
	}
}

//Handles Logout function of the user
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.TokenAuthentication)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	if err != nil {
		log.Println(err)
	}

	err = core.Logout(requestUser)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		res := &models.ErrorMessage{err.Error()}
		err := json.NewEncoder(w).Encode(res)
		models.CheckError(err, w)
	} else {
		w.Header().Set("Content-Type", "application/json")
		res := &models.SuccessMessage{"Logged out successfully"}
		err := json.NewEncoder(w).Encode(res)
		models.CheckError(err, w)
	}
}
