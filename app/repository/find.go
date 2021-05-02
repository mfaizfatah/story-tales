package repository

import (
	"context"
	"log"
	"time"

	"github.com/go-shadow/moment"
	"github.com/mfaizfatah/story-tales/app/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) FindGetAuthorProfile(id int, table string) (*models.AuthorProfile, error) {
	var data models.AuthorProfile
	err := r.db.Table(table).
		Where("id = ?", id).
		Find(&data)
	if err.Error != nil {
		log.Printf("error: %v", err.Error)
		return nil, err.Error
	}
	return &data, nil
}
func (r *repo) FindAllComment(table string, storyid, episodeid int) ([]models.CommentView, error) {
	var data []models.CommentView
	err := r.db.Table(table).
		Where("id_story = ?", storyid).
		Where("id_episodes = ?", episodeid).
		Order("created_at").
		Find(&data)
	if err.Error != nil {
		log.Printf("error: %v", err.Error)
		return nil, err.Error
	}
	return data, nil
}

func (r *repo) FindFollowSt(userID, id int) (*models.UserFollow, error) {
	var data models.UserFollow
	err := r.db.Table("user_follow").
		Where("userfollow_id = ? AND userfollowing_id = ?", userID, id).
		Find(&data)
	if err.Error != nil {
		log.Printf("error: %v", err.Error)
		return nil, err.Error
	}
	return &data, nil
}

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

func (r *repo) FindMyLike(table string, userID int) ([]models.Likes, error) {
	var data []models.Likes
	err := r.db.Table(table).
		Where("id_users = ?", userID).
		Find(&data)
	log.Printf("msg: %v", data)
	if err != nil {
		return data, nil
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
			&data.ID, &data.ID_Author, &data.Title, &data.Sinopsis, &data.Season, &data.Images, &data.FlagOnGoing, &data.FlagCommment,
			&data.Publish_Date,
			&data.Rating,
			&data.Genre, &data.Author,
			&list.ID, &list.Publish_Status, &list.Publish_Date, &list.Eps_Number, &list.Eps_Title, &list.Images_Eps, &list.Like)

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
	rows, err := r.db.Table("detailEpisode").
		Where("id_story = ?", storyid).
		Where("id", episodeid).
		Rows()

	defer rows.Close()
	for rows.Next() {
		var detail models.Detail
		err = rows.Scan(&data.ID, &data.ID_Story, &data.Eps_Number, &data.Eps_Title, &detail.ID, &detail.Page, &detail.Schedule, &detail.Images)
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

func (r *repo) FindAuthorStory(table string, authorID int) ([]models.ResponseAllStory, error) {
	var data []models.ResponseAllStory
	err := r.db.Table(table).
		Where("id_author = ?", authorID).
		Find(&data)
	log.Printf("msg: %v", data)
	if err != nil {
		return data, nil
	}
	return data, nil
}

func (r *repo) FindMyComment(table string, userID int) ([]models.CommentView, error) {
	var data []models.CommentView
	err := r.db.Table(table).
		Where("id_users = ?", userID).
		Find(&data)
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

func (r *repo) MongoFindAll(where interface{}, tablename string, opt *options.FindOptions) (*mongo.Cursor, error) {
	result, err := r.mongo.Collection(tablename).Find(context.TODO(), where, opt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repo) DBFindAll(table string, i, where interface{}, field string, whereValue ...interface{}) error {
	err := r.db.Table(table).Where(where, whereValue...).Select(field).Find(i).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) MongoFindOne(i, where interface{}, TableName string) error {
	result := r.mongo.Collection(TableName).FindOne(context.TODO(), where)
	if result.Err() != nil {
		return result.Err()
	}

	result.Decode(i)
	return nil
}
