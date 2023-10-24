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

	assert.NoError(t, builder.StartObject())           // {
	assert.NoError(t, builder.AddKey("name"))          // "name":
	assert.NoError(t, builder.AddValue("John"))        // "John",
	assert.NoError(t, builder.AddKey("surname"))       // "surname":
	assert.NoError(t, builder.AddValue("Doe"))         // "Doe",
	assert.NoError(t, builder.AddKey("age"))           // "age":
	assert.NoError(t, builder.AddValue(42))            // 42,
	assert.NoError(t, builder.AddKey("employed"))      // "employed":
	assert.NoError(t, builder.AddValue(true))          // true,
	assert.NoError(t, builder.AddKey("tags"))          // "tags":
	assert.NoError(t, builder.StartArray())            // [
	assert.NoError(t, builder.AddValue("test"))        // "test",
	assert.NoError(t, builder.AddValue("data"))        // "data",
	assert.NoError(t, builder.EndObjectOrArray())      // ],
	assert.NoError(t, builder.AddKey("address"))       // "address":
	assert.NoError(t, builder.StartObject())           // {
	assert.NoError(t, builder.AddKey("city"))          // "city":
	assert.NoError(t, builder.AddValue("Purple Town")) // "Purple Town"
	assert.NoError(t, builder.EndObjectOrArray())      // }
	assert.NoError(t, builder.EndObjectOrArray())      // }

	assert.NoError(t, builder.Build().MarshalWrite(os.Stdout))
}

func TestIntegrity(t *testing.T) {
	t.Parallel()

	jstr := `{"age":42,"name":"John","surname":"Doe","address":{"town":"Purple Town"},"tags":[1,true]}`

	root, err := gosds.Unmarshal(jstr)

	assert.NoError(t, err)

	assert.Same(t, root, root.Root())
	assert.Nil(t, root.Parent())
	assert.Equal(t, 0, root.Index())
	assert.Equal(t, "", root.Key())

	root.AsObject().NodeForKey("address").Set("123 East Street")

	assert.NoError(t, root.MarshalWrite(os.Stdout))
}
