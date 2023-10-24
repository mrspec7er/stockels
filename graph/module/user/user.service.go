package user

import (
	"os"
	"stockels/graph/object"
	"stockels/models"
	"stockels/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(email string, password string) (*object.LoginResponse, error) {
	user := models.User{}
	result := &object.LoginResponse{}
	
	err := utils.DB().First(&user, "email = ?", email).Error
	if err != nil {
		return  result, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password));
	if err != nil {
		return  result, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	exportToken, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return  result, err
	}

	result.Token = exportToken

	return result, err

}

func CreateUserService(payload *object.Register) (*object.User, error){
	user := &object.User{}
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 11);
	if err != nil {
		return user, err
	}

	result := models.User{FullName: payload.FullName, Email: payload.Email,Password: string(encryptedPass), IsVerified: false}
	err = utils.DB().Create(&result).Error

	user.CreatedAt = result.CreatedAt.String()

	return &object.User{FullName: result.FullName, Email: result.Email, Password: "Encrypted", IsVerified: result.IsVerified, CreatedAt: result.CreatedAt.String(), UpdatedAt: result.CreatedAt.String()}, err
}