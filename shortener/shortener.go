package shortener

type storage struct {
	m  map[string]string
	d  string
	db string
}

func Init(c string) storage {
	s := storage{nil, "", ""}
	switch c {
	case "mem":
		s.m = make(map[string]string)
		s.db = "memory"
	case "db":
		s.d = "db connection blah blah blah"
		s.db = "postgres"
	}
	println(s.m, s.d, s.db)
	return s
}

func hashUrl(url string) string {

}

func SaveUrl(url string, s storage) string {
	short := hashUrl(url)
	switch s.db {
	case "memory":
		s.m[short] = url
	case "postgres":
	}
	return short
}
