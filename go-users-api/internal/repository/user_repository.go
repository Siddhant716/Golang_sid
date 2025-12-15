package repository


import (
"context"
"github.com/yourname/go-users-api/db/sqlc"
)


type UserRepository struct {
q *sqlc.Queries
}


func NewUserRepository(q *sqlc.Queries) *UserRepository {
return &UserRepository{q: q}
}


func (r *UserRepository) Create(ctx context.Context, name string, dob string) (sqlc.User, error) {
return r.q.CreateUser(ctx, sqlc.CreateUserParams{Name: name, Dob: dob})
}


func (r *UserRepository) GetByID(ctx context.Context, id int32) (sqlc.User, error) {
return r.q.GetUserByID(ctx, id)
}


func (r *UserRepository) List(ctx context.Context) ([]sqlc.User, error) {
return r.q.ListUsers(ctx)
}


func (r *UserRepository) Update(ctx context.Context, id int32, name, dob string) (sqlc.User, error) {
return r.q.UpdateUser(ctx, sqlc.UpdateUserParams{ID: id, Name: name, Dob: dob})
}


func (r *UserRepository) Delete(ctx context.Context, id int32) error {
return r.q.DeleteUser(ctx, id)
}