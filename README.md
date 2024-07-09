A program to help play wordle

At this point I'm looking at creating two things:
1. A library that can be used to eliminate words from a list based on the results of a wordle turn
2. A command line program that uses this library
3. A web app that uses this library

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