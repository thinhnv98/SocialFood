// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// RestaurantReview is an object representing the database table.
type RestaurantReview struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	RestaurantID int         `boil:"restaurant_id" json:"restaurant_id" toml:"restaurant_id" yaml:"restaurant_id"`
	AccountID    int         `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	Rating       null.Int    `boil:"rating" json:"rating,omitempty" toml:"rating" yaml:"rating,omitempty"`
	Text         null.String `boil:"text" json:"text,omitempty" toml:"text" yaml:"text,omitempty"`

	R *restaurantReviewR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L restaurantReviewL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RestaurantReviewColumns = struct {
	ID           string
	RestaurantID string
	AccountID    string
	Rating       string
	Text         string
}{
	ID:           "id",
	RestaurantID: "restaurant_id",
	AccountID:    "account_id",
	Rating:       "rating",
	Text:         "text",
}

var RestaurantReviewTableColumns = struct {
	ID           string
	RestaurantID string
	AccountID    string
	Rating       string
	Text         string
}{
	ID:           "restaurant_reviews.id",
	RestaurantID: "restaurant_reviews.restaurant_id",
	AccountID:    "restaurant_reviews.account_id",
	Rating:       "restaurant_reviews.rating",
	Text:         "restaurant_reviews.text",
}

// Generated where

var RestaurantReviewWhere = struct {
	ID           whereHelperint
	RestaurantID whereHelperint
	AccountID    whereHelperint
	Rating       whereHelpernull_Int
	Text         whereHelpernull_String
}{
	ID:           whereHelperint{field: "\"restaurant_reviews\".\"id\""},
	RestaurantID: whereHelperint{field: "\"restaurant_reviews\".\"restaurant_id\""},
	AccountID:    whereHelperint{field: "\"restaurant_reviews\".\"account_id\""},
	Rating:       whereHelpernull_Int{field: "\"restaurant_reviews\".\"rating\""},
	Text:         whereHelpernull_String{field: "\"restaurant_reviews\".\"text\""},
}

// RestaurantReviewRels is where relationship names are stored.
var RestaurantReviewRels = struct {
	Account    string
	Restaurant string
}{
	Account:    "Account",
	Restaurant: "Restaurant",
}

// restaurantReviewR is where relationships are stored.
type restaurantReviewR struct {
	Account    *Account    `boil:"Account" json:"Account" toml:"Account" yaml:"Account"`
	Restaurant *Restaurant `boil:"Restaurant" json:"Restaurant" toml:"Restaurant" yaml:"Restaurant"`
}

// NewStruct creates a new relationship struct
func (*restaurantReviewR) NewStruct() *restaurantReviewR {
	return &restaurantReviewR{}
}

// restaurantReviewL is where Load methods for each relationship are stored.
type restaurantReviewL struct{}

var (
	restaurantReviewAllColumns            = []string{"id", "restaurant_id", "account_id", "rating", "text"}
	restaurantReviewColumnsWithoutDefault = []string{"restaurant_id", "account_id", "rating", "text"}
	restaurantReviewColumnsWithDefault    = []string{"id"}
	restaurantReviewPrimaryKeyColumns     = []string{"id"}
)

type (
	// RestaurantReviewSlice is an alias for a slice of pointers to RestaurantReview.
	// This should almost always be used instead of []RestaurantReview.
	RestaurantReviewSlice []*RestaurantReview
	// RestaurantReviewHook is the signature for custom RestaurantReview hook methods
	RestaurantReviewHook func(context.Context, boil.ContextExecutor, *RestaurantReview) error

	restaurantReviewQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	restaurantReviewType                 = reflect.TypeOf(&RestaurantReview{})
	restaurantReviewMapping              = queries.MakeStructMapping(restaurantReviewType)
	restaurantReviewPrimaryKeyMapping, _ = queries.BindMapping(restaurantReviewType, restaurantReviewMapping, restaurantReviewPrimaryKeyColumns)
	restaurantReviewInsertCacheMut       sync.RWMutex
	restaurantReviewInsertCache          = make(map[string]insertCache)
	restaurantReviewUpdateCacheMut       sync.RWMutex
	restaurantReviewUpdateCache          = make(map[string]updateCache)
	restaurantReviewUpsertCacheMut       sync.RWMutex
	restaurantReviewUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var restaurantReviewBeforeInsertHooks []RestaurantReviewHook
var restaurantReviewBeforeUpdateHooks []RestaurantReviewHook
var restaurantReviewBeforeDeleteHooks []RestaurantReviewHook
var restaurantReviewBeforeUpsertHooks []RestaurantReviewHook

var restaurantReviewAfterInsertHooks []RestaurantReviewHook
var restaurantReviewAfterSelectHooks []RestaurantReviewHook
var restaurantReviewAfterUpdateHooks []RestaurantReviewHook
var restaurantReviewAfterDeleteHooks []RestaurantReviewHook
var restaurantReviewAfterUpsertHooks []RestaurantReviewHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RestaurantReview) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restaurantReviewBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RestaurantReview) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restaurantReviewBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RestaurantReview) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restaurantReviewBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RestaurantReview) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restaurantReviewBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RestaurantReview) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restaurantReviewAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RestaurantReview) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restaurantReviewAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RestaurantReview) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restaurantReviewAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RestaurantReview) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restaurantReviewAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RestaurantReview) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restaurantReviewAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRestaurantReviewHook registers your hook function for all future operations.
func AddRestaurantReviewHook(hookPoint boil.HookPoint, restaurantReviewHook RestaurantReviewHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		restaurantReviewBeforeInsertHooks = append(restaurantReviewBeforeInsertHooks, restaurantReviewHook)
	case boil.BeforeUpdateHook:
		restaurantReviewBeforeUpdateHooks = append(restaurantReviewBeforeUpdateHooks, restaurantReviewHook)
	case boil.BeforeDeleteHook:
		restaurantReviewBeforeDeleteHooks = append(restaurantReviewBeforeDeleteHooks, restaurantReviewHook)
	case boil.BeforeUpsertHook:
		restaurantReviewBeforeUpsertHooks = append(restaurantReviewBeforeUpsertHooks, restaurantReviewHook)
	case boil.AfterInsertHook:
		restaurantReviewAfterInsertHooks = append(restaurantReviewAfterInsertHooks, restaurantReviewHook)
	case boil.AfterSelectHook:
		restaurantReviewAfterSelectHooks = append(restaurantReviewAfterSelectHooks, restaurantReviewHook)
	case boil.AfterUpdateHook:
		restaurantReviewAfterUpdateHooks = append(restaurantReviewAfterUpdateHooks, restaurantReviewHook)
	case boil.AfterDeleteHook:
		restaurantReviewAfterDeleteHooks = append(restaurantReviewAfterDeleteHooks, restaurantReviewHook)
	case boil.AfterUpsertHook:
		restaurantReviewAfterUpsertHooks = append(restaurantReviewAfterUpsertHooks, restaurantReviewHook)
	}
}

// One returns a single restaurantReview record from the query.
func (q restaurantReviewQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RestaurantReview, error) {
	o := &RestaurantReview{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for restaurant_reviews")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RestaurantReview records from the query.
func (q restaurantReviewQuery) All(ctx context.Context, exec boil.ContextExecutor) (RestaurantReviewSlice, error) {
	var o []*RestaurantReview

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RestaurantReview slice")
	}

	if len(restaurantReviewAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RestaurantReview records in the query.
func (q restaurantReviewQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count restaurant_reviews rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q restaurantReviewQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if restaurant_reviews exists")
	}

	return count > 0, nil
}

// Account pointed to by the foreign key.
func (o *RestaurantReview) Account(mods ...qm.QueryMod) accountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.AccountID),
	}

	queryMods = append(queryMods, mods...)

	query := Accounts(queryMods...)
	queries.SetFrom(query.Query, "\"account\"")

	return query
}

// Restaurant pointed to by the foreign key.
func (o *RestaurantReview) Restaurant(mods ...qm.QueryMod) restaurantQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.RestaurantID),
	}

	queryMods = append(queryMods, mods...)

	query := Restaurants(queryMods...)
	queries.SetFrom(query.Query, "\"restaurant\"")

	return query
}

// LoadAccount allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (restaurantReviewL) LoadAccount(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRestaurantReview interface{}, mods queries.Applicator) error {
	var slice []*RestaurantReview
	var object *RestaurantReview

	if singular {
		object = maybeRestaurantReview.(*RestaurantReview)
	} else {
		slice = *maybeRestaurantReview.(*[]*RestaurantReview)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &restaurantReviewR{}
		}
		args = append(args, object.AccountID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &restaurantReviewR{}
			}

			for _, a := range args {
				if a == obj.AccountID {
					continue Outer
				}
			}

			args = append(args, obj.AccountID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`account`),
		qm.WhereIn(`account.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Account")
	}

	var resultSlice []*Account
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Account")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for account")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for account")
	}

	if len(restaurantReviewAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Account = foreign
		if foreign.R == nil {
			foreign.R = &accountR{}
		}
		foreign.R.RestaurantReviews = append(foreign.R.RestaurantReviews, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AccountID == foreign.ID {
				local.R.Account = foreign
				if foreign.R == nil {
					foreign.R = &accountR{}
				}
				foreign.R.RestaurantReviews = append(foreign.R.RestaurantReviews, local)
				break
			}
		}
	}

	return nil
}

