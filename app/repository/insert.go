package repository

import (
	"context"
	"log"
	"time"

	"github.com/mfaizfatah/story-tales/app/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) Insert(table string, i interface{}) error {
	query := r.db.Table(table).Create(i)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *repo) InsertFollow(*models.UserFollow) error {
	var data models.UserFollow

	err := r.db.Table("user_follow").
		Where("user_follow.userfollow_id = ? AND user_follow.userfollowing_id = ?", data.UserFollowID, data.UserFollowingID).
		Find(&data)

	if err.Error != nil {
		log.Printf("error: %v", err.Error)
		return err.Error
	}
	var countModel = 1

	if countModel == 0 {
		query := r.db.Table("user_follow").Create(&data)
		if query.Error != nil {
			return query.Error
		}
		return nil
	}
	if countModel > 0 {
		query := r.db.Table("user_follow").Delete(&data)
		if query.Error != nil {
			return query.Error
		}
		return nil
	}

	return nil
}

/* func (r *repo) InsertStory(table string, story interface{}, episode []models.Episode, episodeDetail []models.Episodes_Detail) error {
	err := r.db.Table(table).Create(story).Error

	for i := range episode {
		if err := r.db.Table("episodes").Create(episode[i]).Error; err != nil {
			log.Error(err)
		}
	}

	for i := range episodeDetail {
		if err := r.db.Table("episodes_details").Create(episodeDetail[i]).Error; err != nil {
			log.Error(err)
		}
	}

	if err != nil {
		return err
	}
	return nil
} */

func (r *repo) SetRedis(key string, value interface{}, exp time.Duration) error {
	err := r.redis.Set(key, value, exp).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) MongoBulkInsert(tablename string, doc []interface{}, opt *options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	result, err := r.mongo.Collection(tablename).InsertMany(context.Background(), doc, opt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repo) MongoInsert(tablename string, doc interface{}, opt *options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	result, err := r.mongo.Collection(tablename).InsertOne(context.Background(), doc, opt)
	if err != nil {
		return nil, err
	}
	return result, nil
}
