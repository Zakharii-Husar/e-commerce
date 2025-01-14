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
	"github.com/volatiletech/strmangle"
)

// Attachment is an object representing the database table.
type Attachment struct {
	ID             int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	AttachmentName string `boil:"attachment_name" json:"attachment_name" toml:"attachment_name" yaml:"attachment_name"`
	ProdID         int64  `boil:"prod_id" json:"prod_id" toml:"prod_id" yaml:"prod_id"`

	R *attachmentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L attachmentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AttachmentColumns = struct {
	ID             string
	AttachmentName string
	ProdID         string
}{
	ID:             "id",
	AttachmentName: "attachment_name",
	ProdID:         "prod_id",
}

var AttachmentTableColumns = struct {
	ID             string
	AttachmentName string
	ProdID         string
}{
	ID:             "attachments.id",
	AttachmentName: "attachments.attachment_name",
	ProdID:         "attachments.prod_id",
}

// Generated where

var AttachmentWhere = struct {
	ID             whereHelperint64
	AttachmentName whereHelperstring
	ProdID         whereHelperint64
}{
	ID:             whereHelperint64{field: "\"attachments\".\"id\""},
	AttachmentName: whereHelperstring{field: "\"attachments\".\"attachment_name\""},
	ProdID:         whereHelperint64{field: "\"attachments\".\"prod_id\""},
}

// AttachmentRels is where relationship names are stored.
var AttachmentRels = struct {
	Prod string
}{
	Prod: "Prod",
}

// attachmentR is where relationships are stored.
type attachmentR struct {
	Prod *Product `boil:"Prod" json:"Prod" toml:"Prod" yaml:"Prod"`
}

// NewStruct creates a new relationship struct
func (*attachmentR) NewStruct() *attachmentR {
	return &attachmentR{}
}

func (r *attachmentR) GetProd() *Product {
	if r == nil {
		return nil
	}
	return r.Prod
}

// attachmentL is where Load methods for each relationship are stored.
type attachmentL struct{}

var (
	attachmentAllColumns            = []string{"id", "attachment_name", "prod_id"}
	attachmentColumnsWithoutDefault = []string{"attachment_name", "prod_id"}
	attachmentColumnsWithDefault    = []string{"id"}
	attachmentPrimaryKeyColumns     = []string{"id"}
	attachmentGeneratedColumns      = []string{}
)

type (
	// AttachmentSlice is an alias for a slice of pointers to Attachment.
	// This should almost always be used instead of []Attachment.
	AttachmentSlice []*Attachment
	// AttachmentHook is the signature for custom Attachment hook methods
	AttachmentHook func(context.Context, boil.ContextExecutor, *Attachment) error

	attachmentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	attachmentType                 = reflect.TypeOf(&Attachment{})
	attachmentMapping              = queries.MakeStructMapping(attachmentType)
	attachmentPrimaryKeyMapping, _ = queries.BindMapping(attachmentType, attachmentMapping, attachmentPrimaryKeyColumns)
	attachmentInsertCacheMut       sync.RWMutex
	attachmentInsertCache          = make(map[string]insertCache)
	attachmentUpdateCacheMut       sync.RWMutex
	attachmentUpdateCache          = make(map[string]updateCache)
	attachmentUpsertCacheMut       sync.RWMutex
	attachmentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var attachmentAfterSelectMu sync.Mutex
var attachmentAfterSelectHooks []AttachmentHook

var attachmentBeforeInsertMu sync.Mutex
var attachmentBeforeInsertHooks []AttachmentHook
var attachmentAfterInsertMu sync.Mutex
var attachmentAfterInsertHooks []AttachmentHook

var attachmentBeforeUpdateMu sync.Mutex
var attachmentBeforeUpdateHooks []AttachmentHook
var attachmentAfterUpdateMu sync.Mutex
var attachmentAfterUpdateHooks []AttachmentHook

var attachmentBeforeDeleteMu sync.Mutex
var attachmentBeforeDeleteHooks []AttachmentHook
var attachmentAfterDeleteMu sync.Mutex
var attachmentAfterDeleteHooks []AttachmentHook

var attachmentBeforeUpsertMu sync.Mutex
var attachmentBeforeUpsertHooks []AttachmentHook
var attachmentAfterUpsertMu sync.Mutex
var attachmentAfterUpsertHooks []AttachmentHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Attachment) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range attachmentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Attachment) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range attachmentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Attachment) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range attachmentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Attachment) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range attachmentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Attachment) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range attachmentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Attachment) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range attachmentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Attachment) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range attachmentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Attachment) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range attachmentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Attachment) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range attachmentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAttachmentHook registers your hook function for all future operations.
func AddAttachmentHook(hookPoint boil.HookPoint, attachmentHook AttachmentHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		attachmentAfterSelectMu.Lock()
		attachmentAfterSelectHooks = append(attachmentAfterSelectHooks, attachmentHook)
		attachmentAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		attachmentBeforeInsertMu.Lock()
		attachmentBeforeInsertHooks = append(attachmentBeforeInsertHooks, attachmentHook)
		attachmentBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		attachmentAfterInsertMu.Lock()
		attachmentAfterInsertHooks = append(attachmentAfterInsertHooks, attachmentHook)
		attachmentAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		attachmentBeforeUpdateMu.Lock()
		attachmentBeforeUpdateHooks = append(attachmentBeforeUpdateHooks, attachmentHook)
		attachmentBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		attachmentAfterUpdateMu.Lock()
		attachmentAfterUpdateHooks = append(attachmentAfterUpdateHooks, attachmentHook)
		attachmentAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		attachmentBeforeDeleteMu.Lock()
		attachmentBeforeDeleteHooks = append(attachmentBeforeDeleteHooks, attachmentHook)
		attachmentBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		attachmentAfterDeleteMu.Lock()
		attachmentAfterDeleteHooks = append(attachmentAfterDeleteHooks, attachmentHook)
		attachmentAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		attachmentBeforeUpsertMu.Lock()
		attachmentBeforeUpsertHooks = append(attachmentBeforeUpsertHooks, attachmentHook)
		attachmentBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		attachmentAfterUpsertMu.Lock()
		attachmentAfterUpsertHooks = append(attachmentAfterUpsertHooks, attachmentHook)
		attachmentAfterUpsertMu.Unlock()
	}
}

