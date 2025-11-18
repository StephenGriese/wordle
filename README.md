A program to help play wordle

This project provides three ways to use the Wordle helper:
1. **Library** - Core logic to eliminate words based on Wordle game results
2. **Command line program** - Terminal-based interface (see below)
3. **Web application** - Browser-based interface with modern UI (NEW! See [SERVER_QUICKSTART.md](SERVER_QUICKSTART.md))

## 🚀 Quick Start - Web Server

To run the web interface:

```bash
make run-server-dev
```

Then open http://localhost:8080 in your browser. You'll get a clean, Wordle-themed interface where you can:
- Enter missed letters (gray squares)
- Mark correct positions (green squares)
- Mark wrong positions (yellow squares)
- See all possible words instantly

For detailed setup and usage instructions, see [SERVER_QUICKSTART.md](SERVER_QUICKSTART.md).

## How the Library Works

So the library will follow this basic algorithm:

1. Take a list of words
2. For each word, decide whether it is a possible solution
3. To determine if a word is a possible solution perform the following steps:
   1. Check if the word contains any of the missed letters. If so, eliminate it
   2. Check if, at any position, the word contains a letter that is not in the correct position. If so, eliminate it.
   3. If, at any position, the word does not contain a letter that we know is there, eliminate it. 
   
I'm not sure which order of checks is best. Testing will determine. 

The american-english word list comes directly from Linux Mint, and is what I started with when I first started playing wordle.
It has not been modified in any way.