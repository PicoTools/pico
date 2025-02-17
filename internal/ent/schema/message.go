package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/PicoTools/pico/pkg/shared"
)

type Message struct {
	ent.Schema
}

func (Message) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "message",
		},
	}
}

func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("command_id").
			Comment("id of command"),
		field.Enum("type").
			GoType(shared.TaskMessage(0)).
			Comment("type of message"),
		field.String("message").
			MaxLen(shared.TaskGroupMessageMaxLength).
			Comment("message itself"),
		field.Time("created_at").
			Default(time.Now).
			Comment("when message created"),
	}
}

func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("command", Command.Type).
			Ref("message").
			Field("command_id").
			Unique().
			Required(),
	}
}