// One returns a single attachment record from the query.
func (q attachmentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Attachment, error) {
	o := &Attachment{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for attachments")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Attachment records from the query.
func (q attachmentQuery) All(ctx context.Context, exec boil.ContextExecutor) (AttachmentSlice, error) {
	var o []*Attachment

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Attachment slice")
	}

	if len(attachmentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Attachment records in the query.
func (q attachmentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count attachments rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q attachmentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if attachments exists")
	}

	return count > 0, nil
}

// Prod pointed to by the foreign key.
func (o *Attachment) Prod(mods ...qm.QueryMod) productQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProdID),
	}

	queryMods = append(queryMods, mods...)

	return Products(queryMods...)
}

// LoadProd allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (attachmentL) LoadProd(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAttachment interface{}, mods queries.Applicator) error {
	var slice []*Attachment
	var object *Attachment

	if singular {
		var ok bool
		object, ok = maybeAttachment.(*Attachment)
		if !ok {
			object = new(Attachment)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAttachment)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAttachment))
			}
		}
	} else {
		s, ok := maybeAttachment.(*[]*Attachment)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAttachment)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAttachment))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &attachmentR{}
		}
		args[object.ProdID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &attachmentR{}
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
		foreign.R.ProdAttachments = append(foreign.R.ProdAttachments, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProdID == foreign.ID {
				local.R.Prod = foreign
				if foreign.R == nil {
					foreign.R = &productR{}
				}
				foreign.R.ProdAttachments = append(foreign.R.ProdAttachments, local)
				break
			}
		}
	}

	return nil
}

