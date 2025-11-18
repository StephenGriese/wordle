# Feature: Dictionary Reload

## Overview
The Wordle Helper now supports dynamic dictionary reloading to keep the word list up-to-date as new Wordle games are played daily.

## Problem Solved
Previously, the dictionary was loaded once at server startup and included a list of "past words" (already used by NYTimes Wordle) to filter them out. Since new Wordle games are released daily, the past words list becomes stale, and the server would need to be restarted to get the updated list.

## Solution Implemented

### 1. Thread-Safe WordList Manager
Created a new `WordList` type in `dictionary/wordlist.go` that:
- ✅ Manages the word list with thread-safe read/write access
- ✅ Tracks last reload time
- ✅ Provides reload capability without server restart
- ✅ Uses mutex locks to prevent race conditions

### 2. Manual Reload Endpoint
Added `POST /reload` endpoint that:
- ✅ Triggers dictionary reload on demand
- ✅ Fetches latest past words from NYTimes
- ✅ Returns success message with word count
- ✅ Logs reload events for monitoring

### 3. UI Reload Button
Added a reload button in the header that:
- ✅ Triggers reload via HTMX (no page refresh)
- ✅ Shows loading spinner during reload
- ✅ Displays success/error messages
- ✅ Auto-dismissible alerts
- ✅ Mobile responsive

## Files Created/Modified

### New Files
1. **dictionary/wordlist.go** (85 lines)
   - `WordList` struct with mutex for thread safety
   - `NewWordList()` - Constructor with initial load
   - `Reload()` - Refresh dictionary from sources
   - `Words()` - Get current word list (thread-safe)
   - `LastReload()` - Get last reload timestamp
   - `Count()` - Get word count
   - `ShouldReload()` - Check if reload needed based on age

### Modified Files
1. **cmd/server/main.go**
   - Changed from static `[]string` to `*dictionary.WordList`
   - Added `POST /reload` route
   - Passes WordList to handlers

2. **handlers/wordle.go**
   - Added `WordList` interface
   - Updated `HandlePostSolve()` to accept WordList interface
   - Added `HandleReload()` handler
   - Gets words via `wordList.Words()` for each request

3. **web/views/layouts/bootstrap.gohtml**
   - Added reload button with HTMX
   - Added CSS for reload button and messages
   - Added Bootstrap JS for dismissible alerts
   - Added loading spinner

## How It Works

### Startup Flow
```
Server Start
    ↓
Create WordList
    ↓
Load dictionary file (american-english)
    ↓
Fetch past words from NYTimes
    ↓
Remove past words from dictionary
    ↓
Remove custom exclusions (words-to-remove)
    ↓
Store words + timestamp
    ↓
Server ready
```

### Reload Flow
```
User clicks "🔄 Reload Word List" button
    ↓
HTMX sends POST /reload
    ↓
Handler calls wordList.Reload()
    ↓
WordList acquires write lock
    ↓
Reload dictionary (same process as startup)
    ↓
Update words + timestamp
    ↓
Release lock
    ↓
Return success HTML
    ↓
HTMX displays success message
    ↓
Auto-dismiss after user sees it
```

### Thread Safety
When multiple requests happen simultaneously:
- **Reading words** (GET /, POST /wordle/solve): Uses read lock, multiple concurrent reads allowed
- **Reloading** (POST /reload): Uses write lock, exclusive access, blocks other operations temporarily
- **Reload duration**: ~1-2 seconds (fetching from web + processing)

## Usage

### Manual Reload (Web UI)
1. Click the "🔄 Reload Word List" button in the header
2. Wait for success message (1-2 seconds)
3. Continue using the app with updated word list

### Programmatic Reload (API)
```bash
# Trigger reload via curl
curl -X POST http://localhost:8080/reload

# Response HTML (for HTMX)
<div class="alert alert-success...">
  ✅ Dictionary reloaded! 4328 words available (Updated: 2:45 PM)
</div>
```

### Automatic Reload (Future Enhancement)
Currently manual only. Future options:
- Daily scheduled reload at midnight
- Auto-reload on first request each day
- Webhook trigger from external scheduler

## Benefits

### For Users
- ✅ **Always up-to-date** - Get latest past words without server restart
- ✅ **One-click refresh** - Simple button in UI
- ✅ **No downtime** - Reload happens in ~1 second
- ✅ **Visual feedback** - See when dictionary was last updated

### For Operations
- ✅ **No restarts needed** - Update data without downtime
- ✅ **Logged events** - Track reload times in logs
- ✅ **Thread-safe** - Safe for concurrent requests
- ✅ **Low latency** - Read operations unaffected

### For Development
- ✅ **Testable** - WordList interface allows mocking
- ✅ **Extensible** - Easy to add auto-reload schedules
- ✅ **Monitored** - Last reload time tracked
- ✅ **Clean architecture** - Separation of concerns

## API Reference

### POST /reload
Reloads the dictionary from disk and fetches latest past words.

