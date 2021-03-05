# GoToDB

A hello world implementation of a bulk loader into [HelloDB](https://github.com/mrderive/HelloDB).

## Installation

Clone or download `gotodb.go`.

## Configuration

The data file has to be delimited by a character not used by any of the field values. Furthermore, the field values have to be in the same order as defined in the `HelloDB` metadata [configuration file](https://github.com/mrderive/HelloDB/blob/main/README.md#configuration).

For example, `data.txt`:
```
Joe|28|Boston
Bob|35|New York
|35|New York
Susan|35
```

## Usage

Run the script with the appropriate command line arguments:
```
go run gotodb.go <table> <file> <delimeter> <server>
```

For example:
```
go run gotodb.go customers data.txt '|' localhost:27000
```

