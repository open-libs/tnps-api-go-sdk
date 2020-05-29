package client

//"log"

const (
	Path = "/"
)

type Request interface {
	GetPath() string
	ToJsonString() string
	FromJsonString(s string) error
}

type BaseRequest struct {
	Path string `json:"-"`
}

func (r *BaseRequest) GetPath() string {
	return r.Path
}
