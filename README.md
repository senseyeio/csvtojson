# CSV-To-JSON

Reads a CSV file and emits one JSON object for each row of the input CSV.

## Flags

 - `-n` Do not use the first row of the file as headings.  Treat the first row as data.
 - `-t <col1,...>` Use the comma-separated list of values,`<col1,...>`, as column headings. Use of this flag implies `-n`.

Aside from flags, all command line arguments are treated as filenames to read.  If no files are specified, input
is read from STDIN.

All rows in an individual file must contain the same number of columns.

## Example Usage

Reading from STDIN, using the first row as column headings: 

```
$ echo -e "asset,online,mode\nrobot1,true,test\nrobot2,false,sleep" | csvtojson
{"asset":"robot1","mode":"test","online":"true"}
{"asset":"robot2","mode":"sleep","online":"false"}
```

Reading from STDIN, using the first row as data:

```
$ echo -e "robot1,true,test\nrobot2,false,sleep" | csvtojson -n
{"_column0":"robot1","_column1":"true","_column2":"test"}
{"_column0":"robot2","_column1":"false","_column2":"sleep"}
```

