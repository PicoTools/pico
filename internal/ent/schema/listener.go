package schema

import (
	"fmt"
	"net"
	"time"

	"github.com/PicoTools/pico/internal/types"
	"github.com/PicoTools/pico/pkg/shared"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Listener struct {
	ent.Schema
}

func (Listener) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "listener",
		},
	}
}

func (Listener) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Comment("listener ID").
			Unique(),
		field.String("token").
			MinLen(32).
			MaxLen(32).
			Unique().
			Optional().
			Comment("authentication token of listener"),
		field.String("name").
			MaxLen(shared.ListenerNameMaxLength).
			Optional().
			Comment("freehand name of listener"),
		field.String("ip").
			GoType(types.Inet{}).
			SchemaType(map[string]string{
				dialect.Postgres: "inet",
			}).
			Validate(func(s string) error {
				if net.ParseIP(s) == nil {
					return fmt.Errorf("invalid value of ip %q", s)
				}
				return nil
			}).
			Optional().
			Comment("bind ip address of listener"),
		field.Uint16("port").
			Min(1).
			Optional().
			Comment("bind port of listener"),
		field.Uint32("color").
			Default(shared.DefaultObjectColor).
			Comment("color of entity"),
		field.String("note").
			MaxLen(shared.ListenerNoteMaxLength).
			Optional().
			Comment("note of listener"),
		field.Time("last").
			Default(func() time.Time {
				return time.Time{}
			}).
			Comment("last activity of listener"),
	}
}

func (Listener) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("agent", Agent.Type),
	}
}

func (Listener) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
