package persistence

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Nexters/myply/domain/memos"
	"github.com/Nexters/myply/infrastructure/persistence/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName = "memos"

type MemoData struct {
	ID             primitive.ObjectID  `json:"id" bson:"_id"`
	DeviceToken    string              `bson:"deviceToken"`
	YoutubeVideoID string              `bson:"youtubeVideoId"`
	Body           string              `bson:"body"`
	Tags           []string            `bson:"tags"`
	CreatedAt      primitive.Timestamp `bson:"createdAt"`
	UpdatedAt      primitive.Timestamp `bson:"updatedAt"`
}

func (m *MemoData) toEntity() *memos.Memo {
	return &memos.Memo{
		ID:             m.ID.Hex(),
		DeviceToken:    m.DeviceToken,
		YoutubeVideoID: m.YoutubeVideoID,
		Body:           m.Body,
		Tags:           m.Tags,
		CreatedAt:      time.Unix(int64(m.CreatedAt.T), 0),
		UpdatedAt:      time.Unix(int64(m.UpdatedAt.T), 0),
	}
}

type MemoMongoRepository struct {
	conn *mongo.Database
}

func NewMemoRepository(i *db.MongoInstance) memos.Repository {
	return &MemoMongoRepository{conn: i.Db}
}

func (r *MemoMongoRepository) GetMemo(id string) (*memos.Memo, error) {
	coll := r.getCollection()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	md := MemoData{}
	if err = coll.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&md); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, &memos.NotFoundError{Msg: fmt.Sprintf("memo is not found. id=%s", id)}
		default:
			return nil, err
		}
	}

	return md.toEntity(), nil
}

func (r *MemoMongoRepository) GetMemoByVideoID(id string) (*memos.Memo, error) {
	coll := r.getCollection()

	md := MemoData{}
	if err := coll.FindOne(context.Background(), bson.M{"youtubeVideoId": id}).Decode(&md); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, &memos.NotFoundError{Msg: fmt.Sprintf("memo is not found. videoID=%s", id)}
		default:
			return nil, err
		}
	}

	return md.toEntity(), nil
}

func (r *MemoMongoRepository) AddMemo(deviceToken string, videoID string, body string, tags []string) (string, error) {
	coll := r.getCollection()

	memoID := primitive.NewObjectID()
	now := r.now()

	md := MemoData{
		ID:             memoID,
		DeviceToken:    deviceToken,
		YoutubeVideoID: videoID,
		Body:           body,
		Tags:           tags,
		CreatedAt:      *now,
		UpdatedAt:      *now,
	}

	insertResult, err := coll.InsertOne(context.Background(), md)
	if err != nil {
		return "", err
	}

	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *MemoMongoRepository) UpdateBody(id string, body string) error {
	coll := r.getCollection()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "body", Value: body}}}}
	if _, err = coll.UpdateOne(context.Background(), filter, update); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return &memos.NotFoundError{Msg: fmt.Sprintf("memo is not found. id=%s", id)}
		default:
			return err
		}
	}

	return nil
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
