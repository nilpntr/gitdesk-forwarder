package cmd

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/nilpntr/gitdesk-forwarder/internal/config"
	"github.com/nilpntr/gitdesk-forwarder/internal/handlers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	cfgFile  string
	logLevel string
	cfg      config.Config
)

var rootCmd = &cobra.Command{
	Use:   "gitdesk-forwarder",
	Short: "GitDesk Forwarder is a tool to forward service desk issues to other message channels",
	RunE: func(cmd *cobra.Command, args []string) error {
		r := gin.New()
		r.Use(gin.Recovery())
		r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))

		r.GET("/health", func(c *gin.Context) {
			c.String(200, "OK")
		})

		for _, webhook := range cfg.Webhooks {
			r.POST(webhook.ListenPath, func(c *gin.Context) {
				handlers.HandleWebhook(c, webhook)
			})
		}

		if err := r.Run(fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	cobra.OnInitialize(initConfig, initLogger)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/config.yaml)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "log level (default is info)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.SetDefault("botUsername", "support-bot")
	viper.SetDefault("port", 8080)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println("Unable to decode config:", err)
		os.Exit(1)
	}
}

func initLogger() {
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	enc := zap.NewProductionEncoderConfig()
	enc.TimeKey = "timestamp"
	enc.EncodeTime = zapcore.ISO8601TimeEncoder

	zapCfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     enc,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}
	logger := zap.Must(zapCfg.Build())
	logger.Info("Logger initialized 🎉")
	zap.ReplaceGlobals(logger)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
