package cmd

import (
	"github.com/leb4r/trader-go/database"
	"github.com/leb4r/trader-go/database/migrations"
	"github.com/leb4r/trader-go/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

var (
	cfgFile   = ".trader-go.yml"
	dbHandler *gorm.DB
	sugar     *zap.SugaredLogger
)

var rootCmd = &cobra.Command{
	Use:   "trader-go",
	Short: "crypto trading CLI",
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// initialize the logger, this needs to happen first
	cobra.OnInitialize(initLogger)

	// initialize config
	cobra.OnInitialize(initConfig)

	// initialize database
	cobra.OnInitialize(initDatabase)

	// where to find config file
	rootCmd.PersistentFlags().String("config", "", "config file")

	// where database file should be at
	rootCmd.PersistentFlags().String("dbPath", "trader-go.db", "database file")
	if err := viper.BindPFlag("dbPath", rootCmd.PersistentFlags().Lookup("dbPath")); err != nil {
		internal.ThrowError(err)
	}
}

func initDatabase() {
	// open database connection
	db, err := database.Open(viper.GetString("dbPath"))
	if err != nil {
		internal.ThrowError(err)
	}

	// make sure database migrations are up
	if err := migrations.Execute(db); err != nil {
		internal.ThrowError(err)
	}

	// save reference to handler
	dbHandler = db
}

func initLogger() {
	sugar = zap.NewExample().Sugar()
	defer sugar.Sync()
}

func initConfig() {
	if cfgFile != "" {
		// cfgFile is bound to command line flag
		viper.SetConfigFile(cfgFile)
	} else {
		// otherwise, search for a `.trader-go.yaml` in the current directory
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".trader-go")
	}

	// load config from environment
	viper.AutomaticEnv()

	// error means there is no config file
	if err := viper.ReadInConfig(); err == nil {
		sugar.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}
