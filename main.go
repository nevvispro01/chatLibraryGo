package ChatLibraryGo

type AuthType int
type Result int

const (
	SignIn AuthType = iota
	SignUp
)

const (
	Success Result = iota
	Error
)

type AuthSign struct {
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Sign     AuthType `json:"sign"`
}

type AuthAnswer struct {
	Result  Result `json:"result"`
	Message string `json:"message"`
}
