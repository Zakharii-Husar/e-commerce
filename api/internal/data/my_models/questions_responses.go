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

// QuestionsResponse is an object representing the database table.
type QuestionsResponse struct {
	ID           int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	QuestionID   int64  `boil:"question_id" json:"question_id" toml:"question_id" yaml:"question_id"`
	UserID       int64  `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	ResponseText string `boil:"response_text" json:"response_text" toml:"response_text" yaml:"response_text"`

	R *questionsResponseR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L questionsResponseL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var QuestionsResponseColumns = struct {
	ID           string
	QuestionID   string
	UserID       string
	ResponseText string
}{
	ID:           "id",
	QuestionID:   "question_id",
	UserID:       "user_id",
	ResponseText: "response_text",
}

var QuestionsResponseTableColumns = struct {
	ID           string
	QuestionID   string
	UserID       string
	ResponseText string
}{
	ID:           "questions_responses.id",
	QuestionID:   "questions_responses.question_id",
	UserID:       "questions_responses.user_id",
	ResponseText: "questions_responses.response_text",
}

// Generated where

var QuestionsResponseWhere = struct {
	ID           whereHelperint64
	QuestionID   whereHelperint64
	UserID       whereHelperint64
	ResponseText whereHelperstring
}{
	ID:           whereHelperint64{field: "\"questions_responses\".\"id\""},
	QuestionID:   whereHelperint64{field: "\"questions_responses\".\"question_id\""},
	UserID:       whereHelperint64{field: "\"questions_responses\".\"user_id\""},
	ResponseText: whereHelperstring{field: "\"questions_responses\".\"response_text\""},
}

// QuestionsResponseRels is where relationship names are stored.
var QuestionsResponseRels = struct {
	Question string
	User     string
}{
	Question: "Question",
	User:     "User",
}

// questionsResponseR is where relationships are stored.
type questionsResponseR struct {
	Question *Question `boil:"Question" json:"Question" toml:"Question" yaml:"Question"`
	User     *User     `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*questionsResponseR) NewStruct() *questionsResponseR {
	return &questionsResponseR{}
}

func (r *questionsResponseR) GetQuestion() *Question {
	if r == nil {
		return nil
	}
	return r.Question
}

