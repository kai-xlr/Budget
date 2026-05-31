# Budget

A simple CLI budget app written in Go.

## Structure

```
cmd/budget/main.go        # entry point
internal/
├── budget/budget.go      # business logic (payday, trends)
├── models/models.go      # domain types
└── storage/storage.go    # JSON persistence
```

## Usage

```
go run ./cmd/budget/
```

### Menu

1. **Payday Mode** — Enter your paycheck and expenses. Results are saved for trends.
2. **Trends Mode** — View all past payday sessions.
3. **Exit**
