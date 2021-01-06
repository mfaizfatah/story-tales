package repository

import (
	"log"
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

func (r *repo) FindGetOne(storyid int) (*models.ResponseOneStory, error) {
	var data = new(models.ResponseOneStory)
	rows, err := r.db.Table("story").
		Joins("INNER JOIN episodes ON episodes.id_story = story.id ").
		Select("story.id, story.title, story.sinopsis, story.season, story.images, story.flag_ongoing, story.flag_comment, story.id_author , episodes.id, episodes.eps_number, episodes.eps_title").
		Where("story.id = ?", storyid).
		Rows()

	defer rows.Close()
	for rows.Next() {
		var list models.ListEpisode
		err = rows.Scan(&data.ID, &data.Title, &data.Sinopsis, &data.Season, &data.Images, &data.FlagOnGoing, &data.FlagCommment, &data.IDAuthor, &list.ID, &list.Eps_Number, &list.Eps_Title)
		if err != nil {
			log.Panic(err)
		}
		data.ListEpisode = append(data.ListEpisode, list)
	}

	if err != nil {
		log.Panic(err)
		return data, err
	}

	return data, nil
}

func (r *repo) FindAll(table string) ([]models.ResponseAllStory, error) {
	var data []models.ResponseAllStory
	err := r.db.Table(table).Scan(&data)
	log.Printf("msg: %v", data)
	if err != nil {
		return data, nil
	}

	return nil, nil
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
