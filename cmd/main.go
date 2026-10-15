package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bildanjhry/go_shared-lib/utils"
	"github.com/bildanjhry/log-be/internal/handler"
	"github.com/bildanjhry/log-be/internal/repository"
	"github.com/bildanjhry/log-be/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	router := gin.Default()
	ctx := context.Background()
	utils.LoadEnv()

	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	api := router.Group("/api")
	api.GET("/users", userHandler.GetUsers)
	router.Run(":8080")
}
