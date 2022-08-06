package clients

import "github.com/google/wire"

var Set = wire.NewSet(NewYoutubeClient)
