# go-check-my-csv
Application to check my csv

This project help you to organize some information of your .csv <br />
You can put Personal Information each line. <br />
for example, convert this `1;2;3;4;5;6;7;8;9;10;11` to this:
```
	Name 1
	Surname 2
	Lastname 3
	Sex 4
	Age 5
	Address 6
	City 7
	Zipcode 8
	County 9
	Countrycode 10
	Tephone 11
```

Files of this project:
```
README.md         -> This readme
data.csv          -> Some data to test
main.go           -> Source code of this project
```


To use you need to send two arguments:
`go run main.go LINE FILE`

`LINE` -> What line do you want parse <br />
`FILE` -> Choose your .csv file

For example:
`go run main.go 1 data.csv` -> Will be get first line of your file.

To create an executable of your system use: `go build -o check-my-csv` <br/>

Inside `main.go` you can change the structure of Personal Information, add more data you want to get. You can change delimiter `;` to another other.
