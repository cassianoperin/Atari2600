# Atari2600

Initial stage Atari 2600 VCS Emulator writen in GO.


**Horizontal Movement** | **Vertical Movement**
:-------------------------:|:-------------------------:
<img width="430" alt="horizontal" src="https://github.com/cassianoperin/Atari2600/blob/master/Images/HorizontalMovement.gif">  |  <img width="430" alt="vertical" src="https://github.com/cassianoperin/Atari2600/blob/master/Images/VerticalMovement.gif">

**Controller Input** | **NTSC Palette**
:-------------------------:|:-------------------------:
<img width="430" alt="input" src="https://github.com/cassianoperin/Atari2600/blob/master/Images/Input.gif"> | <img width="430" alt="palette" src="https://github.com/cassianoperin/Atari2600/blob/master/Images/Palette.png">

**Players and Scoreboard** | **Scoreboard Colors**
:-------------------------:|:-------------------------:
<img width="430" alt="players" src="https://github.com/cassianoperin/Atari2600/blob/master/Images/PlayersScoreboard.png"> | <img width="430" alt="scoreboard" src="https://github.com/cassianoperin/Atari2600/blob/master/Images/PlayersScoreboardColor.png">

**Background and Playfield** | **Playfield Reflection**
:-------------------------:|:-------------------------:
<img width="430" alt="background" src="https://github.com/cassianoperin/Atari2600/blob/master/Images/Playfield.png"> | <img width="430" alt="reflection" src="https://github.com/cassianoperin/Atari2600/blob/master/Images/PlayfieldReflex.png">


## Emulation status:
| Name  | Status |
| :------------ | :----- |
| CRT NTSC TV Display | OK |
| Graphics: NTSC Color Palette | OK | |
| Graphics: Background | OK |
| Graphics: Playfield | OK |
| Graphics: Scoreboard | OK |
| Graphics: Player 0 and Player 1 | OK |
| Graphics: Player Vertical Movement | OK |
| Graphics: Player Horizontal Movement | OK |
| Graphics: Player Stretch  (NUSIZ0 and NUSIZ1) | OK |
| Graphics: Player Multiply (NUSIZ0 and NUSIZ1) | OK |
| Graphics: Player Inverted (REFP0 and REFP1) | OK |
| Controller Input | OK |
| Memory page boundary cross detection | OK |
| CPU Stack | OK |
| 6502/6507 CPU Opcodes | 45 / 150 |

## Missing:
- X and Y Movement: Need to split object when crossing the boundaries of the screen (test on 103bomber_WithoutADC.bin)
- Graphics: Ball
- Graphics: Missiles
- Object Collisions
- Scoreboard value increment for both players
- Scoreboard multi digit
- Sound
- Controller Input: More than one key interpreted at the same time (Counter of keys pressed and do not update PC?)
- Review progs 1 and 2
- Review 0x95 , X?


## Documentation:


### Opcodes:

https://www.masswerk.at/6502/6502_instruction_set.html#CLD

https://problemkaputt.de/2k6specs.htm

https://dwheeler.com/6502/oneelkruns/asm1step.html


### Addressing:

http://www.obelisk.me.uk/6502/addressing.html#ABY

http://www.emulator101.com/6502-addressing-modes.html


### FLAGS:
http://www.obelisk.me.uk/6502/reference.html#CPY


### 6502
http://www.6502.org/


### NTSC Palette:

http://www.qotile.net/minidig/docs/tia_color.html


### Cycles counting:
https://www.randomterrain.com/atari-2600-memories-guide-to-cycle-counting.html


### Overflow flag:
http://www.righto.com/2012/12/the-6502-overflow-flag-explained.html


### BRK/IRQ/NMI/RESET:
https://www.pagetable.com/?p=410


### General:

https://cdn.hackaday.io/files/1646277043401568/Atari_2600_Programming_for_Newbies_Revised_Edition.pdf

https://www.atariarchives.org/roots/chapter_6.php

https://pt.slideshare.net/chesterbr/atari-2600programming

https://www.randomterrain.com/atari-2600-memories-tutorial-andrew-davie-01.html#basics

https://dwheeler.com/6502/oneelkruns/asm1step.html


### PIXEL:
https://gitter.im/pixellib/Lobby?at=5dbc310c10bd4128a19e5608


### Object draw
https://alienbill.com/2600/playerpalnext.html


### Online debugger
https://8bitworkshop.com/
