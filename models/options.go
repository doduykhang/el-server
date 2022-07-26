// Code generated by SQLBoiler 4.12.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Option is an object representing the database table.
type Option struct {
	ID         uint   `boil:"id" json:"id" toml:"id" yaml:"id"`
	Content    string `boil:"content" json:"content" toml:"content" yaml:"content"`
	Position   uint   `boil:"position" json:"position" toml:"position" yaml:"position"`
	QuestionID uint   `boil:"question_id" json:"questionID" toml:"questionID" yaml:"questionID"`

	R *optionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L optionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OptionColumns = struct {
	ID         string
	Content    string
	Position   string
	QuestionID string
}{
	ID:         "id",
	Content:    "content",
	Position:   "position",
	QuestionID: "question_id",
}

var OptionTableColumns = struct {
	ID         string
	Content    string
	Position   string
	QuestionID string
}{
	ID:         "options.id",
	Content:    "options.content",
	Position:   "options.position",
	QuestionID: "options.question_id",
}

// Generated where

var OptionWhere = struct {
	ID         whereHelperuint
	Content    whereHelperstring
	Position   whereHelperuint
	QuestionID whereHelperuint
}{
	ID:         whereHelperuint{field: "`options`.`id`"},
	Content:    whereHelperstring{field: "`options`.`content`"},
	Position:   whereHelperuint{field: "`options`.`position`"},
	QuestionID: whereHelperuint{field: "`options`.`question_id`"},
}

// OptionRels is where relationship names are stored.
var OptionRels = struct {
	Question string
}{
	Question: "Question",
}

// optionR is where relationships are stored.
type optionR struct {
	Question *Question `boil:"Question" json:"Question" toml:"Question" yaml:"Question"`
}

// NewStruct creates a new relationship struct
func (*optionR) NewStruct() *optionR {
	return &optionR{}
}

func (r *optionR) GetQuestion() *Question {
	if r == nil {
		return nil
	}
	return r.Question
}

// optionL is where Load methods for each relationship are stored.
type optionL struct{}

var (
	optionAllColumns            = []string{"id", "content", "position", "question_id"}
	optionColumnsWithoutDefault = []string{"content", "position", "question_id"}
	optionColumnsWithDefault    = []string{"id"}
	optionPrimaryKeyColumns     = []string{"id"}
	optionGeneratedColumns      = []string{}
)

type (
	// OptionSlice is an alias for a slice of pointers to Option.
	// This should almost always be used instead of []Option.
	OptionSlice []*Option
	// OptionHook is the signature for custom Option hook methods
	OptionHook func(context.Context, boil.ContextExecutor, *Option) error

	optionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	optionType                 = reflect.TypeOf(&Option{})
	optionMapping              = queries.MakeStructMapping(optionType)
	optionPrimaryKeyMapping, _ = queries.BindMapping(optionType, optionMapping, optionPrimaryKeyColumns)
	optionInsertCacheMut       sync.RWMutex
	optionInsertCache          = make(map[string]insertCache)
	optionUpdateCacheMut       sync.RWMutex
	optionUpdateCache          = make(map[string]updateCache)
	optionUpsertCacheMut       sync.RWMutex
	optionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var optionAfterSelectHooks []OptionHook

var optionBeforeInsertHooks []OptionHook
var optionAfterInsertHooks []OptionHook

var optionBeforeUpdateHooks []OptionHook
var optionAfterUpdateHooks []OptionHook

var optionBeforeDeleteHooks []OptionHook
var optionAfterDeleteHooks []OptionHook

var optionBeforeUpsertHooks []OptionHook
var optionAfterUpsertHooks []OptionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Option) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range optionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Option) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range optionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Option) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range optionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Option) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range optionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Option) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range optionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Option) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range optionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Option) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range optionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Option) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range optionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Option) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range optionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOptionHook registers your hook function for all future operations.
func AddOptionHook(hookPoint boil.HookPoint, optionHook OptionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		optionAfterSelectHooks = append(optionAfterSelectHooks, optionHook)
	case boil.BeforeInsertHook:
		optionBeforeInsertHooks = append(optionBeforeInsertHooks, optionHook)
	case boil.AfterInsertHook:
		optionAfterInsertHooks = append(optionAfterInsertHooks, optionHook)
	case boil.BeforeUpdateHook:
		optionBeforeUpdateHooks = append(optionBeforeUpdateHooks, optionHook)
	case boil.AfterUpdateHook:
		optionAfterUpdateHooks = append(optionAfterUpdateHooks, optionHook)
	case boil.BeforeDeleteHook:
		optionBeforeDeleteHooks = append(optionBeforeDeleteHooks, optionHook)
	case boil.AfterDeleteHook:
		optionAfterDeleteHooks = append(optionAfterDeleteHooks, optionHook)
	case boil.BeforeUpsertHook:
		optionBeforeUpsertHooks = append(optionBeforeUpsertHooks, optionHook)
	case boil.AfterUpsertHook:
		optionAfterUpsertHooks = append(optionAfterUpsertHooks, optionHook)
	}
}

