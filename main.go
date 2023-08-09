package main

import (
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/spf13/pflag"
)

var Version string = "default version"

func mustEnv(env string) (string, error) {
	res := os.Getenv(env)
	if res == "" {
		return "", fmt.Errorf("error: %s is must_env, but this env is empty", env)
	}
	return res, nil
}

func main() {
	var optVersion = pflag.BoolP("version", "v", false, "Print version")
	var optOutput = pflag.StringP("out", "o", "", "Output file path")
	pflag.Parse()
	_ = optOutput

	if *optVersion {
		fmt.Fprintln(os.Stdout, Version)
		os.Exit(0)
	}

	args := pflag.Args()
	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	f, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Open: %s\n", err)
	}
	defer f.Close()

	body, err := io.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "io.ReadAll: %s\n", err)
		os.Exit(1)
	}

	funcMap := template.FuncMap{
		"env":      os.Getenv,
		"must_env": mustEnv,
	}

	tmpl, err := template.New("template").Funcs(funcMap).Parse(string(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing: %s\n", err)
		os.Exit(1)
	}

	var outFile *os.File
	if *optOutput != "" {
		outFile, err = os.Create(*optOutput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "os.Create: %s\n", err)
			os.Exit(1)
		}
		defer outFile.Close()
	} else {
		outFile = os.Stdout
	}

	if err := tmpl.Execute(outFile, nil); err != nil {
		fmt.Fprintf(os.Stderr, "execution: %s\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintln(os.Stderr, `Usage: text-env template_file_path`)
	pflag.PrintDefaults()
}
