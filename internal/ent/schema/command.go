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

type Command struct {
	ent.Schema
}

func (Command) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "command",
		},
	}
}

func (Command) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique(),
		field.Uint32("ant_id").
			Comment("ant ID"),
		field.String("cmd").
			MinLen(shared.TaskGroupCmdMinLength).
			MaxLen(shared.TaskGroupCmdMaxLength).
			Comment("command with arguments"),
		field.Bool("visible").
			Comment("is group visible for other operators"),
		field.Int64("author_id").
			Comment("author of group"),
		field.Time("created_at").
			Default(time.Now).
			Comment("when group created"),
		field.Time("closed_at").
			Optional().
			Comment("when group closed"),
	}
}

func (Command) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ant", Ant.Type).
			Ref("command").
			Field("ant_id").
			Unique().
			Required(),
		edge.From("operator", Operator.Type).
			Ref("command").
			Field("author_id").
			Unique().
			Required(),
		edge.To("message", Message.Type),
		edge.To("task", Task.Type),
	}
}
