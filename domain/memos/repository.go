package memos

type Repository interface {
	GetMemo(id string) (*Memo, error)
	SaveMemo(videoId string, body string, deviceToken string) (string, error)
	GetMemoByVideoId(id string) (*Memo, error)
}
