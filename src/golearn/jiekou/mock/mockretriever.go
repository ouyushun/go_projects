package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Post(url string) string {
	return "postt"
}

func (r *Retriever) Get(url string) string {
	return "yyss"
}
