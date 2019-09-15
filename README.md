# Code for a board navigation coding test
Company name omitted to make it a bit harder for people to hit this by searching Github for pre-made solutions.

## About the repository name
Just to push away the idea that's a hint on how amazing the (were)goat is by providing this solution.

Of course the bad pun on "maze/amazing" is there, but it's about a goat going "a' mazing" (which fits kind of nicely on how I looked at the problem).

 _Amazingoat_ sounded nicer to me than _goatamazing_(also there was a nagging association with _Amazing Grace_ popping in my head at the time).

## The board
The board is **always** a matrix of rows and columns.

To allow for different shapes the single cells may be
 * in/on the board (true) 
 * or not (false), like holes
 
Example:

    |  *****  |
    |*  ***  *|
    |**  *  **|
    |***   ***|
  
Thus, making a "V" maze (* = holes)

I came to this view because I first looked at the problem like a mix between _Logo_ (hence the goat as cursor) and a roguelike game (_Moria_, _Nethack_, _Angband_...).

Later, when I started writing this, I saw that, possibly, a more "test oriented" solution would have been to use double linking. Matter of fact I don't remember ever using double linked lists in any of my jobs, but surely it would give the impression that I know some CS stuff.

Still it's too late to change, and I like my solution alright; especially the goat part.  

## Binary things
I decided to discharge the binary constraints except for input-output. The casting and encoding was just wearisome and I don't think it added much, also I believe it's more flexible this way.

Given that Golang int is 32 bits signed, the unsigned int16 and int8 should all fit in.

## Errors
It was not clear how to handle errors. -1,-1 seems to be the solution **only** for the cursor exiting the board, so I decided to just exit (with a log error message) in case of errors (like a badly formed JSON, for example).

## Golang part
It's a simple Golang program. I didn't use any package outside the standard library and no modules.
`go run`, `go build`, `go test` as usual.

## Example
```
echo '{"Width":4,"Height":4, "X":2, "Y":2, "Commands": ["f", "l", "f", "r", "b", "r", "b", "l", "f", "q"]}' | ./amazingoat -json -debug
2019/09/15 09:10:08 goat at [2,2] facing north
2019/09/15 09:10:08 commands to process: [f l f r b r b l f q]
2019/09/15 09:10:08 processing command f
2019/09/15 09:10:08 goat at [2,1] facing north
2019/09/15 09:10:08 processing command l
2019/09/15 09:10:08 goat at [2,1] facing west
2019/09/15 09:10:08 processing command f
2019/09/15 09:10:08 goat at [1,1] facing west
2019/09/15 09:10:08 processing command r
2019/09/15 09:10:08 goat at [1,1] facing north
2019/09/15 09:10:08 processing command b
2019/09/15 09:10:08 goat at [1,2] facing north
2019/09/15 09:10:08 processing command r
2019/09/15 09:10:08 goat at [1,2] facing east
2019/09/15 09:10:08 processing command b
2019/09/15 09:10:08 goat at [0,2] facing east
2019/09/15 09:10:08 processing command l
2019/09/15 09:10:08 goat at [0,2] facing north
2019/09/15 09:10:08 processing command f
2019/09/15 09:10:08 goat at [0,1] facing north
2019/09/15 09:10:08 processing command q
{"X":0,"Y":1}
```