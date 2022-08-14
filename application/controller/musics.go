package controller

import (
	"errors"

	"github.com/Nexters/myply/domain/musics"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type MusicController interface {
	Search() fiber.Handler
	Retrieve() fiber.Handler
}

type musicController struct {
	logger       *zap.SugaredLogger
	musicService musics.Service
}

func NewMusicController(l *zap.SugaredLogger, ms musics.Service) MusicController {
	return &musicController{logger: l, musicService: ms}
}

type Order string

const (
	RecentOrder  Order = "recent"
	PopularOrder       = "count"
)

func (o Order) convert() (musics.Order, error) {
	switch o {
	case RecentOrder:
		return musics.OrderByDate, nil
	case PopularOrder:
		return musics.OrderByViewCount, nil
	default:
		return "", errors.New("unsupported order")
	}
}

type SearchQueryParams struct {
	Q         []string `query:"q"`
	NextToken string   `query:"nextToken"`
}

type RetrieveQueryParams struct {
	Order     Order  `query:"order"`
	NextToken string `query:"nextToken"`
}

type MusicResponse struct {
	YoutubeVideoID string   `json:"youtubeVideoID"`
	ThumbnailURL   string   `json:"thumbnailURL"`
	Title          string   `json:"title"`
	YoutubeTags    []string `json:"youtubeTags"`
	VideoDeepLink  string   `json:"videoDeepLink"`
	IsMemoed       bool     `json:"isMemoed"`
}

type ListMusicData struct {
	Musics        []MusicResponse `json:"musics"`
	NextPageToken string          `json:"nextPageToken,omitempty"`
}

type ListMusicResponse struct {
	BaseResponse
	Data ListMusicData `json:"data"`
}

// @Summary      Search music playlist
// @Description  플레이리스트 검색
// @Tags         musics
// @Accept       json
// @Produce      json
// @Param payload query SearchQueryParams true "query params"
// @Success      200  {object}   ListMusicResponse
// @Failure      500
// @Router       /musics/search [get]
// @Security ApiKeyAuth
func (mc *musicController) Search() fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		var (
			musicListDto *musics.MusicListDto
			err          error
		)
		p := new(SearchQueryParams)
		err = ctx.QueryParser(p)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		musicListDto, err = mc.musicService.GetMusicList(p.Q, p.NextToken)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		musicList := []MusicResponse{}
		for _, m := range *musicListDto.Musics {
			musicList = append(musicList, MusicResponse{
				YoutubeVideoID: m.YoutubeVideoID,
				ThumbnailURL:   m.ThumbnailURL,
				Title:          m.Title,
				YoutubeTags:    m.YoutubeTags,
				VideoDeepLink:  m.DeepLink(),
				IsMemoed:       false, // TODO: memo service
			})
		}
		return ctx.Status(200).JSON(BaseResponse{
			Code: Ok,
			Data: ListMusicData{
				Musics:        musicList,
				NextPageToken: musicListDto.NextPageToken,
			},
		})
	}
}

// @Summary      Retrieve music playlist
// @Description  플레이리스트 조회
// @Tags         musics
// @Accept       json
// @Produce      json
// @Param payload query RetrieveQueryParams true "query params"
// @Success      200  {object}   ListMusicResponse
// @Failure      500
// @Router       /musics [get]
// @Security ApiKeyAuth
func (mc *musicController) Retrieve() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		p := new(RetrieveQueryParams)
		err := ctx.QueryParser(p)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		order, err := p.Order.convert()
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		musicListDto, err := mc.musicService.GetPlayListBy(order, p.NextToken)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		var musicList []MusicResponse
		for _, m := range *musicListDto.Musics {
			musicList = append(musicList, MusicResponse{
				YoutubeVideoID: m.YoutubeVideoID,
				ThumbnailURL:   m.ThumbnailURL,
				Title:          m.Title,
				YoutubeTags:    m.YoutubeTags,
				VideoDeepLink:  m.DeepLink(),
				IsMemoed:       false, // TODO: memo service
			})
		}
		return ctx.Status(200).JSON(BaseResponse{
			Code: Ok,
			Data: ListMusicData{
				Musics:        musicList,
				NextPageToken: musicListDto.NextPageToken,
			},
		})
	}
}
