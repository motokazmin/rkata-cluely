package ui

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"cluely/internal/config"

	"github.com/gorilla/websocket"
)

type Server struct {
	cfg      config.UIConfig
	clients  []*websocket.Conn
	upgrader websocket.Upgrader
	mu       sync.Mutex
}

func NewServer(cfg config.UIConfig) *Server {
	return &Server{
		cfg:     cfg,
		clients: make([]*websocket.Conn, 0),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	}
}

func (s *Server) Start() error {
	if !s.cfg.Enabled {
		log.Println("⏭️  UI Server disabled")
		return nil
	}

	http.HandleFunc("/ws", s.handleWebSocket)
	http.HandleFunc("/", s.handleIndex)
	http.HandleFunc("/health", s.handleHealth)

	addr := fmt.Sprintf(":%d", s.cfg.Port)
	go func() {
		log.Printf("🌐 UI server listening on http://localhost%s", addr)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Printf("❌ UI server error: %v", err)
		}
	}()

	return nil
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("❌ WebSocket upgrade error: %v", err)
		return
	}

	s.mu.Lock()
	s.clients = append(s.clients, conn)
	s.mu.Unlock()

	log.Println("✅ New WebSocket client connected")

	// Отправляем приветствие
	s.sendToClient(conn, map[string]interface{}{
		"type": "info",
		"data": "Connected to Cluely",
	})
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Cluely - AI Assistant</title>
    <style>
        body { 
            font-family: Arial, sans-serif; 
            background: #1a1a1a; 
            color: #fff;
            margin: 0;
            padding: 20px;
        }
        #hints {
            max-width: 800px;
            margin: 0 auto;
        }
        .hint {
            background: #2a2a2a;
            border-left: 4px solid #00ff88;
            padding: 15px;
            margin: 10px 0;
            border-radius: 4px;
            animation: slideIn 0.3s ease;
        }
        @keyframes slideIn {
            from { opacity: 0; transform: translateX(-20px); }
            to { opacity: 1; transform: translateX(0); }
        }
        .status {
            color: #00ff88;
            margin-bottom: 20px;
        }
        h1 { color: #00ff88; }
        .timestamp {
            color: #888;
            font-size: 12px;
        }
    </style>
</head>
<body>
    <h1>🤖 Cluely AI Assistant</h1>
    <div class="status" id="status">Connecting...</div>
    <div id="hints"></div>
    
    <script>
        const ws = new WebSocket('ws://' + window.location.host + '/ws');
        const status = document.getElementById('status');
        const hints = document.getElementById('hints');
        
        ws.onopen = () => {
            status.textContent = '✅ Connected';
        };
        
        ws.onclose = () => {
            status.textContent = '❌ Disconnected';
        };
        
        ws.onmessage = (event) => {
            const msg = JSON.parse(event.data);
            
            if (msg.type === 'hint') {
                const hint = document.createElement('div');
                hint.className = 'hint';
                const time = new Date().toLocaleTimeString();
                hint.innerHTML = '<div class="timestamp">' + time + '</div>' + msg.data;
                hints.insertBefore(hint, hints.firstChild);
                
                // Ограничиваем количество сообщений
                while (hints.children.length > 10) {
                    hints.removeChild(hints.lastChild);
                }
            }
        };
    </script>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func (s *Server) SendHint(hint string) {
	message := map[string]interface{}{
		"type": "hint",
		"data": hint,
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for i := len(s.clients) - 1; i >= 0; i-- {
		if err := s.clients[i].WriteJSON(message); err != nil {
			log.Printf("⚠️  Error sending to client: %v", err)
			s.clients[i].Close()
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
		}
	}
}

func (s *Server) sendToClient(conn *websocket.Conn, message interface{}) {
	if err := conn.WriteJSON(message); err != nil {
		log.Printf("⚠️  Error sending to client: %v", err)
	}
}

func (s *Server) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, client := range s.clients {
		client.Close()
	}
	s.clients = nil
}
