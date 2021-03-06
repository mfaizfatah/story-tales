package repository

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/mfaizfatah/story-tales/app/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

var (
	// ErrRecordNotFound record not found error
	ErrRecordNotFound = errors.New("record not found")
	// ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
	ErrInvalidTransaction = errors.New("no valid transaction")
	// ErrNotImplemented not implemented
	ErrNotImplemented = errors.New("not implemented")
	// ErrMissingWhereClause missing where clause
	ErrMissingWhereClause = errors.New("WHERE conditions required")
	// ErrUnsupportedRelation unsupported relations
	ErrUnsupportedRelation = errors.New("unsupported relations")
	// ErrPrimaryKeyRequired primary keys required
	ErrPrimaryKeyRequired = errors.New("primary key required")
	// ErrModelValueRequired model value required
	ErrModelValueRequired = errors.New("model value required")
	// ErrInvalidData unsupported data
	ErrInvalidData = errors.New("unsupported data")
	// ErrUnsupportedDriver unsupported driver
	ErrUnsupportedDriver = errors.New("unsupported driver")
	// ErrRegistered registered
	ErrRegistered = errors.New("registered")
	// ErrInvalidField invalid field
	ErrInvalidField = errors.New("invalid field")
	// ErrEmptySlice empty slice found
	ErrEmptySlice = errors.New("empty slice found")
	// ErrDryRunModeUnsupported dry run mode unsupported
	ErrDryRunModeUnsupported = errors.New("dry run mode unsupported")
	// ErrConflict for error if data conflict
	ErrConflict = errors.New("data conflict")
	// ErrBadRequest for error bad request
	ErrBadRequest = errors.New("bad request")
	// ErrUnouthorized for error authorization
	ErrUnouthorized = errors.New("unouthorized")
)

// repo struct with value mysqldb connection
type repo struct {
	db    *gorm.DB
	redis *redis.Client
	mongo *mongo.Database
}

// Repo represent the Repository contract
type Repo interface {
	// find
	FindGetAuthorProfile(id int, table string) (*models.AuthorProfile, error)
	FindGetUserProfile(id int, table string) (*models.UserEdit, error)
	FindGetBanner(id int) (*models.BannerDetailRs, error)
	FindAllBanner(table string) ([]models.ListBannerThumbRs, error)
	FindOne(table string, i, where interface{}, field string, whereValue ...interface{}) error
	DBFindAll(table string, i, where interface{}, field string, whereValue ...interface{}) error
	FindMyLike(table string, userID int) ([]models.Likes, error)

	//find Story
	FindGetOneStory(storyid int) (*models.ResponseOneStory, error)
	FindGetDetailEpisode(storyid, episodeid int) (*models.ResponseDetailEpisode, error)
	FindAllStory(table string) ([]models.ResponseAllStory, error)
	FindAuthorStory(table string, authorID int) ([]models.ResponseAllStory, error)
	FindRekomendasiStory(table string) ([]models.ResponseRekomenStory, error)
	FindFavoriteStory(table string, limit, storyid, userid int) ([]models.ResponseFavoriteStory, error)

	// find comment
	FindTopComment(table string, storyID, episodeID int) ([]models.CommentView, error)
	FindAllComment(table string, storyID, episodeID int) ([]models.CommentView, error)
	FindMyComment(table string, authorID int) ([]models.CommentView, error)
	FindCommentLikeSt(commentID, userID int) (*models.CommentLike, error)

	GetTTLRedis(key string) (int64, error)
	FindToken(key string) (string, error)
	FindFollowSt(userID, id int) (*models.UserFollow, error)
	FindGetCountFollower(id int) (*models.UserCountFollower, error)
	FindGetCountFollowing(id int) (*models.UserCountFollowing, error)
	FindListFollower(id int) ([]models.ListFollower, error)
	FindListFollowing(id int) ([]models.ListFollowing, error)

	// insert
	Insert(table string, i interface{}) error
	InsertFollow(*models.UserFollow) error
	/* 	InsertStory(table string, story interface{}, episode []models.Episode, episodeDetail []models.Episodes_Detail) error */
	SetRedis(key string, value interface{}, exp time.Duration) error

	// Update
	Update(tableName string, i interface{}, data map[string]interface{}) error
	UpdateWhere(tableName string, i interface{}, where interface{}, data map[string]interface{}, whereValue ...interface{}) error
	UpdateData(tableName string, i, where interface{}, data interface{}, whereValue ...interface{}) error
	UpdateStory(story *models.Story) error

	// delete
	Delete(table string, i interface{}) error
	DeleteRedis(key string) (int64, error)
	DeleteFavorite(storyid, userid int) error
	DeleteLikes(storyid, episodeid, userid int) error
	DeleteRating(storyid, userid int) error

	// mongo find
	MongoFindAll(where interface{}, tablename string, opt *options.FindOptions) (*mongo.Cursor, error)
	MongoFindOne(i, where interface{}, TableName string) error

	// mongo insert
	MongoBulkInsert(tablename string, doc []interface{}, opt *options.InsertManyOptions) (*mongo.InsertManyResult, error)
	MongoInsert(tablename string, doc interface{}, opt *options.InsertOneOptions) (*mongo.InsertOneResult, error)

	// mongo delete
	MongoDeleteAll(tablename string, where interface{}, opt *options.DeleteOptions) (*mongo.DeleteResult, error)

	// mongo update
	MongoUpdateOne(data, where interface{}, TableName string) error
}

/*NewRepo will create an object that represent the Repository interface (Repo)
 * @parameter
 * db - mysql database connection
 *
 * @represent
 * interface Repo
 *
 * @return
 * repo struct with value db (mysql database connection)
 */
func NewRepo(db *gorm.DB, redis *redis.Client, mongo *mongo.Database) Repo {
	return &repo{db: db, redis: redis, mongo: mongo}
}
