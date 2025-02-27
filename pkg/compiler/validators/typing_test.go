package validators

import (
	"testing"

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/stretchr/testify/assert"
)

func TestSimpleLiteralCasting(t *testing.T) {
	t.Run("BaseCase_Integer", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
			},
			&core.LiteralType{
				Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
			},
		)
		assert.True(t, castable, "Integers should be castable to other integers")
	})

	t.Run("IntegerToFloat", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
			},
			&core.LiteralType{
				Type: &core.LiteralType_Simple{Simple: core.SimpleType_FLOAT},
			},
		)
		assert.False(t, castable, "Integers should not be castable to floats")
	})

	t.Run("FloatToInteger", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_Simple{Simple: core.SimpleType_FLOAT},
			},
			&core.LiteralType{
				Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
			},
		)
		assert.False(t, castable, "Floats should not be castable to integers")
	})

	t.Run("VoidToInteger", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_Simple{Simple: core.SimpleType_NONE},
			},
			&core.LiteralType{
				Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
			},
		)
		assert.True(t, castable, "Floats are nullable")
	})

	t.Run("IgnoreMetadata", func(t *testing.T) {
		s := structpb.Struct{
			Fields: map[string]*structpb.Value{
				"a": {},
			},
		}
		castable := AreTypesCastable(
			&core.LiteralType{
				Type:     &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
				Metadata: &s,
			},
			&core.LiteralType{
				Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
			},
		)
		assert.True(t, castable, "Metadata should be ignored")
	})
}

func TestCollectionCasting(t *testing.T) {
	t.Run("BaseCase_SingleIntegerCollection", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_CollectionType{
					CollectionType: &core.LiteralType{
						Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
					},
				},
			},
			&core.LiteralType{
				Type: &core.LiteralType_CollectionType{
					CollectionType: &core.LiteralType{
						Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
					},
				},
			},
		)
		assert.True(t, castable, "[Integer] should be castable to [Integer].")
	})

	t.Run("SingleIntegerCollectionToSingleFloatCollection", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_CollectionType{
					CollectionType: &core.LiteralType{
						Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
					},
				},
			},
			&core.LiteralType{
				Type: &core.LiteralType_CollectionType{
					CollectionType: &core.LiteralType{
						Type: &core.LiteralType_Simple{Simple: core.SimpleType_FLOAT},
					},
				},
			},
		)
		assert.False(t, castable, "[Integer] should not be castable to [Float]")
	})

	t.Run("MismatchedNestLevels_Scalar", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_CollectionType{
					CollectionType: &core.LiteralType{
						Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
					},
				},
			},
			&core.LiteralType{
				Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
			},
		)
		assert.False(t, castable, "[Integer] should not be castable to Integer")
	})

	t.Run("MismatchedNestLevels_Collections", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_CollectionType{
					CollectionType: &core.LiteralType{
						Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
					},
				},
			},
			&core.LiteralType{
				Type: &core.LiteralType_CollectionType{
					CollectionType: &core.LiteralType{
						Type: &core.LiteralType_CollectionType{
							CollectionType: &core.LiteralType{
								Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
							},
						},
					},
				},
			},
		)
		assert.False(t, castable, "[Integer] should not be castable to [[Integer]]")
	})

	t.Run("Nullable_Collections", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_Simple{
					Simple: core.SimpleType_NONE,
				},
			},
			&core.LiteralType{
				Type: &core.LiteralType_CollectionType{
					CollectionType: &core.LiteralType{
						Type: &core.LiteralType_CollectionType{
							CollectionType: &core.LiteralType{
								Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
							},
						},
					},
				},
			},
		)
		assert.True(t, castable, "Collections are nullable")
	})
}

