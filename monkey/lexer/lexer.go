package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  //入力における現在の位置(現在の文字を指し示す)
	readPosition int  // これから読み込む位置(現在の文字の次)
	ch           byte // 現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	// positionを次に進める
    l.position = l.readPosition
    // readpositonを次に進める
    l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()          //英文字の部分を切り取って、tokのLiteralフィールドにセット
			tok.Type = token.LookupIdent(tok.Literal) //token.LookupIdent()を使って、それがキーワードか識別子か判定し、対応するtokenTypeをセットする
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// tokenTypeと文字を受け取って、対応するトークンを生成する
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Lexerについてのメソッドで、Lexerが現在読んでいる文字が英文字の時には、後に続く英文字の部分を切り出し、Lexerのinputにセットする
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Lexerについてのメソッドで、Lexerが現在読んでいる場所が空文字の時には、そのままreadCharを呼び出して、そこをスキップする
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Lexerについてのメソッドで、Lexerが現在読んでいる場所が数字のときには、後に続く数字の部分を切り出し、Lexerのinputにセットする
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
