package cmd

import (
	"encoding/hex"
	"os"

	"eminentcodex/cryptor/module"
	"github.com/spf13/cobra"
)

var (
	filePath  string
	plainText string
	key       string
)

func init() {
	encrypt.Flags().StringVarP(&filePath, "filepath", "f", "", "Path to file of whose content need to be encrypted")
	encrypt.Flags().StringVarP(&plainText, "text", "t", "", "Plain text that need to be encrypted")
	encrypt.Flags().StringVarP(&key, "key", "k", "", "key that is required to encrypt the text")
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
			encoded []byte
			fp      *os.File
		)

		if key == "" {
			module.WriteToExit(os.Stdout, "Please provide a key")
		}

		if filePath == "" && plainText == "" {
			module.WriteToExit(os.Stdout, "Please provide either filepath or plain text to encrypt")
		}

		if filePath != "" {
			if content, err = module.GetFileContent(filePath); err != nil {
				module.WriteToExit(os.Stdout, err.Error())
			}
		} else {
			content = []byte(plainText)
		}

		encoded, err = module.AESCBCEncrypt(content, []byte(key))
		if err != nil {
			module.WriteToExit(os.Stdout, err.Error())
		}
		if filePath != "" {
			// write to file
			fp, err = os.OpenFile(filePath, os.O_RDWR, 0777)
			if err != nil {
				module.WriteToExit(os.Stdout, err.Error())
			}

			fp.WriteString(hex.EncodeToString(encoded))

			os.Exit(0)
		} else {
			module.WriteTo(os.Stdout, hex.EncodeToString(encoded))
		}
	},
}
