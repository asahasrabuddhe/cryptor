package cmd

import (
	"fmt"
	"os"

	"eminentcodex/cryptor/module"
	"github.com/spf13/cobra"
)

var (
	FilePath  string
	PlainText string
	Key       string
)

func init() {
	encrypt.Flags().StringVarP(&FilePath, "filepath", "f", "", "Path to file of whose content need to be encrypted")
	encrypt.Flags().StringVarP(&PlainText, "text", "t", "", "Plain text that need to be encrypted")
	encrypt.Flags().StringVarP(&Key, "key", "k", "", "Key that is required to encrypt the text")
	encrypt.MarkFlagRequired("key")
	rootCmd.AddCommand(encrypt)
}

var encrypt = &cobra.Command{
	Use:   "encrypt",
	Short: "Performs aes-256 encryption in cbc mode",
	Long:  "Performs aes-256 encryption in cbc mode",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			content []byte
			err     error
		)

		if FilePath == "" && PlainText == "" {
			module.WriteToExit(os.Stdout, "Please provide either filepath or plain text to encrypt")
		}

		if FilePath != "" {
			if content, err = module.GetFileContent(FilePath); err != nil {
				module.WriteToExit(os.Stdout, err.Error())
			}
		} else {
			content = []byte(PlainText)
		}

		fmt.Println(PlainText, content)

	},
}
