package types

type TestRecord struct {
	//flatfile is one indexed, position starts at 1
	AccountNum int    `flatfile:"1,10"`
	FirstName  string `flatfile:"11,5"`
	Age        int    `flatfile:"16,2"`
}
