package schema

import (
	"fmt"
	"net"
	"time"

	"github.com/PicoTools/pico-shared/shared"
	"github.com/PicoTools/pico/internal/types"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Agent struct {
	ent.Schema
}

func (Agent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "agent",
		},
	}
}

func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Unique().
			Comment("agent ID"),
		field.Int64("listener_id").
			Comment("linked listener ID"),
		field.String("ext_ip").
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
			Comment("external IP address of agent"),
		field.String("int_ip").
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
			Comment("internal IP address of agent"),
		field.Enum("os").
			GoType(shared.AgentOs(0)).
			Comment("type of operating system"),
		field.String("os_meta").
			MaxLen(shared.AgentOsMetaMaxLength).
			Optional().
			Comment("metadata of operating system"),
		field.String("hostname").
			MaxLen(shared.AgentHostnameMaxLength).
			Optional().
			Comment("hostname of machine, on which agent deployed"),
		field.String("username").
			MaxLen(shared.AgentUsernameMaxLength).
			Optional().
			Comment("username of agent's process"),
		field.String("domain").
			MaxLen(shared.AgentDomainMaxLength).
			Optional().
			Comment("domain of machine, on which agent deployed"),
		field.Bool("privileged").
			Optional().
			Comment("is agent process is privileged"),
		field.String("process_name").
			MaxLen(shared.AgentProcessNameMaxLength).
			Optional().
			Comment("name of agent process"),
		// used int64 as sqlite unable handle uint64
		field.Int64("pid").
			Optional().
			Comment("process ID of agent"),
		field.Enum("arch").
			GoType(shared.AgentArch(0)).
			Comment("architecture of agent process"),
		field.Uint32("sleep").
			Comment("sleep value of agent"),
		field.Uint8("jitter").
			Comment("jitter value of sleep"),
		field.Time("first").
			Default(time.Now).
			Comment("first checkout timestamp"),
		field.Time("last").
			Default(time.Now).
			Comment("last activity of listener"),
		field.Uint32("caps").
			Comment("capabilities of agent"),
		field.String("note").
			MaxLen(shared.AgentNoteMaxLength).
			Optional().
			Comment("note of agent"),
		field.Uint32("color").
			Default(shared.DefaultObjectColor).
			Comment("color of entity"),
	}
}

func (Agent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("listener", Listener.Type).
			Ref("agent").
			Field("listener_id").
			Unique().
			Required(),
		edge.To("command", Command.Type),
		edge.To("task", Task.Type),
	}
}

func (Agent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
