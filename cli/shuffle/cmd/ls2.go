package cmd

import (
	"context"
	"fmt"

	"github.com/merulis/shuffle/src/app"
	"github.com/merulis/shuffle/src/domain"
	"github.com/spf13/cobra"
)

var ls2Cmd = &cobra.Command{
	Use:   "ls2 <source:/path>",
	Short: "List artifacts using saved source config",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runLs2(args)
	},
}

func init() {
	rootCmd.AddCommand(ls2Cmd)
}

func runLs2(args []string) error {
	target, err := parseTarget(args[0], ":")
	if err != nil {
		return err
	}

	source, sourceConfig, err := resolveSource(target.SourceName)
	if err != nil {
		return err
	}

	uc := app.NewListArtifacts(source)

	items, err := uc.Execute(
		context.Background(),
		domain.NewLocator(target.Path, sourceConfig.Ref),
	)
	if err != nil {
		return fmt.Errorf("ls2: %w", err)
	}

	for _, item := range items {
		fmt.Printf("%-4s %-8d %s\n", item.Type, item.Size, item.Path)
	}

	return nil
}
