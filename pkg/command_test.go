package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CommandCase struct {
	Commands []string
	Res      []byte
}

func TestCommand(t *testing.T) {
	store := NewStore()
	testCases := []CommandCase{
		{
			Commands: []string{"PING"},
			Res:      []byte(pongMsg),
		},
		{
			Commands: []string{"SET", "1", "1", "0"},
			Res:      []byte(okMsg),
		},
		{
			Commands: []string{"GET", "1"},
			Res:      []byte("1\n"),
		},
		{
			Commands: []string{"GET", "2"},
			Res:      []byte(nilMsg),
		},
		{
			Commands: []string{"DEL", "1"},
			Res:      []byte(okMsg),
		},
		{
			Commands: []string{"DEL", "2"},
			Res:      []byte(keyNotfoundMsg),
		},
		{
			Commands: []string{"XXX", "2"},
			Res:      []byte(commandNotfoundMsg),
		},
	}
	for _, tc := range testCases {
		res := HandleCommand(tc.Commands, store)
		assert.Equal(t, res, tc.Res)
	}
}
