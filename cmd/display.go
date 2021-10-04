package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// displayCmd represents the display command
var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "display your currents shows",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("display called with data file: %v\n", viper.GetString("storage-file"))
		execute()
	},
}

func init() {
	rootCmd.AddCommand(displayCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func execute() {
	//load data
	LoadData()
	//display data
	for i := 0; i < len(Shows); i++ {
		fmt.Println(Shows[i])
	}
}
