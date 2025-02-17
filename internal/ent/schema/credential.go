package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/PicoTools/pico/pkg/shared"
)

type Credential struct {
	ent.Schema
}

func (Credential) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "credential",
		},
	}
}

func (Credential) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Comment("credential ID").
			Unique(),
		field.String("username").
			MaxLen(shared.CredentialUsernameMaxLength).
			Optional().
			Comment("username of credential"),
		field.String("secret").
			MaxLen(shared.CredentialSecretMaxLength).
			Optional().
			Comment("password/hash/secret"),
		field.String("realm").
			MaxLen(shared.CredentialRealmMaxLength).
			Optional().
			Comment("realm of host (or zone where credentials are valid)"),
		field.String("host").
			MaxLen(shared.CredentialHostMaxLength).
			Optional().
			Comment("host from which creds has been extracted"),
		field.String("note").
			MaxLen(shared.CredentialNoteMaxLength).
			Optional().
			Comment("note of credential"),
		field.Uint32("color").
			Default(shared.DefaultObjectColor).
			Comment("color of entity"),
	}
}

func (Credential) Edges() []ent.Edge {
	return nil
}

func (Credential) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
