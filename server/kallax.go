// Code generated by https://github.com/src-d/go-kallax. DO NOT EDIT.
// Please, do not touch the code below, and if you do, do it under your own
// risk. Take into account that all the code you write here will be completely
// erased from earth the next time you generate the kallax models.
package main

import (
	"database/sql"
	"fmt"
	"time"

	"gopkg.in/src-d/go-kallax.v1"
	"gopkg.in/src-d/go-kallax.v1/types"
)

var _ types.SQLType
var _ fmt.Formatter

type modelSaveFunc func(*kallax.Store) error

// NewProject returns a new instance of Project.
func NewProject() (record *Project) {
	return new(Project)
}

// GetID returns the primary key of the model.
func (r *Project) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.Id)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Project) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.Id), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "slug":
		return &r.Slug, nil
	case "internal_slug":
		return &r.InternalSlug, nil
	case "name":
		return types.Nullable(&r.Name), nil
	case "description":
		return types.Nullable(&r.Description), nil
	case "user_id":
		return types.Nullable(kallax.VirtualColumn("user_id", r, new(kallax.NumericID))), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Project: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Project) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Id, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "slug":
		return r.Slug, nil
	case "internal_slug":
		return r.InternalSlug, nil
	case "name":
		if r.Name == (*string)(nil) {
			return nil, nil
		}
		return r.Name, nil
	case "description":
		if r.Description == (*string)(nil) {
			return nil, nil
		}
		return r.Description, nil
	case "user_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Project: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Project) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "User":
		return new(User), nil

	}
	return nil, fmt.Errorf("kallax: model Project has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Project) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "User":
		val, ok := rel.(*User)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship User", rel)
		}
		r.User = *val
		return nil

	}
	return fmt.Errorf("kallax: model Project has no relationship %s", field)
}

// ProjectStore is the entity to access the records of the type Project
// in the database.
type ProjectStore struct {
	*kallax.Store
}

// NewProjectStore creates a new instance of ProjectStore
// using a SQL database.
func NewProjectStore(db *sql.DB) *ProjectStore {
	return &ProjectStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *ProjectStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *ProjectStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *ProjectStore) Debug() *ProjectStore {
	return &ProjectStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *ProjectStore) DebugWith(logger kallax.LoggerFunc) *ProjectStore {
	return &ProjectStore{s.Store.DebugWith(logger)}
}

// DisableCacher turns off prepared statements, which can be useful in some scenarios.
func (s *ProjectStore) DisableCacher() *ProjectStore {
	return &ProjectStore{s.Store.DisableCacher()}
}

func (s *ProjectStore) inverseRecords(record *Project) []modelSaveFunc {
	var result []modelSaveFunc

	if !record.User.GetID().IsEmpty() && !record.User.IsSaving() {
		record.AddVirtualColumn("user_id", record.User.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&UserStore{store}).Save(&record.User)
			return err
		})
	}

	return result
}

// Insert inserts a Project in the database. A non-persisted object is
// required for this operation.
func (s *ProjectStore) Insert(record *Project) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	if err := record.BeforeSave(); err != nil {
		return err
	}

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			if err := s.Insert(Schema.Project.BaseSchema, record); err != nil {
				return err
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.Project.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *ProjectStore) Update(record *Project, cols ...kallax.SchemaField) (updated int64, err error) {
	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	record.SetSaving(true)
	defer record.SetSaving(false)

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			updated, err = s.Update(Schema.Project.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

	return s.Store.Update(Schema.Project.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *ProjectStore) Save(record *Project) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *ProjectStore) Delete(record *Project) error {
	return s.Store.Delete(Schema.Project.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *ProjectStore) Find(q *ProjectQuery) (*ProjectResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewProjectResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *ProjectStore) MustFind(q *ProjectQuery) *ProjectResultSet {
	return NewProjectResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *ProjectStore) Count(q *ProjectQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *ProjectStore) MustCount(q *ProjectQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *ProjectStore) FindOne(q *ProjectQuery) (*Project, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// FindAll returns a list of all the rows returned by the given query.
func (s *ProjectStore) FindAll(q *ProjectQuery) ([]*Project, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *ProjectStore) MustFindOne(q *ProjectQuery) *Project {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Project with the data in the database and
// makes it writable.
func (s *ProjectStore) Reload(record *Project) error {
	return s.Store.Reload(Schema.Project.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *ProjectStore) Transaction(callback func(*ProjectStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&ProjectStore{store})
	})
}

// ProjectQuery is the object used to create queries for the Project
// entity.
type ProjectQuery struct {
	*kallax.BaseQuery
}

// NewProjectQuery returns a new instance of ProjectQuery.
func NewProjectQuery() *ProjectQuery {
	return &ProjectQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Project.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *ProjectQuery) Select(columns ...kallax.SchemaField) *ProjectQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *ProjectQuery) SelectNot(columns ...kallax.SchemaField) *ProjectQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *ProjectQuery) Copy() *ProjectQuery {
	return &ProjectQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *ProjectQuery) Order(cols ...kallax.ColumnOrder) *ProjectQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *ProjectQuery) BatchSize(size uint64) *ProjectQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *ProjectQuery) Limit(n uint64) *ProjectQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *ProjectQuery) Offset(n uint64) *ProjectQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *ProjectQuery) Where(cond kallax.Condition) *ProjectQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *ProjectQuery) WithUser() *ProjectQuery {
	q.AddRelation(Schema.User.BaseSchema, "User", kallax.OneToOne, nil)
	return q
}

