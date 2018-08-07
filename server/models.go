package main
import (
	"time"
	"github.com/blainehansen/goqu"
)

var UsersTable = db.From("users")

type usersIdColumn struct {
	column
}
func (c *usersIdColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *usersIdColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *usersIdColumn) Get(val int64) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersIdColumn) Eq(val int64) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersIdColumn) Neq(val int64) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *usersIdColumn) In(val []int64) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *usersIdColumn) NotIn(val []int64) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}

type usersDateCreatedColumn struct {
	column
}
func (c *usersDateCreatedColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *usersDateCreatedColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *usersDateCreatedColumn) Eq(val time.Time) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersDateCreatedColumn) Neq(val time.Time) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *usersDateCreatedColumn) Gt(val time.Time) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *usersDateCreatedColumn) Gte(val time.Time) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *usersDateCreatedColumn) Lt(val time.Time) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *usersDateCreatedColumn) Lte(val time.Time) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *usersDateCreatedColumn) Between(startVal time.Time, endVal time.Time) goqu.RangeExpression {
	return c.column.i.Between(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *usersDateCreatedColumn) NotBetween(startVal time.Time, endVal time.Time) goqu.RangeExpression {
	return c.column.i.NotBetween(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *usersDateCreatedColumn) In(val []time.Time) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *usersDateCreatedColumn) NotIn(val []time.Time) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}

type usersDateUpdatedColumn struct {
	column
}
func (c *usersDateUpdatedColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *usersDateUpdatedColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *usersDateUpdatedColumn) Eq(val time.Time) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersDateUpdatedColumn) Neq(val time.Time) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *usersDateUpdatedColumn) Gt(val time.Time) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *usersDateUpdatedColumn) Gte(val time.Time) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *usersDateUpdatedColumn) Lt(val time.Time) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *usersDateUpdatedColumn) Lte(val time.Time) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *usersDateUpdatedColumn) Between(startVal time.Time, endVal time.Time) goqu.RangeExpression {
	return c.column.i.Between(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *usersDateUpdatedColumn) NotBetween(startVal time.Time, endVal time.Time) goqu.RangeExpression {
	return c.column.i.NotBetween(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *usersDateUpdatedColumn) In(val []time.Time) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *usersDateUpdatedColumn) NotIn(val []time.Time) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}

type usersSlugColumn struct {
	column
}
func (c *usersSlugColumn) Set(val string) SetExpression {
	return SetExpression{ Name: "slug", Value: val }
}
func (c *usersSlugColumn) Eq(val string) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersSlugColumn) Neq(val string) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *usersSlugColumn) Gt(val string) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *usersSlugColumn) Gte(val string) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *usersSlugColumn) Lt(val string) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *usersSlugColumn) Lte(val string) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *usersSlugColumn) In(val []string) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *usersSlugColumn) NotIn(val []string) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}
func (c *usersSlugColumn) Like(val string) goqu.BooleanExpression {
	return c.column.i.Like(val)
}
func (c *usersSlugColumn) NotLike(val string) goqu.BooleanExpression {
	return c.column.i.NotLike(val)
}
func (c *usersSlugColumn) ILike(val string) goqu.BooleanExpression {
	return c.column.i.ILike(val)
}
func (c *usersSlugColumn) NotILike(val string) goqu.BooleanExpression {
	return c.column.i.NotILike(val)
}

type usersUrlSlugColumn struct {
	column
}
func (c *usersUrlSlugColumn) Set(val string) SetExpression {
	return SetExpression{ Name: "url_slug", Value: val }
}
func (c *usersUrlSlugColumn) Eq(val string) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersUrlSlugColumn) Neq(val string) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *usersUrlSlugColumn) Gt(val string) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *usersUrlSlugColumn) Gte(val string) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *usersUrlSlugColumn) Lt(val string) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *usersUrlSlugColumn) Lte(val string) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *usersUrlSlugColumn) In(val []string) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *usersUrlSlugColumn) NotIn(val []string) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}
func (c *usersUrlSlugColumn) Like(val string) goqu.BooleanExpression {
	return c.column.i.Like(val)
}
func (c *usersUrlSlugColumn) NotLike(val string) goqu.BooleanExpression {
	return c.column.i.NotLike(val)
}
func (c *usersUrlSlugColumn) ILike(val string) goqu.BooleanExpression {
	return c.column.i.ILike(val)
}
func (c *usersUrlSlugColumn) NotILike(val string) goqu.BooleanExpression {
	return c.column.i.NotILike(val)
}

