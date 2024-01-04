# mower-project


## Getting started

# Description

The MowItNow company decided to develop an automatic lawn mower, intended for
rectangular surfaces.
The mower can be programmed to cover the entire surface.
The position of the mower is represented by a combination of coordinates (x,y) and a
letter indicating the orientation according to English cardinal notation (N,E,W,S). The lawn is
divided into a grid to simplify navigation.
For example, the position of the mower could be "0, 0, N", which means it is located
in the lower left corner of the lawn, and oriented towards the North.
To control the mower, we send it a simple sequence of letters. Possible letters
are “R”, “L” and “F”. “R” and “L” rotate the mower 90° to the right or left
respectively, without moving it. “F” means that the mower is moved forward one space in the
direction it faces, and without changing its orientation.
If the position after movement is outside the lawn, the mower does not move,
maintains its orientation and processes the next command.
We assume that the box directly north of position (x, y) has coordinates (x,
y+1).
To program the mower, we provide it with an input file constructed as follows:

- The first line corresponds to the coordinates of the upper right corner of the lawn, those
of the lower left corner are assumed to be (0,0)
- The rest of the file allows you to control all the mowers that have been deployed. Each
mower has two lines about it:
- the first line gives the initial position of the mower, as well as its orientation. There
position and orientation are provided as 2 numbers and a letter, separated
by a space
- the second line is a series of instructions instructing the mower to explore the
lawn. Instructions are a sequence of characters without spaces.

Each mower moves sequentially, meaning the second mower does not
only moves when the first has completely executed its series of instructions.
When a mower completes a series of instructions, it communicates its position and its
orientation.

## To run the program

The following commands will run the program:
```
cd mower-project
go mod tidy
go run main.go
```

## Test the program

2 tests files are implemented to insure the program functionality.

Use the following command to run the unit tests :

```
go test -v ./...
```
***

# Details
## Authors and acknowledgment
Joseph EL BACHA 

## License
it's an Open source project.

## Project status
Test project.

## TODO list
No TODO list points for now.