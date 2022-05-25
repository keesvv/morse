package decoder

type Morse map[rune]string

type MorseDecoder struct {
	morse Morse
}

var InternationalMorse = Morse{
	'A': ".-", 'J': ".---", 'S': "...",
	'B': "-...", 'K': "-.-", 'T': "-",
	'C': "-.-.", 'L': ".-..", 'U': "..-",
	'D': "-..", 'M': "--", 'V': "...-",
	'E': ".", 'N': "-.", 'W': ".--",
	'F': "..-.", 'O': "---", 'X': "-..-",
	'G': "--.", 'P': ".--.", 'Y': "-.--",
	'H': "....", 'Q': "--.-", 'Z': "--..",
	'I': "..", 'R': ".-.",

	'1': ".----", '6': "-....",
	'2': "..---", '7': "--...",
	'3': "...--", '8': "---..",
	'4': "....-", '9': "----.",
	'5': ".....", '0': "-----",
}

func NewDecoder(morse Morse) *MorseDecoder {
	return &MorseDecoder{morse}
}

func (d *MorseDecoder) Decode(v []byte) rune {
	for r, code := range d.morse {
		if string(v) == code {
			return r
		}
	}
	return 0
}
