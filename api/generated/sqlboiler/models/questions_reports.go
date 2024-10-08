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

// QuestionsReport is an object representing the database table.
type QuestionsReport struct {
	ID         int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	QuestionID int64     `boil:"question_id" json:"question_id" toml:"question_id" yaml:"question_id"`
	ReporterID int64     `boil:"reporter_id" json:"reporter_id" toml:"reporter_id" yaml:"reporter_id"`
	Timestamp  time.Time `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	// fraud, bad quality etc
	ReportType        string `boil:"report_type" json:"report_type" toml:"report_type" yaml:"report_type"`
	ReportDescription string `boil:"report_description" json:"report_description" toml:"report_description" yaml:"report_description"`
	// eg pending, approved, declined
	ReportStatus int64 `boil:"report_status" json:"report_status" toml:"report_status" yaml:"report_status"`

	R *questionsReportR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L questionsReportL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var QuestionsReportColumns = struct {
	ID                string
	QuestionID        string
	ReporterID        string
	Timestamp         string
	ReportType        string
	ReportDescription string
	ReportStatus      string
}{
	ID:                "id",
	QuestionID:        "question_id",
	ReporterID:        "reporter_id",
	Timestamp:         "timestamp",
	ReportType:        "report_type",
	ReportDescription: "report_description",
	ReportStatus:      "report_status",
}

var QuestionsReportTableColumns = struct {
	ID                string
	QuestionID        string
	ReporterID        string
	Timestamp         string
	ReportType        string
	ReportDescription string
	ReportStatus      string
}{
	ID:                "questions_reports.id",
	QuestionID:        "questions_reports.question_id",
	ReporterID:        "questions_reports.reporter_id",
	Timestamp:         "questions_reports.timestamp",
	ReportType:        "questions_reports.report_type",
	ReportDescription: "questions_reports.report_description",
	ReportStatus:      "questions_reports.report_status",
}

// Generated where

var QuestionsReportWhere = struct {
	ID                whereHelperint64
	QuestionID        whereHelperint64
	ReporterID        whereHelperint64
	Timestamp         whereHelpertime_Time
	ReportType        whereHelperstring
	ReportDescription whereHelperstring
	ReportStatus      whereHelperint64
}{
	ID:                whereHelperint64{field: "\"questions_reports\".\"id\""},
	QuestionID:        whereHelperint64{field: "\"questions_reports\".\"question_id\""},
	ReporterID:        whereHelperint64{field: "\"questions_reports\".\"reporter_id\""},
	Timestamp:         whereHelpertime_Time{field: "\"questions_reports\".\"timestamp\""},
	ReportType:        whereHelperstring{field: "\"questions_reports\".\"report_type\""},
	ReportDescription: whereHelperstring{field: "\"questions_reports\".\"report_description\""},
	ReportStatus:      whereHelperint64{field: "\"questions_reports\".\"report_status\""},
}

// QuestionsReportRels is where relationship names are stored.
var QuestionsReportRels = struct {
	Question string
	Reporter string
}{
	Question: "Question",
	Reporter: "Reporter",
}

// questionsReportR is where relationships are stored.
type questionsReportR struct {
	Question *Question `boil:"Question" json:"Question" toml:"Question" yaml:"Question"`
	Reporter *User     `boil:"Reporter" json:"Reporter" toml:"Reporter" yaml:"Reporter"`
}

// NewStruct creates a new relationship struct
func (*questionsReportR) NewStruct() *questionsReportR {
	return &questionsReportR{}
}

func (r *questionsReportR) GetQuestion() *Question {
	if r == nil {
		return nil
	}
	return r.Question
}

func (r *questionsReportR) GetReporter() *User {
	if r == nil {
		return nil
	}
	return r.Reporter
}

// questionsReportL is where Load methods for each relationship are stored.
type questionsReportL struct{}

var (
	questionsReportAllColumns            = []string{"id", "question_id", "reporter_id", "timestamp", "report_type", "report_description", "report_status"}
	questionsReportColumnsWithoutDefault = []string{"question_id", "reporter_id", "timestamp", "report_type", "report_description", "report_status"}
	questionsReportColumnsWithDefault    = []string{"id"}
	questionsReportPrimaryKeyColumns     = []string{"id"}
	questionsReportGeneratedColumns      = []string{}
)

type (
	// QuestionsReportSlice is an alias for a slice of pointers to QuestionsReport.
	// This should almost always be used instead of []QuestionsReport.
	QuestionsReportSlice []*QuestionsReport
	// QuestionsReportHook is the signature for custom QuestionsReport hook methods
	QuestionsReportHook func(context.Context, boil.ContextExecutor, *QuestionsReport) error

	questionsReportQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	questionsReportType                 = reflect.TypeOf(&QuestionsReport{})
	questionsReportMapping              = queries.MakeStructMapping(questionsReportType)
	questionsReportPrimaryKeyMapping, _ = queries.BindMapping(questionsReportType, questionsReportMapping, questionsReportPrimaryKeyColumns)
	questionsReportInsertCacheMut       sync.RWMutex
	questionsReportInsertCache          = make(map[string]insertCache)
	questionsReportUpdateCacheMut       sync.RWMutex
	questionsReportUpdateCache          = make(map[string]updateCache)
	questionsReportUpsertCacheMut       sync.RWMutex
	questionsReportUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var questionsReportAfterSelectMu sync.Mutex
var questionsReportAfterSelectHooks []QuestionsReportHook

var questionsReportBeforeInsertMu sync.Mutex
var questionsReportBeforeInsertHooks []QuestionsReportHook
var questionsReportAfterInsertMu sync.Mutex
var questionsReportAfterInsertHooks []QuestionsReportHook

var questionsReportBeforeUpdateMu sync.Mutex
var questionsReportBeforeUpdateHooks []QuestionsReportHook
var questionsReportAfterUpdateMu sync.Mutex
var questionsReportAfterUpdateHooks []QuestionsReportHook

var questionsReportBeforeDeleteMu sync.Mutex
var questionsReportBeforeDeleteHooks []QuestionsReportHook
var questionsReportAfterDeleteMu sync.Mutex
var questionsReportAfterDeleteHooks []QuestionsReportHook

var questionsReportBeforeUpsertMu sync.Mutex
var questionsReportBeforeUpsertHooks []QuestionsReportHook
var questionsReportAfterUpsertMu sync.Mutex
var questionsReportAfterUpsertHooks []QuestionsReportHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *QuestionsReport) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsReportAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *QuestionsReport) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsReportBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *QuestionsReport) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsReportAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *QuestionsReport) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsReportBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *QuestionsReport) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsReportAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *QuestionsReport) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsReportBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *QuestionsReport) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsReportAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *QuestionsReport) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsReportBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *QuestionsReport) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range questionsReportAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddQuestionsReportHook registers your hook function for all future operations.
func AddQuestionsReportHook(hookPoint boil.HookPoint, questionsReportHook QuestionsReportHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		questionsReportAfterSelectMu.Lock()
		questionsReportAfterSelectHooks = append(questionsReportAfterSelectHooks, questionsReportHook)
		questionsReportAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		questionsReportBeforeInsertMu.Lock()
		questionsReportBeforeInsertHooks = append(questionsReportBeforeInsertHooks, questionsReportHook)
		questionsReportBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		questionsReportAfterInsertMu.Lock()
		questionsReportAfterInsertHooks = append(questionsReportAfterInsertHooks, questionsReportHook)
		questionsReportAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		questionsReportBeforeUpdateMu.Lock()
		questionsReportBeforeUpdateHooks = append(questionsReportBeforeUpdateHooks, questionsReportHook)
		questionsReportBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		questionsReportAfterUpdateMu.Lock()
		questionsReportAfterUpdateHooks = append(questionsReportAfterUpdateHooks, questionsReportHook)
		questionsReportAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		questionsReportBeforeDeleteMu.Lock()
		questionsReportBeforeDeleteHooks = append(questionsReportBeforeDeleteHooks, questionsReportHook)
		questionsReportBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		questionsReportAfterDeleteMu.Lock()
		questionsReportAfterDeleteHooks = append(questionsReportAfterDeleteHooks, questionsReportHook)
		questionsReportAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		questionsReportBeforeUpsertMu.Lock()
		questionsReportBeforeUpsertHooks = append(questionsReportBeforeUpsertHooks, questionsReportHook)
		questionsReportBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		questionsReportAfterUpsertMu.Lock()
		questionsReportAfterUpsertHooks = append(questionsReportAfterUpsertHooks, questionsReportHook)
		questionsReportAfterUpsertMu.Unlock()
	}
}

// One returns a single questionsReport record from the query.
func (q questionsReportQuery) One(ctx context.Context, exec boil.ContextExecutor) (*QuestionsReport, error) {
	o := &QuestionsReport{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for questions_reports")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all QuestionsReport records from the query.
func (q questionsReportQuery) All(ctx context.Context, exec boil.ContextExecutor) (QuestionsReportSlice, error) {
	var o []*QuestionsReport

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to QuestionsReport slice")
	}

	if len(questionsReportAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all QuestionsReport records in the query.
func (q questionsReportQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count questions_reports rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q questionsReportQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if questions_reports exists")
	}

	return count > 0, nil
}

// Question pointed to by the foreign key.
func (o *QuestionsReport) Question(mods ...qm.QueryMod) questionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.QuestionID),
	}

	queryMods = append(queryMods, mods...)

	return Questions(queryMods...)
}

// Reporter pointed to by the foreign key.
func (o *QuestionsReport) Reporter(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ReporterID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadQuestion allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (questionsReportL) LoadQuestion(ctx context.Context, e boil.ContextExecutor, singular bool, maybeQuestionsReport interface{}, mods queries.Applicator) error {
	var slice []*QuestionsReport
	var object *QuestionsReport

	if singular {
		var ok bool
		object, ok = maybeQuestionsReport.(*QuestionsReport)
		if !ok {
			object = new(QuestionsReport)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeQuestionsReport)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeQuestionsReport))
			}
		}
	} else {
		s, ok := maybeQuestionsReport.(*[]*QuestionsReport)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeQuestionsReport)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeQuestionsReport))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &questionsReportR{}
		}
		args[object.QuestionID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &questionsReportR{}
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
		foreign.R.QuestionsReports = append(foreign.R.QuestionsReports, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.QuestionID == foreign.ID {
				local.R.Question = foreign
				if foreign.R == nil {
					foreign.R = &questionR{}
				}
				foreign.R.QuestionsReports = append(foreign.R.QuestionsReports, local)
				break
			}
		}
	}

	return nil
}

// LoadReporter allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (questionsReportL) LoadReporter(ctx context.Context, e boil.ContextExecutor, singular bool, maybeQuestionsReport interface{}, mods queries.Applicator) error {
	var slice []*QuestionsReport
	var object *QuestionsReport

	if singular {
		var ok bool
		object, ok = maybeQuestionsReport.(*QuestionsReport)
		if !ok {
			object = new(QuestionsReport)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeQuestionsReport)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeQuestionsReport))
			}
		}
	} else {
		s, ok := maybeQuestionsReport.(*[]*QuestionsReport)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeQuestionsReport)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeQuestionsReport))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &questionsReportR{}
		}
		args[object.ReporterID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &questionsReportR{}
			}

			args[obj.ReporterID] = struct{}{}

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
		object.R.Reporter = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.ReporterQuestionsReports = append(foreign.R.ReporterQuestionsReports, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ReporterID == foreign.ID {
				local.R.Reporter = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.ReporterQuestionsReports = append(foreign.R.ReporterQuestionsReports, local)
				break
			}
		}
	}

	return nil
}

// SetQuestion of the questionsReport to the related item.
// Sets o.R.Question to related.
// Adds o to related.R.QuestionsReports.
func (o *QuestionsReport) SetQuestion(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Question) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"questions_reports\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"question_id"}),
		strmangle.WhereClause("\"", "\"", 2, questionsReportPrimaryKeyColumns),
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
		o.R = &questionsReportR{
			Question: related,
		}
	} else {
		o.R.Question = related
	}

	if related.R == nil {
		related.R = &questionR{
			QuestionsReports: QuestionsReportSlice{o},
		}
	} else {
		related.R.QuestionsReports = append(related.R.QuestionsReports, o)
	}

	return nil
}

// SetReporter of the questionsReport to the related item.
// Sets o.R.Reporter to related.
// Adds o to related.R.ReporterQuestionsReports.
func (o *QuestionsReport) SetReporter(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"questions_reports\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"reporter_id"}),
		strmangle.WhereClause("\"", "\"", 2, questionsReportPrimaryKeyColumns),
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

	o.ReporterID = related.ID
	if o.R == nil {
		o.R = &questionsReportR{
			Reporter: related,
		}
	} else {
		o.R.Reporter = related
	}

	if related.R == nil {
		related.R = &userR{
			ReporterQuestionsReports: QuestionsReportSlice{o},
		}
	} else {
		related.R.ReporterQuestionsReports = append(related.R.ReporterQuestionsReports, o)
	}

	return nil
}

// QuestionsReports retrieves all the records using an executor.
func QuestionsReports(mods ...qm.QueryMod) questionsReportQuery {
	mods = append(mods, qm.From("\"questions_reports\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"questions_reports\".*"})
	}

	return questionsReportQuery{q}
}

// FindQuestionsReport retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindQuestionsReport(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*QuestionsReport, error) {
	questionsReportObj := &QuestionsReport{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"questions_reports\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, questionsReportObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from questions_reports")
	}

	if err = questionsReportObj.doAfterSelectHooks(ctx, exec); err != nil {
		return questionsReportObj, err
	}

	return questionsReportObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *QuestionsReport) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no questions_reports provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(questionsReportColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	questionsReportInsertCacheMut.RLock()
	cache, cached := questionsReportInsertCache[key]
	questionsReportInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			questionsReportAllColumns,
			questionsReportColumnsWithDefault,
			questionsReportColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(questionsReportType, questionsReportMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(questionsReportType, questionsReportMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"questions_reports\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"questions_reports\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into questions_reports")
	}

	if !cached {
		questionsReportInsertCacheMut.Lock()
		questionsReportInsertCache[key] = cache
		questionsReportInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the QuestionsReport.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *QuestionsReport) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	questionsReportUpdateCacheMut.RLock()
	cache, cached := questionsReportUpdateCache[key]
	questionsReportUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			questionsReportAllColumns,
			questionsReportPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update questions_reports, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"questions_reports\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, questionsReportPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(questionsReportType, questionsReportMapping, append(wl, questionsReportPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update questions_reports row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for questions_reports")
	}

	if !cached {
		questionsReportUpdateCacheMut.Lock()
		questionsReportUpdateCache[key] = cache
		questionsReportUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q questionsReportQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for questions_reports")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for questions_reports")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o QuestionsReportSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionsReportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"questions_reports\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, questionsReportPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in questionsReport slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all questionsReport")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *QuestionsReport) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no questions_reports provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(questionsReportColumnsWithDefault, o)

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

	questionsReportUpsertCacheMut.RLock()
	cache, cached := questionsReportUpsertCache[key]
	questionsReportUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			questionsReportAllColumns,
			questionsReportColumnsWithDefault,
			questionsReportColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			questionsReportAllColumns,
			questionsReportPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert questions_reports, could not build update column list")
		}

		ret := strmangle.SetComplement(questionsReportAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(questionsReportPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert questions_reports, could not build conflict column list")
			}

			conflict = make([]string, len(questionsReportPrimaryKeyColumns))
			copy(conflict, questionsReportPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"questions_reports\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(questionsReportType, questionsReportMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(questionsReportType, questionsReportMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert questions_reports")
	}

	if !cached {
		questionsReportUpsertCacheMut.Lock()
		questionsReportUpsertCache[key] = cache
		questionsReportUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single QuestionsReport record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *QuestionsReport) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no QuestionsReport provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), questionsReportPrimaryKeyMapping)
	sql := "DELETE FROM \"questions_reports\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from questions_reports")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for questions_reports")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q questionsReportQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no questionsReportQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from questions_reports")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for questions_reports")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o QuestionsReportSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(questionsReportBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionsReportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"questions_reports\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, questionsReportPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from questionsReport slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for questions_reports")
	}

	if len(questionsReportAfterDeleteHooks) != 0 {
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
func (o *QuestionsReport) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindQuestionsReport(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *QuestionsReportSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := QuestionsReportSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), questionsReportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"questions_reports\".* FROM \"questions_reports\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, questionsReportPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in QuestionsReportSlice")
	}

	*o = slice

	return nil
}

// QuestionsReportExists checks if the QuestionsReport row exists.
func QuestionsReportExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"questions_reports\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if questions_reports exists")
	}

	return exists, nil
}

// Exists checks if the QuestionsReport row exists.
func (o *QuestionsReport) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return QuestionsReportExists(ctx, exec, o.ID)
}