type usersNameColumn struct {
	column
}
func (c *usersNameColumn) Set(val string) SetExpression {
	return SetExpression{ Name: "name", Value: val }
}
func (c *usersNameColumn) Clear() SetExpression {
	return SetExpression{ Name: "name", Value: nil }
}
func (c *usersNameColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *usersNameColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *usersNameColumn) Eq(val string) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersNameColumn) Neq(val string) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *usersNameColumn) Gt(val string) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *usersNameColumn) Gte(val string) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *usersNameColumn) Lt(val string) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *usersNameColumn) Lte(val string) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *usersNameColumn) In(val []string) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *usersNameColumn) NotIn(val []string) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}
func (c *usersNameColumn) Like(val string) goqu.BooleanExpression {
	return c.column.i.Like(val)
}
func (c *usersNameColumn) NotLike(val string) goqu.BooleanExpression {
	return c.column.i.NotLike(val)
}
func (c *usersNameColumn) ILike(val string) goqu.BooleanExpression {
	return c.column.i.ILike(val)
}
func (c *usersNameColumn) NotILike(val string) goqu.BooleanExpression {
	return c.column.i.NotILike(val)
}

type usersEmailColumn struct {
	column
}
func (c *usersEmailColumn) Set(val string) SetExpression {
	return SetExpression{ Name: "email", Value: val }
}
func (c *usersEmailColumn) Eq(val string) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersEmailColumn) Neq(val string) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *usersEmailColumn) Gt(val string) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *usersEmailColumn) Gte(val string) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *usersEmailColumn) Lt(val string) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *usersEmailColumn) Lte(val string) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *usersEmailColumn) In(val []string) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *usersEmailColumn) NotIn(val []string) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}
func (c *usersEmailColumn) Like(val string) goqu.BooleanExpression {
	return c.column.i.Like(val)
}
func (c *usersEmailColumn) NotLike(val string) goqu.BooleanExpression {
	return c.column.i.NotLike(val)
}
func (c *usersEmailColumn) ILike(val string) goqu.BooleanExpression {
	return c.column.i.ILike(val)
}
func (c *usersEmailColumn) NotILike(val string) goqu.BooleanExpression {
	return c.column.i.NotILike(val)
}

type usersPasswordColumn struct {
	column
}
func (c *usersPasswordColumn) Set(val []byte) SetExpression {
	return SetExpression{ Name: "password", Value: val }
}
func (c *usersPasswordColumn) Eq(val []byte) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersPasswordColumn) Neq(val []byte) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}

type usersProfilePhotoSlugColumn struct {
	column
}
func (c *usersProfilePhotoSlugColumn) Set(val string) SetExpression {
	return SetExpression{ Name: "profile_photo_slug", Value: val }
}
func (c *usersProfilePhotoSlugColumn) Clear() SetExpression {
	return SetExpression{ Name: "profile_photo_slug", Value: nil }
}
func (c *usersProfilePhotoSlugColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *usersProfilePhotoSlugColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *usersProfilePhotoSlugColumn) Eq(val string) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersProfilePhotoSlugColumn) Neq(val string) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *usersProfilePhotoSlugColumn) Gt(val string) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *usersProfilePhotoSlugColumn) Gte(val string) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *usersProfilePhotoSlugColumn) Lt(val string) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *usersProfilePhotoSlugColumn) Lte(val string) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *usersProfilePhotoSlugColumn) In(val []string) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *usersProfilePhotoSlugColumn) NotIn(val []string) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}
func (c *usersProfilePhotoSlugColumn) Like(val string) goqu.BooleanExpression {
	return c.column.i.Like(val)
}
func (c *usersProfilePhotoSlugColumn) NotLike(val string) goqu.BooleanExpression {
	return c.column.i.NotLike(val)
}
func (c *usersProfilePhotoSlugColumn) ILike(val string) goqu.BooleanExpression {
	return c.column.i.ILike(val)
}
func (c *usersProfilePhotoSlugColumn) NotILike(val string) goqu.BooleanExpression {
	return c.column.i.NotILike(val)
}

type usersForgotPasswordTokenColumn struct {
	column
}
func (c *usersForgotPasswordTokenColumn) Set(val []byte) SetExpression {
	return SetExpression{ Name: "forgot_password_token", Value: val }
}
func (c *usersForgotPasswordTokenColumn) Clear() SetExpression {
	return SetExpression{ Name: "forgot_password_token", Value: nil }
}
func (c *usersForgotPasswordTokenColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *usersForgotPasswordTokenColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *usersForgotPasswordTokenColumn) Eq(val []byte) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersForgotPasswordTokenColumn) Neq(val []byte) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}


