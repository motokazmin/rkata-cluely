package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"cluely/internal/agent"
	"cluely/internal/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("🚀 Starting Cluely Agent...")

	cfg, err := config.Load("default.toml")
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
