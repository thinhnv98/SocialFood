// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRestaurantPhotos(t *testing.T) {
	t.Parallel()

	query := RestaurantPhotos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRestaurantPhotosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RestaurantPhotos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRestaurantPhotosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RestaurantPhotos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RestaurantPhotos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRestaurantPhotosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RestaurantPhotoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RestaurantPhotos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRestaurantPhotosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RestaurantPhotoExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RestaurantPhoto exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RestaurantPhotoExists to return true, but got false.")
	}
}

func testRestaurantPhotosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	restaurantPhotoFound, err := FindRestaurantPhoto(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if restaurantPhotoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRestaurantPhotosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RestaurantPhotos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRestaurantPhotosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RestaurantPhotos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRestaurantPhotosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	restaurantPhotoOne := &RestaurantPhoto{}
	restaurantPhotoTwo := &RestaurantPhoto{}
	if err = randomize.Struct(seed, restaurantPhotoOne, restaurantPhotoDBTypes, false, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}
	if err = randomize.Struct(seed, restaurantPhotoTwo, restaurantPhotoDBTypes, false, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = restaurantPhotoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = restaurantPhotoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RestaurantPhotos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRestaurantPhotosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	restaurantPhotoOne := &RestaurantPhoto{}
	restaurantPhotoTwo := &RestaurantPhoto{}
	if err = randomize.Struct(seed, restaurantPhotoOne, restaurantPhotoDBTypes, false, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}
	if err = randomize.Struct(seed, restaurantPhotoTwo, restaurantPhotoDBTypes, false, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = restaurantPhotoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = restaurantPhotoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RestaurantPhotos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func restaurantPhotoBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RestaurantPhoto) error {
	*o = RestaurantPhoto{}
	return nil
}

func restaurantPhotoAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RestaurantPhoto) error {
	*o = RestaurantPhoto{}
	return nil
}

func restaurantPhotoAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RestaurantPhoto) error {
	*o = RestaurantPhoto{}
	return nil
}

func restaurantPhotoBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RestaurantPhoto) error {
	*o = RestaurantPhoto{}
	return nil
}

func restaurantPhotoAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RestaurantPhoto) error {
	*o = RestaurantPhoto{}
	return nil
}

func restaurantPhotoBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RestaurantPhoto) error {
	*o = RestaurantPhoto{}
	return nil
}

func restaurantPhotoAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RestaurantPhoto) error {
	*o = RestaurantPhoto{}
	return nil
}

func restaurantPhotoBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RestaurantPhoto) error {
	*o = RestaurantPhoto{}
	return nil
}

func restaurantPhotoAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RestaurantPhoto) error {
	*o = RestaurantPhoto{}
	return nil
}

func testRestaurantPhotosHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RestaurantPhoto{}
	o := &RestaurantPhoto{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto object: %s", err)
	}

	AddRestaurantPhotoHook(boil.BeforeInsertHook, restaurantPhotoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	restaurantPhotoBeforeInsertHooks = []RestaurantPhotoHook{}

	AddRestaurantPhotoHook(boil.AfterInsertHook, restaurantPhotoAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	restaurantPhotoAfterInsertHooks = []RestaurantPhotoHook{}

	AddRestaurantPhotoHook(boil.AfterSelectHook, restaurantPhotoAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	restaurantPhotoAfterSelectHooks = []RestaurantPhotoHook{}

	AddRestaurantPhotoHook(boil.BeforeUpdateHook, restaurantPhotoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	restaurantPhotoBeforeUpdateHooks = []RestaurantPhotoHook{}

	AddRestaurantPhotoHook(boil.AfterUpdateHook, restaurantPhotoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	restaurantPhotoAfterUpdateHooks = []RestaurantPhotoHook{}

	AddRestaurantPhotoHook(boil.BeforeDeleteHook, restaurantPhotoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	restaurantPhotoBeforeDeleteHooks = []RestaurantPhotoHook{}

	AddRestaurantPhotoHook(boil.AfterDeleteHook, restaurantPhotoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	restaurantPhotoAfterDeleteHooks = []RestaurantPhotoHook{}

	AddRestaurantPhotoHook(boil.BeforeUpsertHook, restaurantPhotoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	restaurantPhotoBeforeUpsertHooks = []RestaurantPhotoHook{}

	AddRestaurantPhotoHook(boil.AfterUpsertHook, restaurantPhotoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	restaurantPhotoAfterUpsertHooks = []RestaurantPhotoHook{}
}

func testRestaurantPhotosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RestaurantPhotos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRestaurantPhotosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(restaurantPhotoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RestaurantPhotos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRestaurantPhotoToOneRestaurantUsingRestaurant(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RestaurantPhoto
	var foreign Restaurant

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, restaurantPhotoDBTypes, false, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, restaurantDBTypes, false, restaurantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Restaurant struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RestaurantID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Restaurant().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RestaurantPhotoSlice{&local}
	if err = local.L.LoadRestaurant(ctx, tx, false, (*[]*RestaurantPhoto)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Restaurant == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Restaurant = nil
	if err = local.L.LoadRestaurant(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Restaurant == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRestaurantPhotoToOneSetOpRestaurantUsingRestaurant(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RestaurantPhoto
	var b, c Restaurant

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, restaurantPhotoDBTypes, false, strmangle.SetComplement(restaurantPhotoPrimaryKeyColumns, restaurantPhotoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, restaurantDBTypes, false, strmangle.SetComplement(restaurantPrimaryKeyColumns, restaurantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, restaurantDBTypes, false, strmangle.SetComplement(restaurantPrimaryKeyColumns, restaurantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Restaurant{&b, &c} {
		err = a.SetRestaurant(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Restaurant != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RestaurantPhotos[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RestaurantID != x.ID {
			t.Error("foreign key was wrong value", a.RestaurantID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RestaurantID))
		reflect.Indirect(reflect.ValueOf(&a.RestaurantID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RestaurantID != x.ID {
			t.Error("foreign key was wrong value", a.RestaurantID, x.ID)
		}
	}
}

func testRestaurantPhotosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRestaurantPhotosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RestaurantPhotoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRestaurantPhotosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RestaurantPhotos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	restaurantPhotoDBTypes = map[string]string{`ID`: `integer`, `RestaurantID`: `integer`, `Photourl`: `text`, `Description`: `text`, `CreatedOn`: `text`}
	_                      = bytes.MinRead
)

func testRestaurantPhotosUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(restaurantPhotoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(restaurantPhotoAllColumns) == len(restaurantPhotoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RestaurantPhotos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRestaurantPhotosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(restaurantPhotoAllColumns) == len(restaurantPhotoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RestaurantPhoto{}
	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RestaurantPhotos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, restaurantPhotoDBTypes, true, restaurantPhotoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(restaurantPhotoAllColumns, restaurantPhotoPrimaryKeyColumns) {
		fields = restaurantPhotoAllColumns
	} else {
		fields = strmangle.SetComplement(
			restaurantPhotoAllColumns,
			restaurantPhotoPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := RestaurantPhotoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRestaurantPhotosUpsert(t *testing.T) {
	t.Parallel()

	if len(restaurantPhotoAllColumns) == len(restaurantPhotoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RestaurantPhoto{}
	if err = randomize.Struct(seed, &o, restaurantPhotoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RestaurantPhoto: %s", err)
	}

	count, err := RestaurantPhotos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, restaurantPhotoDBTypes, false, restaurantPhotoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RestaurantPhoto struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RestaurantPhoto: %s", err)
	}

	count, err = RestaurantPhotos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
