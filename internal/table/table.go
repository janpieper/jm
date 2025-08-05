package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
)

func New(headers []string) *tablewriter.Table {
	table := tablewriter.NewTable(os.Stdout)
	table.Header(headers)

	table.Configure(func(cfg *tablewriter.Config) {
		// Left align the headers
		cfg.Header.Alignment.Global = tw.AlignLeft

		// Remove padding (except between columns)
		cfg.Header.Padding.Global = tw.Padding{Left: "", Right: ""}
		cfg.Row.Padding.Global = tw.Padding{Left: "", Right: "  "}
	})

	table.Options(tablewriter.WithRendition(
		tw.Rendition{
			// Disable borders
			Borders: tw.Border{Left: tw.Off, Right: tw.Off},

			// Disable separators & lines
			Settings: tw.Settings{
				Separators: tw.Separators{BetweenColumns: tw.Off},
				Lines:      tw.Lines{ShowTop: tw.Off, ShowBottom: tw.Off, ShowHeaderLine: tw.Off},
			},
		},
	))

	return table
}
