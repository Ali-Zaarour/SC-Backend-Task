package dto

// module used to define a specific return value in api

type Responce struct {
	Reply Reply `json:"reply"`
}

type Reply string

const (
	FetchError            Reply = "Problem with get data"
	Created               Reply = "Student Created"
	Deleted               Reply = "Student Deleted"
	ErrorCreate           Reply = "Error in create a new Student"
	ErrorWithReceivedData Reply = "Error in data format"
	ErrorDelete           Reply = "Error in deleting this student"
	ErrorUpdate           Reply = "Error updating data"
)
