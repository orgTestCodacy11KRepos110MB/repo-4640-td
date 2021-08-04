package tgerr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"

	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
)

func TestError(t *testing.T) {
	t.Run("FLOOD_WAIT_0", func(t *testing.T) {
		require.Equal(t, "rpc error code 420: FLOOD_WAIT (0)", tgerr.New(420, "FLOOD_WAIT_0").Error())
	})
	t.Run("FLOOD_WAIT", func(t *testing.T) {
		require.Equal(t, "rpc error code 420: FLOOD_WAIT", tgerr.New(420, "FLOOD_WAIT").Error())
	})
}

func TestErrorParse(t *testing.T) {
	t.Run("FLOOD_WAIT", func(t *testing.T) {
		require.Equal(t, &tgerr.Error{
			Code:     420,
			Message:  "FLOOD_WAIT_359",
			Type:     "FLOOD_WAIT",
			Argument: 359,
		}, tgerr.New(420, "FLOOD_WAIT_359"))
	})
	t.Run("FLOOD_WAIT_0", func(t *testing.T) {
		require.Equal(t, &tgerr.Error{
			Code:     420,
			Message:  "FLOOD_WAIT_0",
			Type:     "FLOOD_WAIT",
			Argument: 0,
		}, tgerr.New(420, "FLOOD_WAIT_0"))
	})
	t.Run("Middle", func(t *testing.T) {
		require.Equal(t, &tgerr.Error{
			Code:     169,
			Message:  "GO_1337_METERS_AWAY",
			Type:     "GO_METERS_AWAY",
			Argument: 1337,
		}, tgerr.New(169, "GO_1337_METERS_AWAY"))
	})
}

func TestHelpers(t *testing.T) {
	err := func() error {
		return tgerr.New(169, "GO_1337_METERS_AWAY")
	}()
	t.Run("Type", func(t *testing.T) {
		assert.True(t, tgerr.Is(err, "GO_METERS_AWAY"))
		assert.True(t, tgerr.Is(err, "FOO", "GO_METERS_AWAY"))
		assert.False(t, tgerr.Is(err, "NOPE"))
		t.Run("AsType", func(t *testing.T) {
			{
				rpcErr, ok := tgerr.AsType(err, "NOPE")
				require.False(t, ok)
				require.Nil(t, rpcErr)
			}
			{
				rpcErr, ok := tgerr.AsType(err, "GO_METERS_AWAY")
				require.True(t, ok)
				require.NotNil(t, rpcErr)
			}
		})
	})
	t.Run("Code", func(t *testing.T) {
		assert.True(t, tgerr.IsCode(err, 169))
		assert.True(t, tgerr.IsCode(err, 1, 169))
		assert.False(t, tgerr.IsCode(err, 168))
	})
	t.Run("Generated", func(t *testing.T) {
		// Ensure that code generation works for errors.
		err := func() error {
			rpcErr := &tgerr.Error{
				Type: tg.ErrAccessTokenExpired,
			}

			return xerrors.Errorf("perform operation: %w", rpcErr)
		}()
		require.True(t, tgerr.Is(err, tg.ErrAccessTokenExpired))
		require.True(t, tg.IsAccessTokenExpired(err))
	})
}
