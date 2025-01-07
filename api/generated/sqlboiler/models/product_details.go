// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// ProductDetail is an object representing the database table.
type ProductDetail struct {
	ID              int64         `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProdID          int64         `boil:"prod_id" json:"prod_id" toml:"prod_id" yaml:"prod_id"`
	ProdDescription string        `boil:"prod_description" json:"prod_description" toml:"prod_description" yaml:"prod_description"`
	ProdWeight      types.Decimal `boil:"prod_weight" json:"prod_weight" toml:"prod_weight" yaml:"prod_weight"`
	ProdWidth       types.Decimal `boil:"prod_width" json:"prod_width" toml:"prod_width" yaml:"prod_width"`
	ProdLength      types.Decimal `boil:"prod_length" json:"prod_length" toml:"prod_length" yaml:"prod_length"`
	ProdHeight      types.Decimal `boil:"prod_height" json:"prod_height" toml:"prod_height" yaml:"prod_height"`
	ProductMaterial string        `boil:"product_material" json:"product_material" toml:"product_material" yaml:"product_material"`

	R *productDetailR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L productDetailL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProductDetailColumns = struct {
	ID              string
	ProdID          string
	ProdDescription string
	ProdWeight      string
	ProdWidth       string
	ProdLength      string
	ProdHeight      string
	ProductMaterial string
}{
	ID:              "id",
	ProdID:          "prod_id",
	ProdDescription: "prod_description",
	ProdWeight:      "prod_weight",
	ProdWidth:       "prod_width",
	ProdLength:      "prod_length",
	ProdHeight:      "prod_height",
	ProductMaterial: "product_material",
}

var ProductDetailTableColumns = struct {
	ID              string
	ProdID          string
	ProdDescription string
	ProdWeight      string
	ProdWidth       string
	ProdLength      string
	ProdHeight      string
	ProductMaterial string
}{
	ID:              "product_details.id",
	ProdID:          "product_details.prod_id",
	ProdDescription: "product_details.prod_description",
	ProdWeight:      "product_details.prod_weight",
	ProdWidth:       "product_details.prod_width",
	ProdLength:      "product_details.prod_length",
	ProdHeight:      "product_details.prod_height",
	ProductMaterial: "product_details.product_material",
}

// Generated where

var ProductDetailWhere = struct {
	ID              whereHelperint64
	ProdID          whereHelperint64
	ProdDescription whereHelperstring
	ProdWeight      whereHelpertypes_Decimal
	ProdWidth       whereHelpertypes_Decimal
	ProdLength      whereHelpertypes_Decimal
	ProdHeight      whereHelpertypes_Decimal
	ProductMaterial whereHelperstring
}{
	ID:              whereHelperint64{field: "\"product_details\".\"id\""},
	ProdID:          whereHelperint64{field: "\"product_details\".\"prod_id\""},
	ProdDescription: whereHelperstring{field: "\"product_details\".\"prod_description\""},
	ProdWeight:      whereHelpertypes_Decimal{field: "\"product_details\".\"prod_weight\""},
	ProdWidth:       whereHelpertypes_Decimal{field: "\"product_details\".\"prod_width\""},
	ProdLength:      whereHelpertypes_Decimal{field: "\"product_details\".\"prod_length\""},
	ProdHeight:      whereHelpertypes_Decimal{field: "\"product_details\".\"prod_height\""},
	ProductMaterial: whereHelperstring{field: "\"product_details\".\"product_material\""},
}

// ProductDetailRels is where relationship names are stored.
var ProductDetailRels = struct {
	Prod string
}{
	Prod: "Prod",
}

// productDetailR is where relationships are stored.
type productDetailR struct {
	Prod *Product `boil:"Prod" json:"Prod" toml:"Prod" yaml:"Prod"`
}

// NewStruct creates a new relationship struct
func (*productDetailR) NewStruct() *productDetailR {
	return &productDetailR{}
}

func (r *productDetailR) GetProd() *Product {
	if r == nil {
		return nil
	}
	return r.Prod
}

// productDetailL is where Load methods for each relationship are stored.
type productDetailL struct{}

var (
	productDetailAllColumns            = []string{"id", "prod_id", "prod_description", "prod_weight", "prod_width", "prod_length", "prod_height", "product_material"}
	productDetailColumnsWithoutDefault = []string{"prod_id", "prod_description", "prod_weight", "prod_width", "prod_length", "prod_height", "product_material"}
	productDetailColumnsWithDefault    = []string{"id"}
	productDetailPrimaryKeyColumns     = []string{"id"}
	productDetailGeneratedColumns      = []string{}
)

type (
	// ProductDetailSlice is an alias for a slice of pointers to ProductDetail.
	// This should almost always be used instead of []ProductDetail.
	ProductDetailSlice []*ProductDetail
	// ProductDetailHook is the signature for custom ProductDetail hook methods
	ProductDetailHook func(context.Context, boil.ContextExecutor, *ProductDetail) error

	productDetailQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	productDetailType                 = reflect.TypeOf(&ProductDetail{})
	productDetailMapping              = queries.MakeStructMapping(productDetailType)
	productDetailPrimaryKeyMapping, _ = queries.BindMapping(productDetailType, productDetailMapping, productDetailPrimaryKeyColumns)
	productDetailInsertCacheMut       sync.RWMutex
	productDetailInsertCache          = make(map[string]insertCache)
	productDetailUpdateCacheMut       sync.RWMutex
	productDetailUpdateCache          = make(map[string]updateCache)
	productDetailUpsertCacheMut       sync.RWMutex
	productDetailUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var productDetailAfterSelectMu sync.Mutex
var productDetailAfterSelectHooks []ProductDetailHook

var productDetailBeforeInsertMu sync.Mutex
var productDetailBeforeInsertHooks []ProductDetailHook
var productDetailAfterInsertMu sync.Mutex
var productDetailAfterInsertHooks []ProductDetailHook

var productDetailBeforeUpdateMu sync.Mutex
var productDetailBeforeUpdateHooks []ProductDetailHook
var productDetailAfterUpdateMu sync.Mutex
var productDetailAfterUpdateHooks []ProductDetailHook

var productDetailBeforeDeleteMu sync.Mutex
var productDetailBeforeDeleteHooks []ProductDetailHook
var productDetailAfterDeleteMu sync.Mutex
var productDetailAfterDeleteHooks []ProductDetailHook

var productDetailBeforeUpsertMu sync.Mutex
var productDetailBeforeUpsertHooks []ProductDetailHook
var productDetailAfterUpsertMu sync.Mutex
var productDetailAfterUpsertHooks []ProductDetailHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProductDetail) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productDetailAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProductDetail) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productDetailBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProductDetail) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productDetailAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProductDetail) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productDetailBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProductDetail) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productDetailAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProductDetail) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productDetailBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProductDetail) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productDetailAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProductDetail) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productDetailBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProductDetail) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productDetailAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProductDetailHook registers your hook function for all future operations.
func AddProductDetailHook(hookPoint boil.HookPoint, productDetailHook ProductDetailHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		productDetailAfterSelectMu.Lock()
		productDetailAfterSelectHooks = append(productDetailAfterSelectHooks, productDetailHook)
		productDetailAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		productDetailBeforeInsertMu.Lock()
		productDetailBeforeInsertHooks = append(productDetailBeforeInsertHooks, productDetailHook)
		productDetailBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		productDetailAfterInsertMu.Lock()
		productDetailAfterInsertHooks = append(productDetailAfterInsertHooks, productDetailHook)
		productDetailAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		productDetailBeforeUpdateMu.Lock()
		productDetailBeforeUpdateHooks = append(productDetailBeforeUpdateHooks, productDetailHook)
		productDetailBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		productDetailAfterUpdateMu.Lock()
		productDetailAfterUpdateHooks = append(productDetailAfterUpdateHooks, productDetailHook)
		productDetailAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		productDetailBeforeDeleteMu.Lock()
		productDetailBeforeDeleteHooks = append(productDetailBeforeDeleteHooks, productDetailHook)
		productDetailBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		productDetailAfterDeleteMu.Lock()
		productDetailAfterDeleteHooks = append(productDetailAfterDeleteHooks, productDetailHook)
		productDetailAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		productDetailBeforeUpsertMu.Lock()
		productDetailBeforeUpsertHooks = append(productDetailBeforeUpsertHooks, productDetailHook)
		productDetailBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		productDetailAfterUpsertMu.Lock()
		productDetailAfterUpsertHooks = append(productDetailAfterUpsertHooks, productDetailHook)
		productDetailAfterUpsertMu.Unlock()
	}
}

// One returns a single productDetail record from the query.
func (q productDetailQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProductDetail, error) {
	o := &ProductDetail{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for product_details")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ProductDetail records from the query.
func (q productDetailQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProductDetailSlice, error) {
	var o []*ProductDetail

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ProductDetail slice")
	}

	if len(productDetailAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ProductDetail records in the query.
func (q productDetailQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count product_details rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q productDetailQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if product_details exists")
	}

	return count > 0, nil
}

// Prod pointed to by the foreign key.
func (o *ProductDetail) Prod(mods ...qm.QueryMod) productQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProdID),
	}

	queryMods = append(queryMods, mods...)

	return Products(queryMods...)
}

// LoadProd allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (productDetailL) LoadProd(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProductDetail interface{}, mods queries.Applicator) error {
	var slice []*ProductDetail
	var object *ProductDetail

	if singular {
		var ok bool
		object, ok = maybeProductDetail.(*ProductDetail)
		if !ok {
			object = new(ProductDetail)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeProductDetail)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeProductDetail))
			}
		}
	} else {
		s, ok := maybeProductDetail.(*[]*ProductDetail)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeProductDetail)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeProductDetail))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &productDetailR{}
		}
		args[object.ProdID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &productDetailR{}
			}

			args[obj.ProdID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`products`),
		qm.WhereIn(`products.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Product")
	}

	var resultSlice []*Product
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Product")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for products")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for products")
	}

	if len(productAfterSelectHooks) != 0 {
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
		object.R.Prod = foreign
		if foreign.R == nil {
			foreign.R = &productR{}
		}
		foreign.R.ProdProductDetails = append(foreign.R.ProdProductDetails, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProdID == foreign.ID {
				local.R.Prod = foreign
				if foreign.R == nil {
					foreign.R = &productR{}
				}
				foreign.R.ProdProductDetails = append(foreign.R.ProdProductDetails, local)
				break
			}
		}
	}

	return nil
}

