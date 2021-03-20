package usecases

import (
	"context"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/models"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	tableSearch = "search"
)

func (r *uc) Searching(ctx context.Context, query string) (context.Context, interface{}, string, int, error) {
	var (
		res []models.SearchModel
		msg string
		st  = http.StatusOK
	)

	res = make([]models.SearchModel, 0)

	where := bson.M{"$text": bson.M{"$search": query}}
	opt := options.FindOptions{}
	opt.SetSort(bson.M{"score": bson.M{"$meta": "textScore"}})
	cursor, err := r.query.MongoFindAll(where, tableSearch, &opt)
	if err != nil {
		msg = "Terjadi kesalahan pada sisi server. Coba beberapa saat lagi, Terima Kasih!"
		return ctx, nil, msg, http.StatusInternalServerError, err
	}

	for cursor.Next(context.TODO()) {
		var i models.SearchModel
		err := cursor.Decode(&i)
		if err != nil {
			msg = "Terjadi kesalahan pada sisi server. Coba beberapa saat lagi, Terima Kasih!"
			return ctx, nil, msg, http.StatusInternalServerError, err
		}
		res = append(res, i)
	}

	return ctx, res, msg, st, nil
}
