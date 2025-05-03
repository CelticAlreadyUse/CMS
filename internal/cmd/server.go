package cmd

import (
	"net/http"

	"github.com/CelticAlreadyUse/CMS/internal/config"
	mysqldb "github.com/CelticAlreadyUse/CMS/internal/database/mysql"
	"github.com/CelticAlreadyUse/CMS/internal/database/redis"
	http_handler "github.com/CelticAlreadyUse/CMS/internal/delivery/http"
	"github.com/CelticAlreadyUse/CMS/internal/repository"
	"github.com/CelticAlreadyUse/CMS/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var startServerCmd = &cobra.Command{
	Use:   "httpsrv",
	Short: "httpsrv is a command to run http server",
	Run: func(cmd *cobra.Command, args []string) {
		mysqlConn := mysqldb.MysqlConnection()
		redisClient := redis.InitConnectRedis()
		// Repository
		userRepo := repository.NewAuthUserRepository(mysqlConn)
		categoryRepo := repository.NewCategoryRepository(mysqlConn)
		newsRepo := repository.NewNewsRepository(mysqlConn)
		commentsRepo := repository.NewcommentRepository(mysqlConn)
		pagesRepo := repository.NewCustomPageRepository(mysqlConn)
		// Usecase
		newsUsecase := usecase.NewNewsUsecase(userRepo,redisClient, newsRepo, categoryRepo, commentsRepo)
		userUsecase := usecase.NewAuthUsecase(userRepo)
		categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
		commentUsecase := usecase.CommmentUsecase(commentsRepo, userRepo)
		pageUsecase := usecase.NewCustomPageUsecase(pagesRepo)
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		categoryHandler := http_handler.InitCateoryHandler(categoryUsecase)
		authHandler := http_handler.NewAuthHandler(userUsecase)
		newsHandler := http_handler.NewNewsHandler(newsUsecase)
		pagesHandler := http_handler.NewCustomPageHandler(pageUsecase)
		commentHandler := http_handler.NewCommentsHandler(commentUsecase)
		categoryHandler.RegisterRoute(r)
		pagesHandler.RegisterRoute(r)
		commentHandler.RegisterRoute(r)
		newsHandler.RegisterRoute(r)
		authHandler.RegisterRoute(r)
		if err := r.Run(":" + config.PORT_HTTP()); err != nil {
			logrus.Fatal("Failed to run HTTP server: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)
}
