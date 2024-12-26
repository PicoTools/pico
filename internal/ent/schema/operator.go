package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/PicoTools/pico-shared/shared"
)

type Operator struct {
	ent.Schema
}

func (Operator) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "operator",
		},
	}
}

func (Operator) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique().
			Comment("operator ID"),
		field.String("username").
			MaxLen(255).
			Comment("username of operator"),
		field.String("token").
			MinLen(32).
			MaxLen(32).
			Optional().
			Comment("access token for operator"),
		field.Uint32("color").
			Default(shared.DefaultObjectColor).
			Comment("color of entity"),
		field.Time("last").
			Default(func() time.Time {
				return time.Time{}
			}).
			Comment("last activity of operator"),
	}
}

func (Operator) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chat", Chat.Type),
		edge.To("command", Command.Type),
	}
}

func (Operator) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
