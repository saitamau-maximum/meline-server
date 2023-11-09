package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"github.com/saitamau-maximum/meline/handler"
	authHandler "github.com/saitamau-maximum/meline/handler/auth"
	"github.com/saitamau-maximum/meline/infra/auth"
	infra "github.com/saitamau-maximum/meline/infra/mysql"
	"github.com/saitamau-maximum/meline/usecase"
)

const (
	HOST = "database"
)

func main() {
	e := echo.New()

	db, err := infra.ConnectDB(HOST)
	if err != nil {
		e.Logger.Error(err)
	}

	bunDB := bun.NewDB(db, mysqldialect.New())
	defer bunDB.Close()

	authRepository := auth.NewAuthRepository()
	userRepository := infra.NewUserRepository(bunDB)
	authInteractor := usecase.NewAuthInteractor(authRepository)
	userInteractor := usecase.NewUserInteractor(userRepository)
	authHandler := authHandler.NewAuthHandler(authInteractor, userInteractor)
	userHandler := handler.NewUserHandler(userInteractor)

	e.GET("/", authHandler.Auth(func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}))
	// auth
	e.GET("/auth/login", authHandler.Login)
	e.GET("/auth/callback", authHandler.CallBack)
	e.POST("/signup", userHandler.CreateUser)

	e.Start(":8000")
}

