#!/bin/bash

# Test script for Wordle server
cd /Users/sgries174@cable.comcast.com/repos/sjg/wordle

echo "=== Building server ==="
go build -o wordle-server cmd/server/main.go
if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
fi
echo "Build successful!"

echo ""
echo "=== Starting server in background ==="
export WORDLE_DICTIONARY=./american-english
export WORDLE_REMOVE=./words-to-remove
export WORDLE_PORT=9090

./wordle-server > server.log 2>&1 &
SERVER_PID=$!
echo "Server started with PID: $SERVER_PID"

echo ""
echo "=== Waiting for server to start ==="
sleep 3

echo ""
echo "=== Server log output ==="
cat server.log

echo ""
echo "=== Testing server ==="
curl -s http://localhost:9090/ | head -30

echo ""
echo "=== Stopping server ==="
kill $SERVER_PID 2>/dev/null

echo "Done!"

