package updates

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

func TestSequenceBox(t *testing.T) {
	var (
		state   int
		updates []update
	)

	box := newSequenceBox(sequenceConfig{
		InitialState: 3,
		Apply: func(ctx context.Context, s int, u []update) error {
			state = s
			updates = append(updates, u...)
			return nil
		},
		Logger: zaptest.NewLogger(t),
	})

	require.Nil(t, box.Handle(update{
		Value: 1,
		State: 2,
		Count: 1,
	}))
	require.Zero(t, state)
	require.Empty(t, updates)
	require.Empty(t, box.pending)

	require.Nil(t, box.Handle(update{
		Value: 1,
		State: 3,
		Count: 1,
	}))
	require.Zero(t, state)
	require.Empty(t, updates)
	require.Empty(t, box.pending)

	require.Nil(t, box.Handle(update{
		Value: 1,
		State: 4,
		Count: 1,
	}))
	require.Equal(t, 4, state)
	require.Equal(t, []update{{1, 4, 1, entities{}, nil}}, updates)
	require.Empty(t, box.pending)
	updates = nil

	require.Nil(t, box.Handle(update{
		Value: 1,
		State: 6,
		Count: 1,
	}))
	require.Equal(t, 4, state)
	require.Empty(t, updates)
	require.Equal(t, []update{{1, 6, 1, entities{}, nil}}, box.pending)

	require.Nil(t, box.Handle(update{
		Value: 2,
		State: 5,
		Count: 1,
	}))
	require.Equal(t, 6, state)
	require.Equal(t, []update{{2, 5, 1, entities{}, nil}, {1, 6, 1, entities{}, nil}}, updates)
	require.Empty(t, box.pending)
	updates = nil

	require.Nil(t, box.Handle(update{
		Value: 3,
		State: 8,
		Count: 1,
	}))
	require.Equal(t, 6, state)
	require.Empty(t, updates)
	require.Equal(t, []update{{3, 8, 1, entities{}, nil}}, box.pending)
	<-box.gapTimeout.C

	require.Equal(t, []gap{{6, 7}}, box.gaps.gaps)
	box.gaps.Clear()
	require.False(t, box.gaps.Has())
}

func TestSequenceBoxApplyPending(t *testing.T) {
	tests := []struct {
		InitialState int
		Pending      []update
		PendingAfter []update
		Applied      []update
	}{
		{
			InitialState: 5,
			Pending: []update{
				{1, 3, 1, entities{}, nil},
				{1, 4, 1, entities{}, nil},
				{1, 1, 1, entities{}, nil},
			},
			PendingAfter: []update{},
			Applied:      []update{},
		},
		{
			InitialState: 5,
			Pending: []update{
				{1, 3, 1, entities{}, nil},
				{1, 8, 1, entities{}, nil},
				{1, 7, 1, entities{}, nil},
				{1, 4, 1, entities{}, nil},
				{1, 1, 1, entities{}, nil},
			},
			PendingAfter: []update{
				{1, 7, 1, entities{}, nil},
				{1, 8, 1, entities{}, nil},
			},
			Applied: []update{},
		},
		{
			InitialState: 5,
			Pending: []update{
				{1, 8, 1, entities{}, nil},
				{1, 7, 1, entities{}, nil},
			},
			PendingAfter: []update{
				{1, 7, 1, entities{}, nil},
				{1, 8, 1, entities{}, nil},
			},
			Applied: []update{},
		},
		{
			InitialState: 5,
			Pending: []update{
				{1, 3, 1, entities{}, nil},
				{1, 6, 1, entities{}, nil},
				{1, 8, 1, entities{}, nil},
				{1, 4, 1, entities{}, nil},
				{1, 1, 1, entities{}, nil},
			},
			PendingAfter: []update{
				{1, 8, 1, entities{}, nil},
			},
			Applied: []update{
				{1, 6, 1, entities{}, nil},
			},
		},
	}

	for _, test := range tests {
		applied := make([]update, 0)
		box := newSequenceBox(sequenceConfig{
			InitialState: test.InitialState,
			Apply: func(_ context.Context, s int, u []update) error {
				applied = append(applied, u...)
				return nil
			},
			Logger: zaptest.NewLogger(t),
		})

		box.pending = test.Pending
		require.NoError(t, box.applyPending(context.TODO()))
		require.Equal(t, test.PendingAfter, box.pending)
		require.Equal(t, test.Applied, applied)
	}
}
