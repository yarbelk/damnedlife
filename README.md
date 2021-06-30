DamnedLife
==========
[![Go Report Card](https://goreportcard.com/badge/github.com/yarbelk/damnedlife)](https://goreportcard.com/report/github.com/yarbelk/damnedlife)


An implementation of Conways game of life, in curses.  for fun.
I built this to get across ncurses in golang, because it had been a decade since I last used
ncurses.  Since its always kinda fun to implement Conways Game of Life; well... This happend

I poke at it every now and then and add features or check that it builds.  Recently
theres an issue with cgo (see below)

Since this is for fun: its licend under the _most_ open licence I could find.

BUILDING
--------

due to a limitation with `cgo` and ncurses libraries,
try

```
export CGO_CFLAGS_ALLOW=".*"
export CGO_LDFLAGS_ALLOW=".*"
go get github.com/rthornton128/goncurses
export CGO_CFLAGS_ALLOW=
export CGO_LDFLAGS_ALLOW=
```

before building if you have an error with unknown flags.

see https://github.com/rthornton128/goncurses/issues/55


Contributing
------------

I'm not going to question your sanity; I mean, this is a project I started.
I'm 100% open to contributions.  I'd like to get some benchmarks up,
implement world state loading so you don't need to recompile for
new start conditions, and decouple the ncurses from the world simulation better.

Once they are decoupled, you can run the simulation as fast as you want
and play it as slow as you want.  And then add controls to it.

```
Simulated generation: 9001
Current Visualised Generation: 53
[Play][Pause][>>][<<]
[Pause Simulation]
```

TODO
====


 - [ ] World state loading from standard formats.  Or at least
       [life 1.06](http://psoup.math.wisc.edu/mcell/ca_files_formats.html)
       formats.

 - [ ] Add an editor?

 - [ ] Good support for unicode... in ncurses...
       (I think I should have ncursesw?)

   [ ] benchmarks

   [ ] decouple simulation from visualizaiton

   [ ] play/pause/speed controls

