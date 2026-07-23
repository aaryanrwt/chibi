package cli

import (
	"fmt"
	"os"

	"chibi/internal/ai"
	"chibi/internal/config"
	"chibi/internal/k8s"
	"chibi/internal/telemetry"
	"chibi/internal/tui"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	debug   bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chibi",
	Short: "Chibi is an AI-powered Kubernetes SRE assistant",
	Long: `Chibi empowers Kubernetes engineers with an intuitive, powerful, 
and intelligent terminal experience that simplifies complex operations 
and accelerates troubleshooting.`,
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		// Initialize Config
		if err := config.InitConfig(cfgFile); err != nil {
			return err
		}
		// Initialize Logger
		if err := telemetry.InitLogger(debug); err != nil {
			return err
		}
		// Initialize Cache
		if err := k8s.InitCache(); err != nil {
			return fmt.Errorf("cache initialization failed: %w", err)
		}
		// Initialize K8s Client (using default context/path for now)
		if err := k8s.InitClient("", ""); err != nil {
			telemetry.Log.Warn(fmt.Sprintf("Failed to initialize k8s client (you may not be connected to a cluster): %v", err))
		}
		return nil
	},
	Run: func(_ *cobra.Command, _ []string) {
		telemetry.Log.Info("Starting Chibi TUI...")

		engine := ai.NewEngine()
		openAIKey := os.Getenv("OPENAI_API_KEY")
		if openAIKey != "" {
			if provider, err := ai.NewOpenAIProvider(openAIKey); err == nil {
				engine.RegisterProvider(provider)
				_ = engine.SetActiveProvider("openai")
			}
		}

		if err := tui.StartTUI(engine); err != nil {
			fmt.Printf("Error running TUI: %v\n", err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.chibi.yaml)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug logging")
}
