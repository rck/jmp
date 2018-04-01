# jmp
`jmp`, or better its alias `j`, can be used to efficiently navigate directories. It learns your most used
directories while you are using it. It adds a weight to every directory and uses them to select the most
likely one. Every time you use `j`, the weight of the destination directory gets increased.

## Usage
[![asciicast](https://asciinema.org/a/Kgn1Sr0RPmKdmNiEtwdPEKiYS.png)](https://asciinema.org/a/Kgn1Sr0RPmKdmNiEtwdPEKiYS)

### Adding a directory
```
cd directory
j .
```

### Displaying weights:
`j -l`

### Deleting a directory
```
cd directory
j -s -1 # set its weight to any negative value
```

### Setting a weight
```
cd directory
j -s 23
```

## Algorithm
So far I don't use any fancy fuzzy matching, all arguments are joined in a simple regex. So far that is good
enough. Let's see.

Weights are limited to `int64`. A weight is capped at the maximum value. If a weight reaches this limit, `j`
tries to normalize the DB.

## This sounds like jump or autojump
Yes, this is not a new idea at all. So what is wrong with these? Nothing, but:

`jump` is used to add "bookmarks" and does not learn. I always forgot how to add new bookmarks and the ones I
had were ugly, like "utils" for my development copy of "drbd-utils", and "usutils", for the "upstream
version.". I remember parts of a path better than bookmarks.

`autojump` is closer to what `jmp` does, but for my taste it has too much "auto magic". If overwrites various
shell commands to learn every directory you `cd` into.

`jmp` is somewhere between these tools, it has "magic", but you as the user control which directories are in
the database.

## Notes
Thanks to `autojump` for the idea and for the completion files, which I shamelessly copied.

Strictly speaking, the DB should use some `flock`-ing, but as long as I have only 10 fingers and 1 brain,
accesses serializes themselves.
