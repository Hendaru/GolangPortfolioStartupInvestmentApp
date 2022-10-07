package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUserService(input RegisterUserInput)(User, error) 
	LoginUserService(input LoginInput)(User, error)
	IsEmailAvailableUserService(input CheckEmailInput)(bool, error)
	SaveAvatarUserService(ID int, fileLocation string) (User, error)
	GetUserByIdUserService(ID int)(User, error)
}

type service struct {
	repository Repository
}

func  NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUserService(input RegisterUserInput)(User, error){
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, err :=bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.SaveUserRepository(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) LoginUserService(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmailUserRepository(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0{
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, errors.New("password salah cookkk")
	}
	return user, nil

}

func (s *service) IsEmailAvailableUserService(input CheckEmailInput)(bool, error){
	email := input.Email
	user, err := s.repository.FindByEmailUserRepository(email)
	if err != nil {
		
		return false, err
	}
	if user.ID ==0 {
		return true, nil
	}
	return false, nil

}
 
func (s *service) SaveAvatarUserService(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindByIDUserRepository(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName =fileLocation
	updatedUser,err := s.repository.UpdateUserRepository(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}

func (s *service) GetUserByIdUserService(ID int)(User, error){
	user,err := s.repository.FindByIDUserRepository(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("No user found with ID")
	}
	return user, nil
}
