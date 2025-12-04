package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type pageData struct {
	Title       string
	Message     string
	Hostname    string
	GeneratedAt string
	HostPort    string
	Example1    string
	Example2    string
}

var homeTemplate = template.Must(template.New("home").Parse(`
<!doctype html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{.Title}}</title>
  <style>
    :root {
      --bg: #0f172a;
      --panel: #0b1222;
      --text: #e2e8f0;
      --muted: #94a3b8;
      --accent: #38bdf8;
      --accent-2: #c084fc;
      --shadow: 0 18px 40px rgba(15, 23, 42, 0.35);
    }
    * { box-sizing: border-box; }
    body {
      margin: 0;
      min-height: 100vh;
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 24px;
      background: radial-gradient(circle at 20% 20%, rgba(56, 189, 248, 0.12), transparent 40%),
                  radial-gradient(circle at 80% 0%, rgba(192, 132, 252, 0.2), transparent 42%),
                  var(--bg);
      color: var(--text);
      font-family: "Space Grotesk", "Segoe UI", system-ui, -apple-system, sans-serif;
    }
    .card {
      width: min(920px, 100%);
      background: linear-gradient(135deg, rgba(15, 23, 42, 0.94), rgba(11, 18, 34, 0.96));
      border: 1px solid rgba(148, 163, 184, 0.18);
      border-radius: 18px;
      box-shadow: var(--shadow);
      padding: 28px 32px;
      backdrop-filter: blur(6px);
    }
    .header {
      display: flex;
      align-items: center;
      gap: 14px;
      margin-bottom: 18px;
    }
    .badge {
      padding: 10px 14px;
      border-radius: 12px;
      background: linear-gradient(135deg, rgba(56, 189, 248, 0.18), rgba(192, 132, 252, 0.18));
      border: 1px solid rgba(148, 163, 184, 0.25);
      color: var(--text);
      font-weight: 600;
      letter-spacing: 0.02em;
    }
    h1 {
      margin: 0;
      font-size: clamp(26px, 4vw, 34px);
      letter-spacing: -0.02em;
    }
    p {
      margin: 8px 0 0;
      color: var(--muted);
      line-height: 1.6;
      max-width: 680px;
    }
    .grid {
      margin-top: 20px;
      display: grid;
      gap: 18px;
      grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
    }
    .tile {
      padding: 18px;
      border-radius: 14px;
      border: 1px solid rgba(148, 163, 184, 0.2);
      background: rgba(15, 23, 42, 0.45);
    }
    .tile h2 {
      margin: 0 0 6px;
      font-size: 18px;
    }
    .tile span {
      color: var(--muted);
      font-size: 15px;
    }
    .actions {
      margin-top: 24px;
      display: flex;
      flex-wrap: wrap;
      gap: 12px;
    }
    button {
      cursor: pointer;
      border: none;
      border-radius: 12px;
      padding: 12px 16px;
      font-size: 15px;
      font-weight: 600;
      color: #0b1222;
      background: linear-gradient(135deg, var(--accent), var(--accent-2));
      box-shadow: 0 12px 25px rgba(56, 189, 248, 0.28);
      transition: transform 120ms ease, box-shadow 120ms ease;
    }
    button:hover { transform: translateY(-1px); box-shadow: 0 14px 26px rgba(56, 189, 248, 0.34); }
    button:active { transform: translateY(0); }
    code {
      background: rgba(148, 163, 184, 0.14);
      padding: 3px 7px;
      border-radius: 6px;
      color: var(--text);
      border: 1px solid rgba(148, 163, 184, 0.22);
    }
    .status {
      padding: 10px 12px;
      border-radius: 12px;
      border: 1px solid rgba(148, 163, 184, 0.2);
      color: var(--muted);
      min-height: 44px;
      background: rgba(15, 23, 42, 0.55);
      display: flex;
      align-items: center;
      gap: 10px;
    }
    .status .dot {
      width: 12px;
      height: 12px;
      border-radius: 50%;
      background: var(--accent);
      box-shadow: 0 0 0 4px rgba(56, 189, 248, 0.14);
      flex-shrink: 0;
    }
    footer {
      margin-top: 22px;
      color: var(--muted);
      font-size: 13px;
    }
    @media (max-width: 640px) {
      .card { padding: 22px; }
      .grid { grid-template-columns: 1fr; }
    }
  </style>
</head>
<body>
  <main class="card">
    <div class="header">
      <div class="badge">Go + Docker Compose</div>
      <h1>{{.Title}}</h1>
    </div>
    <p>{{.Message}}</p>
    <div class="grid">
      <div class="tile">
        <h2>Container Host</h2>
        <span>{{.Hostname}}</span>
      </div>
      <div class="tile">
        <h2>Rendered At</h2>
        <span>{{.GeneratedAt}}</span>
      </div>
      <div class="tile">
        <h2>Health Endpoint</h2>
        <span><code>GET /healthz</code></span>
      </div>
      <div class="tile">
        <h2>Host Port (HOST_PORT)</h2>
        <span>{{.HostPort}}</span>
      </div>
      <div class="tile">
        <h2>Example1</h2>
        <span>{{.Example1}}</span>
      </div>
      <div class="tile">
        <h2>Example2</h2>
        <span>{{.Example2}}</span>
      </div>
    </div>
    <div class="actions">
      <button id="ping">Check Health</button>
      <div class="status" id="status">
        <div class="dot"></div>
        Waiting to ping the server...
      </div>
    </div>
    <footer>Try curling the server inside the container: <code>curl -s localhost:8080/healthz</code></footer>
  </main>
  <script>
    const pingBtn = document.getElementById('ping');
    const status = document.getElementById('status');
    pingBtn.addEventListener('click', async () => {
      status.innerHTML = '<div class="dot"></div>Checking...';
      try {
        const res = await fetch('/healthz');
        const body = await res.json();
        status.innerHTML = '<div class="dot"></div>Status: ' + (body.status || 'unknown');
      } catch (err) {
        status.innerHTML = '<div class="dot" style="background:#f87171;box-shadow:0 0 0 4px rgba(248,113,113,0.2)"></div>Failed: ' + err;
      }
    });
  </script>
</body>
</html>
`))

func main() {
	port := envOrDefault("PORT", "8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/healthz", healthHandler)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           logging(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Starting server on :%s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	data := pageData{
		Title:       "Dockerized Go Webserver",
		Message:     "A tiny Go server running behind Docker Compose, rendering a simple HTML page.",
		Hostname:    hostname,
		GeneratedAt: time.Now().Format(time.RFC1123),
		HostPort:    envOrDefault("HOST_PORT", envOrDefault("PORT", "8080")),
		Example1:    envOrDefault("EXAMPLE1", "not set"),
		Example2:    envOrDefault("EXAMPLE2", "not set"),
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := homeTemplate.Execute(w, data); err != nil {
		log.Printf("template execute error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s from %s in %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
