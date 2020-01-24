package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// Tonnage represents the client for Tonnage table
type Tonnage struct{}

// NewTonnage returns a new NewTonnage database instance
func NewTonnage() *Tonnage {
	return &Tonnage{}
}

// // BasinBalance returns TimeCountObject
// func (u *Tonnage) BasinBalance(c echo.Context, basin string, seg string) (*model.TimeCountObject, error) {
// 	return u.udb.BasinBalance(u.db, basin, seg)
// }

// // DryDockTimeseries returns TimeseriesObject
// func (u *Tonnage) DryDockTimeseries(c echo.Context, seg string, metric string, absolute bool) (*model.TimeseriesObject, error) {
// 	return u.udb.DryDockTimeseries(u.db, seg, metric, absolute)
// }

// // DryDockSummaryStats returns DryDockSummaryStatsObject
// func (u *Tonnage) DryDockSummaryStats(c echo.Context, seg string) (*model.DryDockSummaryStatsObject, error) {
// 	return u.udb.DryDockSummaryStats(u.db, seg)
// }

// BasinBalance returns TimeCountObject
func (u *Tonnage) BasinBalance(db *mongo.Database, basin, seg string) (*model.TimeCountObject, error) {
	var result []model.BasinBalanceFullRow
	collection := db.Collection(model.ColBasinBalance)
	ctx := context.Background()

	cur, err := collection.Find(ctx,
		bson.M{
			"basin":   primitive.Regex{Pattern: basin, Options: ""},
			"segment": primitive.Regex{Pattern: seg, Options: ""},
		})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &result)
	return model.ConvertBasinTimeCount(result), err
}

// DryDockTimeseries returns TimeseriesObject
func (u *Tonnage) DryDockTimeseries(db *mongo.Database, seg, metric string, absolute bool) (*model.TimeseriesObject, error) {
	var result []model.DryDockFullRow
	collection := db.Collection(model.ColDryDock)
	ctx := context.Background()

	cur, err := collection.Find(ctx,
		bson.M{
			"segment": primitive.Regex{Pattern: seg, Options: ""},
		})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &result)
	return model.ConvertDryDockTimeseries(result, metric, absolute), err
}

// DryDockSummaryStats returns DryDockSummaryStatsObject
func (u *Tonnage) DryDockSummaryStats(db *mongo.Database, seg string) (*model.DryDockSummaryStatsObject, error) {
	var result []model.DryDockSummaryStatsFullRow
	collection := db.Collection(model.ColDryDockStays)
	ctx := context.Background()

	cur, err := collection.Find(ctx,
		bson.M{
			"segment": primitive.Regex{Pattern: seg, Options: ""},
		})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &result)
	return model.ConvertDryDockSummaryStats(result, seg), err
}
