package gosds_test

import (
	"os"
	"testing"

	"github.com/adrienaury/gosds"
	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	t.Parallel()

	builder := gosds.Builder{}

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
