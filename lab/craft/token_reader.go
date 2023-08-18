package craft

type TokenReader interface {
	/*
		返回Token流中下一个Token，并从流中取出，如果流已经为空，返回nil
	*/
	Read() Token
	/*
		返回Token流中下一个Token，不从流中取出，如果流已经为空，返回nil
	*/
	Peek() Token
	/*
		Token流回退一步，恢复原来的Token
	*/
	UnRead()
	/*
		获取Token流当前的读取位置
	*/
	GetPosition() int
	/*
		设置Token流当前的读取位置
	*/
	SetPosition(pos int)
}
