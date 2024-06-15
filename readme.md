# Nonograms

A tool for solving nonograms

### Input file

The first line contains a number x that shows the width of the nonogram grid. In next x lines, each line contains column values from top to bottom. In the remaining lines, each line contains row values from left to right.

### Flags

To give a custom file use flag -file="*filename*" (defaults to "input.txt").

To give initial values of the nonogram grid use -pins="*x:y=z*". x and y are the coordinates of the square and z is the value (1 for filled and 2 for not filled). You can give multiple pins by seperating them by semicolon: -pins="*1:2=1;1:3=1;1:4=0*".

### Examples

Examples are in [examples folder](examples/) of this repository.

### License

Copyright © 2023, pijuspie
