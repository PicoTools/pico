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

type Ant struct {
	ent.Schema
}

func (Ant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "ant",
		},
	}
}

func (Ant) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Unique().
			Comment("ant ID"),
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
			Comment("external IP address of ant"),
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
			Comment("internal IP address of ant"),
		field.Enum("os").
			GoType(shared.AntOs(0)).
			Comment("type of operating system"),
		field.String("os_meta").
			MaxLen(shared.AntOsMetaMaxLength).
			Optional().
			Comment("metadata of operating system"),
		field.String("hostname").
			MaxLen(shared.AntHostnameMaxLength).
			Optional().
			Comment("hostname of machine, on which ant deployed"),
		field.String("username").
			MaxLen(shared.AntUsernameMaxLength).
			Optional().
			Comment("username of ant's process"),
		field.String("domain").
			MaxLen(shared.AntDomainMaxLength).
			Optional().
			Comment("domain of machine, on which ant deployed"),
		field.Bool("privileged").
			Optional().
			Comment("is ant process is privileged"),
		field.String("process_name").
			MaxLen(shared.AntProcessNameMaxLength).
			Optional().
			Comment("name of ant process"),
		// used int64 as sqlite unable handle uint64
		field.Int64("pid").
			Optional().
			Comment("process ID of ant"),
		field.Enum("arch").
			GoType(shared.AntArch(0)).
			Comment("architecture of ant process"),
		field.Uint32("sleep").
			Comment("sleep value of ant"),
		field.Uint8("jitter").
			Comment("jitter value of sleep"),
		field.Time("first").
			Default(time.Now).
			Comment("first checkout timestamp"),
		field.Time("last").
			Default(time.Now).
			Comment("last activity of listener"),
		field.Uint32("caps").
			Comment("capabilities of ant"),
		field.String("note").
			MaxLen(shared.AntNoteMaxLength).
			Optional().
			Comment("note of ant"),
		field.Uint32("color").
			Default(shared.DefaultObjectColor).
			Comment("color of entity"),
	}
}

func (Ant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("listener", Listener.Type).
			Ref("ant").
			Field("listener_id").
			Unique().
			Required(),
		edge.To("command", Command.Type),
		edge.To("task", Task.Type),
	}
}

func (Ant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
