# live-code-stats

Sends real-time coding stats to a local Go server, which broadcasts them to an OBS browser source overlay.

## What it tracks

- **Lines** — line count of the current file (updates on save)
- **Errors** — error diagnostics in the current file
- **KPM** — keystrokes per minute (rolling 60s window)
- **File** — current file name

## Requirements

The [live-code-stats server](https://github.com/kiing-dom/live-code-stats) must be running locally before the extension will send any data.

```bash
go run cmd/live-code-stats/main.go
```

## Usage

Once both the server and extension are running, open any file in VSCode and start coding. Stats are sent automatically — no configuration needed.
