# Account CLI

This is a demo Account CLI


## Run

Make sure `account` , `emailer`, `greeter` services are **UP** before running `Account CLI`

Check services are **UP** with `micro list services` command

1. Start Services
    ```bash
    make run-emailer
    make run-greeter
    make run-account
    ```

2. Run Account CLI
    ```bash
    go run cmd/account/main.go
    go run cmd/account/main.go -username=crazy -email=a@b.com -limit=7
    ```
