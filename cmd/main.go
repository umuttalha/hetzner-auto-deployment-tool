package main

import (
	"flag"
	"log"
	"os"

	"github.com/umuttalha/go-cli-tool/internal/config"
	"github.com/umuttalha/go-cli-tool/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	// Define command-line flags
	cfg := &config.Config{}

	// Server configuration flags
	flag.StringVar(&cfg.ServerType, "server-type", "cx21", "Hetzner server type (e.g., cx22, cx32, cx42)")
	flag.StringVar(&cfg.ServerImage, "server-image", "ubuntu-24.04", "Server image to use")
	flag.StringVar(&cfg.ServerLocation, "location", "nbg1", "Server location (nbg1, fsn1, hel1)")
	flag.StringVar(&cfg.BackendRepoURL, "repo-url", "", "Backend repository URL")

	// Add a help flag
	helpFlag := flag.Bool("help", false, "Show available options")

	flag.Parse()

	// Show help and available options
	if *helpFlag {
		config.ShowHelp()
		os.Exit(0)
	}

	// Load environment variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load config from environment variables
	cfg.LoadFromEnv()

	// Validate required configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Create server and setup DNS
	if err := server.SetupInfrastructure(*cfg); err != nil {
		log.Fatal(err)
	}
}
