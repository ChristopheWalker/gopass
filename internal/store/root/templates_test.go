package root

import (
	"context"
	"testing"

	"github.com/fatih/color"
	"github.com/gopasspw/gopass/pkg/ctxutil"
	"github.com/gopasspw/gopass/tests/gptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplate(t *testing.T) {
	u := gptest.NewUnitTester(t)

	ctx := context.Background()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	ctx = ctxutil.WithHidden(ctx, true)
	color.NoColor = true

	rs, err := createRootStore(ctx, u)
	require.NoError(t, err)

	tt, err := rs.TemplateTree(ctx)
	require.NoError(t, err)
	assert.Equal(t, "gopass\n", tt.Format(0))

	assert.False(t, rs.HasTemplate(ctx, "foo"))
	_, err = rs.GetTemplate(ctx, "foo")
	assert.Error(t, err)
	assert.Error(t, rs.RemoveTemplate(ctx, "foo"))

	assert.NoError(t, rs.SetTemplate(ctx, "foo", []byte("foobar")))
	assert.True(t, rs.HasTemplate(ctx, "foo"))

	b, err := rs.GetTemplate(ctx, "foo")
	require.NoError(t, err)
	assert.Equal(t, "foobar", string(b))

	_, b, found := rs.LookupTemplate(ctx, "foo/bar")
	assert.True(t, found)
	assert.Equal(t, "foobar", string(b))
	assert.NoError(t, rs.RemoveTemplate(ctx, "foo"))
}