// One returns a single option record from the query.
func (q optionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Option, error) {
	o := &Option{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for options")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Option records from the query.
func (q optionQuery) All(ctx context.Context, exec boil.ContextExecutor) (OptionSlice, error) {
	var o []*Option

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Option slice")
	}

	if len(optionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Option records in the query.
func (q optionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count options rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q optionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if options exists")
	}

	return count > 0, nil
}

// Question pointed to by the foreign key.
func (o *Option) Question(mods ...qm.QueryMod) questionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.QuestionID),
	}

	queryMods = append(queryMods, mods...)

	return Questions(queryMods...)
}

// LoadQuestion allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (optionL) LoadQuestion(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOption interface{}, mods queries.Applicator) error {
	var slice []*Option
	var object *Option

	if singular {
		var ok bool
		object, ok = maybeOption.(*Option)
		if !ok {
			object = new(Option)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeOption)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeOption))
			}
		}
	} else {
		s, ok := maybeOption.(*[]*Option)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeOption)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeOption))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &optionR{}
		}
		args = append(args, object.QuestionID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &optionR{}
			}

			for _, a := range args {
				if a == obj.QuestionID {
					continue Outer
				}
			}

			args = append(args, obj.QuestionID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`questions`),
		qm.WhereIn(`questions.id in ?`, args...),
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

	if len(optionAfterSelectHooks) != 0 {
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
		foreign.R.Options = append(foreign.R.Options, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.QuestionID == foreign.ID {
				local.R.Question = foreign
				if foreign.R == nil {
					foreign.R = &questionR{}
				}
				foreign.R.Options = append(foreign.R.Options, local)
				break
			}
		}
	}

	return nil
}

// SetQuestion of the option to the related item.
// Sets o.R.Question to related.
// Adds o to related.R.Options.
func (o *Option) SetQuestion(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Question) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `options` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"question_id"}),
		strmangle.WhereClause("`", "`", 0, optionPrimaryKeyColumns),
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
		o.R = &optionR{
			Question: related,
		}
	} else {
		o.R.Question = related
	}

	if related.R == nil {
		related.R = &questionR{
			Options: OptionSlice{o},
		}
	} else {
		related.R.Options = append(related.R.Options, o)
	}

	return nil
}

// Options retrieves all the records using an executor.
func Options(mods ...qm.QueryMod) optionQuery {
	mods = append(mods, qm.From("`options`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`options`.*"})
	}

	return optionQuery{q}
}

// FindOption retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOption(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*Option, error) {
	optionObj := &Option{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `options` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, optionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from options")
	}

	if err = optionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return optionObj, err
	}

	return optionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Option) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no options provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(optionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	optionInsertCacheMut.RLock()
	cache, cached := optionInsertCache[key]
	optionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			optionAllColumns,
			optionColumnsWithDefault,
			optionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(optionType, optionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(optionType, optionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `options` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `options` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `options` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, optionPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into options")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == optionMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for options")
	}

CacheNoHooks:
	if !cached {
		optionInsertCacheMut.Lock()
		optionInsertCache[key] = cache
		optionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Option.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Option) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	optionUpdateCacheMut.RLock()
	cache, cached := optionUpdateCache[key]
	optionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			optionAllColumns,
			optionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update options, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `options` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, optionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(optionType, optionMapping, append(wl, optionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update options row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for options")
	}

	if !cached {
		optionUpdateCacheMut.Lock()
		optionUpdateCache[key] = cache
		optionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q optionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for options")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for options")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OptionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), optionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `options` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, optionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in option slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all option")
	}
	return rowsAff, nil
}

var mySQLOptionUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Option) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no options provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(optionColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOptionUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	optionUpsertCacheMut.RLock()
	cache, cached := optionUpsertCache[key]
	optionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			optionAllColumns,
			optionColumnsWithDefault,
			optionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			optionAllColumns,
			optionPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert options, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`options`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `options` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(optionType, optionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(optionType, optionMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for options")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == optionMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(optionType, optionMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for options")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for options")
	}

CacheNoHooks:
	if !cached {
		optionUpsertCacheMut.Lock()
		optionUpsertCache[key] = cache
		optionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Option record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Option) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Option provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), optionPrimaryKeyMapping)
	sql := "DELETE FROM `options` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from options")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for options")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q optionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no optionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from options")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for options")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OptionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(optionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), optionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `options` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, optionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from option slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for options")
	}

	if len(optionAfterDeleteHooks) != 0 {
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
func (o *Option) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOption(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OptionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OptionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), optionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `options`.* FROM `options` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, optionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OptionSlice")
	}

	*o = slice

	return nil
}

// OptionExists checks if the Option row exists.
func OptionExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `options` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if options exists")
	}

	return exists, nil
}
