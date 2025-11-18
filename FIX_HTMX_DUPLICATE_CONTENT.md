# Fix: HTMX Results Display Issue

## Problem
After the first form submission, the results section was showing the entire page header and "How to use" instructions again within the results area, causing duplicate content.

## Root Cause
The POST handler (`HandlePostSolve`) was rendering the full page template (`wordleform.gohtml`) which includes the header and instructions. When HTMX replaced the `#results-section` div, it was injecting the entire page structure into that section.

## Solution
Implemented a partial template approach that detects HTMX requests and renders only the results portion:

### Changes Made

#### 1. Created New Results Template (`web/views/results.gohtml`)
- Extracted just the results display logic into a separate template
- Contains only the results card with word list and count
- No header, no instructions, no form - just results

#### 2. Updated Handler (`handlers/wordle.go`)
- Modified `HandlePostSolve` signature to accept both `view` and `resultsView` parameters
- Added HTMX detection by checking `HX-Request` header
- Renders full page for regular POST (fallback for non-HTMX browsers)
- Renders only results partial for HTMX requests

```go
// Check if this is an HTMX request
isHTMX := r.Header.Get("HX-Request") == "true"

var renderView *views.View
if isHTMX {
    renderView = resultsView // Render just the results partial
} else {
    renderView = view // Render full page (for non-HTMX fallback)
}
```

#### 3. Updated Server Main (`cmd/server/main.go`)
- Created `resultsView` in addition to `formView`
- Passed both views to `HandlePostSolve`

```go
formView := views.NewView("bootstrap", "web/views/wordleform.gohtml")
resultsView := views.NewView("results", "web/views/results.gohtml")
```

### How It Works Now

1. **Initial Page Load (GET /)**
   - Renders full page with header, instructions, form, and empty results section

2. **Form Submission (POST /wordle/solve via HTMX)**
   - HTMX sends request with `HX-Request: true` header
   - Handler detects HTMX request
   - Renders only the results template (`results.gohtml`)
   - HTMX swaps content into `#results-section` div
   - No duplicate headers or instructions!

3. **Fallback (POST /wordle/solve without HTMX)**
   - If JavaScript disabled or HTMX fails
   - Handler renders full page template
   - Page reloads with results (graceful degradation)

### Benefits
✅ Clean HTMX updates - only results change  
✅ No duplicate content  
✅ Progressive enhancement - works without JavaScript  
✅ Better separation of concerns  
✅ More maintainable template structure  

### Testing
```bash
# Start server
make run-server-dev

# Test in browser
# 1. Load page - should see form and instructions once
# 2. Submit form - should see results appear below form
# 3. Submit again - results update in place, no duplicate headers
```

## Files Modified
1. **handlers/wordle.go** - Updated `HandlePostSolve` function signature and logic
2. **cmd/server/main.go** - Added resultsView creation
3. **web/views/results.gohtml** - NEW: Partial template for results only

## Files Unchanged
- **web/views/wordleform.gohtml** - Still contains full form (no changes needed)
- **web/views/layouts/bootstrap.gohtml** - Layout unchanged

## Verification
✅ Code compiles successfully  
✅ No errors detected  
✅ HTMX detection implemented correctly  
✅ Fallback behavior maintained  

The issue is now fixed! When you submit the form, only the results section will update without showing duplicate headers or instructions.

