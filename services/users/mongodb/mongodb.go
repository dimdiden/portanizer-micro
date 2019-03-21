package mongodb

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/dimdiden/portanizer-micro/services/users"
)

type repository struct {
	db     *mongo.Database
	logger log.Logger
}

const collName = "users"
const dbName = "portanizer"

func NewRepository(ctx context.Context, url string, logger log.Logger) (*repository, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	// set email as unique feild
	db := client.Database(dbName)
	collection := db.Collection(collName)
	_, err = collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, err
	}

	return &repository{
		db:     db,
		logger: log.With(logger, "repository", "mongodb"),
	}, nil
}

func (r *repository) InsertUser(ctx context.Context, email, pwd string) (*users.User, error) {
	collection := r.db.Collection(collName)

	hash, err := hashAndSalt(pwd)
	if err != nil {
		return nil, err
	}

	user := &users.User{Email: email, Password: string(hash)}

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			return nil, users.ErrExists
		default:
			level.Error(r.logger).Log("err", err)
			return nil, users.ErrQueryRepository
		}
	}
	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return user, nil
}

func hashAndSalt(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// func comparePasswords(hashPwd string, plainPwd string) error {
// 	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(plainPwd))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
