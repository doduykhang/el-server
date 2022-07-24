// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// UserFolder is an object representing the database table.
type UserFolder struct {
	ID     int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name   string `boil:"name" json:"name" toml:"name" yaml:"name"`
	UserID int    `boil:"user_id" json:"userID" toml:"userID" yaml:"userID"`

	R *userFolderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userFolderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserFolderColumns = struct {
	ID     string
	Name   string
	UserID string
}{
	ID:     "id",
	Name:   "name",
	UserID: "user_id",
}

var UserFolderTableColumns = struct {
	ID     string
	Name   string
	UserID string
}{
	ID:     "user_folder.id",
	Name:   "user_folder.name",
	UserID: "user_folder.user_id",
}

// Generated where

var UserFolderWhere = struct {
	ID     whereHelperint
	Name   whereHelperstring
	UserID whereHelperint
}{
	ID:     whereHelperint{field: "\"user_folder\".\"id\""},
	Name:   whereHelperstring{field: "\"user_folder\".\"name\""},
	UserID: whereHelperint{field: "\"user_folder\".\"user_id\""},
}

// UserFolderRels is where relationship names are stored.
var UserFolderRels = struct {
	User  string
	Words string
}{
	User:  "User",
	Words: "Words",
}

// userFolderR is where relationships are stored.
type userFolderR struct {
	User  *User     `boil:"User" json:"User" toml:"User" yaml:"User"`
	Words WordSlice `boil:"Words" json:"Words" toml:"Words" yaml:"Words"`
}

// NewStruct creates a new relationship struct
func (*userFolderR) NewStruct() *userFolderR {
	return &userFolderR{}
}

func (r *userFolderR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

func (r *userFolderR) GetWords() WordSlice {
	if r == nil {
		return nil
	}
	return r.Words
}

// userFolderL is where Load methods for each relationship are stored.
type userFolderL struct{}

var (
	userFolderAllColumns            = []string{"id", "name", "user_id"}
	userFolderColumnsWithoutDefault = []string{"id", "name", "user_id"}
	userFolderColumnsWithDefault    = []string{}
	userFolderPrimaryKeyColumns     = []string{"id"}
	userFolderGeneratedColumns      = []string{}
)

type (
	// UserFolderSlice is an alias for a slice of pointers to UserFolder.
	// This should almost always be used instead of []UserFolder.
	UserFolderSlice []*UserFolder
	// UserFolderHook is the signature for custom UserFolder hook methods
	UserFolderHook func(context.Context, boil.ContextExecutor, *UserFolder) error

	userFolderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userFolderType                 = reflect.TypeOf(&UserFolder{})
	userFolderMapping              = queries.MakeStructMapping(userFolderType)
	userFolderPrimaryKeyMapping, _ = queries.BindMapping(userFolderType, userFolderMapping, userFolderPrimaryKeyColumns)
	userFolderInsertCacheMut       sync.RWMutex
	userFolderInsertCache          = make(map[string]insertCache)
	userFolderUpdateCacheMut       sync.RWMutex
	userFolderUpdateCache          = make(map[string]updateCache)
	userFolderUpsertCacheMut       sync.RWMutex
	userFolderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userFolderAfterSelectHooks []UserFolderHook

var userFolderBeforeInsertHooks []UserFolderHook
var userFolderAfterInsertHooks []UserFolderHook

var userFolderBeforeUpdateHooks []UserFolderHook
var userFolderAfterUpdateHooks []UserFolderHook

var userFolderBeforeDeleteHooks []UserFolderHook
var userFolderAfterDeleteHooks []UserFolderHook

var userFolderBeforeUpsertHooks []UserFolderHook
var userFolderAfterUpsertHooks []UserFolderHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserFolder) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userFolderAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserFolder) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userFolderBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserFolder) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userFolderAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserFolder) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userFolderBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserFolder) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userFolderAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserFolder) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userFolderBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserFolder) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userFolderAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserFolder) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userFolderBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserFolder) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userFolderAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserFolderHook registers your hook function for all future operations.
func AddUserFolderHook(hookPoint boil.HookPoint, userFolderHook UserFolderHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userFolderAfterSelectHooks = append(userFolderAfterSelectHooks, userFolderHook)
	case boil.BeforeInsertHook:
		userFolderBeforeInsertHooks = append(userFolderBeforeInsertHooks, userFolderHook)
	case boil.AfterInsertHook:
		userFolderAfterInsertHooks = append(userFolderAfterInsertHooks, userFolderHook)
	case boil.BeforeUpdateHook:
		userFolderBeforeUpdateHooks = append(userFolderBeforeUpdateHooks, userFolderHook)
	case boil.AfterUpdateHook:
		userFolderAfterUpdateHooks = append(userFolderAfterUpdateHooks, userFolderHook)
	case boil.BeforeDeleteHook:
		userFolderBeforeDeleteHooks = append(userFolderBeforeDeleteHooks, userFolderHook)
	case boil.AfterDeleteHook:
		userFolderAfterDeleteHooks = append(userFolderAfterDeleteHooks, userFolderHook)
	case boil.BeforeUpsertHook:
		userFolderBeforeUpsertHooks = append(userFolderBeforeUpsertHooks, userFolderHook)
	case boil.AfterUpsertHook:
		userFolderAfterUpsertHooks = append(userFolderAfterUpsertHooks, userFolderHook)
	}
}

