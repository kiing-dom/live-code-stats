# live-code-stats

A real-time coding stats overlay for OBS. Tracks lines of code, errors, and keystrokes per minute as you code in VSCode and displays them as a browser source.

## How it works

```
VSCode extension → Go server → WebSocket → OBS browser source
```

## Requirements

- Go 1.21+
- Node.js
- VSCode
- OBS Studio

## Setup

### 1. Start the server

```bash
go run cmd/live-code-stats/main.go
```

### 2. Install the extension

```bash
cd internal/extension
npm install
npm run compile
npx @vscode/vsce package --no-dependencies
code --install-extension live-code-stats-0.0.1.vsix
```

Restart VSCode after installing.

### 3. Add the overlay to OBS

In OBS, add a **Browser Source** and set the local file to `internal/overlay/index.html`. Make sure custom CSS is empty so the background stays transparent.

## Stats

| Field | Description |
|-------|-------------|
| Lines | Line count of the current file (updates on save) |
| Errors | Error diagnostics in the current file |
| KPM | Keystrokes per minute (rolling 60s window) |
| File | Current file name |

## License

MIT
