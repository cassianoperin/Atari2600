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
| Graphics: Scoreboard Player Color | OK |
| Graphics: Player 0 and Player 1 | OK |
| Graphics: Player Vertical Movement | OK |
| Graphics: Player Horizontal Movement | OK |
| Graphics: Player Stretch  (NUSIZ0 and NUSIZ1) | OK |
| Graphics: Player Multiply (NUSIZ0 and NUSIZ1) | OK |
| Graphics: Player Inverted (REFP0 and REFP1) | OK |
| Controller Input | OK |
| Memory page boundary cross detection | OK |
| CPU Stack | OK |
| 6507 CPU Opcodes | 44 / 56 |

## Missing:
- I'm recreating background, playfield and players now following TIA hardware.


- Improve P0 and P1 scroll (X and Y)
- Graphics: Ball
- Graphics: Missiles
- Scoreboard value increment for both players
- Scoreboard multi digit
- Sound
- Implement ISB "Opcode"
- Correct TIA implementation (I'm drawing first the background, then playfield and then objects after wsync. Not following the beam. Generates some glitches.)
- Improve TIA changed address detection (it's inside STA, STX and STY opcodes)
- Implement TIA Collision Detection:
	a) CXM0P - Not started
	b) CXM1P - Not started
	c) CXP0FB:
		D6 - P0–BL: Not started
		D7 - P0–PF: DONE!
	d) CXP1FB - Not started
	e) CXM0FB - Not started
	f) CXM1FB - Not started
	g) CXBLPF - Not started
	h) CXPPMM:
		D6 - M0–M1: Not started
		D7 - P0-P1: DONE!

## Documentation:


### Opcodes:

https://www.atariarchives.org/alp/appendix_1.php

https://www.masswerk.at/6502/6502_instruction_set.html#CLD

https://problemkaputt.de/2k6specs.htm

https://dwheeler.com/6502/oneelkruns/asm1step.html


### Addressing:

http://www.obelisk.me.uk/6502/addressing.html#ABY

http://www.emulator101.com/6502-addressing-modes.html


### FLAGS:
http://www.obelisk.me.uk/6502/reference.html#CPY


### 6502:
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


### Object draw:
https://alienbill.com/2600/playerpalnext.html


### Online debugger:
https://8bitworkshop.com/