// FindById adds a new filter to the query that will require that
// the Id property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *ProjectQuery) FindById(v ...int64) *ProjectQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Project.Id, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *ProjectQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *ProjectQuery {
	return q.Where(cond(Schema.Project.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *ProjectQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *ProjectQuery {
	return q.Where(cond(Schema.Project.UpdatedAt, v))
}

// FindBySlug adds a new filter to the query that will require that
// the Slug property is equal to the passed value.
func (q *ProjectQuery) FindBySlug(v string) *ProjectQuery {
	return q.Where(kallax.Eq(Schema.Project.Slug, v))
}

// FindByInternalSlug adds a new filter to the query that will require that
// the InternalSlug property is equal to the passed value.
func (q *ProjectQuery) FindByInternalSlug(v string) *ProjectQuery {
	return q.Where(kallax.Eq(Schema.Project.InternalSlug, v))
}

// FindByUser adds a new filter to the query that will require that
// the foreign key of User is equal to the passed value.
func (q *ProjectQuery) FindByUser(v int64) *ProjectQuery {
	return q.Where(kallax.Eq(Schema.Project.UserFK, v))
}

// ProjectResultSet is the set of results returned by a query to the
// database.
type ProjectResultSet struct {
	ResultSet kallax.ResultSet
	last      *Project
	lastErr   error
}

// NewProjectResultSet creates a new result set for rows of the type
// Project.
func NewProjectResultSet(rs kallax.ResultSet) *ProjectResultSet {
	return &ProjectResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *ProjectResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Project.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Project)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Project")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *ProjectResultSet) Get() (*Project, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *ProjectResultSet) ForEach(fn func(*Project) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *ProjectResultSet) All() ([]*Project, error) {
	var result []*Project
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *ProjectResultSet) One() (*Project, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *ProjectResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *ProjectResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewUser returns a new instance of User.
func NewUser() (record *User) {
	return new(User)
}

// GetID returns the primary key of the model.
func (r *User) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.Id)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *User) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.Id), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "slug":
		return &r.Slug, nil
	case "internal_slug":
		return &r.InternalSlug, nil
	case "name":
		return types.Nullable(&r.Name), nil
	case "email":
		return &r.Email, nil
	case "password":
		return types.Slice(&r.Password), nil
	case "profile_photo_slug":
		return types.Nullable(&r.ProfilePhotoSlug), nil
	case "forgot_password_token":
		return types.Slice(r.ForgotPasswordToken), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in User: %s", col)
	}
}

// Value returns the value of the given column.
func (r *User) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.Id, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "slug":
		return r.Slug, nil
	case "internal_slug":
		return r.InternalSlug, nil
	case "name":
		if r.Name == (*string)(nil) {
			return nil, nil
		}
		return r.Name, nil
	case "email":
		return r.Email, nil
	case "password":
		return types.Slice(r.Password), nil
	case "profile_photo_slug":
		if r.ProfilePhotoSlug == (*string)(nil) {
			return nil, nil
		}
		return r.ProfilePhotoSlug, nil
	case "forgot_password_token":
		if r.ForgotPasswordToken == (*byte)(nil) {
			return nil, nil
		}
		return types.Slice(r.ForgotPasswordToken), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in User: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *User) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Projects":
		return new(Project), nil

	}
	return nil, fmt.Errorf("kallax: model User has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *User) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Projects":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Projects = make([]*Project, len(records))
		for i, record := range records {
			rel, ok := record.(*Project)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Projects[i] = rel
		}
		return nil

	}
	return fmt.Errorf("kallax: model User has no relationship %s", field)
}

