# pomodoro

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
