package integration_test

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"net/http"
	"os"
	"testing"

	. "github.com/Eun/go-hit"
)

const (
	serviceURL = "http://backend-face-service-dev:80/recognize"
)

type dataType [256]float64

// HTTP POST: /recognize
func TestRecognize(t *testing.T) {

	Test(t,
		Description("Wrong request type"),
		Get(serviceURL),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().String().Contains(``),
	)

	Test(t,
		Description("Database is empty"),
		Post(serviceURL),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(map[string]dataType{"encoding": generateData(1, 2)}),
		Expect().Status().Equal(http.StatusNotFound),
	)

	coll := connectToDatabase()
	ctx := context.TODO()
	docs := []interface{}{
		bson.D{{Key: "encoding", Value: generateData(1, 1)}},
		bson.D{{Key: "encoding", Value: generateData(2, 3)}},
	}
	result, _ := coll.InsertMany(ctx, docs)
	listIds := result.InsertedIDs

	defer func() {
		for _, key := range listIds {
			_, err := coll.DeleteOne(ctx, bson.D{{Key: "_id", Value: key}})
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()

	Test(t,
		Description("Wrong body (string)"),
		Post(serviceURL),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(map[string]string{"encoding": "hello"}),
		Expect().Status().Equal(http.StatusBadRequest),
	)

	var body [255]float64
	for i := 0; i < 255; i++ {
		body[i] = 1
	}

	Test(t,
		Description("There are no suitable users"),
		Post(serviceURL),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(map[string]dataType{"encoding": generateData(5, 6)}),
		Expect().Status().Equal(http.StatusNotFound),
	)

	Test(t,
		Description("UserId has been sent"),
		Post(serviceURL),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(map[string]dataType{"encoding": generateData(1, 1)}),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().JQ(".user_id").Equal(listIds[0].(primitive.ObjectID).Hex()),
	)

	Test(t,
		Description("Wrong body ([255]float64)"),
		Post(serviceURL),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().JSON(map[string][255]float64{"encoding": body}),
		Expect().Status().Equal(http.StatusBadRequest),
	)

}

func generateData(min, max float64) (data dataType) {
	for i := 0; i < 256; i++ {
		data[i] = min + rand.Float64()*(max-min)
	}
	return
}

func connectToDatabase() (collection *mongo.Collection) {
	ctx := context.TODO()
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@mongodb-service:27017/%s?authSource=admin",
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGODB_DATABASE_NAME"))))
	if err != nil {
		log.Fatalln(err)
	}

	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalln(err)
	}

	collection = client.Database(os.Getenv("MONGODB_DATABASE_NAME")).Collection(os.Getenv("MONGODB_USERS_COLLECTION"))
	return
}