// One returns a single userFolder record from the query.
func (q userFolderQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserFolder, error) {
	o := &UserFolder{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_folder")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserFolder records from the query.
func (q userFolderQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserFolderSlice, error) {
	var o []*UserFolder

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserFolder slice")
	}

	if len(userFolderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserFolder records in the query.
func (q userFolderQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_folder rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userFolderQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_folder exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UserFolder) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// Words retrieves all the word's Words with an executor.
func (o *UserFolder) Words(mods ...qm.QueryMod) wordQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"word_in_folder\" on \"words\".\"id\" = \"word_in_folder\".\"word_id\""),
		qm.Where("\"word_in_folder\".\"folder_id\"=?", o.ID),
	)

	return Words(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userFolderL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserFolder interface{}, mods queries.Applicator) error {
	var slice []*UserFolder
	var object *UserFolder

	if singular {
		object = maybeUserFolder.(*UserFolder)
	} else {
		slice = *maybeUserFolder.(*[]*UserFolder)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userFolderR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userFolderR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
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

	if len(userFolderAfterSelectHooks) != 0 {
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
		foreign.R.UserFolders = append(foreign.R.UserFolders, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserFolders = append(foreign.R.UserFolders, local)
				break
			}
		}
	}

	return nil
}

// LoadWords allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (userFolderL) LoadWords(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserFolder interface{}, mods queries.Applicator) error {
	var slice []*UserFolder
	var object *UserFolder

	if singular {
		object = maybeUserFolder.(*UserFolder)
	} else {
		slice = *maybeUserFolder.(*[]*UserFolder)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userFolderR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userFolderR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("\"words\".\"id\", \"words\".\"word\", \"words\".\"definition\", \"words\".\"example\", \"words\".\"pronouciation\", \"words\".\"type\", \"words\".\"manager_id\", \"a\".\"folder_id\""),
		qm.From("\"words\""),
		qm.InnerJoin("\"word_in_folder\" as \"a\" on \"words\".\"id\" = \"a\".\"word_id\""),
		qm.WhereIn("\"a\".\"folder_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load words")
	}

	var resultSlice []*Word

	var localJoinCols []int
	for results.Next() {
		one := new(Word)
		var localJoinCol int

		err = results.Scan(&one.ID, &one.Word, &one.Definition, &one.Example, &one.Pronouciation, &one.Type, &one.ManagerID, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for words")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice words")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on words")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for words")
	}

	if len(wordAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Words = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &wordR{}
			}
			foreign.R.FolderUserFolders = append(foreign.R.FolderUserFolders, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Words = append(local.R.Words, foreign)
				if foreign.R == nil {
					foreign.R = &wordR{}
				}
				foreign.R.FolderUserFolders = append(foreign.R.FolderUserFolders, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the userFolder to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserFolders.
func (o *UserFolder) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_folder\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userFolderPrimaryKeyColumns),
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
		o.R = &userFolderR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserFolders: UserFolderSlice{o},
		}
	} else {
		related.R.UserFolders = append(related.R.UserFolders, o)
	}

	return nil
}

// AddWords adds the given related objects to the existing relationships
// of the user_folder, optionally inserting them as new records.
// Appends related to o.R.Words.
// Sets related.R.FolderUserFolders appropriately.
func (o *UserFolder) AddWords(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Word) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"word_in_folder\" (\"folder_id\", \"word_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, query)
			fmt.Fprintln(writer, values)
		}
		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &userFolderR{
			Words: related,
		}
	} else {
		o.R.Words = append(o.R.Words, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &wordR{
				FolderUserFolders: UserFolderSlice{o},
			}
		} else {
			rel.R.FolderUserFolders = append(rel.R.FolderUserFolders, o)
		}
	}
	return nil
}