// UserStore is the entity to access the records of the type User
// in the database.
type UserStore struct {
	*kallax.Store
}

// NewUserStore creates a new instance of UserStore
// using a SQL database.
func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *UserStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *UserStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *UserStore) Debug() *UserStore {
	return &UserStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *UserStore) DebugWith(logger kallax.LoggerFunc) *UserStore {
	return &UserStore{s.Store.DebugWith(logger)}
}

// DisableCacher turns off prepared statements, which can be useful in some scenarios.
func (s *UserStore) DisableCacher() *UserStore {
	return &UserStore{s.Store.DisableCacher()}
}

func (s *UserStore) relationshipRecords(record *User) []modelSaveFunc {
	var result []modelSaveFunc

	for i := range record.Projects {
		r := record.Projects[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("user_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&ProjectStore{store}).Save(r)
				return err
			})
		}
	}

	return result
}

// Insert inserts a User in the database. A non-persisted object is
// required for this operation.
func (s *UserStore) Insert(record *User) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	if err := record.BeforeSave(); err != nil {
		return err
	}

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			if err := s.Insert(Schema.User.BaseSchema, record); err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.User.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *UserStore) Update(record *User, cols ...kallax.SchemaField) (updated int64, err error) {
	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	record.SetSaving(true)
	defer record.SetSaving(false)

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {
			updated, err = s.Update(Schema.User.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

	return s.Store.Update(Schema.User.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *UserStore) Save(record *User) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *UserStore) Delete(record *User) error {
	return s.Store.Delete(Schema.User.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *UserStore) Find(q *UserQuery) (*UserResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewUserResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *UserStore) MustFind(q *UserQuery) *UserResultSet {
	return NewUserResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *UserStore) Count(q *UserQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *UserStore) MustCount(q *UserQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *UserStore) FindOne(q *UserQuery) (*User, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// FindAll returns a list of all the rows returned by the given query.
func (s *UserStore) FindAll(q *UserQuery) ([]*User, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *UserStore) MustFindOne(q *UserQuery) *User {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the User with the data in the database and
// makes it writable.
func (s *UserStore) Reload(record *User) error {
	return s.Store.Reload(Schema.User.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *UserStore) Transaction(callback func(*UserStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&UserStore{store})
	})
}

// RemoveProjects removes the given items of the Projects field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Projects` is not empty. This method clears the
// the elements of Projects in a model, it does not retrieve them to know
// what relationships the model has.
func (s *UserStore) RemoveProjects(record *User, deleted ...*Project) error {
	var updated []*Project
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Projects
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Project.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Projects = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Project.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Project.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Projects {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Projects = updated
	return nil
}

// UserQuery is the object used to create queries for the User
// entity.
type UserQuery struct {
	*kallax.BaseQuery
}

// NewUserQuery returns a new instance of UserQuery.
func NewUserQuery() *UserQuery {
	return &UserQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.User.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *UserQuery) Select(columns ...kallax.SchemaField) *UserQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *UserQuery) SelectNot(columns ...kallax.SchemaField) *UserQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *UserQuery) Copy() *UserQuery {
	return &UserQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *UserQuery) Order(cols ...kallax.ColumnOrder) *UserQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *UserQuery) BatchSize(size uint64) *UserQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *UserQuery) Limit(n uint64) *UserQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *UserQuery) Offset(n uint64) *UserQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *UserQuery) Where(cond kallax.Condition) *UserQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *UserQuery) WithProjects(cond kallax.Condition) *UserQuery {
	q.AddRelation(Schema.Project.BaseSchema, "Projects", kallax.OneToMany, cond)
	return q
}

// FindById adds a new filter to the query that will require that
// the Id property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *UserQuery) FindById(v ...int64) *UserQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.User.Id, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *UserQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *UserQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.UpdatedAt, v))
}

// FindBySlug adds a new filter to the query that will require that
// the Slug property is equal to the passed value.
func (q *UserQuery) FindBySlug(v string) *UserQuery {
	return q.Where(kallax.Eq(Schema.User.Slug, v))
}

