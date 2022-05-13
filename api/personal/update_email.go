package personal

import (
	"errors"
	"net/http"
	"yaroslavl-parkings/api"
)

func (d *personalDependencies) updateUserEmail(w http.ResponseWriter, r *http.Request) error {
	if !api.IsAuth(d.sessions, r) {
		return errors.New("not authenticated")
	}

	session, valid := api.GetSessionIfValid(d.sessions, r)
	if !valid {
		return errors.New("not authenticated")
	}

	newEmail := r.PostFormValue("newEmail")
	if newEmail == "" {
		return errors.New("the email provided is invalid")
	}

	err := d.db.UpdateUserEmail(session.UserID, newEmail)
	return err
}