type usersSchema struct {
	dataset *goqu.Dataset
	Id usersIdColumn
	DateCreated usersDateCreatedColumn
	DateUpdated usersDateUpdatedColumn
	Slug usersSlugColumn
	UrlSlug usersUrlSlugColumn
	Name usersNameColumn
	Email usersEmailColumn
	Password usersPasswordColumn
	ProfilePhotoSlug usersProfilePhotoSlugColumn
	ForgotPasswordToken usersForgotPasswordTokenColumn
}

var Users = &usersSchema{
	Id: usersIdColumn{ column { i: goqu.I("users.id") } },
	DateCreated: usersDateCreatedColumn{ column { i: goqu.I("users.date_created") } },
	DateUpdated: usersDateUpdatedColumn{ column { i: goqu.I("users.date_updated") } },
	Slug: usersSlugColumn{ column { i: goqu.I("users.slug") } },
	UrlSlug: usersUrlSlugColumn{ column { i: goqu.I("users.url_slug") } },
	Name: usersNameColumn{ column { i: goqu.I("users.name") } },
	Email: usersEmailColumn{ column { i: goqu.I("users.email") } },
	Password: usersPasswordColumn{ column { i: goqu.I("users.password") } },
	ProfilePhotoSlug: usersProfilePhotoSlugColumn{ column { i: goqu.I("users.profile_photo_slug") } },
	ForgotPasswordToken: usersForgotPasswordTokenColumn{ column { i: goqu.I("users.forgot_password_token") } },
}

var usersKinds = map[string]NestedKind {
	"slug": NestedKind { reflect.String, reflect.Invalid },
	"url_slug": NestedKind { reflect.String, reflect.Invalid },
	"name": NestedKind { reflect.String, reflect.Invalid },
	"email": NestedKind { reflect.String, reflect.Invalid },
	"password": NestedKind { reflect.Int8, reflect.Slice },
	"profile_photo_slug": NestedKind { reflect.String, reflect.Invalid },
	"forgot_password_token": NestedKind { reflect.Int8, reflect.Slice },
}

func (d *usersSchema) Where(expressions ...goqu.Expression) *usersSchema {
	return &usersSchema{ d.dataset.Where(expressions...) }
}

func (d *usersSchema) Select(columns ...DbColumn) *usersSchema {
	return &usersSchema{ d.dataset.Select(makeColumns(columns)...) }
}

func (d *usersSchema) Returning(columns ...DbColumn) *usersSchema {
	return &usersSchema{ d.dataset.Returning(makeColumns(columns)...) }
}

func (d *usersSchema) Update(expressions ...SetExpression) *goqu.CrudExec {
	return d.dataset.Update(makeRecord(expressions))
}

func (d *usersSchema) Insert(expressions ...SetExpression) *goqu.CrudExec {
	return d.dataset.Insert(makeRecord(expressions))
}

func (d *usersSchema) Patch(values map[string]interface{}) (*goqu.CrudExec, bool) {
	if !validatePatch(values, &usersKinds) {
		return nil, false
	}

	return d.dataset.Update(values), true
}

var ProjectsTable = db.From("projects")

type projectsIdColumn struct {
	column
}
func (c *projectsIdColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *projectsIdColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *projectsIdColumn) Get(val int64) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *projectsIdColumn) Eq(val int64) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *projectsIdColumn) Neq(val int64) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *projectsIdColumn) In(val []int64) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *projectsIdColumn) NotIn(val []int64) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}

type projectsDateCreatedColumn struct {
	column
}
func (c *projectsDateCreatedColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *projectsDateCreatedColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *projectsDateCreatedColumn) Eq(val time.Time) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *projectsDateCreatedColumn) Neq(val time.Time) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *projectsDateCreatedColumn) Gt(val time.Time) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *projectsDateCreatedColumn) Gte(val time.Time) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *projectsDateCreatedColumn) Lt(val time.Time) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *projectsDateCreatedColumn) Lte(val time.Time) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *projectsDateCreatedColumn) Between(startVal time.Time, endVal time.Time) goqu.RangeExpression {
	return c.column.i.Between(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *projectsDateCreatedColumn) NotBetween(startVal time.Time, endVal time.Time) goqu.RangeExpression {
	return c.column.i.NotBetween(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *projectsDateCreatedColumn) In(val []time.Time) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *projectsDateCreatedColumn) NotIn(val []time.Time) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}