// FindByInternalSlug adds a new filter to the query that will require that
// the InternalSlug property is equal to the passed value.
func (q *UserQuery) FindByInternalSlug(v string) *UserQuery {
	return q.Where(kallax.Eq(Schema.User.InternalSlug, v))
}

// FindByEmail adds a new filter to the query that will require that
// the Email property is equal to the passed value.
func (q *UserQuery) FindByEmail(v string) *UserQuery {
	return q.Where(kallax.Eq(Schema.User.Email, v))
}

// FindByPassword adds a new filter to the query that will require that
// the Password property contains all the passed values; if no passed values,
// it will do nothing.
func (q *UserQuery) FindByPassword(v ...byte) *UserQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.ArrayContains(Schema.User.Password, values...))
}

// UserResultSet is the set of results returned by a query to the
// database.
type UserResultSet struct {
	ResultSet kallax.ResultSet
	last      *User
	lastErr   error
}

// NewUserResultSet creates a new result set for rows of the type
// User.
func NewUserResultSet(rs kallax.ResultSet) *UserResultSet {
	return &UserResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *UserResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.User.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*User)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *User")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *UserResultSet) Get() (*User, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *UserResultSet) ForEach(fn func(*User) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *UserResultSet) All() ([]*User, error) {
	var result []*User
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *UserResultSet) One() (*User, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *UserResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *UserResultSet) Close() error {
	return rs.ResultSet.Close()
}

type schema struct {
	Project *schemaProject
	User    *schemaUser
}

type schemaProject struct {
	*kallax.BaseSchema
	Id           kallax.SchemaField
	CreatedAt    kallax.SchemaField
	UpdatedAt    kallax.SchemaField
	Slug         kallax.SchemaField
	InternalSlug kallax.SchemaField
	Name         kallax.SchemaField
	Description  kallax.SchemaField
	UserFK       kallax.SchemaField
}

type schemaUser struct {
	*kallax.BaseSchema
	Id                  kallax.SchemaField
	CreatedAt           kallax.SchemaField
	UpdatedAt           kallax.SchemaField
	Slug                kallax.SchemaField
	InternalSlug        kallax.SchemaField
	Name                kallax.SchemaField
	Email               kallax.SchemaField
	Password            kallax.SchemaField
	ProfilePhotoSlug    kallax.SchemaField
	ForgotPasswordToken kallax.SchemaField
}

var Schema = &schema{
	Project: &schemaProject{
		BaseSchema: kallax.NewBaseSchema(
			"projects",
			"__project",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"User": kallax.NewForeignKey("user_id", true),
			},
			func() kallax.Record {
				return new(Project)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("slug"),
			kallax.NewSchemaField("internal_slug"),
			kallax.NewSchemaField("name"),
			kallax.NewSchemaField("description"),
			kallax.NewSchemaField("user_id"),
		),
		Id:           kallax.NewSchemaField("id"),
		CreatedAt:    kallax.NewSchemaField("created_at"),
		UpdatedAt:    kallax.NewSchemaField("updated_at"),
		Slug:         kallax.NewSchemaField("slug"),
		InternalSlug: kallax.NewSchemaField("internal_slug"),
		Name:         kallax.NewSchemaField("name"),
		Description:  kallax.NewSchemaField("description"),
		UserFK:       kallax.NewSchemaField("user_id"),
	},
	User: &schemaUser{
		BaseSchema: kallax.NewBaseSchema(
			"users",
			"__user",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Projects": kallax.NewForeignKey("user_id", false),
			},
			func() kallax.Record {
				return new(User)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("slug"),
			kallax.NewSchemaField("internal_slug"),
			kallax.NewSchemaField("name"),
			kallax.NewSchemaField("email"),
			kallax.NewSchemaField("password"),
			kallax.NewSchemaField("profile_photo_slug"),
			kallax.NewSchemaField("forgot_password_token"),
		),
		Id:                  kallax.NewSchemaField("id"),
		CreatedAt:           kallax.NewSchemaField("created_at"),
		UpdatedAt:           kallax.NewSchemaField("updated_at"),
		Slug:                kallax.NewSchemaField("slug"),
		InternalSlug:        kallax.NewSchemaField("internal_slug"),
		Name:                kallax.NewSchemaField("name"),
		Email:               kallax.NewSchemaField("email"),
		Password:            kallax.NewSchemaField("password"),
		ProfilePhotoSlug:    kallax.NewSchemaField("profile_photo_slug"),
		ForgotPasswordToken: kallax.NewSchemaField("forgot_password_token"),
	},
}
