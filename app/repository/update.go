package repository

import (
	"context"

	"github.com/mfaizfatah/story-tales/app/models"
)

/*Update to database
 * @paremeter
 * i - struct to saving into database
 *
 * @return
 * uint - id after insert into database
 * error
 */
func (r *repo) Update(tableName string, i interface{}, data map[string]interface{}) error {
	query := r.db.Table(tableName).Model(i).Updates(data)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *repo) UpdateWhere(tableName string, i, where interface{}, data map[string]interface{}, whereValue ...interface{}) error {
	query := r.db.Table(tableName).Model(i).Where(where, whereValue...).Updates(data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
func (r *repo) UpdateData(tableName string, i, where interface{}, data interface{}, whereValue ...interface{}) error {
	query := r.db.Table(tableName).Model(i).Where(where, whereValue...).Updates(data)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *repo) MongoUpdateOne(data, where interface{}, TableName string) error {
	_, err := r.mongo.Collection(TableName).UpdateOne(context.TODO(), where, data)
	if err != nil {
		return err
	}
	return nil

}

func (r *repo) UpdateStory(story *models.Story) error {
	err := r.db.Table("story").Where("id = ?", story.ID).Save(story).Error
	if err != nil {
		return err
	}

	return err

}
