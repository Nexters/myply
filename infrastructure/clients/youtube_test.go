package clients

import (
	"os"
	"testing"

	"github.com/Nexters/myply/infrastructure/configs"
)

func setup(t *testing.T) (func(t *testing.T), YoutubeClient) {
	t.Log("setup test case")
	youtubeClient, _ := NewYoutubeClient(&configs.Config{
		YoutubeAPIKey: os.Getenv("YOUTUBE_API_KEY"),
	})

	return func(t *testing.T) {
		t.Log("teardown test case")
	}, youtubeClient
}

// TestSearchPlaylist test real youtube api
func TestSearchPlaylist(t *testing.T) {
	t.Skip()
	teardown, youtubeClient := setup(t)
	defer teardown(t)

	res, _ := youtubeClient.SearchPlaylist("beenzino", "", "")

	t.Logf("%+v", res)
}
