package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/contrib/entgql" // ✅ Correct package for GraphQL annotations
)

// Todo schema definition.
type Todo struct {
	ent.Schema
}

// Fields of the Todo entity.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("title"),
		field.String("description").Optional(),
	}
}

// ✅ Correct Annotations function
func (Todo) Annotations() []ent.Annotation {
	return []ent.Annotation{
		entgql.QueryField(),                       // ✅ Correct
		entgql.Mutations(entgql.MutationCreate()), // ✅ Correct
	}
}
