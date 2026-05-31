# Budget

A budget tracking app in Go with three interface modes.

## Structure

```
cmd/budget/main.go          # launcher entry point
internal/
├── budget/budget.go        # CLI business logic (payday, trends)
├── gui/gui.go              # GUI placeholder
├── models/models.go        # domain types
├── storage/storage.go      # JSON persistence
└── tui/                    # Terminal UI (Bubbletea)
    ├── tui.go              # model, routing, helpers
    ├── menu.go             # main menu screen
    ├── payday.go           # payday flow screens
    └── trends.go           # trends screen
```

## Usage

```
go run ./cmd/budget/
```

### Launcher

```
1. Simple CLI    — original terminal prompt flow
2. Terminal UI   — Bubbletea TUI
3. GUI           — (coming soon)
4. Exit
```

Both CLI and TUI share the same data file (`budget_data.json`).
