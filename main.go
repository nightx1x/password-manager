package main

import (
	"fmt"

	"password-manager/config"
	"password-manager/manager"

	"github.com/spf13/cobra"
)

func main() {
	cfg := config.NewConfig()
	mgr := manager.NewManager(cfg.StoragePath)

	// Add command

	var addCmd = &cobra.Command{
		Use:   "add [name] [password]",
		Short: "Add new password",
		Run: func(cmd *cobra.Command, args []string) {
			if err := mgr.Add(args[0], args[1]); err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Pasword added.")
		},
	}

	// Generate command

	var generateCmd = &cobra.Command{
		Use:   "generate [name] [length]",
		Short: "Generate new password",
		Run: func(cmd *cobra.Command, args []string) {
			length := 12
			if len(args) == 2 {
				fmt.Sscanf(args[1], "%d", &length)
			}
			password, err := mgr.Generate(args[0], length)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Password generated:", password)
		},
	}

	// Delete command

	var deleteCmd = &cobra.Command{
		Use:   "delete [name]",
		Short: "Delete password",
		Run: func(cmd *cobra.Command, args []string) {
			if err := mgr.Delete(args[0]); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Password deleted.")
			}
		},
	}

	// List command

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List of passwords",
		Run: func(cmd *cobra.Command, args []string) {
			entries := mgr.List()
			if len(entries) == 0 {
				fmt.Println("There are no passwords.")
			} else {
				for _, e := range entries {
					fmt.Printf("%s | %s\n", e.Name, e.Password)
				}
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "password-manager"}
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.Execute()

}
