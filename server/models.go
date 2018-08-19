package main
import (
	"time"
	"reflect"
	"github.com/blainehansen/goqu"
	"github.com/iancoleman/strcase"
)

type usersIdColumn struct {
	column
}
func (c *usersIdColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *usersIdColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *usersIdColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *usersIdColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *usersIdColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *usersIdColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *usersDateCreatedColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *usersDateCreatedColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *usersDateCreatedColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *usersDateCreatedColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *usersDateUpdatedColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *usersDateUpdatedColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *usersDateUpdatedColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *usersDateUpdatedColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *usersSlugColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *usersSlugColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *usersSlugColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *usersSlugColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *usersUrlSlugColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *usersUrlSlugColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *usersUrlSlugColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *usersUrlSlugColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *usersNameColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *usersNameColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *usersNameColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *usersNameColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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

type usersBioColumn struct {
	column
}
func (c *usersBioColumn) Set(val string) SetExpression {
	return SetExpression{ Name: "bio", Value: val }
}
func (c *usersBioColumn) Clear() SetExpression {
	return SetExpression{ Name: "bio", Value: nil }
}
func (c *usersBioColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *usersBioColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *usersBioColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *usersBioColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *usersBioColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *usersBioColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
}
func (c *usersBioColumn) Eq(val string) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersBioColumn) Neq(val string) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *usersBioColumn) Gt(val string) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *usersBioColumn) Gte(val string) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *usersBioColumn) Lt(val string) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *usersBioColumn) Lte(val string) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *usersBioColumn) In(val []string) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *usersBioColumn) NotIn(val []string) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}
func (c *usersBioColumn) Like(val string) goqu.BooleanExpression {
	return c.column.i.Like(val)
}
func (c *usersBioColumn) NotLike(val string) goqu.BooleanExpression {
	return c.column.i.NotLike(val)
}
func (c *usersBioColumn) ILike(val string) goqu.BooleanExpression {
	return c.column.i.ILike(val)
}
func (c *usersBioColumn) NotILike(val string) goqu.BooleanExpression {
	return c.column.i.NotILike(val)
}

type usersLocationColumn struct {
	column
}
func (c *usersLocationColumn) Set(val string) SetExpression {
	return SetExpression{ Name: "location", Value: val }
}
func (c *usersLocationColumn) Clear() SetExpression {
	return SetExpression{ Name: "location", Value: nil }
}
func (c *usersLocationColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *usersLocationColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *usersLocationColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *usersLocationColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *usersLocationColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *usersLocationColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
}
func (c *usersLocationColumn) Eq(val string) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersLocationColumn) Neq(val string) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *usersLocationColumn) Gt(val string) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *usersLocationColumn) Gte(val string) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *usersLocationColumn) Lt(val string) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *usersLocationColumn) Lte(val string) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *usersLocationColumn) In(val []string) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *usersLocationColumn) NotIn(val []string) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}
func (c *usersLocationColumn) Like(val string) goqu.BooleanExpression {
	return c.column.i.Like(val)
}
func (c *usersLocationColumn) NotLike(val string) goqu.BooleanExpression {
	return c.column.i.NotLike(val)
}
func (c *usersLocationColumn) ILike(val string) goqu.BooleanExpression {
	return c.column.i.ILike(val)
}
func (c *usersLocationColumn) NotILike(val string) goqu.BooleanExpression {
	return c.column.i.NotILike(val)
}

