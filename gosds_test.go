package gosds_test

import (
	"fmt"
	"testing"

	"github.com/adrienaury/gosds"
	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	t.Parallel()

	root := gosds.NewObject(nil)

	root.SetValueForKey("name", "adrien")

	fmt.Println(root.NodeForKey("name").Parent())
}

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

	fmt.Println(builder.Finalize().MustObject().PrimitiveObject())
}

// func TestSonicBuilder(t *testing.T) {
// 	t.Parallel()

// 	jstr := `{"age":42,"name":"John","surname":"Doe","address":{"town":"Purple Town"},"mail":"john.doe@domain.fr"}`
// 	builder := gosds.NewBuilderSonic()

// 	assert.NoError(t, ast.Preorder(jstr, builder, &ast.VisitorOptions{OnlyNumber: true}))

// 	fmt.Println(builder.Finalize().MustObject().PrimitiveObject())
// }
