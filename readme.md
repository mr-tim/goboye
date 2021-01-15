# Goboye

This is a (work in progress) toy gameboy emulator written in Go.

Currently the emulator has a basic implementation of the CPU (with all the supported opcodes) along with basic 
support for the LCD display.

## Todo

- Successfully run all the blargg test ROMs (currently only half run correctly)
- Support for all interrupts
- SDL based display and input
- Emulate sound hardware

# Running the debugger

To run 

    go run cmd/debugger_ws/main.go
    yarn start

    # open http://localhost:3000 in a browser
