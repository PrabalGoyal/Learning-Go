package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var rootCmd = &cobra.Command{
	Use:   "echo [-n] [string ...]",
	Short: "echo - write arguments to the standard output",
	Long:  `The echo utility writes any specified operands, separated by single blank (' ') characters and followed by a newline ('\n') character, to the standard output.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get docs flag
		generateDocs, err := cmd.Flags().GetBool("docs")
		if err != nil {
			log.Fatal(err)
		}
		// get docs flag
		skipNewline, err := cmd.Flags().GetBool("newline")
		if err != nil {
			log.Fatal(err)
		}
		// get version flag
		version, err := cmd.Flags().GetBool("version")
		if err != nil {
			log.Fatal(err)
		}
		// if version and docs flags is set then throw error
		if version && generateDocs {
			log.Fatal("--version and --docs flags are mutually exclusive")
			return
		}
		if generateDocs {
			createDirIfNotExist("./docs")
			err := doc.GenMarkdownTree(cmd, "./docs")
			if err != nil {
				log.Fatal(err)
			}
		} else if version {
			cmd.Println("0.0.1-alpha")
		} else {
			for i := 0; i < len(args)-1; i++ {
				fmt.Printf("%v ", args[i])
			}
			if !skipNewline {
				fmt.Println(args[len(args)-1])
			} else {
				fmt.Print(args[len(args)-1])
			}
		}
	},
}

func init() {
	rootCmd.Flags().BoolP("newline", "n", false, "Do not print the trailing newline character")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")
	rootCmd.Flags().Bool("docs", false, "Generate markdown docs for all commands")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	rootCmd.Execute()
}
