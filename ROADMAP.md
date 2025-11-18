# Implementation Roadmap

## Quick Start Implementation Plan

### Phase 1: Core Functionality (MVP)

#### Step 1: Update Server Main (cmd/server/main.go)
- [ ] Load dictionary using existing `dictionary.Create()` function
- [ ] Set up routes for GET / and POST /wordle/solve
- [ ] Add static file server for CSS/JS
- [ ] Read configuration from environment variables
- [ ] Proper error handling and logging

**Files to modify:**
- `cmd/server/main.go` (major rewrite)

#### Step 2: Create Handler Functions (handlers/wordle.go)
- [ ] `HandleGetForm()` - render main page
- [ ] `HandlePostSolve()` - process form and return results
- [ ] Helper function to parse form data into wordle types
- [ ] Error handling for invalid inputs

**Files to modify:**
- `handlers/wordle.go` (major additions)

**New types needed:**
```go
type FormData struct {
    Missed string
    Pos0   string
    Pos1   string
    Pos2   string
    Pos3   string
    Pos4   string
}

type ResultData struct {
    Words      []string
    Count      int
    FormData   FormData
    Error      string
}
```

#### Step 3: Create Templates
- [ ] Update `web/views/wordleform.gohtml` with full form
- [ ] Create `web/views/results.gohtml` for partial results
- [ ] Update `web/views/layouts/bootstrap.gohtml` for proper styling

**Files to modify:**
- `web/views/wordleform.gohtml`
- `web/views/layouts/bootstrap.gohtml`

**New files:**
- `web/views/results.gohtml`

#### Step 4: Add Bootstrap CSS
- [ ] Download Bootstrap CSS or use CDN
- [ ] Add to templates

**Files to create/modify:**
- `web/static/css/bootstrap.min.css` (or use CDN)
- `web/views/layouts/bootstrap.gohtml` (add CSS link)

#### Step 5: Testing
- [ ] Manual testing with browser
- [ ] Test various input combinations
- [ ] Verify results match CLI output

### Phase 2: Polish & Enhancement

#### Step 6: Improve UI/UX
- [ ] Add custom CSS for Wordle-themed colors
- [ ] Add helpful placeholder text
- [ ] Add example inputs
- [ ] Responsive design tweaks

**New files:**
- `web/static/css/wordle.css`

#### Step 7: Better Error Handling
- [ ] Validate input format
- [ ] Show specific error messages
- [ ] Graceful degradation without HTMX

#### Step 8: Loading States
- [ ] Add HTMX loading indicators
- [ ] Disable form during submission
- [ ] Show word count prominently

### Phase 3: Optional Enhancements

#### Step 9: Additional Features (Pick and choose)
- [ ] Session-based game history
- [ ] Word frequency/difficulty indicators
- [ ] Best next guess suggestions
- [ ] Export results
- [ ] Dark mode toggle

## Implementation Code Structure

### Recommended File Organization

```
cmd/server/main.go
├─ Load config
├─ Initialize dictionary
├─ Set up logging
├─ Register routes
└─ Start server

handlers/wordle.go
├─ HandleGetForm
│  └─ Render empty form
├─ HandlePostSolve
│  ├─ Parse form data
│  ├─ Validate inputs
│  ├─ Call wordle.MakePossibles
│  └─ Render results
└─ Helper functions
   ├─ parseFormData
   └─ validateInputs

web/views/
├─ layouts/bootstrap.gohtml (base HTML)
├─ wordleform.gohtml (main page)
└─ results.gohtml (results partial)
```

## Key Integration Points

### Using Existing Code

1. **Dictionary Loading** (already works)
   ```go
   words, err := dictionary.Create(stderr, dict, remove)
   ```

2. **Word Filtering** (already works)
   ```go
   possibles := wordle.MakePossibles(words, missed, lettersAt, lettersNotAt)
   ```

3. **Input Parsing** (adapt from usrcmd)
   ```go
   // Transform form inputs to match CLI format
   args := []string{missed, pos0, pos1, pos2, pos3, pos4}
   missed, lettersAt, lettersNotAt, err := usrcmd.readArgs(args)
   ```

### New Code Needed

