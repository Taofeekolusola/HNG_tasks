package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"time"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
    return []ent.Field{
        field.String("model"),
        field.Time("registered_at").
            Default(time.Now), // Default value
    }
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
    return []ent.Edge{
        // Create an inverse-edge called "owner" of type `User`
        // and reference it to the "cars" edge (in User schema)
        // explicitly using the `Ref` method.
        edge.From("owner", User.Type). // ❌ Ensure "User" is imported
            Ref("cars").
            Unique(),
    }
}