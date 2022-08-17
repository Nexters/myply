package memos

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

var (
	createdAt      = time.Unix(time.Now().Unix(), 0)
	updatedAt      = time.Unix(time.Now().Unix(), 0)
	mockYoutubeIDs = []string{"wQ12qs3bmOg", "T--TNFU_Cv8", "TqcrbaZrFhg", "bpW4M5f3_00"}
)

func genMemo(id, youtubeID string) Memo {
	return Memo{
		ID:             id,
		DeviceToken:    "testDeviceToken",
		YoutubeVideoID: youtubeID,
		Body:           "I am a test Body",
		Tags:           []string{"tag1", "tag2", "tag3"},
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}

func genMemos() Memos {
	mockMemos := []Memo{}
	for i, yid := range mockYoutubeIDs {
		mockMemos = append(mockMemos, genMemo(fmt.Sprint(i), yid))
	}
	return mockMemos
}

func setup(t *testing.T) (func(t *testing.T), Memos) {
	t.Log("setup test case")
	memos := genMemos()
	return func(t *testing.T) {
		t.Log("teardown test case")
	}, memos
}

func TestMemosYoutubeVideoIdsSuccess(t *testing.T) {
	teardown, memos := setup(t)
	defer teardown(t)

	if reflect.DeepEqual(memos.YoutubeVideoIDs(), mockYoutubeIDs) == false {
		t.Fail()
	}
}
