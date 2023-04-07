package usecase

import (
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/app/repository"
	"github.com/Paulo-Lopes-Estevao/ci-cd_terraform_aws-ec2/entities"
)

type IUserUsecase interface {
	CreateUser(createUserDto *CreateUserDTO) (*UserResponse, error)
	GetByNameUser(name string) (*UserResponse, error)
	GetAllUser() ([]*UserResponse, error)
}

type userUsecase struct {
	UserRepository repository.IUserRepository
}

func NewUserUseCase(userRepo repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		UserRepository: userRepo,
	}
}

func (u *userUsecase) CreateUser(createUserDto *CreateUserDTO) (*UserResponse, error) {
	user, err := entities.NewUser(createUserDto.Name)
	if err != nil {
		return nil, err
	}

	findbyuser, err := u.UserRepository.FindByName(user.Name)
	if err == nil {
		if err := findbyuser.CheckIfUserExists(user.Name); err != nil {
			return nil, err
		}
	}

	if err := u.UserRepository.Save(user); err != nil {
		return nil, err
	}

	userDTO := UserResponse{}
	return userDTO.FromUser(user), nil

}

func (u *userUsecase) GetByNameUser(name string) (*UserResponse, error) {
	user, err := u.UserRepository.FindByName(name)
	if err != nil {
		return nil, err
	}
	userDTO := UserResponse{}
	return userDTO.FromUser(user), nil

}

func (u *userUsecase) GetAllUser() ([]*UserResponse, error) {
	users, err := u.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	userDTO := UserResponse{}
	userList := userDTO.FromUsers(users)

	return userList, nil

}
