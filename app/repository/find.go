package repository

import (
	"time"

	"github.com/mfaizfatah/story-tales/app/models"
)

func (r *repo) FindOne(table string, i, where interface{}, field string, whereValue ...interface{}) error {
	err := r.db.Table(table).Where(where, whereValue...).Select(field).First(i).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) FindGetOne(table string, i, where interface{}, field string, value string) error {
	err := r.db.Table(table).Where(where, value).First(i).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *repo) FindAll(table string) ([]models.ResponseStory, error) {
	var data []models.ResponseStory
	err := r.db.Table(table).Scan(&data)

	if err != nil {
		return data, nil
	}

	return data, nil
}

func (r *repo) GetTTLRedis(key string) (int64, error) {
	result, err := r.redis.TTL(key).Result()
	if err != nil {
		return 0, err
	}
	exp := int64(result / time.Second)

	return exp, nil
}

func (r *repo) FindToken(key string) (string, error) {
	result, err := r.redis.Get(key).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}
