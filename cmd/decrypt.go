package cmd

import (
	"encoding/hex"
	"os"

	"eminentcodex/cryptor/module"
	"github.com/spf13/cobra"
)

var (
	text string
)

func init() {
	decrypt.Flags().StringVarP(&filePath, "filepath", "f", "", "Path to file of whose content need to be decrypted")
	decrypt.Flags().StringVarP(&text, "text", "t", "", "Plain text that need to be decrypted")
	decrypt.Flags().StringVarP(&key, "key", "k", "", "Key that is required to decrypt the text")
	decrypt.MarkFlagRequired("key")
	rootCmd.AddCommand(decrypt)
}

var decrypt = &cobra.Command{
	Use:   "decrypt",
	Short: "Performs aes-256 decryption in cbc mode",
	Long:  "Performs aes-256 decryption in cbc mode",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			content []byte
			err     error
			decoded []byte
			fp      *os.File
		)

		if key == "" {
			module.WriteToExit(os.Stdout, "Please provide a key")
		}

		if filePath == "" && text == "" {
			module.WriteToExit(os.Stdout, "Please provide either filepath or plain text to decrypt")
		}

		// text and file content are in hex encoded so need to decode them first

		if filePath != "" {
			if content, err = module.GetFileContent(filePath); err != nil {
				module.WriteToExit(os.Stdout, err.Error())
			}
		} else {
			content = []byte(text)
		}

		content, err = hex.DecodeString(string(content))
		if err != nil {
			module.WriteToExit(os.Stdout, err.Error())
		}

		decoded, err = module.AESCBCDecrypt(content, []byte(key))
		if err != nil {
			module.WriteToExit(os.Stdout, err.Error())
		}
		if filePath != "" {
			//write to file
			fp, err = os.OpenFile(filePath, os.O_RDWR, 0777)
			if err != nil {
				module.WriteToExit(os.Stdout, err.Error())
			}
			fp.Truncate(0)
			fp.Write(decoded)
			os.Exit(0)
		} else {
			module.WriteTo(os.Stdout, string(decoded))
		}
	},
}
