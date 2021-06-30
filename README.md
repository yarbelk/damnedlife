DamnedLife
==========
[![Go Report Card](https://goreportcard.com/badge/github.com/yarbelk/damnedlife)](https://goreportcard.com/report/github.com/yarbelk/damnedlife)

An implementation of Conways game of life, in curses.  for fun.


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


TODO
====


 - [ ] World state loading from standard formats.  Or at least
       [life 1.06](http://psoup.math.wisc.edu/mcell/ca_files_formats.html)
       formats.

 - [ ] Add an editor?

 - [ ] Good support for unicode... in ncurses...
       (I think I should have ncursesw?)
