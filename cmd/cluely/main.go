package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"cluely/internal/agent"
	"cluely/internal/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("🚀 Starting Cluely Agent...")

	// Try to find config file
	configPath := findConfigFile()
	if configPath == "" {
		log.Fatalf("❌ Failed to find config file (checked: default.toml, ./configs/default.toml)")
	}

	log.Printf("📋 Loading config from: %s", configPath)

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	agentInstance := agent.New(cfg)

	if err := agentInstance.Start(ctx); err != nil {
		log.Fatalf("❌ Failed to start agent: %v", err)
	}

	log.Println("✅ Cluely Agent started successfully!")
	log.Println("📝 Press Ctrl+C to stop...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	log.Println("\n🛑 Shutting down gracefully...")

	cancel()

	agentInstance.Stop()
	log.Println("👋 Goodbye!")
}

// findConfigFile searches for the config file in multiple locations
func findConfigFile() string {
	// Possible locations to search
	possiblePaths := []string{
		"default.toml",                     // Current directory
		"configs/default.toml",             // Configs subdirectory
		filepath.Join(".", "default.toml"), // Explicit current dir
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	return ""
}
