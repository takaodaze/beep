package csv

func ToRow(s []string) string {
	var r string
	for _, v := range s {
		r += v + ","
	}
	return r[:len(r)-1]
}
