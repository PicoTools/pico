// Code generated by ent, DO NOT EDIT.

package pki

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/PicoTools/pico/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Pki {
	return predicate.Pki(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Pki {
	return predicate.Pki(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Pki {
	return predicate.Pki(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Pki {
	return predicate.Pki(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Pki {
	return predicate.Pki(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Pki {
	return predicate.Pki(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Pki {
	return predicate.Pki(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Pki {
	return predicate.Pki(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Pki {
	return predicate.Pki(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldEQ(FieldDeletedAt, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldEQ(FieldKey, v))
}

// Cert applies equality check predicate on the "cert" field. It's identical to CertEQ.
func Cert(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldEQ(FieldCert, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Pki {
	return predicate.Pki(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Pki {
	return predicate.Pki(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Pki {
	return predicate.Pki(sql.FieldNotNull(FieldDeletedAt))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Pki {
	return predicate.Pki(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Pki {
	return predicate.Pki(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Pki {
	return predicate.Pki(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Pki {
	return predicate.Pki(sql.FieldNotIn(FieldType, vs...))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...[]byte) predicate.Pki {
	return predicate.Pki(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...[]byte) predicate.Pki {
	return predicate.Pki(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldLTE(FieldKey, v))
}

// CertEQ applies the EQ predicate on the "cert" field.
func CertEQ(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldEQ(FieldCert, v))
}

// CertNEQ applies the NEQ predicate on the "cert" field.
func CertNEQ(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldNEQ(FieldCert, v))
}

// CertIn applies the In predicate on the "cert" field.
func CertIn(vs ...[]byte) predicate.Pki {
	return predicate.Pki(sql.FieldIn(FieldCert, vs...))
}

// CertNotIn applies the NotIn predicate on the "cert" field.
func CertNotIn(vs ...[]byte) predicate.Pki {
	return predicate.Pki(sql.FieldNotIn(FieldCert, vs...))
}

// CertGT applies the GT predicate on the "cert" field.
func CertGT(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldGT(FieldCert, v))
}

// CertGTE applies the GTE predicate on the "cert" field.
func CertGTE(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldGTE(FieldCert, v))
}

// CertLT applies the LT predicate on the "cert" field.
func CertLT(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldLT(FieldCert, v))
}

// CertLTE applies the LTE predicate on the "cert" field.
func CertLTE(v []byte) predicate.Pki {
	return predicate.Pki(sql.FieldLTE(FieldCert, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Pki) predicate.Pki {
	return predicate.Pki(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Pki) predicate.Pki {
	return predicate.Pki(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Pki) predicate.Pki {
	return predicate.Pki(sql.NotPredicates(p))
}