func (r *questionsResponseR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// questionsResponseL is where Load methods for each relationship are stored.
type questionsResponseL struct{}

var (
	questionsResponseAllColumns            = []string{"id", "question_id", "user_id", "response_text"}
	questionsResponseColumnsWithoutDefault = []string{"question_id", "user_id", "response_text"}
	questionsResponseColumnsWithDefault    = []string{"id"}
	questionsResponsePrimaryKeyColumns     = []string{"id"}
	questionsResponseGeneratedColumns      = []string{}
)

type (
	// QuestionsResponseSlice is an alias for a slice of pointers to QuestionsResponse.
	// This should almost always be used instead of []QuestionsResponse.
	QuestionsResponseSlice []*QuestionsResponse
	// QuestionsResponseHook is the signature for custom QuestionsResponse hook methods
	QuestionsResponseHook func(context.Context, boil.ContextExecutor, *QuestionsResponse) error

	questionsResponseQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	questionsResponseType                 = reflect.TypeOf(&QuestionsResponse{})
	questionsResponseMapping              = queries.MakeStructMapping(questionsResponseType)
	questionsResponsePrimaryKeyMapping, _ = queries.BindMapping(questionsResponseType, questionsResponseMapping, questionsResponsePrimaryKeyColumns)
	questionsResponseInsertCacheMut       sync.RWMutex
	questionsResponseInsertCache          = make(map[string]insertCache)
	questionsResponseUpdateCacheMut       sync.RWMutex
	questionsResponseUpdateCache          = make(map[string]updateCache)
	questionsResponseUpsertCacheMut       sync.RWMutex
	questionsResponseUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var questionsResponseAfterSelectMu sync.Mutex
var questionsResponseAfterSelectHooks []QuestionsResponseHook

var questionsResponseBeforeInsertMu sync.Mutex
var questionsResponseBeforeInsertHooks []QuestionsResponseHook
var questionsResponseAfterInsertMu sync.Mutex
var questionsResponseAfterInsertHooks []QuestionsResponseHook

var questionsResponseBeforeUpdateMu sync.Mutex
var questionsResponseBeforeUpdateHooks []QuestionsResponseHook
var questionsResponseAfterUpdateMu sync.Mutex
var questionsResponseAfterUpdateHooks []QuestionsResponseHook

var questionsResponseBeforeDeleteMu sync.Mutex
var questionsResponseBeforeDeleteHooks []QuestionsResponseHook
var questionsResponseAfterDeleteMu sync.Mutex
var questionsResponseAfterDeleteHooks []QuestionsResponseHook

var questionsResponseBeforeUpsertMu sync.Mutex
var questionsResponseBeforeUpsertHooks []QuestionsResponseHook
var questionsResponseAfterUpsertMu sync.Mutex
var questionsResponseAfterUpsertHooks []QuestionsResponseHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *QuestionsResponse) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsResponseAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *QuestionsResponse) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsResponseBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *QuestionsResponse) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsResponseAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *QuestionsResponse) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsResponseBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *QuestionsResponse) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsResponseAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *QuestionsResponse) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsResponseBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *QuestionsResponse) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsResponseAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *QuestionsResponse) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsResponseBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *QuestionsResponse) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsResponseAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddQuestionsResponseHook registers your hook function for all future operations.
func AddQuestionsResponseHook(hookPoint boil.HookPoint, questionsResponseHook QuestionsResponseHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		questionsResponseAfterSelectMu.Lock()
		questionsResponseAfterSelectHooks = append(questionsResponseAfterSelectHooks, questionsResponseHook)
		questionsResponseAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		questionsResponseBeforeInsertMu.Lock()
		questionsResponseBeforeInsertHooks = append(questionsResponseBeforeInsertHooks, questionsResponseHook)
		questionsResponseBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		questionsResponseAfterInsertMu.Lock()
		questionsResponseAfterInsertHooks = append(questionsResponseAfterInsertHooks, questionsResponseHook)
		questionsResponseAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		questionsResponseBeforeUpdateMu.Lock()
		questionsResponseBeforeUpdateHooks = append(questionsResponseBeforeUpdateHooks, questionsResponseHook)
		questionsResponseBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		questionsResponseAfterUpdateMu.Lock()
		questionsResponseAfterUpdateHooks = append(questionsResponseAfterUpdateHooks, questionsResponseHook)
		questionsResponseAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		questionsResponseBeforeDeleteMu.Lock()
		questionsResponseBeforeDeleteHooks = append(questionsResponseBeforeDeleteHooks, questionsResponseHook)
		questionsResponseBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		questionsResponseAfterDeleteMu.Lock()
		questionsResponseAfterDeleteHooks = append(questionsResponseAfterDeleteHooks, questionsResponseHook)
		questionsResponseAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		questionsResponseBeforeUpsertMu.Lock()
		questionsResponseBeforeUpsertHooks = append(questionsResponseBeforeUpsertHooks, questionsResponseHook)
		questionsResponseBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		questionsResponseAfterUpsertMu.Lock()
		questionsResponseAfterUpsertHooks = append(questionsResponseAfterUpsertHooks, questionsResponseHook)
		questionsResponseAfterUpsertMu.Unlock()
	}
}

// One returns a single questionsResponse record from the query.
func (q questionsResponseQuery) One(ctx context.Context, exec boil.ContextExecutor) (*QuestionsResponse, error) {
	o := &QuestionsResponse{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for questions_responses")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all QuestionsResponse records from the query.
func (q questionsResponseQuery) All(ctx context.Context, exec boil.ContextExecutor) (QuestionsResponseSlice, error) {
	var o []*QuestionsResponse

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to QuestionsResponse slice")
	}

	if len(questionsResponseAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all QuestionsResponse records in the query.
func (q questionsResponseQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count questions_responses rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q questionsResponseQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if questions_responses exists")
	}

	return count > 0, nil
}

// Question pointed to by the foreign key.
func (o *QuestionsResponse) Question(mods ...qm.QueryMod) questionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.QuestionID),
	}

	queryMods = append(queryMods, mods...)

	return Questions(queryMods...)
}

