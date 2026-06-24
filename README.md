# pomodoro-go

Simple Pomodoro timer that sends desktop notifications when it's time to work or take a break.

## Usage

```text
pomodoro [flags] work_duration pause_duration
```

### Flags

```text
-defaultIconPath string
    default notification icon

-pauseIconPath string
    break notification icon

-pauseText string
    break notification text
    (default "Time for a break!")

-workIconPath string
    work notification icon

-workText string
    work notification text
    (default "Back to work!")
```

### Examples

```bash
pomodoro 25 5
pomodoro 25m 5m
pomodoro 45m 10m
pomodoro 1h 15m
pomodoro 90s 15s
```

Durations support Go's duration format (`90s`, `25m`, `1h`). If no unit is specified, minutes are assumed.

## Installation

### Run with Nix
```bash
nix run github:jim-ww/pomodoro-go 25 5
```
Or add it to your flake.nix.
### Pre-built binaries

Check the [Releases](https://github.com/jim-ww/pomodoro-go/releases) page for static binaries for Linux, macOS, and Windows.

### Build from source

```bash
git clone https://github.com/jim-ww/pomodoro-go.git
cd pomodoro-go
go build -o pomodoro .
sudo mv pomodoro /usr/local/bin/
```

Or install directly:

```bash
go install github.com/jim-ww/pomodoro-go@latest
```

## Donate
If you find this tool useful, consider a small donation:

**Monero (XMR)**
`83YGRqP8uHed6NeegZQeX9ccCxbzoRHHEEi7pTwk4aqdJZEVXXA6NWtetnsEM2v33zFBBt3Rp6DNhU9qhJEGPspU14yN8t7`

## License Notice

This program is free software licensed under the **GNU General Public License v3 (GPLv3)**.

It means this is **free software** — you are free to use, study, share, and modify it however you like (as long as you keep the same freedoms for others).
