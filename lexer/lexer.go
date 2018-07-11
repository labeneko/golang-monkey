package lexer

// Lexer は現在の文字やその位置などを格納するtype
type Lexer struct {
	input        string
	position     int  // 入力における現在の位置（現在の文字を指し示す）
	readPosition int  // これから読み込む位置（現在の文字の次）
	ch           byte // 現在検査中の文字
}

// New はinputをもとにLexerを生成し、一文字読み取ったLexerを返す
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 入力が終端に達している場合、ファイルの終わりを示すために"NULL"
		// を表す0を設定する
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}
