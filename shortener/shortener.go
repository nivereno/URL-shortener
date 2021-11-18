package shortener


func Init(c string) interface{} {
	storage = interface{}
	switch c {
	case "db":
		storage := make(map[string]string)
	case "mem":
		storage := "db connection blah blah blah"
	}
	return storage
}
