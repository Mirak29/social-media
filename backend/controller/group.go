package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	helper "social_network/helper"
	"social_network/models"
	"strconv"
	"strings"
)

type NewGroup struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	NotifStatus bool      `json:"notif_type"`
	SuggestList []NewUser `json:"suggests"`
}
type NewUser struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Isrequested bool   `json:"is_requested"`
}

type GroupDetail struct {
	NbrFollowers int            `json:"nbrfollowers"`
	Events       []models.Event `json:"events"`
	Groupdata    models.Group   `json:"groupdata"`
	PostsData    []models.Post  `json:"postsdata"`
}

func CreateFollowGroup(w http.ResponseWriter, r *http.Request) {
	type GroupFollowData struct {
		GroupID int `json:"groupId"`
		UserID  int `json:"userId"`
	}

	var followData GroupFollowData

	err := json.NewDecoder(r.Body).Decode(&followData)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}

	var user = models.User{}
fmt.Println("in createfollow")
	errr := user.GetUserById(DB, followData.UserID)
	if errr != nil {
		fmt.Println(err)
		helper.ErrorPage(w, 500)
		return
	}

	group, errg := models.GetGroupByID(DB, followData.GroupID)
	if errg != nil {
		fmt.Println("t")
		helper.ErrorPage(w, 500)
		return
	}

	var notification = models.Notification{}
	notification.SenderID = user.ID
	notification.User_id = group.UserID
	notification.Type = "join-Group"
	notification.Group_id = group.ID
	notification.Status = "false"
	notification.FirstName = user.FirstName
	notification.LastName = user.LastName
	notification.Avatar = user.Avatar

	notichecker, errn:=models.GetNotificationByUserIDAndType(DB, notification.SenderID, notification.User_id, notification.Type, notification.Group_id)
	if errn != nil {
		fmt.Println(errn)
		helper.ErrorPage(w, 500)
	}

	if len(notichecker) == 0 {
		ern := notification.CreateNotification(DB)
		if ern != nil {
			fmt.Println(ern)
			helper.ErrorPage(w, 500)
		}
		currentGroup, err := models.GetGroupByID(DB,  group.ID)
		if err != nil {
			helper.ErrorPage(w, 500)
			return
		}
		notification.GroupTitle=currentGroup.Title
		SendSocketNotification(notification, "notification")
	}

}

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	_, userEmail, _ := helper.Auth(DB, r)

	var user = models.User{}
	err := user.GetUserByEmail(DB, userEmail)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var newGroup models.Group

	err = json.NewDecoder(r.Body).Decode(&newGroup)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		fmt.Println("in decode")
		return
	}
	if newGroup.Title == "" || newGroup.Description == "" {
		http.Error(w, "Title and description cannot be empty", http.StatusBadRequest)
		return
	}
	newGroup.UserID = user.ID
	lastInsertId, err := newGroup.CreateGroup(DB)
	if err != nil {
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
		fmt.Println(err)
		fmt.Println("in user")
		return
	}
	er := models.JoinGroup(DB, user.ID, int(lastInsertId))
	if er != nil {
		fmt.Println(er)
		return
	}
	for c, _ := range Clients {
		if c.ID == user.ID {
			c.Groups = append(c.Groups, int(lastInsertId))
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Group created successfully",
		"groupId": lastInsertId,
	})
}

func GetGroupDetail(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	_, userEmail, _ := helper.Auth(DB, r)
	err := user.GetUserByEmail(DB, userEmail)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}

	groupId := strings.Split(r.URL.Path, "/")

	grId, errg := strconv.Atoi(groupId[len(groupId)-1])
	if errg != nil {
		helper.ErrorPage(w, 500)
		return
	}
	gr, err := models.GetGroupByID(DB, grId)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}
	ismem, errm := user.IsGroupmemeber(DB, grId)
	if errm != nil || !ismem {
		helper.ErrorPage(w, 500)
		return
	}
	events, erre := models.GetEventsByGroupId(DB, grId)
	if erre != nil {
		fmt.Println(erre)
		helper.ErrorPage(w, 500)
		return
	}
	nbr, errgm := models.GetGroupMemberCount(DB, grId)
	if errgm != nil {
		helper.ErrorPage(w, 500)
		return
	}
	posts, errp := models.GetGroupPost(DB, grId)
	if errp != nil {
		helper.ErrorPage(w, 500)
		return
	}
	fmt.Println("=>>>>>>>")
	fmt.Println(posts)
	groupapagedata := GroupDetail{}
	groupapagedata.NbrFollowers = nbr
	groupapagedata.Events = events
	groupapagedata.Groupdata = *gr
	groupapagedata.PostsData = posts

	helper.WriteJSON(w, 200, map[string]interface{}{"success": true, "groupDetail": groupapagedata}, nil)
}
