package bptree2

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		n := New()

		ss := []string{}

		for i := 0; i < 10000; i++ {
			ss = append(ss, uuid.NewString())
		}

		for _, s := range ss {
			n.Insert(s)
		}
		for _, s := range ss {
			ev := fmt.Sprintf("d-%s", s)
			av := n.Find(s)
			assert.Equal(t, ev, av)
		}
		assert.Equal(t, "", "")
	})

}
