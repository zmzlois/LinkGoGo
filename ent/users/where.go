// Code generated by ent, DO NOT EDIT.

package users

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/zmzlois/LinkGoGo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldID, id))
}

// ExternalID applies equality check predicate on the "external_id" field. It's identical to ExternalIDEQ.
func ExternalID(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldExternalID, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldUsername, v))
}

// GlobalName applies equality check predicate on the "global_name" field. It's identical to GlobalNameEQ.
func GlobalName(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldGlobalName, v))
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldSlug, v))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldFirstName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldLastName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldEmail, v))
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldAvatar, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldDescription, v))
}

// AccessToken applies equality check predicate on the "access_token" field. It's identical to AccessTokenEQ.
func AccessToken(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldAccessToken, v))
}

// RefreshToken applies equality check predicate on the "refresh_token" field. It's identical to RefreshTokenEQ.
func RefreshToken(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldRefreshToken, v))
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v bool) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldDeleted, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldUpdatedAt, v))
}

// ExternalIDEQ applies the EQ predicate on the "external_id" field.
func ExternalIDEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldExternalID, v))
}

// ExternalIDNEQ applies the NEQ predicate on the "external_id" field.
func ExternalIDNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldExternalID, v))
}

// ExternalIDIn applies the In predicate on the "external_id" field.
func ExternalIDIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldExternalID, vs...))
}

// ExternalIDNotIn applies the NotIn predicate on the "external_id" field.
func ExternalIDNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldExternalID, vs...))
}

// ExternalIDGT applies the GT predicate on the "external_id" field.
func ExternalIDGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldExternalID, v))
}

// ExternalIDGTE applies the GTE predicate on the "external_id" field.
func ExternalIDGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldExternalID, v))
}

// ExternalIDLT applies the LT predicate on the "external_id" field.
func ExternalIDLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldExternalID, v))
}

// ExternalIDLTE applies the LTE predicate on the "external_id" field.
func ExternalIDLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldExternalID, v))
}

// ExternalIDContains applies the Contains predicate on the "external_id" field.
func ExternalIDContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldExternalID, v))
}

// ExternalIDHasPrefix applies the HasPrefix predicate on the "external_id" field.
func ExternalIDHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldExternalID, v))
}

// ExternalIDHasSuffix applies the HasSuffix predicate on the "external_id" field.
func ExternalIDHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldExternalID, v))
}

// ExternalIDEqualFold applies the EqualFold predicate on the "external_id" field.
func ExternalIDEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldExternalID, v))
}

// ExternalIDContainsFold applies the ContainsFold predicate on the "external_id" field.
func ExternalIDContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldExternalID, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldUsername, v))
}

// GlobalNameEQ applies the EQ predicate on the "global_name" field.
func GlobalNameEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldGlobalName, v))
}

// GlobalNameNEQ applies the NEQ predicate on the "global_name" field.
func GlobalNameNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldGlobalName, v))
}

// GlobalNameIn applies the In predicate on the "global_name" field.
func GlobalNameIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldGlobalName, vs...))
}

// GlobalNameNotIn applies the NotIn predicate on the "global_name" field.
func GlobalNameNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldGlobalName, vs...))
}

// GlobalNameGT applies the GT predicate on the "global_name" field.
func GlobalNameGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldGlobalName, v))
}

// GlobalNameGTE applies the GTE predicate on the "global_name" field.
func GlobalNameGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldGlobalName, v))
}

// GlobalNameLT applies the LT predicate on the "global_name" field.
func GlobalNameLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldGlobalName, v))
}

// GlobalNameLTE applies the LTE predicate on the "global_name" field.
func GlobalNameLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldGlobalName, v))
}

// GlobalNameContains applies the Contains predicate on the "global_name" field.
func GlobalNameContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldGlobalName, v))
}

// GlobalNameHasPrefix applies the HasPrefix predicate on the "global_name" field.
func GlobalNameHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldGlobalName, v))
}

// GlobalNameHasSuffix applies the HasSuffix predicate on the "global_name" field.
func GlobalNameHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldGlobalName, v))
}

// GlobalNameEqualFold applies the EqualFold predicate on the "global_name" field.
func GlobalNameEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldGlobalName, v))
}

