package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Region struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	RegionId uint32             `bson:"region_id"`
	Ec2bData []byte             `bson:"ec2b_data"`
	NextUid  uint32             `bson:"next_uid"`
}

func (d *Dao) InsertRegion(region *Region) error {
	db := d.db.Collection("region")
	region.RegionId = 1
	_, err := db.InsertOne(context.TODO(), region)
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) DeleteRegion() error {
	db := d.db.Collection("region")
	_, err := db.DeleteOne(context.TODO(), bson.D{{"region_id", 1}})
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) UpdateRegion(region *Region) error {
	db := d.db.Collection("region")
	region.RegionId = 1
	_, err := db.UpdateMany(
		context.TODO(),
		bson.D{{"region_id", 1}},
		bson.D{{"$set", region}},
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) QueryRegion() (*Region, error) {
	db := d.db.Collection("region")
	result := db.FindOne(
		context.TODO(),
		bson.D{{"region_id", 1}},
	)
	region := new(Region)
	err := result.Decode(region)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return region, nil
}
