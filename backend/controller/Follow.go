package controller

import (
	"net/http"
	"social_network/db"
	"social_network/helper"
	"social_network/models"
	"strconv"
)

func UnFollow(w http.ResponseWriter, r *http.Request) {

	_, _, user_id := helper.Auth(db.DB, r)

	user := models.User{ID: user_id}
	listusers, err := user.GetUnFollow(db.DB, 100)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}
	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "listusers": listusers}, nil)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}
}

func Follow(w http.ResponseWriter, r *http.Request) {

	follow_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}
	_, _, user_id := helper.Auth(db.DB, r)

	follower := models.User{}
	current_user := models.User{}
	errfollow := follower.GetUserById(db.DB, follow_id)
	erruser := current_user.GetUserById(db.DB, user_id)

	if errfollow != nil || erruser != nil {
		helper.ErrorPage(w, 400)
		return
	}

	if follower.IsPublic == 1 {
		errreq := follower.AddFlow(db.DB, user_id)
		if errreq != nil {
			helper.ErrorPage(w, 500)
			return
		}
	}
	err = helper.WriteJSON(w, http.StatusOK, map[string]interface{}{"success": true, "user_id": user_id}, nil)
	if err != nil {
		helper.ErrorPage(w, 400)
		return
	}
}

func AcceptFriend(w http.ResponseWriter, r *http.Request) {

}