// SetWords removes all previously related items of the
// user_folder replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.FolderUserFolders's Words accordingly.
// Replaces o.R.Words with related.
// Sets related.R.FolderUserFolders's Words accordingly.
func (o *UserFolder) SetWords(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Word) error {
	query := "delete from \"word_in_folder\" where \"folder_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeWordsFromFolderUserFoldersSlice(o, related)
	if o.R != nil {
		o.R.Words = nil
	}

	return o.AddWords(ctx, exec, insert, related...)
}

// RemoveWords relationships from objects passed in.
// Removes related items from R.Words (uses pointer comparison, removal does not keep order)
// Sets related.R.FolderUserFolders.
func (o *UserFolder) RemoveWords(ctx context.Context, exec boil.ContextExecutor, related ...*Word) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	query := fmt.Sprintf(
		"delete from \"word_in_folder\" where \"folder_id\" = $1 and \"word_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeWordsFromFolderUserFoldersSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Words {
			if rel != ri {
				continue
			}

			ln := len(o.R.Words)
			if ln > 1 && i < ln-1 {
				o.R.Words[i] = o.R.Words[ln-1]
			}
			o.R.Words = o.R.Words[:ln-1]
			break
		}
	}

	return nil
}

func removeWordsFromFolderUserFoldersSlice(o *UserFolder, related []*Word) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.FolderUserFolders {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.FolderUserFolders)
			if ln > 1 && i < ln-1 {
				rel.R.FolderUserFolders[i] = rel.R.FolderUserFolders[ln-1]
			}
			rel.R.FolderUserFolders = rel.R.FolderUserFolders[:ln-1]
			break
		}
	}
}

// UserFolders retrieves all the records using an executor.
func UserFolders(mods ...qm.QueryMod) userFolderQuery {
	mods = append(mods, qm.From("\"user_folder\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"user_folder\".*"})
	}

	return userFolderQuery{q}
}

// FindUserFolder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserFolder(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*UserFolder, error) {
	userFolderObj := &UserFolder{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_folder\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userFolderObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_folder")
	}

	if err = userFolderObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userFolderObj, err
	}

	return userFolderObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserFolder) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_folder provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userFolderColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userFolderInsertCacheMut.RLock()
	cache, cached := userFolderInsertCache[key]
	userFolderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userFolderAllColumns,
			userFolderColumnsWithDefault,
			userFolderColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userFolderType, userFolderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userFolderType, userFolderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_folder\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_folder\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into user_folder")
	}

	if !cached {
		userFolderInsertCacheMut.Lock()
		userFolderInsertCache[key] = cache
		userFolderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserFolder.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserFolder) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userFolderUpdateCacheMut.RLock()
	cache, cached := userFolderUpdateCache[key]
	userFolderUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userFolderAllColumns,
			userFolderPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_folder, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_folder\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userFolderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userFolderType, userFolderMapping, append(wl, userFolderPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_folder row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_folder")
	}

	if !cached {
		userFolderUpdateCacheMut.Lock()
		userFolderUpdateCache[key] = cache
		userFolderUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userFolderQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_folder")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_folder")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserFolderSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userFolderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_folder\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userFolderPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userFolder slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userFolder")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserFolder) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_folder provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userFolderColumnsWithDefault, o)

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

	userFolderUpsertCacheMut.RLock()
	cache, cached := userFolderUpsertCache[key]
	userFolderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userFolderAllColumns,
			userFolderColumnsWithDefault,
			userFolderColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userFolderAllColumns,
			userFolderPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert user_folder, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userFolderPrimaryKeyColumns))
			copy(conflict, userFolderPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_folder\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userFolderType, userFolderMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userFolderType, userFolderMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert user_folder")
	}

	if !cached {
		userFolderUpsertCacheMut.Lock()
		userFolderUpsertCache[key] = cache
		userFolderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserFolder record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserFolder) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserFolder provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userFolderPrimaryKeyMapping)
	sql := "DELETE FROM \"user_folder\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_folder")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_folder")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userFolderQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userFolderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_folder")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_folder")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserFolderSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userFolderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userFolderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_folder\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userFolderPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userFolder slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_folder")
	}

	if len(userFolderAfterDeleteHooks) != 0 {
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
func (o *UserFolder) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserFolder(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserFolderSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserFolderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userFolderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_folder\".* FROM \"user_folder\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userFolderPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserFolderSlice")
	}

	*o = slice

	return nil
}

// UserFolderExists checks if the UserFolder row exists.
func UserFolderExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_folder\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_folder exists")
	}

	return exists, nil
}
