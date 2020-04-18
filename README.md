# Atari2600

Initial stage Atari 2600 VCS Emulator writen in GO


## Emulation status:
- CRT NTSC TV Display: OK
- Graphics: Background: OK
- Graphics: Playfield: OK
- Graphics: Scoreboard: OK
- Graphics: Player 0 and Player 1: OK
- Graphics: Player Vertical Movement: OK
- 6502/6507 CPU Opcodes: 36/150

## Missing:
- 6vertical.bin program, wrong line draw at the end of screen (first line draw)
- Graphics: Player Horizontal Movement - Measurement with CPU cycles
- Graphics: Ball
- Graphics: Missiles
- Input
- Sound

## Documentation:

### General:

https://cdn.hackaday.io/files/1646277043401568/Atari_2600_Programming_for_Newbies_Revised_Edition.pdf

https://www.atariarchives.org/roots/chapter_6.php

https://pt.slideshare.net/chesterbr/atari-2600programming

https://www.randomterrain.com/atari-2600-memories-tutorial-andrew-davie-01.html#basics


### Opcodes:

https://www.masswerk.at/6502/6502_instruction_set.html#CLD

https://problemkaputt.de/2k6specs.htm


### Addressing:

http://www.obelisk.me.uk/6502/addressing.html#ABY

http://www.emulator101.com/6502-addressing-modes.html


### FLAGS:
http://www.obelisk.me.uk/6502/reference.html#CPY


### NTSC Palette:

http://www.qotile.net/minidig/docs/tia_color.html

### Cycles counting:
https://www.randomterrain.com/atari-2600-memories-guide-to-cycle-counting.html


### Overflow flag:
http://www.righto.com/2012/12/the-6502-overflow-flag-explained.html


### BRK/IRQ/NMI/RESET:
https://www.pagetable.com/?p=410


### PIXEL:
https://gitter.im/pixellib/Lobby?at=5dbc310c10bd4128a19e5608
