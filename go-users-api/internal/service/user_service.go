package service


import (
"context"
"time"


"github.com/yourname/go-users-api/internal/repository"
"github.com/yourname/go-users-api/db/sqlc"
)


type UserService struct {
repo *repository.UserRepository
}


func NewUserService(r *repository.UserRepository) *UserService {
return &UserService{repo: r}
}


func calculateAge(dob time.Time) int {
now := time.Now()
age := now.Year() - dob.Year()
if now.YearDay() < dob.YearDay() {
age--
}
return age
}


func (s *UserService) GetUser(ctx context.Context, id int32) (sqlc.User, int, error) {
user, err := s.repo.GetByID(ctx, id)
if err != nil {
return user, 0, err
}
age := calculateAge(user.Dob)
return user, age, nil
}