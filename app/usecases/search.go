package usecases

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	tableSearch = "search"
)

func (r *uc) Searching(ctx context.Context, query, genre string) (context.Context, interface{}, string, int, error) {
	var (
		res   []models.SearchModel
		msg   string
		st    = http.StatusOK
		where interface{}
	)

	res = make([]models.SearchModel, 0)

	/*
		where := bson.M{"$text": bson.M{"$search": query}}
		opt := options.FindOptions{}
		opt.SetSort(bson.M{"score": bson.M{"$meta": "textScore"}})
	*/
	char := "^" + query
	if genre != "" {
		where = bson.M{"$and": []bson.M{{"genre": genre}, {"$or": []bson.M{{"title": bson.M{"$regex": char, "$options": "i"}}, {"author": bson.M{"$regex": char, "$options": "i"}}}}}}
	} else {
		where = bson.M{"$or": []bson.M{{"title": bson.M{"$regex": char, "$options": "i"}}, {"author": bson.M{"$regex": char, "$options": "i"}}}}
	}

	cursor, err := r.query.MongoFindAll(where, tableSearch, nil)
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

func (r *uc) GenerateDocument(ctx context.Context) (context.Context, interface{}, string, int, error) {
	var (
		res interface{}
		msg string
		st  = http.StatusOK
		wg  sync.WaitGroup
	)

	delResult := make(chan *mongo.DeleteResult)
	delErr := make(chan error)
	go func() {
		defer close(delResult)
		del, err := r.query.MongoDeleteAll(tableSearch, bson.M{}, nil)
		delResult <- del
		delErr <- err
	}()

	listStory, err := r.query.FindRekomendasiStory(tableStory)
	if err != nil {
		msg = "Terjadi kesalahan pada sisi server. Coba beberapa saat lagi, Terima Kasih!"
		return ctx, nil, msg, http.StatusInternalServerError, err
	}
	if len(listStory) == 0 {
		res = make([]interface{}, 0)
		return ctx, nil, msg, http.StatusOK, err
	}

	docs := make([]interface{}, len(listStory))
	for i, story := range listStory {
		wg.Add(i)
		go func(i int, story models.ResponseRekomenStory) {
			defer wg.Done()
			doc := models.SearchModel{}
			doc.IDStory = story.ID
			doc.Title = story.Title
			doc.Author = story.Author
			doc.Genre = story.Genre
			doc.CreatedAt = time.Now()
			docs[i] = doc
		}(i, story)
	}

	deleteResult := <-delResult
	err = <-delErr
	if err != nil {
		ctx = logger.Logf(ctx, "delete all documents error() => %v", err)
	}

	ctx = logger.Logf(ctx, "delete result() => %v", deleteResult)

	result, err := r.query.MongoBulkInsert(tableSearch, docs, nil)
	if err != nil {
		msg = "Terjadi kesalahan pada sisi server. Coba beberapa saat lagi, Terima Kasih!"
		return ctx, nil, msg, http.StatusInternalServerError, err
	}

	res = fmt.Sprintf("Successfully insert %v documents", len(result.InsertedIDs))

	return ctx, res, msg, st, nil
}