type usersLinksColumn struct {
	column
}
func (c *usersLinksColumn) Set(val string) SetExpression {
	return SetExpression{ Name: "links", Value: val }
}
func (c *usersLinksColumn) Clear() SetExpression {
	return SetExpression{ Name: "links", Value: nil }
}
func (c *usersLinksColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *usersLinksColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *usersLinksColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *usersLinksColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *usersLinksColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *usersLinksColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
}
func (c *usersLinksColumn) Eq(val string) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersLinksColumn) Neq(val string) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}
func (c *usersLinksColumn) Gt(val string) goqu.BooleanExpression {
	return c.column.i.Gt(val)
}
func (c *usersLinksColumn) Gte(val string) goqu.BooleanExpression {
	return c.column.i.Gte(val)
}
func (c *usersLinksColumn) Lt(val string) goqu.BooleanExpression {
	return c.column.i.Lt(val)
}
func (c *usersLinksColumn) Lte(val string) goqu.BooleanExpression {
	return c.column.i.Lte(val)
}
func (c *usersLinksColumn) In(val []string) goqu.BooleanExpression {
	return c.column.i.In(val)
}
func (c *usersLinksColumn) NotIn(val []string) goqu.BooleanExpression {
	return c.column.i.NotIn(val)
}
func (c *usersLinksColumn) Like(val string) goqu.BooleanExpression {
	return c.column.i.Like(val)
}
func (c *usersLinksColumn) NotLike(val string) goqu.BooleanExpression {
	return c.column.i.NotLike(val)
}
func (c *usersLinksColumn) ILike(val string) goqu.BooleanExpression {
	return c.column.i.ILike(val)
}
func (c *usersLinksColumn) NotILike(val string) goqu.BooleanExpression {
	return c.column.i.NotILike(val)
}

type usersEmailColumn struct {
	column
}
func (c *usersEmailColumn) Set(val string) SetExpression {
	return SetExpression{ Name: "email", Value: val }
}
func (c *usersEmailColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *usersEmailColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *usersEmailColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *usersEmailColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *usersPasswordColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *usersPasswordColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *usersPasswordColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *usersPasswordColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *usersProfilePhotoSlugColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *usersProfilePhotoSlugColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *usersProfilePhotoSlugColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *usersProfilePhotoSlugColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *usersForgotPasswordTokenColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *usersForgotPasswordTokenColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *usersForgotPasswordTokenColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *usersForgotPasswordTokenColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
}
func (c *usersForgotPasswordTokenColumn) Eq(val []byte) goqu.BooleanExpression {
	return c.column.i.Eq(val)
}
func (c *usersForgotPasswordTokenColumn) Neq(val []byte) goqu.BooleanExpression {
	return c.column.i.Neq(val)
}

type usersGeneralSearchVectorColumn struct {
	column
}
func (c *usersGeneralSearchVectorColumn) Search(val string) goqu.LiteralExpression {
	return goqu.L(`users.general_search_vector @@ to_tsquery('pg_catalog.english', ?)`, val)
}


type usersDataset struct {
	*goqu.Dataset
}

type usersSchema struct {
	Table *goqu.Dataset
	Query *usersDataset
	Id usersIdColumn
	DateCreated usersDateCreatedColumn
	DateUpdated usersDateUpdatedColumn
	Slug usersSlugColumn
	UrlSlug usersUrlSlugColumn
	Name usersNameColumn
	Bio usersBioColumn
	Location usersLocationColumn
	Links usersLinksColumn
	Email usersEmailColumn
	Password usersPasswordColumn
	ProfilePhotoSlug usersProfilePhotoSlugColumn
	ForgotPasswordToken usersForgotPasswordTokenColumn
	GeneralSearchVector usersGeneralSearchVectorColumn
}

