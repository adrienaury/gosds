package gosds_test

import (
	"os"
	"testing"

	"github.com/adrienaury/gosds"
	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	t.Parallel()

	builder := gosds.NewBuilder()

	assert.NoError(t, builder.StartObject(""))
	assert.NoError(t, builder.AddString("John", "name"))
	assert.NoError(t, builder.AddString("Doe", "surname"))
	assert.NoError(t, builder.AddInt(42, "age"))
	assert.NoError(t, builder.AddBool(true, "employed"))
	assert.NoError(t, builder.StartArray("tags"))
	assert.NoError(t, builder.AddString("test", ""))
	assert.NoError(t, builder.AddString("data", ""))
	assert.NoError(t, builder.EndObjectOrArray())
	assert.NoError(t, builder.StartObject("address"))
	assert.NoError(t, builder.AddString("Purple Town", "city"))
	assert.NoError(t, builder.EndObjectOrArray())
	assert.NoError(t, builder.EndObjectOrArray())

	assert.NoError(t, builder.Finalize().MarshalWrite(os.Stdout))
}

// func TestSonicBuilder(t *testing.T) {
// 	t.Parallel()

// 	jstr := `{"age":42,"name":"John","surname":"Doe","address":{"town":"Purple Town"},"tags":[1,true]}`
// 	builder := gosds.NewBuilderSonic()

// 	assert.NoError(t, ast.Preorder(jstr, builder, &ast.VisitorOptions{OnlyNumber: true}))
// 	assert.NoError(t, gosds.EncodeJSON(builder.Finalize(), os.Stdout))
// }
