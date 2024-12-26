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

type Chat struct {
	ent.Schema
}

func (Chat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "chat",
		},
	}
}

func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Comment("when message created"),
		field.Int64("author_id").
			Optional().
			Comment("creator of message"),
		field.String("message").
			MinLen(shared.ChatMessageMinLength).
			MaxLen(shared.ChatMessageMaxLength).
			Comment("message itself"),
		field.Bool("is_server").
			Default(false).
			Comment("is message created by server"),
	}
}

func (Chat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("operator", Operator.Type).
			Ref("chat").
			Field("author_id").
			Unique(),
	}
}
