package repository

import (
	"time"
)

func (r *repo) Insert(table string, i interface{}) error {
	query := r.db.Table(table).Create(i)
	if query.Error != nil {
		return query.Error
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
