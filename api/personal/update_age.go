package personal

import (
	"errors"
	"net/http"
	"strconv"
	"yaroslavl-parkings/api"
)

func (d *personalDependencies) updateUserAge(w http.ResponseWriter, r *http.Request) error {
	if !api.IsAuth(d.sessions, r) {
		return errors.New("not authenticated")
	}

	session, valid := api.GetSessionIfValid(d.sessions, r)
	if !valid {
		return errors.New("not authenticated")
	}

	newAge := r.PostFormValue("newAge")
	if newAge == "" {
		return errors.New("no age was provided")
	}

	newAgeInt, err := strconv.Atoi(newAge)
	if err != nil {
		return errors.New("the age must be a number")
	}

	err = d.db.UpdateUserAge(session.UserID, uint(newAgeInt))
	return err
}
