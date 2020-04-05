# Atari2600

Initial stage Atari 2600 VCS Emulator writen in GO

Documentation:

Opcodes:

https://www.masswerk.at/6502/6502_instruction_set.html#CLD

https://problemkaputt.de/2k6specs.htm

http://www.obelisk.me.uk/6502/addressing.html#ABY

https://cdn.hackaday.io/files/1646277043401568/Atari_2600_Programming_for_Newbies_Revised_Edition.pdf

https://www.atariarchives.org/roots/chapter_6.php

FLAGS:
http://www.obelisk.me.uk/6502/reference.html#CPY

NTSC Palette:

http://www.qotile.net/minidig/docs/tia_color.html


Missing
- Input
- sound
- opcodes
- beam index:
 BNE:   ** add 1 to cycles if branch occurs on same page
     add 2 to cycles if branch occurs to different page
 LDA: *  add 1 to cycles if page boundery is crossed
 - drawing in the OVERSCAN
