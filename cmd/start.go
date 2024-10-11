package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/kiapanahi/gooler/internal/traffic"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the gooler generator",
	Run: func(cmd *cobra.Command, args []string) {
		ifaceName := viper.GetString("interface")
		targets := viper.GetStringSlice("targets")

		log.Default().Println("Starting gooler with interface", ifaceName, "and targets", targets)

		iface, err := net.InterfaceByName(ifaceName)
		if err != nil {
			log.Default().Fatal(err)
		}

		config := traffic.GeneratorConfig{
			Interface: iface,
			Targets:   targets,
		}
		err = traffic.GenerateTraffic(config)
		if err != nil {
			log.Default().Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
