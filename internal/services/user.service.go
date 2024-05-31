package services

import repo "github.com/tuanbui-n9/project-4/internal/repositories"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserById(id string) string {
	return us.userRepo.GetUserById(id)
}
