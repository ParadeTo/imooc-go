package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url        string
	ParserFunc ParserFunc
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url 	string
	Type 	string
	Id 		string
	Payload interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type Deduplicate interface {
	IsDuplicate(url string)	bool
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