// SetProd of the attachment to the related item.
// Sets o.R.Prod to related.
// Adds o to related.R.ProdAttachments.
func (o *Attachment) SetProd(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Product) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"attachments\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"prod_id"}),
		strmangle.WhereClause("\"", "\"", 2, attachmentPrimaryKeyColumns),
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
		o.R = &attachmentR{
			Prod: related,
		}
	} else {
		o.R.Prod = related
	}

	if related.R == nil {
		related.R = &productR{
			ProdAttachments: AttachmentSlice{o},
		}
	} else {
		related.R.ProdAttachments = append(related.R.ProdAttachments, o)
	}

	return nil
}

// Attachments retrieves all the records using an executor.
func Attachments(mods ...qm.QueryMod) attachmentQuery {
	mods = append(mods, qm.From("\"attachments\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"attachments\".*"})
	}

	return attachmentQuery{q}
}

// FindAttachment retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAttachment(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Attachment, error) {
	attachmentObj := &Attachment{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"attachments\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, attachmentObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from attachments")
	}

	if err = attachmentObj.doAfterSelectHooks(ctx, exec); err != nil {
		return attachmentObj, err
	}

	return attachmentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Attachment) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no attachments provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(attachmentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	attachmentInsertCacheMut.RLock()
	cache, cached := attachmentInsertCache[key]
	attachmentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			attachmentAllColumns,
			attachmentColumnsWithDefault,
			attachmentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(attachmentType, attachmentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(attachmentType, attachmentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"attachments\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"attachments\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into attachments")
	}

	if !cached {
		attachmentInsertCacheMut.Lock()
		attachmentInsertCache[key] = cache
		attachmentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Attachment.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Attachment) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	attachmentUpdateCacheMut.RLock()
	cache, cached := attachmentUpdateCache[key]
	attachmentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			attachmentAllColumns,
			attachmentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update attachments, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"attachments\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, attachmentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(attachmentType, attachmentMapping, append(wl, attachmentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update attachments row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for attachments")
	}

	if !cached {
		attachmentUpdateCacheMut.Lock()
		attachmentUpdateCache[key] = cache
		attachmentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q attachmentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for attachments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for attachments")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AttachmentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), attachmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"attachments\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, attachmentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in attachment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all attachment")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Attachment) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no attachments provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(attachmentColumnsWithDefault, o)

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

	attachmentUpsertCacheMut.RLock()
	cache, cached := attachmentUpsertCache[key]
	attachmentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			attachmentAllColumns,
			attachmentColumnsWithDefault,
			attachmentColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			attachmentAllColumns,
			attachmentPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert attachments, could not build update column list")
		}

		ret := strmangle.SetComplement(attachmentAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(attachmentPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert attachments, could not build conflict column list")
			}

			conflict = make([]string, len(attachmentPrimaryKeyColumns))
			copy(conflict, attachmentPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"attachments\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(attachmentType, attachmentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(attachmentType, attachmentMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert attachments")
	}

	if !cached {
		attachmentUpsertCacheMut.Lock()
		attachmentUpsertCache[key] = cache
		attachmentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Attachment record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Attachment) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Attachment provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), attachmentPrimaryKeyMapping)
	sql := "DELETE FROM \"attachments\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from attachments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for attachments")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q attachmentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no attachmentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from attachments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for attachments")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AttachmentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(attachmentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), attachmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"attachments\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, attachmentPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from attachment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for attachments")
	}

	if len(attachmentAfterDeleteHooks) != 0 {
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
func (o *Attachment) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAttachment(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AttachmentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AttachmentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), attachmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"attachments\".* FROM \"attachments\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, attachmentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AttachmentSlice")
	}

	*o = slice

	return nil
}

// AttachmentExists checks if the Attachment row exists.
func AttachmentExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"attachments\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if attachments exists")
	}

	return exists, nil
}

// Exists checks if the Attachment row exists.
func (o *Attachment) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AttachmentExists(ctx, exec, o.ID)
}
