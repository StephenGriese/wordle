# Change: Use american-english Dictionary File

## Summary
Updated the server configuration to use the `american-english` dictionary file instead of `wordle.txt`, matching the CLI's standard configuration.

## Rationale
The CLI version reads the `WORDLE_DICTIONARY` environment variable and typically uses the `american-english` file. To maintain consistency between the CLI and web server, and to follow the established project convention, the server's default configuration now uses the same dictionary file.

## Changes Made

### 1. Makefile
- ✅ Updated `run-server` target error message to reference `american-english`
- ✅ Updated `run-server-dev` target to use `./american-english` instead of `./wordle.txt`

### 2. Documentation Files Updated
All documentation has been updated to reflect the correct dictionary file:

- ✅ **SERVER_QUICKSTART.md** (3 locations)
  - Manual setup example
  - Build and run example
  - Troubleshooting section

- ✅ **IMPLEMENTATION_SUMMARY.md** (2 locations)
  - Manual start example
  - Build binary example

- ✅ **SUCCESS_SUMMARY.md** (1 location)
  - Manual setup option

- ✅ **ARCHITECTURE.md** (4 locations)
  - Data files diagram
  - Environment variables table
  - File locations tree
  - Docker example

- ✅ **ROADMAP.md** (3 locations)
  - Local development example
  - Docker example
  - Makefile example

- ✅ **DESIGN.md** (1 location)
  - Local deployment example

- ✅ **test-server.sh** (1 location)
  - Test script environment variables

### 3. Code Files
**No code changes were needed!** The server code already reads from the `WORDLE_DICTIONARY` environment variable, so it automatically uses whatever file is specified. Only the documentation and default configuration values needed updating.

## What This Means

### Before
```bash
export WORDLE_DICTIONARY=./wordle.txt
make run-server-dev
```

### After
```bash
export WORDLE_DICTIONARY=./american-english
make run-server-dev
```

Or simply:
```bash
make run-server-dev  # Now uses american-english by default
```

## Backward Compatibility
✅ **Fully backward compatible!** Users can still use any dictionary file they want by setting the environment variable:

```bash
# Can still use wordle.txt if desired
export WORDLE_DICTIONARY=./wordle.txt
go run cmd/server/main.go

# Or any other dictionary
export WORDLE_DICTIONARY=/path/to/my/custom-dictionary.txt
go run cmd/server/main.go
```

The change only affects the **default** value used by the `make run-server-dev` convenience target.

## Files Modified
- ✅ Makefile
- ✅ SERVER_QUICKSTART.md
- ✅ IMPLEMENTATION_SUMMARY.md
- ✅ SUCCESS_SUMMARY.md
- ✅ ARCHITECTURE.md
- ✅ ROADMAP.md
- ✅ DESIGN.md
- ✅ test-server.sh

## Verification
✅ Code compiles successfully  
✅ No runtime code changes needed  
✅ All documentation consistent  
✅ Makefile targets updated  
✅ Test script updated  

## Usage

Now when you run:
```bash
make run-server-dev
```

The server will use the `american-english` dictionary file, matching the CLI behavior and using the comprehensive word list that's included in the project.

This provides consistency across both interfaces (CLI and web) and uses the more complete dictionary that comes directly from Linux Mint.

