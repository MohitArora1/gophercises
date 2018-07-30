package cmd

import (
	"fmt"

	"github.com/MohitArora1/gophercises/secret/vault"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set command will put api key into secrets",
	Run: func(cmd *cobra.Command, args []string) {
		v := vault.GetVault(encodingKey, secretsPath())
		err := v.Set(args[0], args[1])
		if err != nil {
			fmt.Printf("%v\n", err.Error())
		}
		fmt.Println("saved key success")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