func TestMapCasting(t *testing.T) {
	t.Run("BaseCase_SingleIntegerMap", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_MapValueType{
					MapValueType: &core.LiteralType{
						Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
					},
				},
			},
			&core.LiteralType{
				Type: &core.LiteralType_MapValueType{
					MapValueType: &core.LiteralType{
						Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
					},
				},
			},
		)
		assert.True(t, castable, "{k: Integer} should be castable to {k: Integer}.")
	})

	t.Run("ScalarIntegerMapToScalarFloatMap", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_MapValueType{
					MapValueType: &core.LiteralType{
						Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
					},
				},
			},
			&core.LiteralType{
				Type: &core.LiteralType_MapValueType{
					MapValueType: &core.LiteralType{
						Type: &core.LiteralType_Simple{Simple: core.SimpleType_FLOAT},
					},
				},
			},
		)
		assert.False(t, castable, "{k: Integer} should not be castable to {k: Float}")
	})

	t.Run("ScalarStructToStruct", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_Simple{
					Simple: core.SimpleType_STRUCT,
				},
			},
			&core.LiteralType{
				Type: &core.LiteralType_Simple{
					Simple: core.SimpleType_STRUCT,
				},
			},
		)
		assert.True(t, castable, "castable from Struct to struct")
	})

	t.Run("MismatchedMapNestLevels_Scalar", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_MapValueType{
					MapValueType: &core.LiteralType{
						Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
					},
				},
			},
			&core.LiteralType{
				Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
			},
		)
		assert.False(t, castable, "{k: Integer} should not be castable to Integer")
	})

	t.Run("MismatchedMapNestLevels_Maps", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_MapValueType{
					MapValueType: &core.LiteralType{
						Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
					},
				},
			},
			&core.LiteralType{
				Type: &core.LiteralType_MapValueType{
					MapValueType: &core.LiteralType{
						Type: &core.LiteralType_MapValueType{
							MapValueType: &core.LiteralType{
								Type: &core.LiteralType_Simple{Simple: core.SimpleType_INTEGER},
							},
						},
					},
				},
			},
		)
		assert.False(t, castable, "{k: Integer} should not be castable to {k: {k: Integer}}")
	})
}

func TestSchemaCasting(t *testing.T) {
	genericSchema := &core.LiteralType{
		Type: &core.LiteralType_Schema{
			Schema: &core.SchemaType{
				Columns: []*core.SchemaType_SchemaColumn{},
			},
		},
	}
	subsetIntegerSchema := &core.LiteralType{
		Type: &core.LiteralType_Schema{
			Schema: &core.SchemaType{
				Columns: []*core.SchemaType_SchemaColumn{
					{
						Name: "a",
						Type: core.SchemaType_SchemaColumn_INTEGER,
					},
				},
			},
		},
	}
	supersetIntegerAndFloatSchema := &core.LiteralType{
		Type: &core.LiteralType_Schema{
			Schema: &core.SchemaType{
				Columns: []*core.SchemaType_SchemaColumn{
					{
						Name: "a",
						Type: core.SchemaType_SchemaColumn_INTEGER,
					},
					{
						Name: "b",
						Type: core.SchemaType_SchemaColumn_FLOAT,
					},
				},
			},
		},
	}
	mismatchedSubsetSchema := &core.LiteralType{
		Type: &core.LiteralType_Schema{
			Schema: &core.SchemaType{
				Columns: []*core.SchemaType_SchemaColumn{
					{
						Name: "a",
						Type: core.SchemaType_SchemaColumn_FLOAT,
					},
				},
			},
		},
	}

	t.Run("BaseCase_GenericSchema", func(t *testing.T) {
		castable := AreTypesCastable(genericSchema, genericSchema)
		assert.True(t, castable, "Schema() should be castable to Schema()")
	})

	t.Run("GenericSchemaToNonGeneric", func(t *testing.T) {
		castable := AreTypesCastable(genericSchema, subsetIntegerSchema)
		assert.False(t, castable, "Schema() should not be castable to Schema(a=Integer)")
	})

	t.Run("NonGenericSchemaToGeneric", func(t *testing.T) {
		castable := AreTypesCastable(subsetIntegerSchema, genericSchema)
		assert.True(t, castable, "Schema(a=Integer) should be castable to Schema()")
	})

	t.Run("SupersetToSubsetTypedSchema", func(t *testing.T) {
		castable := AreTypesCastable(supersetIntegerAndFloatSchema, subsetIntegerSchema)
		assert.True(t, castable, "Schema(a=Integer, b=Float) should be castable to Schema(a=Integer)")
	})

	t.Run("SubsetToSupersetSchema", func(t *testing.T) {
		castable := AreTypesCastable(subsetIntegerSchema, supersetIntegerAndFloatSchema)
		assert.False(t, castable, "Schema(a=Integer) should not be castable to Schema(a=Integer, b=Float)")
	})

	t.Run("MismatchedColumns", func(t *testing.T) {
		castable := AreTypesCastable(subsetIntegerSchema, mismatchedSubsetSchema)
		assert.False(t, castable, "Schema(a=Integer) should not be castable to Schema(a=Float)")
	})

	t.Run("MismatchedColumnsFlipped", func(t *testing.T) {
		castable := AreTypesCastable(mismatchedSubsetSchema, subsetIntegerSchema)
		assert.False(t, castable, "Schema(a=Float) should not be castable to Schema(a=Integer)")
	})

	t.Run("SchemasAreNullable", func(t *testing.T) {
		castable := AreTypesCastable(
			&core.LiteralType{
				Type: &core.LiteralType_Simple{
					Simple: core.SimpleType_NONE,
				},
			},
			subsetIntegerSchema)
		assert.True(t, castable, "Schemas are nullable")
	})
}
