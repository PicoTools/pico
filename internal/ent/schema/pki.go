package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Pki struct {
	ent.Schema
}

func (Pki) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "pki",
		},
	}
}

func (Pki) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values("ca", "listener", "operator", "management").
			Comment("type of certificate blob (ca, listener, operator)"),
		field.Bytes("key").
			Comment("certificate key"),
		field.Bytes("cert").
			Comment("certificate chain"),
	}
}

func (Pki) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Pki) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
