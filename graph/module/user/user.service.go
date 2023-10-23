package userModule

import (
	"os"
	"stockels/models"
	"stockels/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) (models.User, error) {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 11);
	if err != nil {
		return models.User{}, err
	}

	userEntries := models.User{FullName: user.FullName,Email: user.Email,Password: string(encryptedPass),IsVerified: false}
	err = utils.DB().Create(&userEntries).Error

	return userEntries, err
}

func Login(email string, password string) (string, error) {
	user := models.User{}
	
	err := utils.DB().First(&user, "email = ?", email).Error
	if err != nil {
		return  "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password));
	if err != nil {
		return  "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	exportToken, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	return exportToken, err

}