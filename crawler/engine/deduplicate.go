package engine

type SimpleDeDuplicate struct {
	VisitedUrls map[string]bool
}


func NewSimpleDeDuplicate() *SimpleDeDuplicate {
	return &SimpleDeDuplicate{make(map[string]bool)}
}

func (d *SimpleDeDuplicate)IsDuplicate(url string) bool {
	if d.VisitedUrls[url] {
		return true
	}

	d.VisitedUrls[url] = true
	return false
}
