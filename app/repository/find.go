package repository

import (
	"log"
	"time"

	"github.com/mfaizfatah/story-tales/app/models"
)

func (r *repo) FindGetBanner(id int) (*models.BannerDetailRs, error) {
	var data models.BannerDetailRs

	err := r.db.Table("banner").
		Where("banner.id = ?", id).
		Find(&data)

	if err.Error != nil {
		log.Printf("error: %v", err.Error)
		return nil, err.Error
	}

	return &data, nil
}

func (r *repo) FindAllBanner(table string) ([]models.ListBannerThumbRs, error) {
	var data []models.ListBannerThumbRs

	err := r.db.Table(table).
		Where("status = ?", 1).
		Order("sequence").
		Find(&data)

	if err.Error != nil {
		log.Printf("error: %v", err.Error)
		return nil, err.Error
	}

	return data, nil
}

func (r *repo) FindOne(table string, i, where interface{}, field string, whereValue ...interface{}) error {
	err := r.db.Table(table).Where(where, whereValue...).Select(field).First(i).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) FindGetOneStory(storyid int) (*models.ResponseOneStory, error) {
	var data = new(models.ResponseOneStory)
	rows, err := r.db.Table("story").
		Joins("LEFT JOIN episodes ON episodes.id_story = story.id ").
		Select("story.id, story.title, story.sinopsis, story.season, story.images, story.flag_ongoing, story.flag_comment, story.id_author , COALESCE(episodes.id,0), COALESCE(episodes.eps_number,0), COALESCE(episodes.eps_title,'')").
		Where("story.id = ?", storyid).
		Rows()

	defer rows.Close()
	for rows.Next() {
		var list models.ListEpisode
		err = rows.Scan(&data.ID, &data.Title, &data.Sinopsis, &data.Season, &data.Images, &data.FlagOnGoing, &data.FlagCommment, &data.IDAuthor, &list.ID, &list.Eps_Number, &list.Eps_Title)
		if err != nil {
			log.Panic(err)
		}
		if list.ID != 0 {
			data.ListEpisode = append(data.ListEpisode, list)
		} else {
			data.ListEpisode = nil
		}

	}

	if err != nil {
		log.Panic(err)
		return data, err
	}

	return data, nil
}

func (r *repo) FindGetDetailEpisode(storyid, episodeid int) (*models.ResponseDetailEpisode, error) {
	var data = new(models.ResponseDetailEpisode)
	rows, err := r.db.Table("episodes").
		Joins("LEFT JOIN episodes_details ON episodes_details.id_episodes = episodes.id ").
		Select("episodes.id, episodes.eps_number, episodes.eps_title, COALESCE(episodes_details.id,0), COALESCE(episodes_details.page,0), COALESCE(episodes_details.schedule,''), COALESCE(episodes_details.images,'')").
		Where("episodes.id_story = ?", storyid).
		Where("episodes.id = ?", episodeid).
		Rows()

	defer rows.Close()
	for rows.Next() {
		var detail models.Detail
		err = rows.Scan(&data.ID, &data.Eps_Number, &data.Eps_Title, &detail.ID, &detail.Page, &detail.Schedule, &detail.Images)
		if err != nil {
			log.Panic(err)
		}
		if detail.ID != 0 {
			data.Detail = append(data.Detail, detail)
		} else {
			data.Detail = nil
		}

	}

	if err != nil {
		log.Panic(err)
		return data, err
	}

	return data, nil
}

func (r *repo) FindAllStory(table string) ([]models.ResponseAllStory, error) {
	var data []models.ResponseAllStory
	err := r.db.Table(table).
		Scan(&data)
	log.Printf("msg: %v", data)
	if err != nil {
		return data, nil
	}

	return data, nil
}

func (r *repo) FindRekomendasiStory(table string) ([]models.ResponseRekomenStory, error) {
	var data []models.ResponseRekomenStory
	err := r.db.Table(table).
		Select("story.id, story.title, story.images, genre.genre, author.name AS author").
		Joins("LEFT JOIN story_genre ON story.id = story_genre.id_story").
		Joins("LEFT JOIN genre ON genre.id = story_genre.id_genre").
		Joins("LEFT JOIN author ON author.id = story.id_author").
		Limit(10).
		Scan(&data)
	log.Printf("msg: %v", data)
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
