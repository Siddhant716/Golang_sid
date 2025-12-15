package main


import (
"database/sql"
"log"


"github.com/gofiber/fiber/v2"
_ "github.com/lib/pq"


"github.com/yourname/go-users-api/config"
"github.com/yourname/go-users-api/db/sqlc"
"github.com/yourname/go-users-api/internal/handler"
"github.com/yourname/go-users-api/internal/logger"
"github.com/yourname/go-users-api/internal/repository"
"github.com/yourname/go-users-api/internal/routes"
"github.com/yourname/go-users-api/internal/service"
)


func main() {
cfg := config.Load()
logg := logger.New()
defer logg.Sync()


db, err := sql.Open("postgres", cfg.DBUrl)
if err != nil {
log.Fatal(err)
}


queries := sqlc.New(db)
repo := repository.NewUserRepository(queries)
svc := service.NewUserService(repo)
h := handler.NewUserHandler(svc)


app := fiber.New()
routes.Register(app, h)


log.Fatal(app.Listen(":8080"))
}