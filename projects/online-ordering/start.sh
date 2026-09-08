#!/bin/bash

set -e

echo "Starting Order API..."
go run main.go &
API_PID=$!

echo "Starting Outbox Publisher..."
go run ./cmd/outbox-producer &
OUTBOX_PID=$!

echo "Starting Outbox Consumer..."
go run ./cmd/outbox-consumer &
WORKER_PID=$!

cleanup() {
    echo ""
    echo "Stopping services..."

    kill $API_PID $OUTBOX_PID $WORKER_PID 2>/dev/null || true
    wait
}

trap cleanup SIGINT SIGTERM EXIT

echo "All services are running."
echo "Press Ctrl+C to stop everything."

wait