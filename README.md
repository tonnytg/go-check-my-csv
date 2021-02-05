# go-check-my-csv
Application to check my csv

This project help you to organize some information of your .csv
for example, convert this `1,2,3,4,5` to this:
```
  line1: 1
  line1: 2
  line1: 3
  line1: 4
  line1: 5
```

To use you need send two arguments:
`go run main.go LINE FILE`

For example:
`go run main.go 1 data.csv`

To create executable of your system use: `go build -o NAME`
Then execute on linux `./NAME 1 data.csv`
or execute on Windows `NAME.exe 1 data.csv`