package utils

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NotFound() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "resource not found"
	return e
}
