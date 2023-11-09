
Data must always be a pipe separated string of integers and may 
come from either a file or an environment variable.

If more than 73 data points are included, only the final 73 data 
points will be displayed. This keeps the full chart within 80 characters.

Environment variable usage:

```powershell
> $env:test = "23|20|15|12|12|15|15|6|5|0"
> chart -var test
 Chart scale is 5
 4.00 ┼ ╭─╮
 3.00 ┤ │ ╰╮ ╭─╮
 2.00 ┤ │  ╰─╯ │
 1.00 ┤ │      ╰─╮
 0.00 ┼─╯        ╰
```

Data file usage:

```
> ./chart -file .\example.csv
Chart scale is 2
 7.00 ┼               ╭─
 6.00 ┤             ╭─╯
 5.00 ┤           ╭─╯
 4.00 ┤         ╭─╯
 3.00 ┤       ╭─╯
 2.00 ┤     ╭─╯
 1.00 ┤   ╭─╯
 0.00 ┼───╯
 ```

## Installation

```sh
go install github.com/edthedev/chart@latest
```
