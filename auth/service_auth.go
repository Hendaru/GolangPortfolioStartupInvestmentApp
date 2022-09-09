package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface{
	GenerateToken(userID int) (string, error)
	
}

type jwtService struct {

}
var SECREAT_KEY = []byte("BWASTARTUP_s3cr3T_K3Y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error){
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err:= token.SignedString(SECREAT_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodeToken string) (*jwt.Token, error) {
	token , err := jwt.Parse(encodeToken, func(token *jwt.Token)(interface{}, error) {
		_, ok :=token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(SECREAT_KEY), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}