package cmd

import (
	"net/http"

	"github.com/CelticAlreadyUse/CMS/internal/config"
	mysqldb "github.com/CelticAlreadyUse/CMS/internal/database/mysql"
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

		// Repository
		userRepo := repository.NewAuthUserRepository(mysqlConn)
		categoryRepo := repository.NewCategoryRepository(mysqlConn)
		newsRepo := repository.NewNewsRepository(mysqlConn)
	
		// Usecase
		newsUsecase := usecase.NewNewsUsecase(newsRepo, categoryRepo)
		userUsecase := usecase.NewAuthUsecase(userRepo)
		categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		categoryHandler := http_handler.InitCateoryHandler(categoryUsecase)
		authHandler := http_handler.NewAuthHandler(userUsecase)
		NewsHandler := http_handler.NewNewsHandler(newsUsecase)
		categoryHandler.RegisterRoute(r)
		NewsHandler.RegisterRoute(r)
		authHandler.RegisterRoute(r)

		if err := r.Run(":" + config.PORT_HTTP()); err != nil {
			logrus.Fatal("Failed to run HTTP server: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)
}
