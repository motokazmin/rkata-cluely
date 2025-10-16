package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Audio  AudioConfig  `toml:"audio"`
	Vision VisionConfig `toml:"vision"`
	AI     AIConfig     `toml:"ai"`
	UI     UIConfig     `toml:"ui"`
}

type AudioConfig struct {
	Enabled           bool              `toml:"enabled"`
	DeviceName        string            `toml:"device_name"`
	SampleRate        int               `toml:"sample_rate"`
	BufferSize        int               `toml:"buffer_size"`
	SilenceThreshold  float64           `toml:"silence_threshold"`
	TranscriberType   string            `toml:"transcriber_type"`
	TranscriberConfig map[string]string `toml:"transcriber_config"`
}

type VisionConfig struct {
	Enabled       bool              `toml:"enabled"`
	MonitoredApps []string          `toml:"monitored_apps"`
	HotKey        string            `toml:"hot_key"`
	OCREngine     string            `toml:"ocr_engine"`
	OCRConfig     map[string]string `toml:"ocr_config"`
}

type AIConfig struct {
	Provider  string `toml:"provider"`
	OllamaURL string `toml:"ollama_url"`
	Model     string `toml:"model"`
	PromptDir string `toml:"prompt_dir"`
}

type UIConfig struct {
	Enabled     bool    `toml:"enabled"`
	Port        int     `toml:"port"`
	Opacity     float64 `toml:"opacity"`
	Position    string  `toml:"position"`
	MaxMessages int     `toml:"max_messages"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
