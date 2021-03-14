package repository

import (
	"log"
	"time"

	"github.com/go-shadow/moment"
	"github.com/mfaizfatah/story-tales/app/models"
)

func (r *repo) FindListFollower(id int) ([]models.ListFollower, error) {
	var data []models.ListFollower

	err := r.db.Table("followerView").
		Where("followerView.userfollowing_id = ?", id).
		Order("date_updated desc").
		Find(&data)

	if err.Error != nil {
		log.Printf("error: %v", err.Error)
		return nil, err.Error
	}

	return data, nil
}

func (r *repo) FindListFollowing(id int) ([]models.ListFollowing, error) {
	var data []models.ListFollowing

	err := r.db.Table("followingView").
		Where("followingView.userfollow_id = ?", id).
		Order("date_updated desc").
		Find(&data)

	if err.Error != nil {
		log.Printf("error: %v", err.Error)
		return nil, err.Error
	}

	return data, nil
}

func (r *repo) FindGetCountFollower(id int) (*models.UserCountFollower, error) {
	var data models.UserCountFollower

	err := r.db.Table("user_follow").
		Where("user_follow.userfollowing_id = ?", id).
		Count(&data.Count)

	if err.Error != nil {
		log.Printf("error: %v", err.Error)
		return nil, err.Error
	}

	return &data, nil
}

func (r *repo) FindGetCountFollowing(id int) (*models.UserCountFollowing, error) {
	var data models.UserCountFollowing

	err := r.db.Table("user_follow").
		Where("user_follow.userfollow_id = ?", id).
		Count(&data.Count)

	if err.Error != nil {
		log.Printf("error: %v", err.Error)
		return nil, err.Error
	}

	return &data, nil
}

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
	currentDate := moment.New().Now().GetTime()

	err := r.db.Table(table).
		Where("status = ?", 1).
		Where("validUntil >= ?", currentDate).
		Order("sequence desc").
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
	rows, err := r.db.Table("detailStoryOneView").
		Where("id = ?", storyid).
		Rows()

	defer rows.Close()
	for rows.Next() {
		var list models.ListEpisode
		err = rows.Scan(
			&data.ID, &data.Title, &data.Sinopsis, &data.Season, &data.Images, &data.FlagOnGoing, &data.FlagCommment,
			&data.Rating,
			&data.Genre, &data.Author,
			&list.ID, &list.Eps_Number, &list.Eps_Title, &list.Images_Eps, &list.Like)

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
	err := r.db.Table("rekomendasiView").
		Scan(&data)
	log.Printf("msg: %v", data)
	if err != nil {
		return data, nil
	}

	return data, nil
}

func (r *repo) FindFavoriteStory(table string, limit, storyid, userid int) ([]models.ResponseFavoriteStory, error) {
	var data []models.ResponseFavoriteStory

	if storyid != 0 && limit != 0 {
		err := r.db.Table("favoriteView").
			Where("id_users = ?", userid).
			Where("id > ?", storyid).
			Limit(limit).
			Scan(&data)
		if err != nil {
			return data, nil
		}
	} else if storyid != 0 && limit == 0 {
		err := r.db.Table("favoriteView").
			Where("id_users = ?", userid).
			Where("id = ?", storyid).
			Scan(&data)
		if err != nil {
			return data, nil
		}
	} else {
		err := r.db.Table("favoriteView").
			Where("id_users = ?", userid).
			Limit(limit).
			Scan(&data)
		if err != nil {
			return data, nil
		}
	}

	log.Printf("msg: %v", data)

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
