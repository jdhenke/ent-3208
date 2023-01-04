package schema

import "entgo.io/ent"

// Field holds the schema definition for the Field entity.
type Field struct {
	ent.Schema
}

// Fields of the Field.
func (Field) Fields() []ent.Field {
	return nil
}

// Edges of the Field.
func (Field) Edges() []ent.Edge {
	return nil
}
