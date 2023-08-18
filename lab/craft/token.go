package craft

type Token interface {
	/*
		Token的类型
	*/
	getType() string

	/*
		Token的文本值
	*/
	getText() string
}

type SimpleToken struct {
	tokenType string
	text      string
}

func (t SimpleToken) getType() string {
	return t.tokenType
}

func (t SimpleToken) getText() string {
	return t.text
}
