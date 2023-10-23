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

	assert.NoError(t, builder.StartObject())
	assert.NoError(t, builder.AddKey("name"))
	assert.NoError(t, builder.AddValue("John"))
	assert.NoError(t, builder.AddKey("surname"))
	assert.NoError(t, builder.AddValue("Doe"))
	assert.NoError(t, builder.AddKey("age"))
	assert.NoError(t, builder.AddValue(42))
	assert.NoError(t, builder.AddKey("employed"))
	assert.NoError(t, builder.AddValue(true))
	assert.NoError(t, builder.AddKey("tags"))
	assert.NoError(t, builder.StartArray())
	assert.NoError(t, builder.AddValue("test"))
	assert.NoError(t, builder.AddValue("data"))
	assert.NoError(t, builder.EndObjectOrArray())
	assert.NoError(t, builder.AddKey("address"))
	assert.NoError(t, builder.StartObject())
	assert.NoError(t, builder.AddKey("city"))
	assert.NoError(t, builder.AddValue("Purple Town"))
	assert.NoError(t, builder.EndObjectOrArray())
	assert.NoError(t, builder.EndObjectOrArray())

	assert.NoError(t, builder.Build().MarshalWrite(os.Stdout))
}
