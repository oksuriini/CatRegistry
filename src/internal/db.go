package internal

import (
	"CatRegistry/src/internal/cat"
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConnection struct {
	Connection     *mongo.Client
	Uri            string
	authentication options.Credential
}

func CreateMongoDBConnection(uri string) *MongoDBConnection {
	connection, err := createConnectionStruct(uri)
	if err != nil {
		return nil
	}
	return &MongoDBConnection{
		Uri:        uri,
		Connection: connection,
	}
}

func createConnectionStruct(uri string) (*mongo.Client, error) {
	credentOpts := options.Credential{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	clientOpts := options.Client().ApplyURI(uri).SetAuth(credentOpts)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *MongoDBConnection) InsertCat(elGato cat.Cat) error {
	collect := c.Connection.Database("animals").Collection("cats")
	_, err := collect.InsertOne(context.TODO(), elGato)
	if err != nil {
		return err
	}
	fmt.Println("Inserted cat into database")
	return nil
}

func (c *MongoDBConnection) GetCatsByFilter(filter cat.Cat) ([]cat.Cat, error) {
	filterCat := bson.D{
		{"$or", bson.A{
			bson.D{{"name", filter.Name}},
			bson.D{{"breed", filter.Breed}},
			bson.D{{"age", filter.Age}},
			bson.D{{"lives", filter.Lives}},
		}},
	}
	var gottenCats []cat.Cat
	collect := c.Connection.Database("animals").Collection("cats")
	cursor, err := collect.Find(context.TODO(), filterCat)
	if err != nil {
		fmt.Println("Error in finding cats")
		fmt.Println(err)
	}
	cursor.All(context.TODO(), &gottenCats)
	fmt.Println(gottenCats)
	return gottenCats, nil
}

func (c *MongoDBConnection) GetAllCats() []cat.Cat {
	var allCats []cat.Cat
	collect := c.Connection.Database("animals").Collection("cats")
	cursor, err := collect.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Error while searching cats, returning zero cats")
		fmt.Println(err)
		return nil
	}
	err = cursor.All(context.TODO(), &allCats)
	if err != nil {
		fmt.Println("Error while getting cats, returning zero cats")
		fmt.Println(err)
		return nil
	}
	return allCats
}
