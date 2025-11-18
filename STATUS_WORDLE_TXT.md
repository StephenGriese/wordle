# Status: wordle.txt File

## Quick Answer
**No, `wordle.txt` is NO LONGER USED by the application.**

## Current Status

### What the App Uses Now
- ✅ **`american-english`** - Primary dictionary (5905 words from Linux Mint)
- ✅ **`words-to-remove`** - Custom exclusion list (14 words)
- ✅ **NYTimes past words** - Fetched from web automatically (1600+ words)

### What wordle.txt Is
- A separate word list file that exists in the repository
- Was potentially used earlier in development
- **Currently not referenced by any code**
- Could be used if user sets `WORDLE_DICTIONARY=./wordle.txt`

## Code References

### No Active Code References
The app **never hardcodes** `wordle.txt` anywhere. All references found:

1. **Documentation only** (CHANGE_USE_AMERICAN_ENGLISH.md)
   - Historical note about switching FROM wordle.txt TO american-english
   - Example showing users CAN use it if they want

2. **One outdated diagram** (ARCHITECTURE.md)
   - Shows `wordle.txt` in file tree
   - Should be updated to `american-english`

### How Dictionary is Selected
```go
// cmd/server/main.go
dict := os.Getenv("WORDLE_DICTIONARY")
// Uses whatever file you specify in environment variable
```

Current defaults:
- **Makefile**: Uses `./american-english`
- **Documentation**: Uses `./american-english`
- **Server**: Uses whatever `WORDLE_DICTIONARY` points to

## Can You Delete wordle.txt?

### Safe to Delete? 
**YES** - if you don't plan to use it

### Should You Delete It?
**Your choice:**

**Option 1: Keep it**
- Users could use it as alternative dictionary
- Some users might prefer it
- Zero harm in keeping it
- Takes minimal space

**Option 2: Delete it**
- Cleaner repository
- Removes confusion
- App doesn't need it
- Can always restore from git history

**Option 3: Document it**
- Keep it but add README note
- Explain it's alternative dictionary
- Show how to use it

## If You Want to Use wordle.txt

You CAN still use it anytime by setting:

```bash
# Local development
export WORDLE_DICTIONARY=./wordle.txt
make run-server

# Heroku deployment
heroku config:set WORDLE_DICTIONARY=./wordle.txt
```

The app is **flexible** - it will use whatever dictionary file you specify.

## Recommendation

### For Your Repository
1. **Update ARCHITECTURE.md** to show `american-english` instead of `wordle.txt`
2. **Either:**
   - Delete `wordle.txt` (cleaner)
   - OR keep it and add note in README: "Alternative dictionary available in wordle.txt"

### Current Setup is Good
Your current configuration using `american-english` is the right choice because:
- ✅ Matches CLI behavior
- ✅ Linux Mint standard word list
- ✅ Comprehensive (5905 words)
- ✅ Well-maintained source

## Files to Update

### 1. ARCHITECTURE.md (Line 80)
**Current:**
```
│  • wordle.txt (dictionary)                                   │
```

**Should be:**
```
│  • american-english (dictionary)                             │
```

### 2. Optionally: Add to README.md
```markdown
## Dictionary Files

- **american-english** - Primary dictionary (5905 words from Linux Mint)
- **words-to-remove** - Custom exclusion list
- **wordle.txt** - Alternative dictionary (available but not used by default)
```

## Summary

- ❌ **Not used** by default configuration
- ❌ **Not referenced** in any code
- ❌ **Not needed** for app to work
- ✅ **Can be used** if user explicitly sets WORDLE_DICTIONARY
- ✅ **Safe to delete** if you don't need it
- ✅ **american-english is now the standard**

**Bottom line:** The app uses `american-english` now. The `wordle.txt` file is just sitting there unused (unless you explicitly configure it).

