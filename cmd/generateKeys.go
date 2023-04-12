package cmd

import (
	"bufio"
	"fmt"
	"github.com/RustamRR/job-rest-api/internal/utils"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var generateKeysCmd = &cobra.Command{
	Use:   "generateKeys",
	Short: "Генерация ключей шифрования",
	Long:  `Генерация приватного и публичного ключей`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Вы уверены, что хотите сгенерировать public.pem и private.pem? (y/n):")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if "y" == text {
			err := utils.GenerateKeys()
			if err != nil {
				fmt.Println(fmt.Errorf("ошибка генерации ключей:%v", err))
				return
			}
			fmt.Println("ключи сгенерированы")
			return
		}

		fmt.Println("Отмена операции")

	},
}

func init() {
	rootCmd.AddCommand(generateKeysCmd)
}
