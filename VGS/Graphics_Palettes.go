package VGS

func NTSC(color byte) (float64, float64, float64) {

	var (
		x	byte = 0
		y	byte = 0
		r	float64 = 0
		g	float64 = 0
		b	float64 = 0
	)

	// Map the first hexadecimal byte
	x = color >> 4
	// Map the second hexadecimal byte
	y = color & 0xF

	// fmt.Printf("Byte 01: 0x%X | Byte 02: 0x%X\n", x, y)

	switch x {
		//MSB
		case 0x0:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 0 	; g = 0 	; b = 0
			} else if ( y == 0x2 || y == 0x3) {
				r = 64 	; g = 64  ; b = 64
			} else if ( y == 0x4 || y == 0x5) {
				r = 108	; g = 108 ; b = 108
			} else if ( y == 0x6 || y == 0x7) {
				r = 144	; g = 144 ; b = 144
			} else if ( y == 0x8 || y == 0x9) {
				r = 176	; g = 176 ; b = 176
			} else if ( y == 0xA || y == 0xB) {
				r = 200	; g = 200 ; b = 200
			} else if ( y == 0xC || y == 0xD) {
				r = 220	; g = 220 ; b = 220
			} else if ( y == 0xE || y == 0xF) {
				r = 236	; g = 236 ; b = 236
			}

		case 0x1:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 68 ;	g = 68 ;	b = 0
			} else if ( y == 0x2 || y == 0x3) {
				r = 100;	g = 100;	b = 16
			} else if ( y == 0x4 || y == 0x5) {
				r = 132;	g = 132;	b = 36
			} else if ( y == 0x6 || y == 0x7) {
				r = 160;	g = 160;	b = 52
			} else if ( y == 0x8 || y == 0x9) {
				r = 184;	g = 184;	b = 64
			} else if ( y == 0xA || y == 0xB) {
				r = 208;	g = 208;	b = 80
			} else if ( y == 0xC || y == 0xD) {
				r = 232;	g = 232;	b = 92
			} else if ( y == 0xE || y == 0xF) {
				r = 252;	g = 252;	b = 104
			}

		case 0x2:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 112;	g = 40 ;	b = 0
			} else if ( y == 0x2 || y == 0x3) {
				r = 132;	g = 68 ;	b = 20
			} else if ( y == 0x4 || y == 0x5) {
				r = 152;	g = 92 ;	b = 40
			} else if ( y == 0x6 || y == 0x7) {
				r = 172;	g = 120;	b = 60
			} else if ( y == 0x8 || y == 0x9) {
				r = 188;	g = 140;	b = 76
			} else if ( y == 0xA || y == 0xB) {
				r = 204;	g = 160;	b = 92
			} else if ( y == 0xC || y == 0xD) {
				r = 220;	g = 180;	b = 104
			} else if ( y == 0xE || y == 0xF) {
				r = 236;	g = 200;	b = 120
			}

		case 0x3:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 132;	g = 24 ;	b = 0
			} else if ( y == 0x2 || y == 0x3) {
				r = 152;	g = 52 ;	b = 24
			} else if ( y == 0x4 || y == 0x5) {
				r = 172;	g = 80 ;	b = 48
			} else if ( y == 0x6 || y == 0x7) {
				r = 192;	g = 104;	b = 72
			} else if ( y == 0x8 || y == 0x9) {
				r = 208;	g = 128;	b = 92
			} else if ( y == 0xA || y == 0xB) {
				r = 224;	g = 148;	b = 112
			} else if ( y == 0xC || y == 0xD) {
				r = 236;	g = 168;	b = 128
			} else if ( y == 0xE || y == 0xF) {
				r = 252;	g = 188;	b = 148
			}

		case 0x4:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 136;	g = 0  ;	b = 0
			} else if ( y == 0x2 || y == 0x3) {
				r = 156;	g = 32 ;	b = 32
			} else if ( y == 0x4 || y == 0x5) {
				r = 176;	g = 60 ;	b = 60
			} else if ( y == 0x6 || y == 0x7) {
				r = 192;	g = 88 ;	b = 88
			} else if ( y == 0x8 || y == 0x9) {
				r = 208;	g = 112;	b = 112
			} else if ( y == 0xA || y == 0xB) {
				r = 224;	g = 136;	b = 136
			} else if ( y == 0xC || y == 0xD) {
				r = 236;	g = 160;	b = 160
			} else if ( y == 0xE || y == 0xF) {
				r = 252;	g = 180;	b = 180
			}

		case 0x5:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 120;	g = 0  ;	b = 92
			} else if ( y == 0x2 || y == 0x3) {
				r = 140;	g = 32 ;	b = 116
			} else if ( y == 0x4 || y == 0x5) {
				r = 160;	g = 60 ;	b = 136
			} else if ( y == 0x6 || y == 0x7) {
				r = 176;	g = 88 ;	b = 156
			} else if ( y == 0x8 || y == 0x9) {
				r = 192;	g = 112;	b = 176
			} else if ( y == 0xA || y == 0xB) {
				r = 208;	g = 132;	b = 192
			} else if ( y == 0xC || y == 0xD) {
				r = 220;	g = 156;	b = 208
			} else if ( y == 0xE || y == 0xF) {
				r = 236;	g = 176;	b = 224
			}

		case 0x6:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 72 ;	g = 0  ;	b = 120
			} else if ( y == 0x2 || y == 0x3) {
				r = 96 ;	g = 32 ;	b = 144
			} else if ( y == 0x4 || y == 0x5) {
				r = 120;	g = 60 ;	b = 164
			} else if ( y == 0x6 || y == 0x7) {
				r = 140;	g = 88 ;	b = 184
			} else if ( y == 0x8 || y == 0x9) {
				r = 160;	g = 112;	b = 204
			} else if ( y == 0xA || y == 0xB) {
				r = 180;	g = 132;	b = 220
			} else if ( y == 0xC || y == 0xD) {
				r = 196;	g = 156;	b = 236
			} else if ( y == 0xE || y == 0xF) {
				r = 212;	g = 176;	b = 252
			}

		case 0x7:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 20 ;	g = 0  ;	b = 132
			} else if ( y == 0x2 || y == 0x3) {
				r = 48 ;	g = 32 ;	b = 152
			} else if ( y == 0x4 || y == 0x5) {
				r = 76 ;	g = 60 ;	b = 172
			} else if ( y == 0x6 || y == 0x7) {
				r = 104;	g = 88 ;	b = 192
			} else if ( y == 0x8 || y == 0x9) {
				r = 124;	g = 112;	b = 208
			} else if ( y == 0xA || y == 0xB) {
				r = 148;	g = 136;	b = 224
			} else if ( y == 0xC || y == 0xD) {
				r = 168;	g = 160;	b = 236
			} else if ( y == 0xE || y == 0xF) {
				r = 188;	g = 180;	b = 252
			}

		case 0x8:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 0  ;	g = 0  ;	b = 136
			} else if ( y == 0x2 || y == 0x3) {
				r = 28 ;	g = 32 ;	b = 156
			} else if ( y == 0x4 || y == 0x5) {
				r = 56 ;	g = 64 ;	b = 176
			} else if ( y == 0x6 || y == 0x7) {
				r = 80 ;	g = 92 ;	b = 192
			} else if ( y == 0x8 || y == 0x9) {
				r = 104;	g = 116;	b = 208
			} else if ( y == 0xA || y == 0xB) {
				r = 124;	g = 140;	b = 224
			} else if ( y == 0xC || y == 0xD) {
				r = 144;	g = 164;	b = 236
			} else if ( y == 0xE || y == 0xF) {
				r = 164;	g = 184;	b = 252
			}

		case 0x9:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 0  ;	g = 24 ;	b = 124
			} else if ( y == 0x2 || y == 0x3) {
				r = 28 ;	g = 56 ;	b = 144
			} else if ( y == 0x4 || y == 0x5) {
				r = 56 ;	g = 84 ;	b = 168
			} else if ( y == 0x6 || y == 0x7) {
				r = 80 ;	g = 112;	b = 188
			} else if ( y == 0x8 || y == 0x9) {
				r = 104;	g = 136;	b = 204
			} else if ( y == 0xA || y == 0xB) {
				r = 124;	g = 156;	b = 220
			} else if ( y == 0xC || y == 0xD) {
				r = 144;	g = 180;	b = 236
			} else if ( y == 0xE || y == 0xF) {
				r = 164;	g = 200;	b = 252
			}

		case 0xA:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 0  ;	g = 44 ;	b = 92
			} else if ( y == 0x2 || y == 0x3) {
				r = 28 ;	g = 76 ;	b = 120
			} else if ( y == 0x4 || y == 0x5) {
				r = 56 ;	g = 104;	b = 144
			} else if ( y == 0x6 || y == 0x7) {
				r = 80 ;	g = 132;	b = 172
			} else if ( y == 0x8 || y == 0x9) {
				r = 104;	g = 156;	b = 192
			} else if ( y == 0xA || y == 0xB) {
				r = 124;	g = 180;	b = 212
			} else if ( y == 0xC || y == 0xD) {
				r = 144;	g = 204;	b = 232
			} else if ( y == 0xE || y == 0xF) {
				r = 164;	g = 224;	b = 252
			}

		case 0xB:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 0  ;	g = 60 ;	b = 44
			} else if ( y == 0x2 || y == 0x3) {
				r = 28 ;	g = 92 ;	b = 72
			} else if ( y == 0x4 || y == 0x5) {
				r = 56 ;	g = 124;	b = 100
			} else if ( y == 0x6 || y == 0x7) {
				r = 80 ;	g = 156;	b = 128
			} else if ( y == 0x8 || y == 0x9) {
				r = 104;	g = 180;	b = 148
			} else if ( y == 0xA || y == 0xB) {
				r = 124;	g = 208;	b = 172
			} else if ( y == 0xC || y == 0xD) {
				r = 144;	g = 228;	b = 192
			} else if ( y == 0xE || y == 0xF) {
				r = 164;	g = 252;	b = 212
			}

		case 0xC:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 0  ;	g = 60 ;	b = 0
			} else if ( y == 0x2 || y == 0x3) {
				r = 32 ;	g = 92 ;	b = 32
			} else if ( y == 0x4 || y == 0x5) {
				r = 64 ;	g = 124;	b = 64
			} else if ( y == 0x6 || y == 0x7) {
				r = 92 ;	g = 156;	b = 92
			} else if ( y == 0x8 || y == 0x9) {
				r = 116;	g = 180;	b = 116
			} else if ( y == 0xA || y == 0xB) {
				r = 140;	g = 208;	b = 140
			} else if ( y == 0xC || y == 0xD) {
				r = 164;	g = 228;	b = 164
			} else if ( y == 0xE || y == 0xF) {
				r = 184;	g = 252;	b = 184
			}

		case 0xD:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 20 ;	g = 56 ;	b = 0
			} else if ( y == 0x2 || y == 0x3) {
				r = 52 ;	g = 92 ;	b = 28
			} else if ( y == 0x4 || y == 0x5) {
				r = 80 ;	g = 124;	b = 56
			} else if ( y == 0x6 || y == 0x7) {
				r = 108;	g = 152;	b = 80
			} else if ( y == 0x8 || y == 0x9) {
				r = 132;	g = 180;	b = 104
			} else if ( y == 0xA || y == 0xB) {
				r = 156;	g = 204;	b = 124
			} else if ( y == 0xC || y == 0xD) {
				r = 180;	g = 228;	b = 144
			} else if ( y == 0xE || y == 0xF) {
				r = 200;	g = 252;	b = 164
			}

		case 0xE:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 44 ;	g = 48 ;	b = 0
			} else if ( y == 0x2 || y == 0x3) {
				r = 76 ;	g = 80 ;	b = 28
			} else if ( y == 0x4 || y == 0x5) {
				r = 104;	g = 112;	b = 52
			} else if ( y == 0x6 || y == 0x7) {
				r = 132;	g = 140;	b = 76
			} else if ( y == 0x8 || y == 0x9) {
				r = 156;	g = 168;	b = 100
			} else if ( y == 0xA || y == 0xB) {
				r = 180;	g = 192;	b = 120
			} else if ( y == 0xC || y == 0xD) {
				r = 204;	g = 212;	b = 136
			} else if ( y == 0xE || y == 0xF) {
				r = 224;	g = 236;	b = 156
			}

		case 0xF:
			//LSB
			if ( y == 0x0 || y == 0x1) {
				r = 68 ;	g = 40 ;	b = 0
			} else if ( y == 0x2 || y == 0x3) {
				r = 100;	g = 72 ;	b = 24
			} else if ( y == 0x4 || y == 0x5) {
				r = 132;	g = 104;	b = 48
			} else if ( y == 0x6 || y == 0x7) {
				r = 160;	g = 132;	b = 68
			} else if ( y == 0x8 || y == 0x9) {
				r = 184;	g = 156;	b = 88
			} else if ( y == 0xA || y == 0xB) {
				r = 208;	g = 180;	b = 108
			} else if ( y == 0xC || y == 0xD) {
				r = 232;	g = 204;	b = 124
			} else if ( y == 0xE || y == 0xF) {
				r = 252;	g = 224;	b = 140
			}


		default:
			// fmt.Printf("\n\tPalette color not mapped!\n")

	}

	return r, g, b

}
