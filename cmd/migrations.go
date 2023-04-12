package cmd

import (
	"bufio"
	"fmt"
	"github.com/RustamRR/job-rest-api/internal/store/postgrestore"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var migrationsCmd = &cobra.Command{
	Use:   "migrations",
	Short: "Выполнить миграции",
	Long:  `Эта команда выполнит миграции к базе`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Вы действительно хотите выполнить миграции к базе? (y/n):")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if "y" == text {
			if err := migrate(); err != nil {
				fmt.Println(fmt.Errorf("ошибка выполнения миграций:%v", err))
			}
			fmt.Println("Миграции выполнены")
			return
		}
		fmt.Println("отмена операции")
	},
}

func init() {
	rootCmd.AddCommand(migrationsCmd)
}

func migrate() error {
	config := viper.GetViper()
	db, err := gorm.Open(
		postgres.Open(config.GetString("dsn")),
	)

	if err != nil {
		return err
	}

	store := postgrestore.New(db)
	if err := store.Migrate(); err != nil {
		return err
	}

	return nil
}
