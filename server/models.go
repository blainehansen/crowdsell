package main
import (
	"time"
	"github.com/blainehansen/goqu"
)

type primaryColumn struct {
	c string
	i goqu.IdentifierExpression
}
func (c primaryColumn) Identifier() goqu.IdentifierExpression {
	return c.i
}
func (c *primaryColumn) Get(val int64) goqu.BooleanExpression {
	return c.i.Eq(val)
}
func (c *primaryColumn) Eq(val int64) goqu.BooleanExpression {
	return c.i.Eq(val)
}
func (c *primaryColumn) Neq(val int64) goqu.BooleanExpression {
	return c.i.Neq(val)
}
func (c *primaryColumn) In(val []int64) goqu.BooleanExpression {
	return c.i.In(val)
}
func (c *primaryColumn) NotIn(val []int64) goqu.BooleanExpression {
	return c.i.NotIn(val)
}
func (c *primaryColumn) As(val string) goqu.AliasedExpression {
	return c.i.As(val)
}
func (c *primaryColumn) Asc() goqu.OrderedExpression {
	return c.i.Asc()
}
func (c *primaryColumn) Desc() goqu.OrderedExpression {
	return c.i.Desc()
}
func (c *primaryColumn) Distinct() goqu.SqlFunctionExpression {
	return c.i.Distinct()
}

type textColumn struct {
	c string
	i goqu.IdentifierExpression
}
func (c textColumn) Identifier() goqu.IdentifierExpression {
	return c.i
}
func (c *textColumn) Set(val string) SetExpression {
	return SetExpression{ Name: c.c, Value: val }
}
func (c *textColumn) Eq(val string) goqu.BooleanExpression {
	return c.i.Eq(val)
}
func (c *textColumn) Neq(val string) goqu.BooleanExpression {
	return c.i.Neq(val)
}
func (c *textColumn) Gt(val string) goqu.BooleanExpression {
	return c.i.Gt(val)
}
func (c *textColumn) Gte(val string) goqu.BooleanExpression {
	return c.i.Gte(val)
}
func (c *textColumn) Lt(val string) goqu.BooleanExpression {
	return c.i.Lt(val)
}
func (c *textColumn) Lte(val string) goqu.BooleanExpression {
	return c.i.Lte(val)
}
func (c *textColumn) In(val []string) goqu.BooleanExpression {
	return c.i.In(val)
}
func (c *textColumn) NotIn(val []string) goqu.BooleanExpression {
	return c.i.NotIn(val)
}
func (c *textColumn) Like(val string) goqu.BooleanExpression {
	return c.i.Like(val)
}
func (c *textColumn) NotLike(val string) goqu.BooleanExpression {
	return c.i.NotLike(val)
}
func (c *textColumn) ILike(val string) goqu.BooleanExpression {
	return c.i.ILike(val)
}
func (c *textColumn) NotILike(val string) goqu.BooleanExpression {
	return c.i.NotILike(val)
}
func (c *textColumn) IsNull() goqu.BooleanExpression {
	return c.i.IsNull()
}
func (c *textColumn) IsNotNull() goqu.BooleanExpression {
	return c.i.IsNotNull()
}
func (c *textColumn) As(val string) goqu.AliasedExpression {
	return c.i.As(val)
}
func (c *textColumn) Asc() goqu.OrderedExpression {
	return c.i.Asc()
}
func (c *textColumn) Desc() goqu.OrderedExpression {
	return c.i.Desc()
}
func (c *textColumn) Distinct() goqu.SqlFunctionExpression {
	return c.i.Distinct()
}

type booleanColumn struct {
	c string
	i goqu.IdentifierExpression
}
func (c booleanColumn) Identifier() goqu.IdentifierExpression {
	return c.i
}
func (c *booleanColumn) Set(val bool) SetExpression {
	return SetExpression{ Name: c.c, Value: val }
}
func (c *booleanColumn) Is(val bool) goqu.BooleanExpression {
	return c.i.Is(val)
}
func (c *booleanColumn) True() goqu.BooleanExpression {
	return c.i.IsTrue()
}
func (c *booleanColumn) False() goqu.BooleanExpression {
	return c.i.IsFalse()
}
func (c *booleanColumn) IsNull() goqu.BooleanExpression {
	return c.i.IsNull()
}
func (c *booleanColumn) IsNotNull() goqu.BooleanExpression {
	return c.i.IsNotNull()
}
func (c *booleanColumn) As(val string) goqu.AliasedExpression {
	return c.i.As(val)
}
func (c *booleanColumn) Asc() goqu.OrderedExpression {
	return c.i.Asc()
}
func (c *booleanColumn) Desc() goqu.OrderedExpression {
	return c.i.Desc()
}
func (c *booleanColumn) Distinct() goqu.SqlFunctionExpression {
	return c.i.Distinct()
}

