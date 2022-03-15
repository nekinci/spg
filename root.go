package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "spg",
	Short: "spg is a simple profile generator for spring boot",
	Long:  `spg is a simple profile generator for spring boot.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a profile",
	Long:  `Generate a profile`,
	Run: func(cmd *cobra.Command, args []string) {
		profile, err := cmd.PersistentFlags().GetString("profile")
		output, err := cmd.PersistentFlags().GetString("output")
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		RunGenerate(args, profile, output)
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure spg",
	Long:  `Configure spg`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You must be select a command. You can use 'spg config help' to see the list of commands")
	},
}

var setConfigCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration value. Example: spg config set [config path]",
	Long:  "Set a configuration value. Example: spg config set [config path]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 || len(args) == 0 {
			fmt.Println("Invalid number of arguments")
			os.Exit(1)
		}
		HandleConfig("set", args[0])
	},
}

var unsetConfigCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a configuration value. The command clears all config values.",
	Long:  "Unset a configuration value. The command clears all config values.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("Invalid number of arguments")
			os.Exit(1)
		}
		HandleConfig("unset", "")
	},
}

var printConfigCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the configuration",
	Long:  "Print the configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("Invalid number of arguments")
			os.Exit(1)
		}

		HandleConfig("print", "")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of spg",
	Long:  `All software has versions. This is spg's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: 0.1.0")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	configCmd.AddCommand(setConfigCmd)
	configCmd.AddCommand(unsetConfigCmd)
	configCmd.AddCommand(printConfigCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
	generateCmd.PersistentFlags().StringP("output", "o", "application-result.yml", "Output file")
	generateCmd.PersistentFlags().StringP("profile", "p", "", "Profile to use. Example: test")
}

func initConfig() {
	// Do nothing yet...
}
