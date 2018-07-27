package cmd

import (
	"fmt"

	"github.com/MohitArora1/gophercises/secret/cipher"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get command will return api key from secrets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
		hex, err := cipher.Encrypt(args[0], args[1])
		if err != nil {
			panic(err)
		}
		value, err := cipher.Decrypt(args[0], hex)
		if err != nil {
			panic(err)
		}
		fmt.Printf("value is %s\n", value)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
