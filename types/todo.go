package types

type Todo struct {
	ID    int
	Title string
	Done  bool
}

type TodoDeleteReq struct {
	ID int `json:"id" binding:"required"`
}
