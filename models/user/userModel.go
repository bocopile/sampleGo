package user

import (
	"database/sql"
	"sampleGo/handler/database/mysql"
	"sampleGo/handler/error"
)

// User는 사용자 정보를 표현하는 구조체입니다.
type User struct {
	ID   string
	Name string
}

// UserModel은 User 관련 작업을 수행하는 구조체입니다.
type UserModel struct {
	DB *sql.DB
}

// NewUserModel은 새로운 UserModel 인스턴스를 생성합니다.
func NewUserModel() *UserModel {
	dbManager := mysql.GetDBManager()
	return &UserModel{DB: dbManager.GetDB()}
}

// User 정보 조회
func (m *UserModel) GetUserByID(userID string) (*User, error) {
	var user User
	err := m.DB.QueryRow("SELECT id, name FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewAppError(2)
		}
		return nil, err
	}
	return &user, nil
}