// SetProd of the productDetail to the related item.
// Sets o.R.Prod to related.
// Adds o to related.R.ProdProductDetails.
func (o *ProductDetail) SetProd(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Product) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"product_details\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"prod_id"}),
		strmangle.WhereClause("\"", "\"", 2, productDetailPrimaryKeyColumns),
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

	o.ProdID = related.ID
	if o.R == nil {
		o.R = &productDetailR{
			Prod: related,
		}
	} else {
		o.R.Prod = related
	}

	if related.R == nil {
		related.R = &productR{
			ProdProductDetails: ProductDetailSlice{o},
		}
	} else {
		related.R.ProdProductDetails = append(related.R.ProdProductDetails, o)
	}

	return nil
}

// ProductDetails retrieves all the records using an executor.
func ProductDetails(mods ...qm.QueryMod) productDetailQuery {
	mods = append(mods, qm.From("\"product_details\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"product_details\".*"})
	}

	return productDetailQuery{q}
}

// FindProductDetail retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProductDetail(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*ProductDetail, error) {
	productDetailObj := &ProductDetail{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"product_details\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, productDetailObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from product_details")
	}

	if err = productDetailObj.doAfterSelectHooks(ctx, exec); err != nil {
		return productDetailObj, err
	}

	return productDetailObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProductDetail) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no product_details provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productDetailColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	productDetailInsertCacheMut.RLock()
	cache, cached := productDetailInsertCache[key]
	productDetailInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			productDetailAllColumns,
			productDetailColumnsWithDefault,
			productDetailColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(productDetailType, productDetailMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(productDetailType, productDetailMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"product_details\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"product_details\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into product_details")
	}

	if !cached {
		productDetailInsertCacheMut.Lock()
		productDetailInsertCache[key] = cache
		productDetailInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ProductDetail.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProductDetail) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	productDetailUpdateCacheMut.RLock()
	cache, cached := productDetailUpdateCache[key]
	productDetailUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			productDetailAllColumns,
			productDetailPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update product_details, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"product_details\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, productDetailPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(productDetailType, productDetailMapping, append(wl, productDetailPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update product_details row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for product_details")
	}

	if !cached {
		productDetailUpdateCacheMut.Lock()
		productDetailUpdateCache[key] = cache
		productDetailUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q productDetailQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for product_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for product_details")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProductDetailSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"product_details\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, productDetailPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in productDetail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all productDetail")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ProductDetail) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no product_details provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productDetailColumnsWithDefault, o)

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

	productDetailUpsertCacheMut.RLock()
	cache, cached := productDetailUpsertCache[key]
	productDetailUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			productDetailAllColumns,
			productDetailColumnsWithDefault,
			productDetailColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			productDetailAllColumns,
			productDetailPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert product_details, could not build update column list")
		}

		ret := strmangle.SetComplement(productDetailAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(productDetailPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert product_details, could not build conflict column list")
			}

			conflict = make([]string, len(productDetailPrimaryKeyColumns))
			copy(conflict, productDetailPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"product_details\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(productDetailType, productDetailMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(productDetailType, productDetailMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert product_details")
	}

	if !cached {
		productDetailUpsertCacheMut.Lock()
		productDetailUpsertCache[key] = cache
		productDetailUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ProductDetail record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProductDetail) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ProductDetail provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), productDetailPrimaryKeyMapping)
	sql := "DELETE FROM \"product_details\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from product_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for product_details")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q productDetailQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no productDetailQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from product_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for product_details")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProductDetailSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(productDetailBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"product_details\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productDetailPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from productDetail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for product_details")
	}

	if len(productDetailAfterDeleteHooks) != 0 {
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
func (o *ProductDetail) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProductDetail(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductDetailSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProductDetailSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"product_details\".* FROM \"product_details\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productDetailPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProductDetailSlice")
	}

	*o = slice

	return nil
}

// ProductDetailExists checks if the ProductDetail row exists.
func ProductDetailExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"product_details\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if product_details exists")
	}

	return exists, nil
}

// Exists checks if the ProductDetail row exists.
func (o *ProductDetail) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProductDetailExists(ctx, exec, o.ID)
}