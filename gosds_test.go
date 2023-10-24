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

	assert.NoError(t, builder.StartObject())                      // {
	assert.NoError(t, builder.AddKeyValue("name", "John"))        // "name": "John",
	assert.NoError(t, builder.AddKeyValue("surname", "Doe"))      // "surname": "Doe",
	assert.NoError(t, builder.AddKeyValue("age", 42))             // "age": 42,
	assert.NoError(t, builder.AddKeyValue("employed", true))      // "employed": true,
	assert.NoError(t, builder.AddKey("tags"))                     // "tags":
	assert.NoError(t, builder.StartArray())                       // [
	assert.NoError(t, builder.AddValue("test"))                   // "test",
	assert.NoError(t, builder.AddValue("data"))                   // "data",
	assert.NoError(t, builder.EndObjectOrArray())                 // ],
	assert.NoError(t, builder.AddKey("address"))                  // "address":
	assert.NoError(t, builder.StartObject())                      // {
	assert.NoError(t, builder.AddKeyValue("city", "Purple Town")) //  "city": "Purple Town"
	assert.NoError(t, builder.EndObjectOrArray())                 // }
	assert.NoError(t, builder.EndObjectOrArray())                 // }

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
