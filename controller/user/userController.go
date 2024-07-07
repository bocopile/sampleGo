package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sampleGo/models/user"
)

type UserController struct {
	userModel *user.UserModel
}

// NewUserController는 새로운 UserController 인스턴스를 생성합니다.
func NewUserController() *UserController {
	return &UserController{
		userModel: user.NewUserModel(),
	}
}

// GetUserByIDHandler는 주어진 ID를 가진 사용자 정보를 반환하는 HTTP 핸들러입니다.
func (ctrl *UserController) GetUserByIDHandler(c *gin.Context) {
	// URL 매개변수에서 ID 추출
	id := c.Param("id")

	// 파라미터가 존재하는지 확인
	if id == "" {
		log.Printf("Missing user ID parameter")
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID parameter is required"})
		return
	}

	// 모델을 사용하여 사용자 정보 가져오기
	user, err := ctrl.userModel.GetUserByID(id)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
		return
	}

	// 사용자 정보 반환
	c.JSON(http.StatusOK, user)
}
