package schema

import (
	"context"
	"fmt"
	"time"

	gen "github.com/PicoTools/pico/internal/ent"
	"github.com/PicoTools/pico/internal/ent/hook"
	"github.com/PicoTools/pico/internal/ent/intercept"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Comment("Time when entity was created").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			Comment("Time when entity was updated").
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Optional().
			Comment("Time when entity was soft-deleted"),
	}
}

func (s TimeMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(c context.Context, q intercept.Query) error {
			if i, _ := c.Value(softDeleteKey{}).(bool); i {
				// i == true -> enable select soft-deleted records
				return nil
			}
			s.P(q)
			return nil
		}),
	}
}

func (s TimeMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					if del, _ := ctx.Value(softDeleteKey{}).(bool); del {
						// del == true -> record permanent delete
						return next.Mutate(ctx, m)
					}
					mx, ok := m.(interface {
						SetOp(ent.Op)
						Client() *gen.Client
						SetDeletedAt(time.Time)
						WhereP(...func(*sql.Selector))
					})
					if !ok {
						return nil, fmt.Errorf("unexpected mutation type %T", m)
					}
					s.P(mx)
					mx.SetOp(ent.OpUpdate)
					mx.SetDeletedAt(time.Now())
					return mx.Client().Mutate(ctx, m)
				})
			},
			ent.OpDeleteOne|ent.OpDelete,
		),
	}
}

// P adds predicate for level of storing
func (s TimeMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		// [2] - "deleted_at" in array []ent.Field{}
		sql.FieldIsNull(s.Fields()[2].Descriptor().Name),
	)
}

type softDeleteKey struct{}

// SkipSoftDelete returns context for permanent record deletion
func SkipSoftDelete(c context.Context) context.Context {
	return context.WithValue(c, softDeleteKey{}, true)
}

// WithSoftDeleted returns contetx with soft-deleted 
func WithSoftDeleted(c context.Context) context.Context {
	return context.WithValue(c, softDeleteKey{}, true)
}
