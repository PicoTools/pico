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

type Task struct {
	ent.Schema
}

func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "task",
		},
	}
}

func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique(),
		field.Int64("command_id").
			Comment("id of command"),
		field.Uint32("ant_id").
			Comment("id of ant"),
		field.Time("created_at").
			Default(time.Now).
			Comment("time when task created"),
		field.Time("pushed_at").
			Optional().
			Comment("time when task pushed to the ant"),
		field.Time("done_at").
			Optional().
			Comment("time when task results received"),
		field.Enum("status").
			GoType(shared.TaskStatus(0)).
			Comment("status of task"),
		field.Enum("cap").
			GoType(shared.Capability(0)).
			Comment("capability to execute"),
		field.Int("args_id").
			Comment("capability arguments"),
		field.Int("output_id").
			Optional().
			Comment("task output"),
		field.Bool("output_big").
			Optional().
			Comment("is output bigger than constant value"),
	}
}

func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("command", Command.Type).
			Ref("task").
			Field("command_id").
			Unique().
			Required(),
		edge.From("ant", Ant.Type).
			Ref("task").
			Field("ant_id").
			Unique().
			Required(),
		edge.From("blobber_args", Blobber.Type).
			Ref("task_args").
			Field("args_id").
			Unique().
			Required(),
		edge.From("blobber_output", Blobber.Type).
			Ref("task_output").
			Field("output_id").
			Unique(),
	}
}
