package cmd

import (
	_ "embed" //needed for embedding files
	"fmt"
	"os"

	"github.com/slurdge/goeland/config"
	"github.com/slurdge/goeland/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goeland run",
	Short: "goeland is a simple rss to email program.",
	Long: `goeland is a simple rss to email program.
	
It was inspired by rss2email, but is an alternative with some cool features, such as filters.
The simple way to use it is to type goeland run, then customize the create config.toml file.
To obtain a list of all the filter, type: goeland help run`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//      Run: func(cmd *cobra.Command, args []string) { },
}

func fatalErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fatalErr(err)
	}
}

//go:embed asset/config.default.toml
var defaultConfig []byte

func createDefaultConfig(cfgFile string) {
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		var configFile *os.File
		if configFile, err = os.Create(cfgFile); err != nil {
			fatalErr(fmt.Errorf("cannot open config.toml for writing"))
		}
		configFile.Write(defaultConfig)
	}
}

func initConfig() {
	createDefaultConfig(cfgFile)
	config.ReadDefaultConfig("goeland", cfgFile)
	log.SetDefaultLogger(log.NewLogger(viper.GetViper()))
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.toml", "config file (default is config.toml)")
	rootCmd.PersistentFlags().String("loglevel", "none", "Log level")
	viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
}
