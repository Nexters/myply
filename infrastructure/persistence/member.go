package persistence

import (
	"context"
	"errors"

	"github.com/Nexters/myply/domain/member"
	"github.com/Nexters/myply/infrastructure/configs"
	"github.com/Nexters/myply/infrastructure/persistence/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type memberData struct {
	DeviceToken string   `json:"id" bson:"_id"`
	Name        string   `json:"name"`
	Keywords    []string `json:"keywords"`
}

type memberRepository struct {
	mongo  *db.MongoInstance
	config *configs.Config
}

const (
	memberCollectionName = "membmers"
)

func NewMemberRepository(mongo *db.MongoInstance, config *configs.Config) member.MemberRepository {
	return &memberRepository{mongo: mongo, config: config}
}

func (mr *memberRepository) Get(deviceToken string) (*member.Member, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mr.config.MongoTTL)
	defer cancel()

	collection := mr.mongo.Db.Collection(memberCollectionName)

	var result memberData
	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: deviceToken}}).Decode(&result); err != nil {
		return nil, err
	}

	return &member.Member{
		DeviceToken: result.DeviceToken,
		Name:        result.Name,
		Keywords:    result.Keywords,
	}, nil
}

func (mr *memberRepository) Create(entity member.Member) error {
	ctx, cancel := context.WithTimeout(context.Background(), mr.config.MongoTTL)
	defer cancel()

	collection := mr.mongo.Db.Collection(memberCollectionName)

	var result memberData
	findErr := collection.FindOne(ctx, bson.D{{Key: "_id", Value: entity.DeviceToken}}).Decode(&result)

	switch findErr {
	case mongo.ErrNoDocuments:
		data := memberData{
			DeviceToken: entity.DeviceToken,
			Name:        entity.Name,
			Keywords:    entity.Keywords,
		}
		_, insertErr := collection.InsertOne(ctx, data)

		return insertErr
	case nil:
		return errors.New("409: already account exist") // TODO: define domain exceptions
	default:
		return findErr
	}
}

func (mr *memberRepository) Update(
	deviceToken string,
	name *string,
	keywords []string,
) error {
	if name == nil && keywords == nil {
		return errors.New("invalid args")
	}

	ctx, cancel := context.WithTimeout(context.Background(), mr.config.MongoTTL)
	defer cancel()

	collection := mr.mongo.Db.Collection(memberCollectionName)
	filter := bson.D{{Key: "_id", Value: deviceToken}}
	updateSet := bson.D{}
	if name != nil {
		updateSet = append(updateSet, bson.E{Key: "name", Value: name})
	}
	if keywords != nil {
		updateSet = append(updateSet, bson.E{Key: "keywords", Value: keywords})
	}
	update := bson.D{{Key: "$set", Value: updateSet}}

	_, err := collection.UpdateOne(ctx, filter, update)

	return err
}
