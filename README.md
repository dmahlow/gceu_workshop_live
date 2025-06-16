# Desktop Automation CLI

Command-line tool for desktop automation using robotgo.

## Installation

```bash
go install github.com/dmahlow/desktop-automation/cmd/desktop-automation@latest
```

## Build from Source

```bash
git clone https://github.com/dmahlow/desktop-automation.git
cd desktop-automation
task build
```

## Commands

### Move Mouse

```bash
# Instant movement
desktop-automation move 800 600

# Smooth movement with default 1 second duration
desktop-automation move --smooth 800 600

# Smooth movement with custom duration
desktop-automation move --smooth --duration 5.0 800 600
```

### Click

```bash
desktop-automation click 100 200
```

### Type Text

```bash
# Type text instantly
desktop-automation type "Hello World"

# Type with delay between characters
desktop-automation type --delay 50 "Slow typing"
```

## Requirements

- Go 1.23+
- Platform-specific dependencies for robotgo

### macOS

Requires accessibility permissions. Grant access in System Preferences > Security & Privacy > Privacy > Accessibility.

## Development

```bash
task build    # Build binary
task run      # Build and run
task clean    # Clean artifacts
task deps     # Download dependencies
task test     # Run tests
task install  # Install to GOPATH/bin
```
