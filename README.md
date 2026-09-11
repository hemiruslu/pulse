# pulse

A fast, concurrent website health and latency checker written in Go.

## Overview
`pulse` inspects a list of network endpoints simultaneously without sequential blocking overhead. By leveraging Go's lightweight runtime concurrency primitives, it minimizes total execution time to the duration of the slowest single request.

## Features
- **Concurrent Execution:** Dispatches non-blocking goroutines for concurrent host validation.
- **Thread-safe Communication:** Aggregates status payloads through Go channels.
- **Latency Measurement:** Tracks round-trip request duration.
- **Fault-Tolerant:** Handles network drops and timeouts gracefully.

## Getting Started

### Prerequisites
- [Go 1.20+](https://go.dev/dl/)

### Installation & Run

```bash
# Clone the repository
git clone [https://github.com/hemiruslu/pulse.git](https://github.com/hemiruslu/pulse.git)

# Navigate to project directory
cd pulse

# Run the tool
go run main.go