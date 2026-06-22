package clih

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/merulis/shuffle/internal/app/entity"
)

func PrintSorces(w io.Writer, sources []entity.Source) error {
	if len(sources) < 1 {
		_, err := fmt.Fprintln(w, "No sources configured")
		return err
	}

	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	if _, err := fmt.Fprintln(tw, "NAME\tTYPE\tOWNER\tREPO\tREF"); err != nil {
		return err
	}

	for _, source := range sources {
		if _, err := fmt.Fprintf(
			tw,
			"%s\t%s\t%s\t%s\t%s\n",
			source.Name,
			source.Type,
			source.Owner,
			source.Repo,
			valueOrDash(source.Ref),
		); err != nil {
			return err
		}
	}

	return tw.Flush()
}

func valueOrDash(value string) string {
	if value == "" {
		return "-"
	}
	return value
}
