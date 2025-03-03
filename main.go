package ChatLibraryGo

type AuthType int

const (
	SigIn AuthType = iota
	SignUp
)

type AuthSignIn struct {
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Sign     AuthType `json:"sign"`
}
