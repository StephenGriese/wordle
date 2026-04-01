# Wordle Helper Documentation Index

This index lists the current, maintained documentation for the project.

## Start Here

1. `SERVER_QUICKSTART.md` - Run and use the web app locally
2. `HEROKU_DEPLOYMENT.md` - Deploy and configure on Heroku
3. `README.md` - Project overview and CLI/web entry points

## Core Technical Docs

- `ARCHITECTURE.md` - Current architecture and component flow
- `DESIGN.md` - Design choices and trade-offs
- `IMPLEMENTATION_SUMMARY.md` - Implementation details and key files
- `ROADMAP.md` - Future enhancements
- `CHECKLIST.md` - Validation checklist

## Historical / Change Notes

- `CHANGE_USE_AMERICAN_ENGLISH.md` - Dictionary source switch notes
- `FIX_HTMX_DUPLICATE_CONTENT.md` - HTMX rendering fix notes
- `GOMPONENTS_CONVERSION.md` - Migration notes from templates to gomponents
- `STATUS_WORDLE_TXT.md` - Status of old `wordle.txt` usage
- `SUCCESS_SUMMARY.md` - Earlier project milestone summary

## Code Map (Current)

- `cmd/server/main.go` - Web server entry point
- `handlers/wordle.go` - HTTP handlers
- `components/*.go` - Gomponents UI
- `dictionary/*.go` - Dictionary loading and word list management
- `cmd/cli/main.go` - CLI entry point
- `wordle/*.go` - Solver logic shared by CLI and web

## Quick Verify

```bash
go test ./...
make run-server-dev
```

Then open `http://localhost:8080`.
