package memos

type Repository interface {
	GetMemo(id string) (*Memo, error)
	AddMemo(deviceToken string, videoId string, body string, tagIds []string) (string, error)
	GetMemoByVideoId(id string) (*Memo, error)
	UpdateBody(id string, body string) (*Memo, error)
}
