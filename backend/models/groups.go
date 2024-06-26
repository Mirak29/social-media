package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Group struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (g *Group) CreateGroup(db *sql.DB) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO \"Group\" (User_id, title, description) VALUES (?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("could not prepare statement: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(g.UserID, g.Title, g.Description)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %v", err)
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}

	return lastInsertId, nil
}

func GetGroupByID(db *sql.DB, groupID int) (*Group, error) {
	stmt, err := db.Prepare("SELECT id, User_id, title, description FROM \"Group\" WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get group by ID statement: %v", err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(groupID)

	var g Group

	err = row.Scan(&g.ID, &g.UserID, &g.Title, &g.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	return &g, nil
}

func GetGroupByUserid(db *sql.DB, userID int) (*Group, error) {
	stmt, err := db.Prepare("SELECT id, User_id, title, description FROM \"Group\" WHERE User_id = ?")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get group by ID statement: %v", err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(userID)

	var g Group

	err = row.Scan(&g.ID, &g.UserID, &g.Title, &g.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	return &g, nil
}

func GetAllGroups(db *sql.DB) ([]Group, error) {
	rows, err := db.Query("SELECT id, User_id, title, description FROM \"Group\"")
	if err != nil {
		return nil, fmt.Errorf("failed to query all groups: %v", err)
	}
	defer rows.Close()

	var groups []Group
	for rows.Next() {
		var g Group
		err := rows.Scan(&g.ID, &g.UserID, &g.Title, &g.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		groups = append(groups, g)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return groups, nil
}

func GetMemberGroupsByUserID(db *sql.DB, userID int) ([]Group, error) {
	stmt, err := db.Prepare(`
		SELECT g.id, g.User_id, g.title, g.description
		FROM "Group" g
		JOIN Joinner j ON g.id = j.Group_id
		WHERE j.User_id = ?
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare member groups statement: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var groups []Group
	for rows.Next() {
		var g Group
		err := rows.Scan(&g.ID, &g.UserID, &g.Title, &g.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		groups = append(groups, g)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return groups, nil
}

func GetNotJoinedGroupsByUserID(db *sql.DB, userID int) ([]Group, error) {
	stmt, err := db.Prepare(`
		SELECT g.id, g.User_id, g.title, g.description
		FROM "Group" g
		LEFT JOIN Joinner j ON g.id = j.Group_id AND j.User_id = ?
		WHERE j.User_id IS NULL
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare not joined groups statement: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var groups []Group
	for rows.Next() {
		var g Group
		err := rows.Scan(&g.ID, &g.UserID, &g.Title, &g.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		groups = append(groups, g)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return groups, nil
}

func JoinGroup(db *sql.DB, userID int, groupID int) error {
	stmt, err := db.Prepare("INSERT INTO Joinner (Group_id, User_id) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare join group statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(groupID, userID)
	if err != nil {
		return fmt.Errorf("failed to execute join group statement: %v", err)
	}

	return nil
}

func GetGroupMemberCount(db *sql.DB, groupID int) (int, error) {
	stmt, err := db.Prepare("SELECT COUNT(*) FROM Joinner WHERE Group_id = ?")
	if err != nil {
		return  0, fmt.Errorf("failed to prepare member count statement: %v", err)
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(groupID).Scan(&count)
	if err != nil {
		return  0, fmt.Errorf("failed to execute member count query: %v", err)
	}

	return count, nil
}

func (G *Group) GetUserOfGroup(DB *sql.DB)([]int,error){
	users:=[]int{}
	req:=`SELECT "User_id" FROM "Joinner" WHERE "Group_id"=?`
	row,err:=DB.Query(req,G.ID)
	if err!=nil {
		return users,err
	}
	usr:=0
	for row.Next(){
		row.Scan(&usr)
		users = append(users, usr)
	}
	return users,nil
}

func (G *Group) GetGroupMembers(DB *sql.DB) ([]User, error) {
    users := []User{}
    query := `
        SELECT u.id, u.firstname, u.lastname, u.avatar
        FROM "User" u
        JOIN "Joinner" j ON u.id = j.User_id
        WHERE j.Group_id = ?
    `
    rows, err := DB.Query(query, G.ID)
    if err != nil {
        return users, fmt.Errorf("failed to query group members: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Avatar)
        if err != nil {
            return users, fmt.Errorf("failed to scan row: %v", err)
        }
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        return users, fmt.Errorf("error during rows iteration: %v", err)
    }

    return users, nil
}

func GetGroupWithMembers(db *sql.DB, groupID int) (*Group, []User, error) {
	// Récupérer les informations du groupe
	group, err := GetGroupByID(db, groupID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get group by ID: %v", err)
	}

	// Récupérer les membres du groupe
	members, err := group.GetGroupMembers(db)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get group members: %v", err)
	}

	return group, members, nil
}