type byteaColumn struct {
	c string
	i goqu.IdentifierExpression
}
func (c byteaColumn) Identifier() goqu.IdentifierExpression {
	return c.i
}
func (c *byteaColumn) Set(val []byte) SetExpression {
	return SetExpression{ Name: c.c, Value: val }
}
func (c *byteaColumn) Eq(val []byte) goqu.BooleanExpression {
	return c.i.Eq(val)
}
func (c *byteaColumn) Neq(val []byte) goqu.BooleanExpression {
	return c.i.Neq(val)
}
func (c *byteaColumn) As(val string) goqu.AliasedExpression {
	return c.i.As(val)
}
func (c *byteaColumn) Asc() goqu.OrderedExpression {
	return c.i.Asc()
}
func (c *byteaColumn) Desc() goqu.OrderedExpression {
	return c.i.Desc()
}
func (c *byteaColumn) Distinct() goqu.SqlFunctionExpression {
	return c.i.Distinct()
}

type bigintColumn struct {
	c string
	i goqu.IdentifierExpression
}
func (c bigintColumn) Identifier() goqu.IdentifierExpression {
	return c.i
}
func (c *bigintColumn) Set(val int64) SetExpression {
	return SetExpression{ Name: c.c, Value: val }
}
func (c *bigintColumn) Eq(val int64) goqu.BooleanExpression {
	return c.i.Eq(val)
}
func (c *bigintColumn) Neq(val int64) goqu.BooleanExpression {
	return c.i.Neq(val)
}
func (c *bigintColumn) Gt(val int64) goqu.BooleanExpression {
	return c.i.Gt(val)
}
func (c *bigintColumn) Gte(val int64) goqu.BooleanExpression {
	return c.i.Gte(val)
}
func (c *bigintColumn) Lt(val int64) goqu.BooleanExpression {
	return c.i.Lt(val)
}
func (c *bigintColumn) Lte(val int64) goqu.BooleanExpression {
	return c.i.Lte(val)
}
func (c *bigintColumn) Between(startVal int64, endVal int64) goqu.RangeExpression {
	return c.i.Between(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *bigintColumn) NotBetween(startVal int64, endVal int64) goqu.RangeExpression {
	return c.i.NotBetween(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *bigintColumn) In(val []int64) goqu.BooleanExpression {
	return c.i.In(val)
}
func (c *bigintColumn) NotIn(val []int64) goqu.BooleanExpression {
	return c.i.NotIn(val)
}
func (c *bigintColumn) IsNull() goqu.BooleanExpression {
	return c.i.IsNull()
}
func (c *bigintColumn) IsNotNull() goqu.BooleanExpression {
	return c.i.IsNotNull()
}
func (c *bigintColumn) As(val string) goqu.AliasedExpression {
	return c.i.As(val)
}
func (c *bigintColumn) Asc() goqu.OrderedExpression {
	return c.i.Asc()
}
func (c *bigintColumn) Desc() goqu.OrderedExpression {
	return c.i.Desc()
}
func (c *bigintColumn) Distinct() goqu.SqlFunctionExpression {
	return c.i.Distinct()
}

type timestamptzColumn struct {
	c string
	i goqu.IdentifierExpression
}
func (c timestamptzColumn) Identifier() goqu.IdentifierExpression {
	return c.i
}
func (c *timestamptzColumn) Set(val time.Time) SetExpression {
	return SetExpression{ Name: c.c, Value: val }
}
func (c *timestamptzColumn) Eq(val time.Time) goqu.BooleanExpression {
	return c.i.Eq(val)
}
func (c *timestamptzColumn) Neq(val time.Time) goqu.BooleanExpression {
	return c.i.Neq(val)
}
func (c *timestamptzColumn) Gt(val time.Time) goqu.BooleanExpression {
	return c.i.Gt(val)
}
func (c *timestamptzColumn) Gte(val time.Time) goqu.BooleanExpression {
	return c.i.Gte(val)
}
func (c *timestamptzColumn) Lt(val time.Time) goqu.BooleanExpression {
	return c.i.Lt(val)
}
func (c *timestamptzColumn) Lte(val time.Time) goqu.BooleanExpression {
	return c.i.Lte(val)
}
func (c *timestamptzColumn) Between(startVal time.Time, endVal time.Time) goqu.RangeExpression {
	return c.i.Between(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *timestamptzColumn) NotBetween(startVal time.Time, endVal time.Time) goqu.RangeExpression {
	return c.i.NotBetween(goqu.RangeVal{ Start: startVal, End: endVal })
}
func (c *timestamptzColumn) In(val []time.Time) goqu.BooleanExpression {
	return c.i.In(val)
}
func (c *timestamptzColumn) NotIn(val []time.Time) goqu.BooleanExpression {
	return c.i.NotIn(val)
}
func (c *timestamptzColumn) IsNull() goqu.BooleanExpression {
	return c.i.IsNull()
}
func (c *timestamptzColumn) IsNotNull() goqu.BooleanExpression {
	return c.i.IsNotNull()
}
func (c *timestamptzColumn) As(val string) goqu.AliasedExpression {
	return c.i.As(val)
}
func (c *timestamptzColumn) Asc() goqu.OrderedExpression {
	return c.i.Asc()
}
func (c *timestamptzColumn) Desc() goqu.OrderedExpression {
	return c.i.Desc()
}
func (c *timestamptzColumn) Distinct() goqu.SqlFunctionExpression {
	return c.i.Distinct()
}


// var UsersTable = CreateTable("users")

type usersSchema struct {
	*goqu.Dataset
	Id primaryColumn
	DateCreated timestamptzColumn
	DateUpdated timestamptzColumn
	Slug textColumn
	UrlSlug textColumn
	Name textColumn
	Email textColumn
	Password byteaColumn
	ProfilePhotoSlug textColumn
	ForgotPasswordToken byteaColumn
}

var Users = &usersSchema{
	db.From("users"),
	Id: primaryColumn{ c: "id", i: goqu.I("id") },
	DateCreated: timestamptzColumn{ c: "date_created", i: goqu.I("date_created") },
	DateUpdated: timestamptzColumn{ c: "date_updated", i: goqu.I("date_updated") },
	Slug: textColumn{ c: "slug", i: goqu.I("slug") },
	UrlSlug: textColumn{ c: "url_slug", i: goqu.I("url_slug") },
	Name: textColumn{ c: "name", i: goqu.I("name") },
	Email: textColumn{ c: "email", i: goqu.I("email") },
	Password: byteaColumn{ c: "password", i: goqu.I("password") },
	ProfilePhotoSlug: textColumn{ c: "profile_photo_slug", i: goqu.I("profile_photo_slug") },
	ForgotPasswordToken: byteaColumn{ c: "forgot_password_token", i: goqu.I("forgot_password_token") },
}

type ServerUser struct {
	Id int64
	DateCreated time.Time
	DateUpdated time.Time
	Slug string
	UrlSlug string
	Name string
	Email string
	Password []byte
	ProfilePhotoSlug string
	ForgotPasswordToken []byte
}

type OwnerPatchUser struct {
	Name string
	Email string
}

type OwnerReadUser struct {
	Slug string
	UrlSlug string
	Name string
	Email string
	ProfilePhotoSlug string
}

type PublicReadUser struct {
	UrlSlug string
	Name string
	ProfilePhotoSlug string
}

var ProjectsTable = CreateTable("projects")

type projectsSchema struct {
	Id primaryColumn
	DateCreated timestamptzColumn
	DateUpdated timestamptzColumn
	Slug textColumn
	UrlSlug textColumn
	Name textColumn
	Description textColumn
	UserId bigintColumn
}

var Projects = &projectsSchema{
	Id: primaryColumn{ c: "id", i: goqu.I("id") },
	DateCreated: timestamptzColumn{ c: "date_created", i: goqu.I("date_created") },
	DateUpdated: timestamptzColumn{ c: "date_updated", i: goqu.I("date_updated") },
	Slug: textColumn{ c: "slug", i: goqu.I("slug") },
	UrlSlug: textColumn{ c: "url_slug", i: goqu.I("url_slug") },
	Name: textColumn{ c: "name", i: goqu.I("name") },
	Description: textColumn{ c: "description", i: goqu.I("description") },
	UserId: bigintColumn{ c: "user_id", i: goqu.I("user_id") },
}

type ServerProject struct {
	Id int64
	DateCreated time.Time
	DateUpdated time.Time
	Slug string
	UrlSlug string
	Name string
	Description string
	UserId int64
}

type OwnerPatchProject struct {
	Name string
	Description string
}

type OwnerReadProject struct {
	Slug string
	UrlSlug string
	Name string
	Description string
}

type PublicReadProject struct {
	UrlSlug string
	Name string
	Description string
}
