package main

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
	"spoolman-go/pkg/fsutil"
	"spoolman-go/pkg/term"
)

var rootCmd = &cobra.Command{
	Use:   "entgen [ -s schemadir]",
	Short: "Generate Ent schema from the specified directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		term.Printfln(term.Debug, "Path to schema %s", schemaPath)
		schemaPath = fsutil.Abs(schemaPath)
		term.Printfln(term.Debug, "Absolute path %s", schemaPath)
		return genSchema(schemaPath)
	},
}

var schemaPath string

var genConfig = &gen.Config{
	Features: []gen.Feature{gen.FeatureVersionedMigration},
}

func printMessage(out *pterm.PrefixPrinter, msg string, args ...any) {
	out.Printfln(msg, args...)
}

func genSchema(schemaPath string) error {
	if term.Verbose {
		return genSchemaWithProgress(schemaPath)
	}
	return entc.Generate(schemaPath, genConfig)
}

func genSchemaWithProgress(schemaPath string) error {
	endChan := make(chan error, 1)
	spinner, err := pterm.DefaultSpinner.WithWriter(os.Stdout).WithShowTimer(true).WithText("Running ent generate").Start()
	if err != nil {
		return err
	}
	go func() {
		endChan <- entc.Generate(schemaPath, genConfig)
	}()
	err = <-endChan
	err2 := spinner.Stop()
	if err2 != nil {
		term.Printfln(term.Error, "Failed to stop the spinner %s", err2)
	}
	return err
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		pterm.Fatal.Printfln("Failed to generate ent schema %s", err)
	} else {
		pterm.Success.Printfln("Generate the ent schema successfully")
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&term.Verbose, "verbose", "v", false, "Output verbose messages")
	rootCmd.Flags().StringVarP(&schemaPath, "schema", "s", "./ent/schema", "Path to ent schema")

}
