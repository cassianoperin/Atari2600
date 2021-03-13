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
| Decimal Mode (BCD) - ADC opcode | OK |
| TIA RO Memory Mirrors (64 bytes) | OK |
| 6507 CPU Opcodes | 46 / 56 |


## Missing:
- Timers
- BCD Mode SBC
- Player are not well centralized (horizontal movement)
- Graphics: Improve P0 and P1 scroll (X and Y)
- Graphics: Ball
- Graphics: Missiles
- Scoreboard value increment for both players
- Scoreboard multi digit
- Sound
- Implement ISB "Opcode"
- Implement TIA Collision Detection:
	a) CXM0P:
		D6 - M0-P0: Not started
		D7 - M0-P1: Not started
	b) CXM1P:
		D6 - M1-P1: Not started
		D7 - M1-P0: Not started
	c) CXP0FB:
		D6 - P0–BL: Not started
		D7 - P0–PF: DONE
	d) CXP1FB:
		D6 - P1-BL: Not started
		D7 - P1-PF: DONE
	e) CXM0FB:
		D6 - M0-BL: Not started
		D7 - M0-PF: Not started
	f) CXM1FB:
		D6 - M1-BL: Not started
		D7 - M1-PF: Not started
	g) CXBLPF:
		D6 - -----
		D7 - BL-PF: Not started
	h) CXPPMM:
		D6 - M0–M1: Not started
		D7 - P0-P1: DONE

## Keys:
- Key 6: Console Color switch
- Key 7: Console Game Select switch
- Key 8: Console Reset switch
- Key 9: Enable/Disable Debug
- Key 0: Reset emulation


## Documentation:


### Opcodes:

https://www.masswerk.at/6502/6502_instruction_set.html

http://www.obelisk.me.uk/6502/reference.html

https://sites.google.com/site/6502asembly/6502-instruction-set/ror

https://www.atariarchives.org/alp/appendix_1.php

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


### Memory Map:
https://atariage.com/forums/topic/192418-mirrored-memory/?tab=comments#comment-2439795
https://problemkaputt.de/2k6specs.htm#videohttps://problemkaputt.de/2k6specs.htm#memoryandiomap


### Timer
https://atariage.com/forums/topic/133686-please-explain-riot-timmers/


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
