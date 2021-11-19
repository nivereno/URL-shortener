package shortener

type storage struct {
	m  map[string]string
	d  string
	db string
}

var s = storage{nil, "", ""}

func Init(c string) storage {
	switch c {
	case "mem":
		s.m = make(map[string]string)
		s.db = "memory"
	case "db":
		s.d = "db connection blah blah blah"
		s.db = "postgres"
	}
	return s
}

func SaveUrl(url string) string {
	shorturl := hashUrl(url)
	switch s.db {
	case "memory":
		for k, v := range s.m {
			if v == url {
				shorturl = k
			} else if k == shorturl {
				SaveUrl(url)
			}
		}
		s.m[shorturl] = url

	case "postgres":
	}
	return shorturl
}

func hashUrl(url string) string {
	shorturl := url
	return shorturl
}

func LookupUrl(shorturl string) string {
	var fullurl string
	switch s.db {
	case "memory":
		if s.m[shorturl] == "" {
			fullurl = "Url does not exist"
		} else {
			fullurl = s.m[shorturl]
		}
	case "postgres":
	}
	return fullurl
}