**Request:**
```http
POST /reload HTTP/1.1
Host: localhost:8080
```

**Response (Success):**
```html
HTTP/1.1 200 OK
Content-Type: text/html

<div class="alert alert-success alert-dismissible fade show mt-2" role="alert">
  ✅ Dictionary reloaded! 4328 words available (Updated: 2:45 PM)
  <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
</div>
```

**Response (Error):**
```html
HTTP/1.1 500 Internal Server Error
Content-Type: text/html

<div class="alert alert-danger alert-dismissible fade show mt-2" role="alert">
  <strong>Error:</strong> Failed to reload dictionary: connection timeout
  <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
</div>
```

**Logs:**
```
time=2025-11-18T14:30:00.000-05:00 level=INFO msg="Reloading dictionary"
Reloading dictionary from ./american-english
Loaded 5905 words from ./american-english
Loaded 1613 past words
...
Dictionary reloaded: 4328 words available (last reload: 2025-11-18T14:30:01-05:00)
time=2025-11-18T14:30:01.123-05:00 level=INFO msg="Dictionary reloaded successfully" word_count=4328
```

## Configuration

### Environment Variables
No new environment variables needed. Uses existing:
- `WORDLE_DICTIONARY` - Dictionary file path
- `WORDLE_REMOVE` - Custom exclusions file path

### Future Configuration Options
Could add (not implemented yet):
- `WORDLE_AUTO_RELOAD=true` - Enable automatic daily reload
- `WORDLE_RELOAD_HOUR=0` - Hour to auto-reload (0-23)
- `WORDLE_MAX_AGE=24h` - Max age before reload needed

## Testing

### Manual Testing
1. Start server: `make run-server-dev`
2. Open browser: http://localhost:8080
3. Submit a query, note word count
4. Click "🔄 Reload Word List"
5. Verify success message appears
6. Submit same query, verify it still works

### Concurrent Testing
```bash
# Terminal 1: Start server
make run-server-dev

# Terminal 2: Send concurrent requests
for i in {1..5}; do
  curl -X POST http://localhost:8080/reload &
done
wait

# All should succeed without race conditions
```

### Error Testing
```bash
# Test with invalid dictionary path
export WORDLE_DICTIONARY=/nonexistent/file
make run-server-dev
# Should fail at startup (not silently)

# Test reload with network issues (requires mock)
# Disconnect network, click reload
# Should show error message
```

## Performance Impact

### Memory
- Before: `[]string` slice (~500KB for 4000 words)
- After: `WordList` struct with mutex (~500KB + 48 bytes overhead)
- **Impact: Negligible**

### CPU
- Read operations: No change (direct slice access)
- Reload operation: ~1-2 seconds once manually triggered
- **Impact: Negligible (only during reload)**

### Latency
- Normal requests: No added latency
- During reload: Requests wait for write lock (~1-2 sec max)
- **Impact: Minimal (rare reload operation)**

## Future Enhancements

### Automatic Daily Reload
```go
// In main.go, add ticker
go func() {
    ticker := time.NewTicker(24 * time.Hour)
    for range ticker.C {
        if wordList.ShouldReload(24 * time.Hour) {
            logger.Info("Auto-reloading dictionary")
            wordList.Reload()
        }
    }
}()
```

### Reload on First Request of Day
```go
// In HandlePostSolve, before getting words
if wordList.ShouldReload(24 * time.Hour) {
    logger.Info("Auto-reloading stale dictionary")
    go wordList.Reload() // Async, don't block user
}
```

### Metrics/Monitoring
```go
// Track reload statistics
type ReloadMetrics struct {
    TotalReloads int
    LastDuration time.Duration
    AvgDuration  time.Duration
}
```

### Webhook Notifications
```go
// Notify external service on reload
func (wl *WordList) Reload() error {
    // ... existing reload logic ...
    
    // Send webhook
    notifyWebhook("dictionary_reloaded", wl.Count())
}
```

## Troubleshooting

### "Failed to reload dictionary"
**Cause:** Network error fetching past words or file read error

**Solution:**
1. Check internet connection
2. Verify dictionary file exists
3. Check logs for specific error
4. Try again (transient network issues)

### Reload button doesn't work
**Cause:** JavaScript/HTMX not loaded

**Solution:**
1. Check browser console for errors
2. Verify HTMX script loaded
3. Check Bootstrap JS loaded (for alerts)
4. Try hard refresh (Ctrl+F5)

### Word list not updating
**Cause:** Reload succeeded but using cached results

**Solution:**
1. Hard refresh browser page
2. Clear browser cache
3. Check logs to verify reload actually happened
4. Restart browser

## Summary

The dictionary reload feature provides:
- ✅ **No downtime updates** - Reload without restart
- ✅ **Thread-safe operation** - Safe concurrent access
- ✅ **User-friendly UI** - One-click button with feedback
- ✅ **Production ready** - Proper error handling and logging
- ✅ **Extensible design** - Easy to add auto-reload later

The implementation is clean, performant, and ready for production use!

