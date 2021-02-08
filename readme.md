# Goboye

This is a (work in progress) toy gameboy emulator written in Go.

Currently, it's possible to run (and play!) Tetris - but many other ROMs probably won't work.

## Current Features
- Full support for all CPU opcodes
- Some interrupts - VBlank and timer
- SDL graphics and input
- Websocket based debugger

## TODO
- Implement remaining interrupts
- Emulate sound hardware
- Emulate other functionality to get more games running
- Game boy colour support


## Dependencies

SDL is used for the UI - so you'll need the libraries installed:

On a mac:

    brew install sdl2


On ubuntu:

    apt install libsdl2-dev


## Running the SDL UI

    go run cmd/goboye/main.go -rom /path/to/rom.gb

This will start the emulator - key bindings are as follows:
- Up/Left/Down/Right: W/A/S/D
- Button A/Button B: P/O
- Start/Select: M/N
- Show framerate: F
- Quit: Esc

## Running the debugger

To run 

    go run cmd/debugger_ws/main.go -rom /path/to/rom.gb
    yarn start

    # open http://localhost:3000 in a browser
