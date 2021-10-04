package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var Shows []Show

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "keetrashow",
	Short: "keep track of your shows !",
	Long:  `TODO`,
	Run:   displayCmd.Run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.keetrashow.yaml)")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		// Search config in home directory with name ".keetrashow" (without extension).
		viper.AddConfigPath(home)

		viper.SetConfigType("json")
		viper.SetConfigName(".keetrashow")

	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func LoadData() {
	Shows = append(Shows, Show{
		title:         "The wire",
		remote_api_id: "noideayet",
		seasons: []Season{
			{
				season:   1,
				episodes: 13,
			}, {
				season:   2,
				episodes: 12,
			}, {
				season:   3,
				episodes: 12,
			}, {
				season:   4,
				episodes: 13,
			}, {
				season:   5,
				episodes: 10,
			},
		},
		last_watched: Episode{
			season:  1,
			episode: 7,
		},
		finished: false,
	})
}
