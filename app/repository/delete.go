package repository

import (
	"context"

	"github.com/mfaizfatah/story-tales/app/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) DeleteRedis(key string) (int64, error) {
	result, err := r.redis.Del(key).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (r *repo) DeleteLikes(storyid, episodeid, userid int) error {
	var data = new(models.Likes)
	result := r.db.
		Where("id_story = ?", storyid).
		Where("id_episodes = ?", episodeid).
		Where("id_users = ?", userid).
		Delete(&data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) DeleteRating(storyid, userid int) error {
	var data = new(models.Rating)
	result := r.db.
		Where("id_story = ?", storyid).
		Where("id_users = ?", userid).
		Delete(&data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) DeleteFavorite(storyid, userid int) error {
	var data = new(models.User_Favorite)
	result := r.db.
		Where("favorite_story = ?", storyid).
		Where("id_users = ?", userid).
		Delete(&data)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) MongoDeleteAll(tablename string, where interface{}, opt *options.DeleteOptions) (*mongo.DeleteResult, error) {
	result, err := r.mongo.Collection(tablename).DeleteMany(context.TODO(), where, opt)
	if err != nil {
		return nil, err
	}
	return result, nil
}
