package responses

/*

In Go, map[string]interface{} is a map where the keys are strings and the values can be of any type, as they are of type interface{}.

Here's a breakdown:

map[string]: This indicates that the keys of the map are of type string.
interface{}: This is an empty interface, which means the values associated with the keys can be of any type. An empty interface can hold values of any type because every type in Go implements at least zero methods.

*/

// ceates a UserResponse struct with Status, Message, and Data property to represent the API response type.
type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
