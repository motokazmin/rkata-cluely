# 🤖 Cluely - AI Assistant for IT Team Leads

**MVP Version - Minimum Viable Product**

Cluely is a local, confidential AI assistant for IT team leads and incident commanders. It works as a "virtual co-pilot" during critical incidents by analyzing audio, visual context, and providing real-time smart suggestions.

## 🚀 Quick Start (MVP Mode)

### Prerequisites

- **Go 1.21+** - [Download](https://go.dev/dl/)
- **Windows 11** (MVP target platform)
- **Make** - [For Windows](https://www.gnu.org/software/make/) or use `go run` directly

### Installation & Running

#### Option 1: Using Makefile (Recommended)

```bash
# Build the application
make build

# This will:
# - Compile to bin/cluely.exe
# - Copy configs/default.toml to bin/default.toml
# - Display instructions

# Run the application
make run

# Or if already built
bin/cluely.exe
```

#### Option 2: Direct Go Commands

```bash
# Download dependencies
go mod download

# Run directly from source
go run ./cmd/cluely/main.go
```

#### Option 3: Build manually

```bash
# Build the binary
go build -o bin/cluely.exe ./cmd/cluely

# Copy config
copy configs/default.toml bin/default.toml

# Run
bin/cluely.exe
```

### Available Make Commands

```bash
make help         # Show all available commands
make build        # Compile to bin/cluely.exe (with config)
make run          # Build and run automatically
make run-binary   # Run if already built
make clean        # Remove bin directory
make fmt          # Format Go code
make lint         # Run go vet
make test         # Run tests
make deps         # Download/update dependencies
make dev-setup    # Clean setup for development
```

3. **Expected Output:**
```
🔨 Building Cluely MVP...
📋 Copying configuration...
✅ Build complete: bin/cluely.exe

📝 To run the application:
   cd bin
   cluely.exe

🌐 Then open: http://localhost:8080
```

4. **Run the application:**
   - After building, run: `bin/cluely.exe` from the project root
   - Or use: `make run-binary`

5. **Open the UI:**
   - Go to: http://localhost:8080
   - You should see a dark interface with "Cluely AI Assistant" header
   - Status should show "✅ Connected"

6. **Watch the magic (in mock mode):**
   - Every 5 seconds: Audio module simulates a transcript
   - Every 10 seconds: Vision module simulates a screenshot
   - AI analyzes the context and sends hints to UI
   - You'll see AI suggestions appearing in real-time

### Configuration (Mock Mode)

Default config is already in mock mode. No setup required!

File: `configs/default.toml`
```toml
[audio]
enabled = true
transcriber_type = "mock"  # Mock mode - no Azure Speech needed

[vision]
enabled = true
ocr_engine = "mock"  # Mock mode - no Tesseract needed

[ai]
provider = "mock"  # Mock mode - no Ollama needed
```

## 🏗️ MVP Architecture

```
┌─────────────────────────────────────────────────────────┐
│                  Cluely Agent Core                       │
├─────────────────────────────────────────────────────────┤
│                                                           │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────┐  │
│  │ Audio Module │    │ Vision Module│    │AI Module │  │
│  │              │    │              │    │          │  │
│  │ - Transcriber│    │ - OCREngine  │    │-Provider │  │
│  │ - Channels   │    │ - Channels   │    │-Analyzer │  │
│  └────┬─────────┘    └────┬─────────┘    └────┬─────┘  │
│       │                    │                    │        │
│       └────────┬───────────┴────────┬──────────┘        │
│               │                    │                    │
│          Processing Loop           │                    │
│          (goroutine)                │                    │
│               │                    │                    │
│               └────────┬───────────┘                    │
│                       │                                 │
│                    ┌──▼──────┐                         │
│                    │UI Server │                         │
│                    │(WebSocket)                         │
│                    └──────────┘                         │
└─────────────────────────────────────────────────────────┘
                         │
                    http://localhost:8080
                         │
                    ┌────▼─────┐
                    │HTML/JS UI │
                    │  (Browser)│
                    └──────────┘
```

## 📊 What MVP Does

1. **Simulates Audio Input** (every 5 seconds)
   - "CPU at 95%, suspected memory leak"
   - "Need to rollback to previous commit"
   - etc.

2. **Simulates Visual Context** (every 10 seconds)
   - ERROR logs
   - Metrics (CPU: 95%, Memory: 8.2GB/16GB)
   - Pod statuses, HTTP errors
   - etc.

3. **AI Analysis** (Mock Provider)
   - Analyzes transcripts and screenshots
   - Generates smart hints based on patterns
   - Suggests actionable tasks
   - Provides warnings

4. **Real-Time UI**
   - WebSocket connection to backend
   - Displays hints with timestamps
   - Auto-scrolls to latest
   - Keeps last 10 messages

## 🔧 Development Notes

### Project Structure
```
cluely/
├── cmd/
│   └── cluely/
│       └── main.go              # Entry point
├── internal/
│   ├── agent/
│   │   └── agent.go             # Orchestrator
│   ├── audio/
│   │   ├── module.go            # Audio module
│   │   ├── transcriber.go       # Transcriber interface
│   │   └── mock_transcriber.go  # Mock implementation
│   ├── vision/
│   │   ├── module.go            # Vision module
│   │   ├── ocr.go               # OCR interface
│   │   └── mock_ocr.go          # Mock implementation
│   ├── ai/
│   │   ├── module.go            # AI module
│   │   ├── provider.go          # AI provider interface
│   │   ├── ollama.go            # Ollama implementation
│   │   └── mock_provider.go     # Mock implementation
│   ├── config/
│   │   └── config.go            # Config structures
│   └── ui/
│       └── server.go            # WebSocket + HTTP server
├── configs/
│   └── default.toml             # Configuration file
├── prompts/                     # AI prompts directory
│   ├── incident_analysis.txt
│   ├── task_generation.txt
│   └── context_summary.txt
├── go.mod                       # Go module definition
└── README.md                    # This file
```

### Key Interfaces (Pluggable)

#### 1. Transcriber (Audio Module)
```go
type Transcriber interface {
    Transcribe(ctx context.Context, audioData []byte) (string, error)
    Initialize() error
    Close() error
}
```
- **Mock**: Simulates transcripts
- **Azure**: Real Azure Speech Services (future)
- **Vosk**: Offline transcription (future)

#### 2. OCREngine (Vision Module)
```go
type OCREngine interface {
    ExtractText(ctx context.Context, imageData []byte) (string, error)
    Initialize() error
    Close() error
}
```
- **Mock**: Simulates OCR text
- **Tesseract**: Real OCR (future)
- **PaddleOCR**: Alternative OCR (future)

#### 3. AIProvider (AI Module)
```go
type AIProvider interface {
    Analyze(ctx context.Context, input AnalysisInput) (AnalysisOutput, error)
    Health(ctx context.Context) error
}
```
- **Mock**: Smart simulated hints
- **Ollama**: Local LLM (future)
- **Cloud**: OpenAI/Azure (future)

## 🧪 Testing the MVP

### Scenario 1: Watch Auto-Generated Hints
```
Just run: go run ./cmd/cluely/main.go
Wait for:
- Audio hints (CPU, Memory, Database, API issues)
- Vision hints (Errors, Metrics, Pod statuses)
- AI provides smart suggestions
```

### Scenario 2: Verify UI
```
1. Open http://localhost:8080
2. You should see "✅ Connected"
3. Hints appear every 5-10 seconds
4. Each hint shows timestamp
```

### Scenario 3: Check Logging
```
Console shows:
- Module startup messages
- Simulated audio transcripts
- Simulated OCR text
- AI analysis results
- UI communication
```

## 📈 MVP to Production (Next Stages)

### Stage 2: Real Audio & Screens
- Replace Mock Transcriber with Azure Speech
- Implement actual audio capture (VB-Cable)
- Replace Mock OCR with Tesseract
- Implement screenshot capture

### Stage 3: Real AI
- Replace Mock AI with Ollama integration
- Deploy local LLM models
- Implement prompt engineering

### Stage 4: Advanced Features
- Entity extraction (IDs, URLs, error codes)
- Voice commands
- Export to JIRA/Linear
- Cross-platform support (Linux, macOS)

## 🛠️ Commands

```bash
# Build for Windows
go build -o bin/cluely.exe ./cmd/cluely

# Run with output to file
go run ./cmd/cluely/main.go > logs/output.log 2>&1

# Run tests (when implemented)
go test ./...

# Format code
go fmt ./...

# Lint
go vet ./...
```

## 🔐 Confidentiality

MVP works **100% offline** in mock mode:
- ✅ No cloud calls
- ✅ No data sent anywhere
- ✅ All processing local
- ✅ No credentials needed

## 📝 Notes

- This is the **MVP** version for testing the architecture
- Mock mode ensures fast iteration without external dependencies
- All modules are pluggable - easy to swap implementations
- Future versions will add real providers (Azure, Ollama, Tesseract)

## 🤝 Contributing

The architecture follows SOLID principles:
- **D**ependency Inversion: All dependencies behind interfaces
- **I**nterface Segregation: Small, focused interfaces
- **S**ingle Responsibility: Each module has one job
- **O**pen/Closed: Easy to extend, hard to break
- **L**iskov Substitution: Implementations are interchangeable

## 📞 Support

For issues or questions, check:
1. `docs/roadmap.md` - Full specification
2. Console logs - Detailed execution trace
3. `configs/default.toml` - Configuration

---

**Ready to launch?**
```bash
go run ./cmd/cluely/main.go
```

Then open: http://localhost:8080

Happy debugging! 🚀
