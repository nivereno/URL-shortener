package shortener

type Storage struct {
	m  map[string]string
	d  string
	db bool
}

func Init(c string) Storage {
	s := Storage{nil, "", false}
	switch c {
	case "mem":
		s.m = make(map[string]string)
		s.db = false
	case "db":
		s.d = "db connection blah blah blah"
		s.db = true
	}
	println(s.m, s.d, s.db)
	return s
}
