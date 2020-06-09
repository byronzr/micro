package handlers

// request method set
type GET struct{}
type HEAD struct{}
type POST struct{}
type PUT struct{}
type PATCH struct{}
type DELETE struct{}
type TRACE struct{}
type OPTIONS struct{}

func init() {
	methods := []interface{}{
		GET{}, HEAD{}, POST{}, PUT{}, PATCH{}, DELETE{}, TRACE{}, OPTIONS{},
	}
	for _, i := range methods {
		RegisterHandler(i)
	}

}
