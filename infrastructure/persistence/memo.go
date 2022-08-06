package persistence

import (
	"context"
	"errors"
	"github.com/Nexters/myply/domain/memos"
	"github.com/Nexters/myply/infrastructure/persistence/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var collectionName = "memos"

type MemoData struct {
	ID             primitive.ObjectID   `json:"id" bson:"_id"`
	DeviceToken    string               `json:"deviceToken"`
	YoutubeVideoId string               `json:"youtubeVideoId"`
	Body           string               `json:"body"`
	TagIds         []primitive.ObjectID `json:"tagIds"`
	CreatedAt      primitive.Timestamp  `json:"createdAt"`
	UpdatedAt      primitive.Timestamp  `json:"updatedAt"`
}

func (m *MemoData) toEntity() *memos.Memo {
	var tagIds []string
	for _, id := range m.TagIds {
		tagIds = append(tagIds, id.Hex())
	}

	return &memos.Memo{
		Id:             m.ID.Hex(),
		DeviceToken:    m.DeviceToken,
		YoutubeVideoId: m.YoutubeVideoId,
		Body:           m.Body,
		TagIds:         tagIds,
		CreatedAt:      time.Unix(int64(m.CreatedAt.T), 0),
		UpdatedAt:      time.Unix(int64(m.UpdatedAt.T), 0),
	}
}

type MemoMongoRepository struct {
	conn *mongo.Database
}

func NewMemoRepository(i *db.MongoInstance) *memos.Repository {
	var r memos.Repository
	r = &MemoMongoRepository{conn: i.Db}
	return &r
}

func (r *MemoMongoRepository) GetMemo(id string) (*memos.Memo, error) {
	coll := r.getCollection()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	md := MemoData{}
	if err = coll.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&md); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, memos.NotFoundException
		default:
			return nil, err
		}
	}

	return md.toEntity(), nil
}

func (r *MemoMongoRepository) GetMemoByVideoId(id string) (*memos.Memo, error) {
	coll := r.getCollection()

	md := MemoData{}
	if err := coll.FindOne(context.Background(), bson.M{"youtubeVideoId": id}).Decode(&md); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, memos.NotFoundException
		default:
			return nil, err
		}
	}

	return md.toEntity(), nil
}

func (r *MemoMongoRepository) SaveMemo(videoId string, body string, deviceToken string) (string, error) {
	coll := r.getCollection()

	memoId := primitive.NewObjectID()
	now := r.now()
	md := MemoData{
		ID:             memoId,
		DeviceToken:    deviceToken,
		YoutubeVideoId: videoId,
		Body:           body,
		TagIds:         nil, // TODO: change to real data
		CreatedAt:      *now,
		UpdatedAt:      *now,
	}

	insertResult, err := coll.InsertOne(context.Background(), md)
	if err != nil {
		return "", err
	}

	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *MemoMongoRepository) getCollection() *mongo.Collection {
	return r.conn.Collection(collectionName)
}

func (r *MemoMongoRepository) now() *primitive.Timestamp {
	return &primitive.Timestamp{
		T: uint32(time.Now().Unix()),
		I: 0,
	}
}
