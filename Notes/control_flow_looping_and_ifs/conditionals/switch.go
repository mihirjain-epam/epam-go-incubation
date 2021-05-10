package conditionals

type HTTPRequest struct {
	Method string
}

func Testing_switch() {
	r := HTTPRequest{Method: "HEAD"}

	switch r.Method {
	case "GET":
		println("in GET")
		fallthrough // to allow to go to below case also, in go we have implicit break between switch cases
	case "POST":
		println("in POST")
	case "DELETE":
		println("in delete")
	case "PUT":
		println("in PUT")
	default:
		println("Unhandled method")
	}

}
