# Wordle Helper Web Server - Quick Start Guide

## Starting the Server

### Option 1: Using Make (Easiest)
```bash
make run-server-dev
```

This will automatically:
- Set environment variables to use local dictionary files
- Start the server on http://localhost:8080

### Option 2: Manual Setup
```bash
# Set required environment variables
export WORDLE_DICTIONARY=./american-english
export WORDLE_REMOVE=./words-to-remove
export WORDLE_PORT=8080  # optional, defaults to 8080

# Run the server
go run cmd/server/main.go
```

### Option 3: Build and Run Binary
```bash
# Build the server
make target/server

# Run it
export WORDLE_DICTIONARY=./american-english
export WORDLE_REMOVE=./words-to-remove
./target/local/bin/wordle-server
```

## Accessing the Application

Once the server is running, open your browser to:
```
http://localhost:8080
```

## Using the Web Interface

### 1. Enter Missed Letters
In the "Missed Letters" field, enter all letters that appeared **gray** (not in the word).

Example: If you tried "CRANE" and C, N, E were gray, enter: `cne`

### 2. Fill Position Boxes
For each of the 5 position boxes, enter:

- **Green square (correct position)**: Just the letter
  - Example: `a` means 'a' is definitely at this position

- **Yellow square (wrong position)**: Minus sign + letters
  - Example: `-r` means 'r' is in the word but NOT at this position
  - Example: `-abc` means a, b, and c are all in the word but NOT at this position

- **Unknown**: Leave empty or use `.`

### 3. Click "Find Possible Words"
The results will appear below showing all matching words.

## Example Walkthrough

Let's say you're playing Wordle and:

**Turn 1: You tried "CRANE"**
- C: Gray ⬛
- R: Yellow 🟨 (in word, wrong position)
- A: Green 🟩 (correct position)
- N: Gray ⬛
- E: Gray ⬛

**How to enter this:**
- Missed Letters: `cne`
- Position 1: `.` (unknown)
- Position 2: `-r` (r is in word but not here)
- Position 3: `a` (correct!)
- Position 4: `.` (unknown)
- Position 5: `.` (unknown)

**Turn 2: You tried "SMART" from the results**
- S: Gray ⬛
- M: Yellow 🟨
- A: Green 🟩
- R: Yellow 🟨
- T: Gray ⬛

**How to update:**
- Missed Letters: `cnest` (added s and t)
- Position 1: `.` (unknown)
- Position 2: `-rm` (r and m are in word but not here)
- Position 3: `a` (still correct!)
- Position 4: `.` (unknown)
- Position 5: `-r` (r is in word but not here)

Continue narrowing down until you find the answer!

## Features

✅ **Real-time filtering** - Uses HTMX for instant results without page reload
✅ **Mobile friendly** - Works on phones and tablets
✅ **Smart filtering** - Uses your existing dictionary and removes past words
✅ **Clear instructions** - Built-in help text
✅ **Error handling** - Helpful error messages if input is invalid

## Configuration Options

### Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `WORDLE_DICTIONARY` | Yes | - | Path to dictionary file |
| `WORDLE_REMOVE` | No | - | Path to words-to-remove file |
| `WORDLE_PORT` | No | 8080 | Server port |
| `WORDLE_HOST` | No | localhost | Server host |

### Custom Port
```bash
export WORDLE_PORT=3000
make run-server
```

### Different Dictionary
```bash
export WORDLE_DICTIONARY=/path/to/your/words.txt
export WORDLE_REMOVE=/path/to/remove-list.txt
make run-server
```

## Troubleshooting

### "Missing WORDLE_DICTIONARY environment variable"
Make sure you've set the environment variable before starting the server:
```bash
export WORDLE_DICTIONARY=./american-english
```

### "Failed to load dictionary"
- Check that the dictionary file exists
- Verify the path is correct
- Ensure you have read permissions

### Server won't start (port in use)
If port 8080 is already in use:
```bash
export WORDLE_PORT=3000
make run-server
```

### Page doesn't load styling
- Make sure you're in the project root directory when running the server
- The `web/static` directory needs to be accessible
- Check browser console for errors

### Results don't match CLI
The web version uses the same logic as CLI, so results should match exactly. If they don't:
- Verify you're entering the constraints the same way
- Check that position inputs use the correct format (`a` for green, `-a` for yellow)

## Development

### Running Tests
```bash
go test ./...
```

### Checking for Errors
```bash
make staticcheck
make lint
```

### Building Both CLI and Server
```bash
make build          # Builds CLI
make target/server  # Builds server
```

## What's Next?

Now that your server is running, you can:
1. Open http://localhost:8080 in your browser
2. Start solving Wordles!
3. Share the server on your local network (change WORDLE_HOST to your IP)

Enjoy solving Wordles faster! 🎯