1. **HTTP Handlers** - Bridge between HTTP and wordle logic
2. **Form Parsing** - Convert HTTP form to wordle types
3. **Templates** - HTML views for browser
4. **Static Assets** - CSS and JS files

## Sample Implementation Snippets

### Main Server Structure
```go
func main() {
    // Load config
    dict := os.Getenv("WORDLE_DICTIONARY")
    remove := os.Getenv("WORDLE_REMOVE")
    port := os.Getenv("WORDLE_PORT")
    if port == "" {
        port = "8080"
    }
    
    // Initialize dictionary
    words, err := dictionary.Create(os.Stderr, dict, remove)
    if err != nil {
        log.Fatal(err)
    }
    
    // Set up logging
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    
    // Create views
    indexView := views.NewView("bootstrap", "web/views/wordleform.gohtml")
    
    // Routes
    http.HandleFunc("GET /", handlers.HandleGetForm(logger, indexView))
    http.HandleFunc("POST /wordle/solve", handlers.HandlePostSolve(logger, indexView, words))
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
    
    // Start server
    logger.Info("Starting server", "port", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}
```

### Form Parsing
```go
func parseFormToWordleInputs(r *http.Request) (string, []wordle.LetterAt, []wordle.LettersNotAt, error) {
    missed := r.FormValue("missed")
    positions := []string{
        r.FormValue("pos0"),
        r.FormValue("pos1"),
        r.FormValue("pos2"),
        r.FormValue("pos3"),
        r.FormValue("pos4"),
    }
    
    // Use existing usrcmd logic
    args := append([]string{missed}, positions...)
    return usrcmd.ReadUserCommand(strings.Join(args, " "))
}
```

### Template Data
```go
type PageData struct {
    FormData  FormData
    Results   []string
    WordCount int
    Error     string
}
```

## Development Workflow

### Local Development
```bash
# Terminal 1: Set up environment
export WORDLE_DICTIONARY=./american-english
export WORDLE_REMOVE=./words-to-remove
export WORDLE_PORT=8080

# Run server
go run cmd/server/main.go

# Terminal 2: Test
open http://localhost:8080
```

### Testing Checklist
- [ ] Empty form loads correctly
- [ ] Form submission works
- [ ] Results display correctly
- [ ] Error handling works
- [ ] Static files load
- [ ] Mobile responsive
- [ ] HTMX works without page reload
- [ ] Same results as CLI for identical inputs

## Deployment Options

### Option 1: Local Binary
```bash
go build -o wordle-server cmd/server/main.go
./wordle-server
```

### Option 2: Docker
```dockerfile
FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go build -o server cmd/server/main.go

FROM debian:stable-slim
WORKDIR /app
COPY --from=builder /app/server .
COPY web ./web
COPY american-english .
COPY words-to-remove .
ENV WORDLE_DICTIONARY=./american-english
ENV WORDLE_REMOVE=./words-to-remove
EXPOSE 8080
CMD ["./server"]
```

### Option 3: Makefile Target
```makefile
.PHONY: run-server
run-server:
	@export WORDLE_DICTIONARY=./american-english && \
	export WORDLE_REMOVE=./words-to-remove && \
	go run cmd/server/main.go
```

## Dependencies

### Already Have
- ✅ `golang.org/x/net` (if needed)
- ✅ `github.com/maragudk/gomponents` (optional, for alternative rendering)
- ✅ Go standard library (html/template, net/http)

### May Want to Add
- `github.com/go-chi/chi` (optional, better routing)
- `github.com/rs/zerolog` (optional, better logging)
- Consider: Keep it simple with stdlib for now

## Success Metrics

1. ✅ Server starts without errors
2. ✅ Form loads in browser
3. ✅ Submit returns results
4. ✅ Results match CLI output
5. ✅ Response time < 100ms
6. ✅ Works on mobile browser
7. ✅ No JavaScript errors in console

## Next Steps

Would you like me to:
1. **Implement the MVP** - I can write all the code for Phase 1
2. **Start with specific files** - Pick which file to start with
3. **Create a working example** - Implement one end-to-end flow
4. **Review existing server code** - Improve what's already there

Let me know which approach you prefer!

