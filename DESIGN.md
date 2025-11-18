# Wordle Helper Web Application Design

## Overview
A web-based interface for the Wordle Helper that allows users to input Wordle game constraints and receive a filtered list of possible words through a browser.

## Architecture

### Tech Stack
- **Backend**: Go standard library `net/http`
- **Frontend**: HTML templates with HTMX for interactivity
- **Styling**: Bootstrap CSS (already referenced in templates)
- **No external database needed** - uses in-memory word list

### Project Structure
```
cmd/
  server/
    main.go              # HTTP server entry point
handlers/
  wordle.go              # HTTP handlers for wordle endpoints
views/
  view.go                # Template rendering utilities
web/
  static/
    js/
      htmx-1.9.11.js     # HTMX library
    css/
      bootstrap.min.css  # Bootstrap CSS (to be added)
    favicon.ico
  views/
    wordleform.gohtml    # Main form for inputting constraints
    results.gohtml       # Partial template for results
    layouts/
      bootstrap.gohtml   # Base layout
```

## User Interface Design

### Main Page Layout
1. **Header**: "Wordle Helper" title with instructions
2. **Input Form**:
   - **Missed Letters**: Single text input for letters that aren't in the word
   - **Position Inputs**: 5 text inputs (one per position) for:
     - Known letters at position (e.g., "a" = 'a' is at this position)
     - Letters NOT at position (prefix with "-", e.g., "-abc" = a, b, c are in word but not here)
     - Empty/dot (.) for unknown positions
3. **Submit Button**: Triggers the word filtering
4. **Results Area**: Displays possible words in a grid layout

### User Interaction Flow
1. User enters constraints based on Wordle feedback:
   - Green letters → position inputs (e.g., "a" in position 1)
   - Yellow letters → position inputs with "-" prefix (e.g., "-a" in position 1)
   - Gray letters → missed letters input
2. User clicks "Find Words"
3. HTMX sends POST request to server
4. Server returns filtered word list
5. Results display in-place without page reload

### Example Input Format (matching CLI)
- **Missed**: "xyz" (letters not in word)
- **Position 1**: "a" (letter 'a' is at position 1)
- **Position 2**: "-bc" (letters 'b' and 'c' are in word but not at position 2)
- **Position 3**: "." (unknown)
- **Position 4**: "." (unknown)
- **Position 5**: "e" (letter 'e' is at position 5)

## API Endpoints

### GET /
- **Purpose**: Serve the main page with the empty form
- **Response**: HTML page with form

### POST /wordle/solve
- **Purpose**: Process wordle constraints and return possible words
- **Request Body** (form data):
  ```
  missed: string
  pos0: string
  pos1: string
  pos2: string
  pos3: string
  pos4: string
  ```
- **Response**: HTML fragment with word list (for HTMX) or full page
- **Error Handling**: Display user-friendly error messages

### GET /static/*
- **Purpose**: Serve static assets (JS, CSS, images)
- **Response**: Static files

## Backend Implementation

### Main Server (cmd/server/main.go)
```go
- Load dictionary from environment variables
- Initialize word list
- Set up HTTP routes
- Start server on configurable port
```

### Handler Functions (handlers/wordle.go)
```go
- HandleGetForm: Render main form page
- HandlePostSolve: Process constraints and return results
- Parse form inputs into wordle constraints
- Call wordle.MakePossibles()
- Render results template
```

### Data Flow
1. HTTP request → Handler
2. Handler parses form data
3. Handler calls `usrcmd.ReadUserCommand()` or similar parsing
4. Handler calls `wordle.MakePossibles()`
5. Handler renders template with results
6. Response sent to browser

## Configuration

### Environment Variables
- `WORDLE_DICTIONARY`: Path to dictionary file (required)
- `WORDLE_REMOVE`: Path to words-to-remove file (optional)
- `WORDLE_PORT`: Server port (default: 8080)
- `WORDLE_HOST`: Server host (default: localhost)

### Command Line Arguments (optional)
- `--port`: Override port
- `--host`: Override host

## Frontend Implementation

### Form Design (Bootstrap classes)
```html
- Card layout for clean appearance
- Input groups for each position
- Help text explaining the format
- Responsive grid for results
- Color coding: green for known, yellow for present, gray for absent
```

