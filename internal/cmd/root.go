package cmd

import (
	"fmt"
	"os"

	"github.com/CelticAlreadyUse/CMS/internal/config"
	"github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
	Use:   "cms",
	Short: "CMS API with JWT authentication and category management",
	Long: `Content Management System API with user authentication using JWT 
and CRUD operations for categories.`,
}

func initConfig() {
	config.InitLoadWithViper()
}
func Execute() {
	initConfig()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

