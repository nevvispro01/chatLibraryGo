package ChatLibraryGo

type AuthType int

const (
	SignIn AuthType = iota
	SignUp
)

type AuthSign struct {
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Sign     AuthType `json:"sign"`
}
