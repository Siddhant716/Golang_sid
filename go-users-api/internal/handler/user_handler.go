package handler


import (
"strconv"


"github.com/gofiber/fiber/v2"
"github.com/go-playground/validator/v10"
"github.com/yourname/go-users-api/internal/models"
"github.com/yourname/go-users-api/internal/service"
)


type UserHandler struct {
svc *service.UserService
v *validator.Validate
}


func NewUserHandler(s *service.UserService) *UserHandler {
return &UserHandler{svc: s, v: validator.New()}
}


func (h *UserHandler) GetUser(c *fiber.Ctx) error {
id, _ := strconv.Atoi(c.Params("id"))
user, age, err := h.svc.GetUser(c.Context(), int32(id))
if err != nil {
return fiber.ErrNotFound
}


return c.JSON(fiber.Map{
"id": user.ID,
"name": user.Name,
"dob": user.Dob.Format("2006-01-02"),
"age": age,
})
}