package craft

import (
	"fmt"
	"strings"
)

type SimpleLexer struct {
	tokenText *strings.Builder // 临时保存token的文本
	tokens    []Token          // 保存解析出来的Token
	token     SimpleToken      // 当前正在解析的Token
}

// 是否是字母
func isAlpha(ch int) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

// 是否是数字
func isDigit(ch int) bool {
	return ch >= '0' && ch <= '9'
}

// 是否是空白字符
func isBlank(ch int) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

type DfaState int

const (
	DfaStateInitial DfaState = iota

	DfaStateIf
	DfaStateId_if1
	DfaStateId_if2
	DfaStateElse
	DfaStateId_else1
	DfaStateId_else2
	DfaStateId_else3
	DfaStateId_else4
	DfaStateInt
	DfaStateId_int1
	DfaStateId_int2
	DfaStateId_int3
	DfaStateId
	DfaStateGT
	DfaStateGE

	DfaStateAssignment

	DfaStatePlus
	DfaStateMinus
	DfaStateStar
	DfaStateSlash

	DfaStateSemiColon
	DfaStateLeftParen
	DfaStateRightParen

	DfaStateIntLiteral
)

func (l *SimpleLexer) initToken(ch int) DfaState {
	if l.tokenText.Len() > 0 {
		l.token.text = l.tokenText.String()
		l.tokens = append(l.tokens, l.token)

		l.tokenText.Reset()
		l.token = SimpleToken{}
	}

	newState := DfaStateInitial
	if isAlpha(ch) {
		if ch == 'i' {
			newState = DfaStateId_int1
		} else {
			newState = DfaStateId
		}
		l.token.tokenType = TokenTypeIdentifier.String()
		l.tokenText.WriteByte(byte(ch))
	} else if isDigit(ch) {
		newState = DfaStateIntLiteral
		l.token.tokenType = TokenTypeIntLiteral.String()
		l.tokenText.WriteByte(byte(ch))
	} else if ch == '>' {
		newState = DfaStateGT
		l.token.tokenType = TokenTypeGT.String()
		l.tokenText.WriteByte(byte(ch))
	} else if ch == '+' {
		newState = DfaStatePlus
		l.token.tokenType = TokenTypePlus.String()
		l.tokenText.WriteByte(byte(ch))
	} else if ch == '-' {
		newState = DfaStateMinus
		l.token.tokenType = TokenTypeMinus.String()
		l.tokenText.WriteByte(byte(ch))
	} else if ch == '*' {
		newState = DfaStateStar
		l.token.tokenType = TokenTypeStar.String()
		l.tokenText.WriteByte(byte(ch))
	} else if ch == '/' {
		newState = DfaStateSlash
		l.token.tokenType = TokenTypeSlash.String()
		l.tokenText.WriteByte(byte(ch))
	} else if ch == ';' {
		newState = DfaStateSemiColon
		l.token.tokenType = TokenTypeSemiColon.String()
		l.tokenText.WriteByte(byte(ch))
	} else if ch == '(' {
		newState = DfaStateLeftParen
		l.token.tokenType = TokenTypeLeftParen.String()
		l.tokenText.WriteByte(byte(ch))
	} else if ch == ')' {
		newState = DfaStateRightParen
		l.token.tokenType = TokenTypeRightParen.String()
		l.tokenText.WriteByte(byte(ch))
	} else if ch == '=' {
		newState = DfaStateAssignment
		l.token.tokenType = TokenTypeAssignment.String()
		l.tokenText.WriteByte(byte(ch))
	} else {
		newState = DfaStateInitial // skip all unknown patterns
	}
	return newState
}

/*
解析字符串，形成token。
这是一个有限状态自动机，在不同的状态中迁移。
*/
func tokenize(code string) SimpleTokenReader {
	lexer := &SimpleLexer{
		tokenText: new(strings.Builder),
		tokens:    make([]Token, 0),
		token:     SimpleToken{},
	}
	var ch int
	state := DfaStateInitial
	for _, c := range code {
		ch = int(c)
		switch state {
		case DfaStateInitial:
			state = lexer.initToken(ch) //重新确定后续状态
		case DfaStateId:
			if isAlpha(ch) || isDigit(ch) {
				lexer.tokenText.WriteByte(byte(ch)) // 保持标识符状态
			} else {
				state = lexer.initToken(ch) // 退出标识符状态，并保存token
			}
		case DfaStateGT:
			if ch == '=' {
				lexer.token.tokenType = TokenTypeGE.String() // 转换成GE
				state = DfaStateGE
				lexer.tokenText.WriteByte(byte(ch))
			} else {
				state = lexer.initToken(ch)
			}
		case DfaStateGE, DfaStateAssignment, DfaStatePlus, DfaStateMinus, DfaStateStar, DfaStateSlash, DfaStateSemiColon, DfaStateLeftParen, DfaStateRightParen:
			state = lexer.initToken(ch)
		case DfaStateIntLiteral:
			if isDigit(ch) {
				lexer.tokenText.WriteByte(byte(ch)) // 继续保持在数字字面量状态
			} else {
				state = lexer.initToken(ch)
			}
		case DfaStateId_int1:
			if ch == 'n' {
				state = DfaStateId_int2
				lexer.tokenText.WriteByte(byte(ch))
			} else if isDigit(ch) || isAlpha(ch) {
				state = DfaStateId
				lexer.tokenText.WriteByte(byte(ch))
			} else {
				state = lexer.initToken(ch)
			}
		case DfaStateId_int2:
			if ch == 't' {
				state = DfaStateId_int3
				lexer.tokenText.WriteByte(byte(ch))
			} else if isDigit(ch) || isAlpha(ch) {
				state = DfaStateId
				lexer.tokenText.WriteByte(byte(ch))
			} else {
				state = lexer.initToken(ch)
			}
		case DfaStateId_int3:
			if isBlank(ch) {
				lexer.token.tokenType = TokenTypeInt.String()
				state = lexer.initToken(ch)
			} else {
				state = DfaStateId
				lexer.tokenText.WriteByte(byte(ch))
			}
		}
	}
	// 把最后一个token送进去
	if lexer.tokenText.Len() > 0 {
		lexer.initToken(ch)
	}
	return NewSimpleTokenReader(lexer.tokens)
}

type SimpleTokenReader struct {
	tokens []Token
	pos    int
}

func NewSimpleTokenReader(tokens []Token) SimpleTokenReader {
	return SimpleTokenReader{tokens: tokens}
}

func (r *SimpleTokenReader) Read() Token {
	if r.pos < len(r.tokens) {
		r.pos++
		return r.tokens[r.pos-1]
	}
	return nil
}

func (r *SimpleTokenReader) Peek() Token {
	if r.pos < len(r.tokens) {
		return r.tokens[r.pos]
	}
	return nil
}

func (r *SimpleTokenReader) UnRead() {
	if r.pos > 0 {
		r.pos--
	}
}

func (r *SimpleTokenReader) GetPosition() int {
	return r.pos
}

func (r *SimpleTokenReader) SetPosition(pos int) {
	if pos >= 0 && pos < len(r.tokens) {
		r.pos = pos
	}
}

func dump(reader SimpleTokenReader) {
	fmt.Println("text\ttype")
	token := reader.Read()
	for token != nil {
		fmt.Println(token.getText(), "\t", token.getType())
		token = reader.Read()
	}
}