### HTMX Integration
```html
- hx-post="/wordle/solve" on form
- hx-target="#results" for in-place updates
- hx-indicator for loading state
- No JavaScript required for basic functionality
```

## Error Handling

### User Input Errors
- Invalid format → Show format help
- Invalid letters → Show valid letter set
- Empty inputs → Provide defaults

### Server Errors
- Dictionary not found → Clear error message
- File read errors → Log and display generic error
- Parse errors → Return validation messages

## Performance Considerations

1. **Dictionary Loading**: Load once at startup, keep in memory
2. **Word Filtering**: Already optimized in wordle package
3. **Concurrent Requests**: Go handles this naturally with goroutines
4. **Static Assets**: Serve with proper caching headers
5. **Results Pagination**: If word list > 500, implement pagination

## Future Enhancements

### Phase 2 Features
1. **Word Suggestions**: Highlight best next guess based on information theory
2. **Game History**: Track previous guesses in the session
3. **Statistics**: Show how many words remain after each guess
4. **Past Words Filter**: Display when a word was previously used
5. **Dark Mode**: Theme toggle

### Phase 3 Features
1. **Multiple Game Modes**: Support for Quordle, Octordle
2. **Word Patterns**: Visual pattern matching
3. **Difficulty Indicator**: Show word rarity/difficulty
4. **Mobile App**: Progressive Web App (PWA)
5. **API Mode**: JSON responses for programmatic access

## Testing Strategy

### Unit Tests
- Handler input parsing
- Form validation
- Template rendering (with test data)

### Integration Tests
- Full request/response cycle
- Dictionary loading
- Error scenarios

### Manual Testing Checklist
- [ ] Form displays correctly
- [ ] All input combinations work
- [ ] Results display properly
- [ ] Error messages are clear
- [ ] Mobile responsiveness
- [ ] Browser compatibility (Chrome, Firefox, Safari)

## Security Considerations

1. **Input Validation**: Sanitize all user inputs
2. **XSS Prevention**: HTML escape template outputs (Go does this by default)
3. **Rate Limiting**: Consider adding for public deployment
4. **HTTPS**: Use reverse proxy (nginx) for production
5. **Environment Variables**: Never expose in client code

## Deployment

### Local Development
```bash
export WORDLE_DICTIONARY=./wordle.txt
export WORDLE_REMOVE=./words-to-remove
go run cmd/server/main.go
```

### Production Deployment Options
1. **Binary**: Compile and run with systemd
2. **Docker**: Containerize with multi-stage build
3. **Cloud**: Deploy to GCP Cloud Run, AWS ECS, or Fly.io

## Success Criteria

1. ✅ User can input Wordle constraints through web form
2. ✅ Results match CLI output for same inputs
3. ✅ Response time < 100ms for typical queries
4. ✅ Works on mobile and desktop browsers
5. ✅ Clear, intuitive user interface
6. ✅ No page reloads required (HTMX)

## Timeline Estimate

- **Phase 1** (Core Functionality): 4-6 hours
  - Set up server routes and handlers
  - Create form template
  - Implement solve endpoint
  - Basic styling with Bootstrap
  
- **Phase 2** (Polish): 2-3 hours
  - Improve UI/UX
  - Add error handling
  - Responsive design
  - Testing
  
- **Phase 3** (Enhancements): As needed
  - Additional features from enhancement list

## Open Questions

1. Should we support both template-based and gomponents-based rendering?
   - **Recommendation**: Stick with html/template (already working)
   
2. Do we need user authentication?
   - **Recommendation**: No, for v1. It's a stateless tool.
   
3. Should we track analytics?
   - **Recommendation**: Optional, can add Google Analytics later.
   
4. What's the deployment target?
   - **Recommendation**: Start with local/personal use, can expand later.

## Implementation Priority

1. **Must Have** (MVP):
   - Working form input
   - Word filtering logic integration
   - Results display
   - Basic styling

2. **Should Have** (v1.0):
   - Good error handling
   - Mobile responsive
   - Loading indicators
   - Clear instructions

3. **Nice to Have** (v1.1+):
   - Game history
   - Word suggestions
   - Statistics
   - Dark mode

