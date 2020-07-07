package models

//Response is a struct of model JsonResponse
type Response struct {
	Messages string
	Status   int
	Data     interface{}
}