type projectsDateUpdatedColumn struct {
	column
}
func (c *projectsDateUpdatedColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *projectsDateUpdatedColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *projectsDateUpdatedColumn) Eq(val time.Time) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *projectsDateUpdatedColumn) Neq(val time.Time) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *projectsDateUpdatedColumn) Gt(val time.Time) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *projectsDateUpdatedColumn) Gte(val time.Time) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *projectsDateUpdatedColumn) Lt(val time.Time) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *projectsDateUpdatedColumn) Lte(val time.Time) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *projectsDateUpdatedColumn) Between(startVal time.Time, endVal time.Time) goqu.RangeExpression {
	return c.column.i.Between(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *projectsDateUpdatedColumn) NotBetween(startVal time.Time, endVal time.Time) goqu.RangeExpression {
	return c.column.i.NotBetween(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *projectsDateUpdatedColumn) In(val []time.Time) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *projectsDateUpdatedColumn) NotIn(val []time.Time) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}

type projectsSlugColumn struct {
	column
}
func (c *projectsSlugColumn) Set(val string) SetExpression {
	return SetExpression{ Name: "slug", Value: val }
}
func (c *projectsSlugColumn) Eq(val string) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *projectsSlugColumn) Neq(val string) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *projectsSlugColumn) Gt(val string) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *projectsSlugColumn) Gte(val string) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *projectsSlugColumn) Lt(val string) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *projectsSlugColumn) Lte(val string) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *projectsSlugColumn) In(val []string) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *projectsSlugColumn) NotIn(val []string) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}
func (c *projectsSlugColumn) Like(val string) goqu.BooleanExpression {
	return c.column.i.Like(val)
}
func (c *projectsSlugColumn) NotLike(val string) goqu.BooleanExpression {
	return c.column.i.NotLike(val)
}
func (c *projectsSlugColumn) ILike(val string) goqu.BooleanExpression {
	return c.column.i.ILike(val)
}
func (c *projectsSlugColumn) NotILike(val string) goqu.BooleanExpression {
	return c.column.i.NotILike(val)
}

type projectsUrlSlugColumn struct {
	column
}
func (c *projectsUrlSlugColumn) Set(val string) SetExpression {
	return SetExpression{ Name: "url_slug", Value: val }
}
func (c *projectsUrlSlugColumn) Eq(val string) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *projectsUrlSlugColumn) Neq(val string) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *projectsUrlSlugColumn) Gt(val string) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *projectsUrlSlugColumn) Gte(val string) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *projectsUrlSlugColumn) Lt(val string) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *projectsUrlSlugColumn) Lte(val string) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *projectsUrlSlugColumn) In(val []string) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *projectsUrlSlugColumn) NotIn(val []string) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}
func (c *projectsUrlSlugColumn) Like(val string) goqu.BooleanExpression {
	return c.column.i.Like(val)
}
func (c *projectsUrlSlugColumn) NotLike(val string) goqu.BooleanExpression {
	return c.column.i.NotLike(val)
}
func (c *projectsUrlSlugColumn) ILike(val string) goqu.BooleanExpression {
	return c.column.i.ILike(val)
}
func (c *projectsUrlSlugColumn) NotILike(val string) goqu.BooleanExpression {
	return c.column.i.NotILike(val)
}

type projectsNameColumn struct {
	column
}
func (c *projectsNameColumn) Set(val string) SetExpression {
	return SetExpression{ Name: "name", Value: val }
}
func (c *projectsNameColumn) Clear() SetExpression {
	return SetExpression{ Name: "name", Value: nil }
}
func (c *projectsNameColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *projectsNameColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *projectsNameColumn) Eq(val string) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *projectsNameColumn) Neq(val string) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *projectsNameColumn) Gt(val string) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *projectsNameColumn) Gte(val string) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *projectsNameColumn) Lt(val string) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *projectsNameColumn) Lte(val string) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *projectsNameColumn) In(val []string) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *projectsNameColumn) NotIn(val []string) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}
func (c *projectsNameColumn) Like(val string) goqu.BooleanExpression {
	return c.column.i.Like(val)
}
func (c *projectsNameColumn) NotLike(val string) goqu.BooleanExpression {
	return c.column.i.NotLike(val)
}
func (c *projectsNameColumn) ILike(val string) goqu.BooleanExpression {
	return c.column.i.ILike(val)
}
func (c *projectsNameColumn) NotILike(val string) goqu.BooleanExpression {
	return c.column.i.NotILike(val)
}