var Users = &usersSchema{
	Table: db.From("users"),
	Query: &usersDataset{ db.From("users") },
	Id: usersIdColumn{ column { i: goqu.I("users.id") } },
	DateCreated: usersDateCreatedColumn{ column { i: goqu.I("users.date_created") } },
	DateUpdated: usersDateUpdatedColumn{ column { i: goqu.I("users.date_updated") } },
	Slug: usersSlugColumn{ column { i: goqu.I("users.slug") } },
	UrlSlug: usersUrlSlugColumn{ column { i: goqu.I("users.url_slug") } },
	Name: usersNameColumn{ column { i: goqu.I("users.name") } },
	Bio: usersBioColumn{ column { i: goqu.I("users.bio") } },
	Location: usersLocationColumn{ column { i: goqu.I("users.location") } },
	Links: usersLinksColumn{ column { i: goqu.I("users.links") } },
	Email: usersEmailColumn{ column { i: goqu.I("users.email") } },
	Password: usersPasswordColumn{ column { i: goqu.I("users.password") } },
	ProfilePhotoSlug: usersProfilePhotoSlugColumn{ column { i: goqu.I("users.profile_photo_slug") } },
	ForgotPasswordToken: usersForgotPasswordTokenColumn{ column { i: goqu.I("users.forgot_password_token") } },
	GeneralSearchVector: usersGeneralSearchVectorColumn{ column { i: goqu.I("users.general_search_vector") } },
}

var usersKinds = map[string]NestedKind {
	"name": NestedKind { Outer: reflect.String, Inner: reflect.Invalid },
	"bio": NestedKind { Outer: reflect.String, Inner: reflect.Invalid },
	"location": NestedKind { Outer: reflect.String, Inner: reflect.Invalid },
	"links": NestedKind { Outer: reflect.String, Inner: reflect.Invalid },
}


func (d *usersDataset) Where(expressions ...goqu.Expression) *usersDataset {
	return &usersDataset{ d.Dataset.Where(expressions...) }
}

func (d *usersDataset) Select(columns ...DbColumn) *usersDataset {
	return &usersDataset{ d.Dataset.Select(makeColumns(columns)...) }
}

func (d *usersDataset) Returning(columns ...DbColumn) *usersDataset {
	return &usersDataset{ d.Dataset.Returning(makeColumns(columns)...) }
}

func (d *usersDataset) Update(expressions ...SetExpression) *goqu.CrudExec {
	return d.Dataset.Update(makeRecord(expressions))
}

func (d *usersDataset) Insert(expressions ...SetExpression) *goqu.CrudExec {
	return d.Dataset.Insert(makeRecord(expressions))
}

func (d *usersDataset) Patch(values map[string]interface{}) *patchExec {
	var realValues = make(map[string]interface{})
	for key, value := range values {
		realValues[strcase.ToSnake(key)] = value
	}

	p := patchExec{
		d.Dataset.Update(realValues),
		validatePatch(&realValues, &usersKinds),
		realValues,
	}

	return &p
}

type ServerUser struct {
	Id int64
	DateCreated time.Time
	DateUpdated time.Time
	Slug string
	UrlSlug string
	Name string
	Bio string
	Location string
	Links string
	Email string
	Password []byte
	ProfilePhotoSlug string
	ForgotPasswordToken []byte
}

type OwnerPatchUser struct {
	Name string
	Bio string
	Location string
	Links string
}

type OwnerReadUser struct {
	Slug string
	UrlSlug string
	Name string
	Bio string
	Location string
	Links string
	Email string
	ProfilePhotoSlug string
}

type PublicReadUser struct {
	UrlSlug string
	Name string
	Bio string
	Location string
	Links string
	ProfilePhotoSlug string
}

