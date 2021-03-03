# [gobase](https://github.com/cancue/gobase)-demo

## 0. Setup

### Go
```bash
$ go build
$ go run app.go
```

## 1. Usages
### New Stage
```bash
$ vi config/{stage}.yaml
$ STAGE={stage} && go run app.go
```

### Example
```
┌──────────────────────────────────────────────┐
│ entity ─── usecase ─── controller ─── router │
│                  └──── repository            │
└──────────────────────────────────────────────┘

┌── config
├── domain
│   ├── entity
│   │   └── account
│   └── usecase
│       └── account
│           ├── account.go
│           ├── dto.go
│           └── gateway.go
├── interface
│   ├── controller
│   │   └── account
│   └── router
└── repository
    ├── account
    └── db
        ├── orm
        │   └── account
        └── migration
```