type projectsDescriptionColumn struct {
	column
}
func (c *projectsDescriptionColumn) Set(val string) SetExpression {
	return SetExpression{ Name: "description", Value: val }
}
func (c *projectsDescriptionColumn) Clear() SetExpression {
	return SetExpression{ Name: "description", Value: nil }
}
func (c *projectsDescriptionColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *projectsDescriptionColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *projectsDescriptionColumn) Eq(val string) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *projectsDescriptionColumn) Neq(val string) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *projectsDescriptionColumn) Gt(val string) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *projectsDescriptionColumn) Gte(val string) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *projectsDescriptionColumn) Lt(val string) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *projectsDescriptionColumn) Lte(val string) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *projectsDescriptionColumn) In(val []string) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *projectsDescriptionColumn) NotIn(val []string) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}
func (c *projectsDescriptionColumn) Like(val string) goqu.BooleanExpression {
	return c.column.i.Like(val)
}
func (c *projectsDescriptionColumn) NotLike(val string) goqu.BooleanExpression {
	return c.column.i.NotLike(val)
}
func (c *projectsDescriptionColumn) ILike(val string) goqu.BooleanExpression {
	return c.column.i.ILike(val)
}
func (c *projectsDescriptionColumn) NotILike(val string) goqu.BooleanExpression {
	return c.column.i.NotILike(val)
}

type projectsUserIdColumn struct {
	column
}
func (c *projectsUserIdColumn) Set(val int64) SetExpression {
	return SetExpression{ Name: "user_id", Value: val }
}
func (c *projectsUserIdColumn) Eq(val int64) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *projectsUserIdColumn) Neq(val int64) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *projectsUserIdColumn) Gt(val int64) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *projectsUserIdColumn) Gte(val int64) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *projectsUserIdColumn) Lt(val int64) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *projectsUserIdColumn) Lte(val int64) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *projectsUserIdColumn) Between(startVal int64, endVal int64) goqu.RangeExpression {
	return c.column.i.Between(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *projectsUserIdColumn) NotBetween(startVal int64, endVal int64) goqu.RangeExpression {
	return c.column.i.NotBetween(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *projectsUserIdColumn) In(val []int64) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *projectsUserIdColumn) NotIn(val []int64) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}


type projectsSchema struct {
	dataset *goqu.Dataset
	Id projectsIdColumn
	DateCreated projectsDateCreatedColumn
	DateUpdated projectsDateUpdatedColumn
	Slug projectsSlugColumn
	UrlSlug projectsUrlSlugColumn
	Name projectsNameColumn
	Description projectsDescriptionColumn
	UserId projectsUserIdColumn
}

var Projects = &projectsSchema{
	Id: projectsIdColumn{ column { i: goqu.I("projects.id") } },
	DateCreated: projectsDateCreatedColumn{ column { i: goqu.I("projects.date_created") } },
	DateUpdated: projectsDateUpdatedColumn{ column { i: goqu.I("projects.date_updated") } },
	Slug: projectsSlugColumn{ column { i: goqu.I("projects.slug") } },
	UrlSlug: projectsUrlSlugColumn{ column { i: goqu.I("projects.url_slug") } },
	Name: projectsNameColumn{ column { i: goqu.I("projects.name") } },
	Description: projectsDescriptionColumn{ column { i: goqu.I("projects.description") } },
	UserId: projectsUserIdColumn{ column { i: goqu.I("projects.user_id") } },
}

var projectsKinds = map[string]NestedKind {
	"slug": NestedKind { reflect.String, reflect.Invalid },
	"url_slug": NestedKind { reflect.String, reflect.Invalid },
	"name": NestedKind { reflect.String, reflect.Invalid },
	"description": NestedKind { reflect.String, reflect.Invalid },
	"user_id": NestedKind { reflect.Int64, reflect.Invalid },
}

func (d *projectsSchema) Where(expressions ...goqu.Expression) *projectsSchema {
	return &projectsSchema{ d.dataset.Where(expressions...) }
}

func (d *projectsSchema) Select(columns ...DbColumn) *projectsSchema {
	return &projectsSchema{ d.dataset.Select(makeColumns(columns)...) }
}

func (d *projectsSchema) Returning(columns ...DbColumn) *projectsSchema {
	return &projectsSchema{ d.dataset.Returning(makeColumns(columns)...) }
}

func (d *projectsSchema) Update(expressions ...SetExpression) *goqu.CrudExec {
	return d.dataset.Update(makeRecord(expressions))
}

func (d *projectsSchema) Insert(expressions ...SetExpression) *goqu.CrudExec {
	return d.dataset.Insert(makeRecord(expressions))
}

func (d *projectsSchema) Patch(values map[string]interface{}) (*goqu.CrudExec, bool) {
	if !validatePatch(values, &projectsKinds) {
		return nil, false
	}

	return d.dataset.Update(values), true
}