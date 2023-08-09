package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"text/template"
)

func mustEnv(env string) (string, error) {
	res := os.Getenv(env)
	if res == "" {
		return "", fmt.Errorf("error: %s is must_env, but this env is empty", env)
	}
	return res, nil
}

func main() {
	args := os.Args
	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	f, err := os.Open(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Open: %s", err)
	}
	defer f.Close()

	body, err := io.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "io.ReadAll: %s", err)
		os.Exit(1)
	}

	funcMap := template.FuncMap{
		"env":      os.Getenv,
		"must_env": mustEnv,
	}

	tmpl, err := template.New("template").Funcs(funcMap).Parse(string(body))
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing: %s", err)
		os.Exit(1)
	}

	if err := tmpl.Execute(os.Stdout, nil); err != nil {
		fmt.Fprintf(os.Stderr, "execution: %s", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintln(os.Stderr, `Usage: text-env template_file_path`)
	flag.PrintDefaults()
}
