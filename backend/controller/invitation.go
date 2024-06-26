package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	helper "social_network/helper"
	"social_network/models"
	"strconv"
)

func GetAllNotjoinedGroups(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	_, userEmail, _ := helper.Auth(DB, r)

	err := user.GetUserByEmail(DB, userEmail)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}

	groups, err := models.GetNotJoinedGroupsByUserID(DB, user.ID)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}
	var groupss []NewGroup
	for i := 0; i < len(groups); i++ {
		not, errn := models.GetNotificationByUserIDAndType(DB, user.ID, groups[i].UserID, "join-Group", groups[i].ID)
		if errn != nil {
			helper.ErrorPage(w, 500)
			return
		}

		notifStatus := false
		if len(not) == 0 {
			notifStatus = false
		} else {
			notifStatus = true
		}

		var gr = NewGroup{
			ID:          groups[i].ID,
			UserID:      groups[i].UserID,
			Title:       groups[i].Title,
			Description: groups[i].Description,
			NotifStatus: notifStatus,
		}

		groupss = append(groupss, gr)
	}

	groups2, err := models.GetMemberGroupsByUserID(DB, user.ID)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}
	var groupss2 []NewGroup
	for i := 0; i < len(groups2); i++ {
		var friendss []NewUser

		frs, er := user.GetFollowedAndFollowersNotInGroup(DB, groups2[i].ID)
		if er != nil {
			helper.ErrorPage(w, 500)
			return
		}
		for j := 0; j < len(frs); j++ {

			notfollow, errn := models.GetNotificationByUserIDAndType(DB, user.ID, frs[j].ID, "invited-to-join-Group", groups2[i].ID)
			if errn != nil {
				helper.ErrorPage(w, 500)
				return
			}
			requested := false

			if len(notfollow) == 0 {
				requested = false
			} else {
				requested = true
			}
			var fr = NewUser{
				ID:          frs[j].ID,
				FirstName:   frs[j].FirstName,
				LastName:    frs[j].LastName,
				Email:       frs[j].Email,
				Avatar:      frs[j].Avatar,
				Isrequested: requested,
			}
			friendss = append(friendss, fr)
		}
		var group = NewGroup{
			ID:          groups2[i].ID,
			UserID:      groups2[i].UserID,
			Title:       groups2[i].Title,
			Description: groups2[i].Description,
			SuggestList: friendss,
		}
		groupss2 = append(groupss2, group)

	}

	responseMap := map[string]interface{}{
		"joined":    groupss2,
		"Notjoined": groupss,
		"userid":    user.ID,
	}

	jsonData, err := json.Marshal(responseMap)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func CreateInvitationGroup(w http.ResponseWriter, r *http.Request) {

	type GroupInvitationData struct {
		GroupID string `json:"groupId"`
		UserID  string `json:"userId"`
		UserIDs []int  `json:"userIds"`
	}
	
	var groupData GroupInvitationData

	err := json.NewDecoder(r.Body).Decode(&groupData)
	fmt.Println("INVITATIONdata" , groupData)
	if err != nil {
		helper.ErrorPage(w, 500)
		return
	}

	var user = models.User{}
	userID, eru := strconv.Atoi(groupData.UserID)
	groupID, erg := strconv.Atoi(groupData.GroupID)
	if eru != nil || erg != nil {
		fmt.Println("w")
		helper.ErrorPage(w, 500)
		return
	}

	errr := user.GetUserById(DB, userID)
	if errr != nil {
		helper.ErrorPage(w, 500)
		return
	}

	mem, errm := user.IsGroupmemeber(DB, groupID)
	if errm != nil || !mem {
		fmt.Println(errm)
		helper.ErrorPage(w, 500)
		return
	}

	_, _, user_id := helper.Auth(DB, r)
	userSendingInvitations := models.User{}
	errGettingUser := userSendingInvitations.GetUserById(DB, user_id)
	if errGettingUser != nil {
		helper.ErrorPage(w, 500)
		return
	}

	for i := 0; i < len(groupData.UserIDs); i++ {
		var notification = models.Notification{}
		notification.SenderID = userSendingInvitations.ID
		notification.User_id = groupData.UserIDs[i]
		notification.Type = "invited-to-join-Group"
		notification.Group_id = groupID
		notification.Status = "false"
		currrentUser := models.User{}
		currrentUser.GetUserById(DB, groupData.UserIDs[i])

		notification.FirstName = userSendingInvitations.FirstName
		notification.LastName = userSendingInvitations.LastName
		notification.Avatar = userSendingInvitations.Avatar

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
			SendSocketNotification(notification, "notification")
		}

		
	}
}
