package cmd

import (
	"fmt"

	"github.com/merulis/shuffle/src/config"
	"github.com/spf13/cobra"
)

var sourceAdd = &cobra.Command{
	Use:   "add <name>",
	Short: "Add a configured source",
	Args:  cobra.ExactArgs(1),
	RunE:  runSourceList,
}

func init() {
	sourceCmd.AddCommand(sourceAdd)

	sourceAdd.Flags().String("type", "github", "source type")
	sourceAdd.Flags().String("owner", "", "source owner")
	sourceAdd.Flags().String("repo", "", "source repo")
	sourceAdd.Flags().String("ref", "", "source ref")
}

func runSourceList(cmd *cobra.Command, args []string) error {
	name := args[0]

	sourceType, err := cmd.Flags().GetString("type")
	if err != nil {
		return fmt.Errorf("read type flag: %w", err)
	}

	owner, err := cmd.Flags().GetString("owner")
	if err != nil {
		return fmt.Errorf("read owner flag: %w", err)
	}

	repo, err := cmd.Flags().GetString("repo")
	if err != nil {
		return fmt.Errorf("read repo flag: %w", err)
	}

	ref, err := cmd.Flags().GetString("ref")
	if err != nil {
		return fmt.Errorf("read ref flag: %w", err)
	}

	service := NewConfigService()

	err = service.Add(config.SourceConfig{
		Name:  name,
		Type:  config.SourceType(sourceType),
		Owner: owner,
		Repo:  repo,
		Ref:   ref,
	})
	if err != nil {
		return fmt.Errorf("add source: %w", err)
	}

	fmt.Println("source added: ", name)
	return nil
}
