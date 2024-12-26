package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Blobber struct {
	ent.Schema
}

func (Blobber) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "blobber",
		},
	}
}

func (Blobber) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("hash").
			Unique().
			Comment("non-cryptographic hash of blob"),
		field.Bytes("blob").
			Comment("blob to store"),
		field.Int("size").
			Comment("real size of blob"),
	}
}

func (Blobber) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("task_args", Task.Type),
		edge.To("task_output", Task.Type),
	}
}

func (Blobber) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
