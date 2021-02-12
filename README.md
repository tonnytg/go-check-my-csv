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

Files of this project:
```
README.md         -> This readme
check-my-csv      -> Linux binary create by go build
check-my-csv.exe  -> Windows exec create by go build on Windows
data.csv          -> Some data to test
main.go           -> Source code of this project
```


To use you need to send two arguments:
`go run main.go LINE FILE`

For example:
`go run main.go 1 data.csv`

To create an executable of your system use: `go build -o check-my-csv` <br/>
Then execute on linux `./check-my-csv 1 data.csv` <br/>
or execute on Windows `check-my-csv.exe 1 data.csv` 


This project can help your to struc each column