// GlobalNameContainsFold applies the ContainsFold predicate on the "global_name" field.
func GlobalNameContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldGlobalName, v))
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldSlug, v))
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldSlug, v))
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldSlug, vs...))
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldSlug, vs...))
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldSlug, v))
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldSlug, v))
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldSlug, v))
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldSlug, v))
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldSlug, v))
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldSlug, v))
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldSlug, v))
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldSlug, v))
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldSlug, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldFirstName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldLastName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldEmail, v))
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldAvatar, v))
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldAvatar, v))
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldAvatar, vs...))
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldAvatar, vs...))
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldAvatar, v))
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldAvatar, v))
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldAvatar, v))
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldAvatar, v))
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldAvatar, v))
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldAvatar, v))
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldAvatar, v))
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldAvatar, v))
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldAvatar, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldDescription, v))
}

// AccessTokenEQ applies the EQ predicate on the "access_token" field.
func AccessTokenEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldAccessToken, v))
}

// AccessTokenNEQ applies the NEQ predicate on the "access_token" field.
func AccessTokenNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldAccessToken, v))
}

// AccessTokenIn applies the In predicate on the "access_token" field.
func AccessTokenIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldAccessToken, vs...))
}

// AccessTokenNotIn applies the NotIn predicate on the "access_token" field.
func AccessTokenNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldAccessToken, vs...))
}

// AccessTokenGT applies the GT predicate on the "access_token" field.
func AccessTokenGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldAccessToken, v))
}

// AccessTokenGTE applies the GTE predicate on the "access_token" field.
func AccessTokenGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldAccessToken, v))
}

// AccessTokenLT applies the LT predicate on the "access_token" field.
func AccessTokenLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldAccessToken, v))
}

// AccessTokenLTE applies the LTE predicate on the "access_token" field.
func AccessTokenLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldAccessToken, v))
}

// AccessTokenContains applies the Contains predicate on the "access_token" field.
func AccessTokenContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldAccessToken, v))
}

// AccessTokenHasPrefix applies the HasPrefix predicate on the "access_token" field.
func AccessTokenHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldAccessToken, v))
}

// AccessTokenHasSuffix applies the HasSuffix predicate on the "access_token" field.
func AccessTokenHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldAccessToken, v))
}

// AccessTokenEqualFold applies the EqualFold predicate on the "access_token" field.
func AccessTokenEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldAccessToken, v))
}

// AccessTokenContainsFold applies the ContainsFold predicate on the "access_token" field.
func AccessTokenContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldAccessToken, v))
}

// RefreshTokenEQ applies the EQ predicate on the "refresh_token" field.
func RefreshTokenEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldRefreshToken, v))
}

// RefreshTokenNEQ applies the NEQ predicate on the "refresh_token" field.
func RefreshTokenNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldRefreshToken, v))
}

// RefreshTokenIn applies the In predicate on the "refresh_token" field.
func RefreshTokenIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldRefreshToken, vs...))
}

// RefreshTokenNotIn applies the NotIn predicate on the "refresh_token" field.
func RefreshTokenNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldRefreshToken, vs...))
}

// RefreshTokenGT applies the GT predicate on the "refresh_token" field.
func RefreshTokenGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldRefreshToken, v))
}

// RefreshTokenGTE applies the GTE predicate on the "refresh_token" field.
func RefreshTokenGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldRefreshToken, v))
}

// RefreshTokenLT applies the LT predicate on the "refresh_token" field.
func RefreshTokenLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldRefreshToken, v))
}

// RefreshTokenLTE applies the LTE predicate on the "refresh_token" field.
func RefreshTokenLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldRefreshToken, v))
}

// RefreshTokenContains applies the Contains predicate on the "refresh_token" field.
func RefreshTokenContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldRefreshToken, v))
}

// RefreshTokenHasPrefix applies the HasPrefix predicate on the "refresh_token" field.
func RefreshTokenHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldRefreshToken, v))
}

// RefreshTokenHasSuffix applies the HasSuffix predicate on the "refresh_token" field.
func RefreshTokenHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldRefreshToken, v))
}

// RefreshTokenEqualFold applies the EqualFold predicate on the "refresh_token" field.
func RefreshTokenEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldRefreshToken, v))
}

// RefreshTokenContainsFold applies the ContainsFold predicate on the "refresh_token" field.
func RefreshTokenContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldRefreshToken, v))
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v bool) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldDeleted, v))
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v bool) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldDeleted, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUsersLinks applies the HasEdge predicate on the "users_links" edge.
func HasUsersLinks() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UsersLinksTable, UsersLinksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersLinksWith applies the HasEdge predicate on the "users_links" edge with a given conditions (other predicates).
func HasUsersLinksWith(preds ...predicate.Links) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newUsersLinksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Users) predicate.Users {
	return predicate.Users(sql.NotPredicates(p))
}
