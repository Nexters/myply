package memos

type Repository interface {
	GetMemo(id string) (*Memo, error)
	GetMemoByUniqueKey(videoID, deviceToken string) (*Memo, error)
	GetMemos(deviceToken string) (Memos, error)
	AddMemo(deviceToken string, videoID string, body string, tags []string) (string, error)
	GetMemoByVideoID(id string) (*Memo, error)
	UpdateBody(id string, body string) error
	DeleteMemo(id string) error
}
