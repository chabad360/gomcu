package gomcu

// Char is an ASCII character between 0x21 (33) and 0x60 (96).
// Note: if the character is above 0x40 (64) subtract 0x40 from it, otherwise the wrong character will be displayed.
type Char byte

const (
	SymbolAt Char = iota
	CharA
	CharB
	CharC
	CharD
	CharE
	CharF
	CharG
	CharH
	CharI
	CharJ
	CharK
	CharL
	CharM
	CharN
	CharO
	CharP
	CharQ
	CharR
	CharS
	CharT
	CharU
	CharV
	CharW
	CharX
	CharY
	CharZ
	SymbolOpenBracket
	SymbolBackSlash
	SymbolCloseBracket
	SymbolAccent
	SymbolUnderscore
	SymbolSpace
	SymbolExclamation
	SymbolDoubleQuote
	SymbolNumberSign
	SymbolDollar
	SymbolPercentSign
	SymbolAmpersand
	SymbolSingleQuote
	SymbolOpenParen
	SymbolCloseParen
	SymbolAsterisk
	SymbolPlus
	SymbolComma
	SymbolHyphen
	SymbolPeriod
	SymbolForwardSlash
	Char0
	Char1
	Char2
	Char3
	Char4
	Char5
	Char6
	Char7
	Char8
	Char9
	SymbolColon
	SymbolSemiColon
	SymbolLessThan
	SymbolEqual
	SymbolGreaterThan
	SymbolQuestion
	// Add DigitDot to the character you want to display to show the dot found on each 7-segment display.
	DigitDot
)

var (
	Letters = map[rune]Char{
		'@':  0x00,
		'A':  0x01,
		'B':  0x02,
		'C':  0x03,
		'D':  0x04,
		'E':  0x05,
		'F':  0x06,
		'G':  0x07,
		'H':  0x08,
		'I':  0x09,
		'J':  0x0A,
		'K':  0x0B,
		'L':  0x0C,
		'M':  0x0D,
		'N':  0x0E,
		'O':  0x0F,
		'P':  0x10,
		'Q':  0x11,
		'R':  0x12,
		'S':  0x13,
		'T':  0x14,
		'U':  0x15,
		'V':  0x16,
		'W':  0x17,
		'X':  0x18,
		'Y':  0x19,
		'Z':  0x1A,
		'[':  0x1B,
		'\\': 0x1C,
		']':  0x1D,
		'^':  0x1E,
		'_':  0x1F,
		' ':  0x20,
		'!':  0x21,
		'"':  0x22,
		'#':  0x23,
		'$':  0x24,
		'%':  0x25,
		'&':  0x26,
		'\'': 0x27,
		'(':  0x28,
		')':  0x29,
		'*':  0x2A,
		'+':  0x2B,
		',':  0x2C,
		'-':  0x2D,
		'.':  0x2E,
		'/':  0x2F,
		'0':  0x30,
		'1':  0x31,
		'2':  0x32,
		'3':  0x33,
		'4':  0x34,
		'5':  0x35,
		'6':  0x36,
		'7':  0x37,
		'8':  0x38,
		'9':  0x39,
		':':  0x3A,
		';':  0x3B,
		'<':  0x3C,
		'=':  0x3D,
		'>':  0x3E,
		'?':  0x3F,
	}
)
