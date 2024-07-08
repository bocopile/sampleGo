package models

import (
	"log"
)

// User는 사용자 정보를 표현하는 구조체입니다.
type User struct {
	NO   int64  `json:"no"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetUserByID(userID string) ([]*User, error) {
	var users []*User
	user := User{}

	rows, err := db.Query("SELECT * FROM user where id=?")
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&user.NO, &user.ID, &user.Name)
		if err != nil {
			log.Fatalf(err.Error())
		}
		users = append(users, &user)

	}

	return users, err
}
