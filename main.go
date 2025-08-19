package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/singleservingfriend/klo/parser"
	"github.com/singleservingfriend/klo/transpiler"
)

func main() {
	app := &cli.App{
		Name:        "klo",
		Usage:       "A minimalist programming language built on Go",
		Version:     "0.1.0",
		Description: "klo transpiles simple, Python-like syntax to Go and executes it",
		Action:      runKloFile,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "transpile",
				Aliases: []string{"t"},
				Usage:   "Only transpile to Go, don't execute",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Usage:   "Show verbose output",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output file for transpiled Go code",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Show version information",
				Action: func(c *cli.Context) error {
					fmt.Printf("klo version %s\n", c.App.Version)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func runKloFile(c *cli.Context) error {
	if c.NArg() < 1 {
		return cli.ShowAppHelp(c)
	}

	filePath := c.Args().Get(0)
	if !strings.HasSuffix(filePath, ".klo") {
		return fmt.Errorf("file must have .klo extension")
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filePath)
	}

	// Read the klo source code
	source, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	if c.Bool("verbose") {
		fmt.Printf("Parsing klo file: %s\n", filePath)
		fmt.Printf("File content: %q\n", string(source))
	}

	// Parse the klo code
	ast, err := parser.Parse(string(source))
	if err != nil {
		return fmt.Errorf("parse error: %v", err)
	}

	if c.Bool("verbose") {
		fmt.Println("Transpiling to Go...")
	}

	// Generate Go code
	goCode := transpiler.GenerateGoCode(ast)

	// Determine output file
	outputFile := c.String("output")
	if outputFile == "" {
		// Create temporary file
		baseName := strings.TrimSuffix(filepath.Base(filePath), ".klo")
		// Avoid _test suffix which Go treats as test files
		if strings.HasSuffix(baseName, "_test") {
			baseName = strings.TrimSuffix(baseName, "_test") + "_example"
		}
		outputFile = fmt.Sprintf("klo_temp_%s.go", baseName)
	}

	// Write Go code to file
	err = os.WriteFile(outputFile, []byte(goCode), 0644)
	if err != nil {
		return fmt.Errorf("error writing Go file: %v", err)
	}

	if c.Bool("verbose") {
		fmt.Printf("Generated Go code written to: %s\n", outputFile)
	}

	// If only transpiling, stop here
	if c.Bool("transpile") {
		fmt.Printf("Transpiled %s to %s\n", filePath, outputFile)
		return nil
	}

	// Execute the Go code
	if c.Bool("verbose") {
		fmt.Println("Executing Go code...")
	}

	cmd := exec.Command("go", "run", outputFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	// Clean up temporary file if it was auto-generated
	if c.String("output") == "" {
		os.Remove(outputFile)
	}

	if err != nil {
		return fmt.Errorf("execution error: %v", err)
	}

	return nil
}
