<p align="center">
  <h3 align="center">DolphinSay</h3>
  <p align="center">
    <a href="https://github.com/scraly/gophersay/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/scraly/gophersay.svg?logo=github&style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/scraly/gophersay"><img src="https://goreportcard.com/badge/github.com/scraly/gophersay" alt="Code Status" /></a>
  </p>
</p>

# About

Welcome in DolphinSay!

DolphinSay is inspired by Cowsay program.

DolphinSay allow you to display a message said by a cute random Dolphin.

# Installation

For MacOS:

```
brew tap scraly/tools
brew install gophersay
```

# Pre-requisites

Install Go in 1.16 version minimum.

# Build the app

`$ go build -o bin/dolphinsay main.go`

or

`$ task build`

# Run the app

`$ ./bin/dolphinsay`

or

`$ task run`

# Test the app

```
$ ./bin/dolphinsay Hello Dolphin lovers!
 ----------------------
< Hello Dolphin lovers! >
 ----------------------
        \
         \
                                    _
                               _.-~~.)
         _.--~~~~~---....__  .' . .,'
       ,'. . . . . . . . . .~- ._ (
      ( .. .g. . . . . . . . . . .~-._
   .~__.-~    ~`. . . . . . . . . . . -.
   `----..._      ~-=~~-. . . . . . . . ~-.
             ~-._   `-._ ~=_~~--. . . . . .~.
              | .~-.._  ~--._-.    ~-. . . . ~-.
               \ .(   ~~--.._~'       `. . . . .~-.                ,
                `._\         ~~--.._    `. . . . . ~-.    .- .   ,'/
_  . _ . -~\        _ ..  _          ~~--.`_. . . . . ~-_     ,-','`  .
             ` ._           ~                ~--. . . . .~=.-'. /. `
       - . -~            -. _ . - ~ - _   - ~     ~--..__~ _,. /   \  - ~
              . __ ..                   ~-               ~~_. (  `
)`. _ _               `-       ..  - .    . - ~ ~ .    \    ~-` ` `  `. _

```