// LoadRestaurant allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (restaurantReviewL) LoadRestaurant(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRestaurantReview interface{}, mods queries.Applicator) error {
	var slice []*RestaurantReview
	var object *RestaurantReview

	if singular {
		object = maybeRestaurantReview.(*RestaurantReview)
	} else {
		slice = *maybeRestaurantReview.(*[]*RestaurantReview)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &restaurantReviewR{}
		}
		args = append(args, object.RestaurantID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &restaurantReviewR{}
			}

			for _, a := range args {
				if a == obj.RestaurantID {
					continue Outer
				}
			}

			args = append(args, obj.RestaurantID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`restaurant`),
		qm.WhereIn(`restaurant.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Restaurant")
	}

	var resultSlice []*Restaurant
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Restaurant")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for restaurant")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for restaurant")
	}

	if len(restaurantReviewAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Restaurant = foreign
		if foreign.R == nil {
			foreign.R = &restaurantR{}
		}
		foreign.R.RestaurantReviews = append(foreign.R.RestaurantReviews, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RestaurantID == foreign.ID {
				local.R.Restaurant = foreign
				if foreign.R == nil {
					foreign.R = &restaurantR{}
				}
				foreign.R.RestaurantReviews = append(foreign.R.RestaurantReviews, local)
				break
			}
		}
	}

	return nil
}

// SetAccount of the restaurantReview to the related item.
// Sets o.R.Account to related.
// Adds o to related.R.RestaurantReviews.
func (o *RestaurantReview) SetAccount(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Account) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"restaurant_reviews\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"account_id"}),
		strmangle.WhereClause("\"", "\"", 2, restaurantReviewPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AccountID = related.ID
	if o.R == nil {
		o.R = &restaurantReviewR{
			Account: related,
		}
	} else {
		o.R.Account = related
	}

	if related.R == nil {
		related.R = &accountR{
			RestaurantReviews: RestaurantReviewSlice{o},
		}
	} else {
		related.R.RestaurantReviews = append(related.R.RestaurantReviews, o)
	}

	return nil
}

// SetRestaurant of the restaurantReview to the related item.
// Sets o.R.Restaurant to related.
// Adds o to related.R.RestaurantReviews.
func (o *RestaurantReview) SetRestaurant(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Restaurant) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"restaurant_reviews\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"restaurant_id"}),
		strmangle.WhereClause("\"", "\"", 2, restaurantReviewPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RestaurantID = related.ID
	if o.R == nil {
		o.R = &restaurantReviewR{
			Restaurant: related,
		}
	} else {
		o.R.Restaurant = related
	}

	if related.R == nil {
		related.R = &restaurantR{
			RestaurantReviews: RestaurantReviewSlice{o},
		}
	} else {
		related.R.RestaurantReviews = append(related.R.RestaurantReviews, o)
	}

	return nil
}

// RestaurantReviews retrieves all the records using an executor.
func RestaurantReviews(mods ...qm.QueryMod) restaurantReviewQuery {
	mods = append(mods, qm.From("\"restaurant_reviews\""))
	return restaurantReviewQuery{NewQuery(mods...)}
}

// FindRestaurantReview retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRestaurantReview(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RestaurantReview, error) {
	restaurantReviewObj := &RestaurantReview{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"restaurant_reviews\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, restaurantReviewObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from restaurant_reviews")
	}

	if err = restaurantReviewObj.doAfterSelectHooks(ctx, exec); err != nil {
		return restaurantReviewObj, err
	}

	return restaurantReviewObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RestaurantReview) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no restaurant_reviews provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(restaurantReviewColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	restaurantReviewInsertCacheMut.RLock()
	cache, cached := restaurantReviewInsertCache[key]
	restaurantReviewInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			restaurantReviewAllColumns,
			restaurantReviewColumnsWithDefault,
			restaurantReviewColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(restaurantReviewType, restaurantReviewMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(restaurantReviewType, restaurantReviewMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"restaurant_reviews\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"restaurant_reviews\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into restaurant_reviews")
	}

	if !cached {
		restaurantReviewInsertCacheMut.Lock()
		restaurantReviewInsertCache[key] = cache
		restaurantReviewInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RestaurantReview.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RestaurantReview) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	restaurantReviewUpdateCacheMut.RLock()
	cache, cached := restaurantReviewUpdateCache[key]
	restaurantReviewUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			restaurantReviewAllColumns,
			restaurantReviewPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update restaurant_reviews, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"restaurant_reviews\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, restaurantReviewPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(restaurantReviewType, restaurantReviewMapping, append(wl, restaurantReviewPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update restaurant_reviews row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for restaurant_reviews")
	}

	if !cached {
		restaurantReviewUpdateCacheMut.Lock()
		restaurantReviewUpdateCache[key] = cache
		restaurantReviewUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q restaurantReviewQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for restaurant_reviews")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for restaurant_reviews")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RestaurantReviewSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), restaurantReviewPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"restaurant_reviews\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, restaurantReviewPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in restaurantReview slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all restaurantReview")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RestaurantReview) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no restaurant_reviews provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(restaurantReviewColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	restaurantReviewUpsertCacheMut.RLock()
	cache, cached := restaurantReviewUpsertCache[key]
	restaurantReviewUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			restaurantReviewAllColumns,
			restaurantReviewColumnsWithDefault,
			restaurantReviewColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			restaurantReviewAllColumns,
			restaurantReviewPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert restaurant_reviews, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(restaurantReviewPrimaryKeyColumns))
			copy(conflict, restaurantReviewPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"restaurant_reviews\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(restaurantReviewType, restaurantReviewMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(restaurantReviewType, restaurantReviewMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert restaurant_reviews")
	}

	if !cached {
		restaurantReviewUpsertCacheMut.Lock()
		restaurantReviewUpsertCache[key] = cache
		restaurantReviewUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RestaurantReview record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RestaurantReview) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RestaurantReview provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), restaurantReviewPrimaryKeyMapping)
	sql := "DELETE FROM \"restaurant_reviews\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from restaurant_reviews")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for restaurant_reviews")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q restaurantReviewQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no restaurantReviewQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from restaurant_reviews")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for restaurant_reviews")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RestaurantReviewSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(restaurantReviewBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), restaurantReviewPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"restaurant_reviews\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, restaurantReviewPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from restaurantReview slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for restaurant_reviews")
	}

	if len(restaurantReviewAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RestaurantReview) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRestaurantReview(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RestaurantReviewSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RestaurantReviewSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), restaurantReviewPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"restaurant_reviews\".* FROM \"restaurant_reviews\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, restaurantReviewPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RestaurantReviewSlice")
	}

	*o = slice

	return nil
}

// RestaurantReviewExists checks if the RestaurantReview row exists.
func RestaurantReviewExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"restaurant_reviews\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if restaurant_reviews exists")
	}

	return exists, nil
}
