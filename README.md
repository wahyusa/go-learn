# About

This is simple CLI App to show data based on `os.Args` inserted when running.

## Installation

`git clone https://github.com/wahyusa/go-learn.git`

## Usage

```bash
go run main.go // Err msg need argument

go run main.go 0 // Works fine as expected

go run main.go 9999 // Err msg out of range

go run main.go abc // Err args is not number
```