// User pointed to by the foreign key.
func (o *QuestionsResponse) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadQuestion allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (questionsResponseL) LoadQuestion(ctx context.Context, e boil.ContextExecutor, singular bool, maybeQuestionsResponse interface{}, mods queries.Applicator) error {
	var slice []*QuestionsResponse
	var object *QuestionsResponse

	if singular {
		var ok bool
		object, ok = maybeQuestionsResponse.(*QuestionsResponse)
		if !ok {
			object = new(QuestionsResponse)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeQuestionsResponse)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeQuestionsResponse))
			}
		}
	} else {
		s, ok := maybeQuestionsResponse.(*[]*QuestionsResponse)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeQuestionsResponse)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeQuestionsResponse))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &questionsResponseR{}
		}
		args[object.QuestionID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &questionsResponseR{}
			}

			args[obj.QuestionID] = struct{}{}

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
		qm.From(`questions`),
		qm.WhereIn(`questions.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Question")
	}

	var resultSlice []*Question
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Question")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for questions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for questions")
	}

	if len(questionAfterSelectHooks) != 0 {
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
		object.R.Question = foreign
		if foreign.R == nil {
			foreign.R = &questionR{}
		}
		foreign.R.QuestionsResponses = append(foreign.R.QuestionsResponses, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.QuestionID == foreign.ID {
				local.R.Question = foreign
				if foreign.R == nil {
					foreign.R = &questionR{}
				}
				foreign.R.QuestionsResponses = append(foreign.R.QuestionsResponses, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (questionsResponseL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeQuestionsResponse interface{}, mods queries.Applicator) error {
	var slice []*QuestionsResponse
	var object *QuestionsResponse

	if singular {
		var ok bool
		object, ok = maybeQuestionsResponse.(*QuestionsResponse)
		if !ok {
			object = new(QuestionsResponse)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeQuestionsResponse)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeQuestionsResponse))
			}
		}
	} else {
		s, ok := maybeQuestionsResponse.(*[]*QuestionsResponse)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeQuestionsResponse)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeQuestionsResponse))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &questionsResponseR{}
		}
		args[object.UserID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &questionsResponseR{}
			}

			args[obj.UserID] = struct{}{}

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
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.QuestionsResponses = append(foreign.R.QuestionsResponses, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.QuestionsResponses = append(foreign.R.QuestionsResponses, local)
				break
			}
		}
	}

	return nil
}

// SetQuestion of the questionsResponse to the related item.
// Sets o.R.Question to related.
// Adds o to related.R.QuestionsResponses.
func (o *QuestionsResponse) SetQuestion(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Question) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"questions_responses\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"question_id"}),
		strmangle.WhereClause("\"", "\"", 2, questionsResponsePrimaryKeyColumns),
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

	o.QuestionID = related.ID
	if o.R == nil {
		o.R = &questionsResponseR{
			Question: related,
		}
	} else {
		o.R.Question = related
	}

	if related.R == nil {
		related.R = &questionR{
			QuestionsResponses: QuestionsResponseSlice{o},
		}
	} else {
		related.R.QuestionsResponses = append(related.R.QuestionsResponses, o)
	}

	return nil
}

// SetUser of the questionsResponse to the related item.
// Sets o.R.User to related.
// Adds o to related.R.QuestionsResponses.
func (o *QuestionsResponse) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"questions_responses\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, questionsResponsePrimaryKeyColumns),
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

	o.UserID = related.ID
	if o.R == nil {
		o.R = &questionsResponseR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			QuestionsResponses: QuestionsResponseSlice{o},
		}
	} else {
		related.R.QuestionsResponses = append(related.R.QuestionsResponses, o)
	}

	return nil
}

// QuestionsResponses retrieves all the records using an executor.
func QuestionsResponses(mods ...qm.QueryMod) questionsResponseQuery {
	mods = append(mods, qm.From("\"questions_responses\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"questions_responses\".*"})
	}

	return questionsResponseQuery{q}
}

// FindQuestionsResponse retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindQuestionsResponse(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*QuestionsResponse, error) {
	questionsResponseObj := &QuestionsResponse{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"questions_responses\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, questionsResponseObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from questions_responses")
	}

	if err = questionsResponseObj.doAfterSelectHooks(ctx, exec); err != nil {
		return questionsResponseObj, err
	}

	return questionsResponseObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *QuestionsResponse) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no questions_responses provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(questionsResponseColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	questionsResponseInsertCacheMut.RLock()
	cache, cached := questionsResponseInsertCache[key]
	questionsResponseInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			questionsResponseAllColumns,
			questionsResponseColumnsWithDefault,
			questionsResponseColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(questionsResponseType, questionsResponseMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(questionsResponseType, questionsResponseMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"questions_responses\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"questions_responses\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into questions_responses")
	}

	if !cached {
		questionsResponseInsertCacheMut.Lock()
		questionsResponseInsertCache[key] = cache
		questionsResponseInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the QuestionsResponse.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *QuestionsResponse) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	questionsResponseUpdateCacheMut.RLock()
	cache, cached := questionsResponseUpdateCache[key]
	questionsResponseUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			questionsResponseAllColumns,
			questionsResponsePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update questions_responses, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"questions_responses\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, questionsResponsePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(questionsResponseType, questionsResponseMapping, append(wl, questionsResponsePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update questions_responses row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for questions_responses")
	}

	if !cached {
		questionsResponseUpdateCacheMut.Lock()
		questionsResponseUpdateCache[key] = cache
		questionsResponseUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q questionsResponseQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for questions_responses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for questions_responses")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o QuestionsResponseSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionsResponsePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"questions_responses\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, questionsResponsePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in questionsResponse slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all questionsResponse")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *QuestionsResponse) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no questions_responses provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(questionsResponseColumnsWithDefault, o)

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

	questionsResponseUpsertCacheMut.RLock()
	cache, cached := questionsResponseUpsertCache[key]
	questionsResponseUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			questionsResponseAllColumns,
			questionsResponseColumnsWithDefault,
			questionsResponseColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			questionsResponseAllColumns,
			questionsResponsePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert questions_responses, could not build update column list")
		}

		ret := strmangle.SetComplement(questionsResponseAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(questionsResponsePrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert questions_responses, could not build conflict column list")
			}

			conflict = make([]string, len(questionsResponsePrimaryKeyColumns))
			copy(conflict, questionsResponsePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"questions_responses\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(questionsResponseType, questionsResponseMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(questionsResponseType, questionsResponseMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert questions_responses")
	}

	if !cached {
		questionsResponseUpsertCacheMut.Lock()
		questionsResponseUpsertCache[key] = cache
		questionsResponseUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single QuestionsResponse record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *QuestionsResponse) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no QuestionsResponse provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), questionsResponsePrimaryKeyMapping)
	sql := "DELETE FROM \"questions_responses\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from questions_responses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for questions_responses")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q questionsResponseQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no questionsResponseQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from questions_responses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for questions_responses")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o QuestionsResponseSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(questionsResponseBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionsResponsePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"questions_responses\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, questionsResponsePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from questionsResponse slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for questions_responses")
	}

	if len(questionsResponseAfterDeleteHooks) != 0 {
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
func (o *QuestionsResponse) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindQuestionsResponse(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *QuestionsResponseSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := QuestionsResponseSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionsResponsePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"questions_responses\".* FROM \"questions_responses\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, questionsResponsePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in QuestionsResponseSlice")
	}

	*o = slice

	return nil
}

// QuestionsResponseExists checks if the QuestionsResponse row exists.
func QuestionsResponseExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"questions_responses\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if questions_responses exists")
	}

	return exists, nil
}

// Exists checks if the QuestionsResponse row exists.
func (o *QuestionsResponse) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return QuestionsResponseExists(ctx, exec, o.ID)
}
