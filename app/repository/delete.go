package repository

import (
	"github.com/mfaizfatah/story-tales/app/models"
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

func (r *repo) DeleteRating(storyid, episodeid, userid int) error {
	var data = new(models.Rating)
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
