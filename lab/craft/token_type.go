package craft

type TokenType int

const (
	TokenTypePlus  TokenType = iota // +
	TokenTypeMinus                  // -
	TokenTypeStar                   // *
	TokenTypeSlash                  // /

	TokenTypeGE // >=
	TokenTypeGT // >
	TokenTypeEQ // ==
	TokenTypeLE // <=
	TokenTypeLT // <

	TokenTypeSemiColon  // ;
	TokenTypeLeftParen  // (
	TokenTypeRightParen // )

	TokenTypeAssignment // =

	TokenTypeIf
	TokenTypeElse

	TokenTypeInt

	TokenTypeIdentifier //标识符

	TokenTypeIntLiteral    //整型字面量
	TokenTypeStringLiteral //字符串字面量
)

func (t TokenType) String() string {
	switch t {
	case TokenTypePlus:
		return "+"
	case TokenTypeMinus:
		return "-"
	case TokenTypeStar:
		return "*"
	case TokenTypeSlash:
		return "/"
	case TokenTypeGE:
		return ">="
	case TokenTypeGT:
		return ">"
	case TokenTypeEQ:
		return "=="
	case TokenTypeLE:
		return "<="
	case TokenTypeLT:
		return "<"
	case TokenTypeSemiColon:
		return ";"
	case TokenTypeLeftParen:
		return "("
	case TokenTypeRightParen:
		return ")"
	case TokenTypeAssignment:
		return "="
	case TokenTypeIf:
		return "If"
	case TokenTypeElse:
		return "Else"
	case TokenTypeInt:
		return "Int"
	case TokenTypeIdentifier:
		return "Identifier"
	case TokenTypeIntLiteral:
		return "IntLiteral"
	case TokenTypeStringLiteral:
		return "StringLiteral"
	default:
		return "Unknown"
	}
}
