# Im using windows + elvish
set shell := ["elvish.exe", "-c"]

# first item in here can be executed directly with `just` command without args
run:
    go run main.go