type projectsIdColumn struct {
	column
}
func (c *projectsIdColumn) IsNull() goqu.BooleanExpression {
	return c.column.i.IsNull()
}
func (c *projectsIdColumn) IsNotNull() goqu.BooleanExpression {
	return c.column.i.IsNotNull()
}
func (c *projectsIdColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *projectsIdColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *projectsIdColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *projectsIdColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *projectsDateCreatedColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *projectsDateCreatedColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *projectsDateCreatedColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *projectsDateCreatedColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *projectsDateUpdatedColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *projectsDateUpdatedColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *projectsDateUpdatedColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *projectsDateUpdatedColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *projectsSlugColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *projectsSlugColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *projectsSlugColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *projectsSlugColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *projectsUrlSlugColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *projectsUrlSlugColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *projectsUrlSlugColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *projectsUrlSlugColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *projectsNameColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *projectsNameColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *projectsNameColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *projectsNameColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *projectsDescriptionColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *projectsDescriptionColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *projectsDescriptionColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *projectsDescriptionColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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
func (c *projectsUserIdColumn) As(val string) goqu.AliasedExpression {
	return c.column.i.As(val)
}
func (c *projectsUserIdColumn) Asc() goqu.OrderedExpression {
	return c.column.i.Asc()
}
func (c *projectsUserIdColumn) Desc() goqu.OrderedExpression {
	return c.column.i.Desc()
}
func (c *projectsUserIdColumn) Distinct() goqu.SqlFunctionExpression {
	return c.column.i.Distinct()
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

type projectsGeneralSearchVectorColumn struct {
	column
}
func (c *projectsGeneralSearchVectorColumn) Search(val string) goqu.LiteralExpression {
	return goqu.L(`projects.general_search_vector @@ to_tsquery('pg_catalog.english', ?)`, val)
}


type projectsDataset struct {
	*goqu.Dataset
}

type projectsSchema struct {
	Table *goqu.Dataset
	Query *projectsDataset
	Id projectsIdColumn
	DateCreated projectsDateCreatedColumn
	DateUpdated projectsDateUpdatedColumn
	Slug projectsSlugColumn
	UrlSlug projectsUrlSlugColumn
	Name projectsNameColumn
	Description projectsDescriptionColumn
	UserId projectsUserIdColumn
	GeneralSearchVector projectsGeneralSearchVectorColumn
}

var Projects = &projectsSchema{
	Table: db.From("projects"),
	Query: &projectsDataset{ db.From("projects") },
	Id: projectsIdColumn{ column { i: goqu.I("projects.id") } },
	DateCreated: projectsDateCreatedColumn{ column { i: goqu.I("projects.date_created") } },
	DateUpdated: projectsDateUpdatedColumn{ column { i: goqu.I("projects.date_updated") } },
	Slug: projectsSlugColumn{ column { i: goqu.I("projects.slug") } },
	UrlSlug: projectsUrlSlugColumn{ column { i: goqu.I("projects.url_slug") } },
	Name: projectsNameColumn{ column { i: goqu.I("projects.name") } },
	Description: projectsDescriptionColumn{ column { i: goqu.I("projects.description") } },
	UserId: projectsUserIdColumn{ column { i: goqu.I("projects.user_id") } },
	GeneralSearchVector: projectsGeneralSearchVectorColumn{ column { i: goqu.I("projects.general_search_vector") } },
}

var projectsKinds = map[string]NestedKind {
	"name": NestedKind { Outer: reflect.String, Inner: reflect.Invalid },
	"description": NestedKind { Outer: reflect.String, Inner: reflect.Invalid },
}


func (d *projectsDataset) Where(expressions ...goqu.Expression) *projectsDataset {
	return &projectsDataset{ d.Dataset.Where(expressions...) }
}

func (d *projectsDataset) Select(columns ...DbColumn) *projectsDataset {
	return &projectsDataset{ d.Dataset.Select(makeColumns(columns)...) }
}

func (d *projectsDataset) Returning(columns ...DbColumn) *projectsDataset {
	return &projectsDataset{ d.Dataset.Returning(makeColumns(columns)...) }
}

func (d *projectsDataset) Update(expressions ...SetExpression) *goqu.CrudExec {
	return d.Dataset.Update(makeRecord(expressions))
}

func (d *projectsDataset) Insert(expressions ...SetExpression) *goqu.CrudExec {
	return d.Dataset.Insert(makeRecord(expressions))
}

func (d *projectsDataset) Patch(values map[string]interface{}) *patchExec {
	var realValues = make(map[string]interface{})
	for key, value := range values {
		realValues[strcase.ToSnake(key)] = value
	}

	p := patchExec{
		d.Dataset.Update(realValues),
		validatePatch(&realValues, &projectsKinds),
		realValues,
	}

	return &p
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