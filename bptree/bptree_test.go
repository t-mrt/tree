package bptree

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestInsert(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		bp := NewBP()
		bp.Insert("1", "a")
		assert.Equal(t, [3]string{"1", "", ""}, bp.root.(*BPLeafNode).key)

		bp.Insert("3", "c")
		assert.Equal(t, [3]string{"1", "3", ""}, bp.root.(*BPLeafNode).key)

		bp.Insert("2", "b")
		assert.Equal(t, [4]string{"", "2", "", ""}, bp.root.(*BPNonLeafNode).key)
		assert.Equal(t, [3]string{"1", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"2", "3", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPLeafNode).key)

		bp.Insert("4", "d")
		assert.Equal(t, [4]string{"", "2", "3", ""}, bp.root.(*BPNonLeafNode).key)
		assert.Equal(t, [3]string{"1", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"2", "", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"3", "4", ""}, bp.root.(*BPNonLeafNode).children[2].(*BPLeafNode).key)
	})

	t.Run("1", func(t *testing.T) {
		bp := NewBP()
		bp.Insert("1", "a")
		assert.Equal(t, [3]string{"1", "", ""}, bp.root.(*BPLeafNode).key)

		bp.Insert("3", "c")
		assert.Equal(t, [3]string{"1", "3", ""}, bp.root.(*BPLeafNode).key)

		bp.Insert("5", "e")
		assert.Equal(t, [4]string{"", "3", "", ""}, bp.root.(*BPNonLeafNode).key)
		assert.Equal(t, [3]string{"1", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"3", "5", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPLeafNode).key)

		bp.Insert("4", "d")
		assert.Equal(t, [4]string{"", "3", "4", ""}, bp.root.(*BPNonLeafNode).key)

		assert.Equal(t, [3]string{"1", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"3", "", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"4", "5", ""}, bp.root.(*BPNonLeafNode).children[2].(*BPLeafNode).key)

		bp.Insert("6", "f")
		assert.Equal(t, [4]string{"", "4", "", ""}, bp.root.(*BPNonLeafNode).key)
		assert.Equal(t, [4]string{"", "3", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).key)
		assert.Equal(t, [4]string{"", "5", "", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).key)
		assert.Equal(t, [3]string{"1", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"3", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).children[1].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"4", "", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"5", "6", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[1].(*BPLeafNode).key)

		bp.Insert("7", "g")
		assert.Equal(t, [4]string{"", "4", "", ""}, bp.root.(*BPNonLeafNode).key)
		assert.Equal(t, [4]string{"", "3", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).key)
		assert.Equal(t, [4]string{"", "5", "6", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).key)
		assert.Equal(t, [3]string{"1", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"3", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).children[1].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"4", "", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"5", "", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[1].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"6", "7", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[2].(*BPLeafNode).key)
	})

	t.Run("3", func(t *testing.T) {
		bp := NewBP()

		ss := []string{
			"97bab201-75ea-4a8c-8aa5-e9b52fcadd97",
			"7ab22432-15c5-4007-bd78-79480792e876",
			"d07d58a3-571b-491a-99da-76d1cfb02b4f",
			"4ea9752f-c3c2-4ffa-a7bb-3e6afc53916d",
			"4f60874e-966f-4635-ad0f-2bd7bd5be78f",
			"367f94e4-896b-4cff-8b87-af2850604c30",
			"18312213-28ec-4633-8b63-ec77983fc172",
			"915af9e1-a59d-4710-af01-95db786161be",
		}

		for _, s := range ss {
			bp.Insert(s, "dummy")
		}

		assert.Equal(t, [4]string{"", "4f60874e-966f-4635-ad0f-2bd7bd5be78f", "", ""}, bp.root.(*BPNonLeafNode).key)

		assert.Equal(t, [4]string{"", "367f94e4-896b-4cff-8b87-af2850604c30", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).key)
		assert.Equal(t, [4]string{"", "7ab22432-15c5-4007-bd78-79480792e876", "97bab201-75ea-4a8c-8aa5-e9b52fcadd97", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).key)

		assert.Equal(t, [3]string{"18312213-28ec-4633-8b63-ec77983fc172", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"367f94e4-896b-4cff-8b87-af2850604c30", "4ea9752f-c3c2-4ffa-a7bb-3e6afc53916d", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).children[1].(*BPLeafNode).key)

		assert.Equal(t, [3]string{"4f60874e-966f-4635-ad0f-2bd7bd5be78f", "", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"7ab22432-15c5-4007-bd78-79480792e876", "915af9e1-a59d-4710-af01-95db786161be", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[1].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"97bab201-75ea-4a8c-8aa5-e9b52fcadd97", "d07d58a3-571b-491a-99da-76d1cfb02b4f", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[2].(*BPLeafNode).key)
	})

	t.Run("4", func(t *testing.T) {
		bp := NewBP()

		bp.Insert("5", "dummy")
		assert.Equal(t, [3]string{"5", "", ""}, bp.root.(*BPLeafNode).key)

		bp.Insert("7", "dummy")
		assert.Equal(t, [3]string{"5", "7", ""}, bp.root.(*BPLeafNode).key)

		bp.Insert("3", "dummy")
		assert.Equal(t, [4]string{"", "5", "", ""}, bp.root.(*BPNonLeafNode).key)
		assert.Equal(t, [3]string{"3", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"5", "7", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPLeafNode).key)

		bp.Insert("1", "dummy")
		assert.Equal(t, [4]string{"", "5", "", ""}, bp.root.(*BPNonLeafNode).key)
		assert.Equal(t, [3]string{"1", "3", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"5", "7", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPLeafNode).key)

		bp.Insert("0", "dummy")
		assert.Equal(t, [4]string{"", "1", "5", ""}, bp.root.(*BPNonLeafNode).key)
		assert.Equal(t, [3]string{"0", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"1", "3", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"5", "7", ""}, bp.root.(*BPNonLeafNode).children[2].(*BPLeafNode).key)

		bp.Insert("2", "dummy")
		assert.Equal(t, [4]string{"", "2", "", ""}, bp.root.(*BPNonLeafNode).key)
		assert.Equal(t, [4]string{"", "1", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).key)
		assert.Equal(t, [4]string{"", "5", "", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).key)
		assert.Equal(t, [3]string{"0", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"1", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).children[1].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"2", "3", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"5", "7", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[1].(*BPLeafNode).key)

		bp.Insert("4", "dummy")
		assert.Equal(t, [4]string{"", "2", "", ""}, bp.root.(*BPNonLeafNode).key)
		assert.Equal(t, [4]string{"", "1", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).key)
		assert.Equal(t, [4]string{"", "3", "5", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).key)
		assert.Equal(t, [3]string{"0", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"1", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).children[1].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"2", "", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"3", "4", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[1].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"5", "7", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[2].(*BPLeafNode).key)

		bp.Insert("6", "dummy")
		assert.Equal(t, [4]string{"", "2", "5", ""}, bp.root.(*BPNonLeafNode).key)
		assert.Equal(t, [4]string{"", "1", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).key)
		assert.Equal(t, [4]string{"", "3", "", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).key)
		assert.Equal(t, [4]string{"", "6", "", ""}, bp.root.(*BPNonLeafNode).children[2].(*BPNonLeafNode).key)
		assert.Equal(t, [3]string{"0", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"1", "", ""}, bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).children[1].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"2", "", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"3", "4", ""}, bp.root.(*BPNonLeafNode).children[1].(*BPNonLeafNode).children[1].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"5", "", ""}, bp.root.(*BPNonLeafNode).children[2].(*BPNonLeafNode).children[0].(*BPLeafNode).key)
		assert.Equal(t, [3]string{"6", "7", ""}, bp.root.(*BPNonLeafNode).children[2].(*BPNonLeafNode).children[1].(*BPLeafNode).key)

		n := bp.root.(*BPNonLeafNode).children[0].(*BPNonLeafNode).children[0].(*BPLeafNode)
		s := []string{}
		for {
			if n == nil {
				break
			}

			for i, v := range n.key {
				if i == n.Length() {
					break
				}
				s = append(s, v)
			}
			n = n.next
		}
		assert.Equal(t, []string{"0", "1", "2", "3", "4", "5", "6", "7"}, s)
	})
}

func TestFind(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		bp := NewBP()

		s := uuid.NewString()
		ev := fmt.Sprintf("v-%s", s)
		bp.Insert(s, ev)

		av, err := bp.Find(s)
		assert.Nil(t, err)

		assert.Equal(t, ev, av)
	})

	t.Run("2", func(t *testing.T) {
		bp := NewBP()

		ss := []string{}
		for i := 0; i < 10000; i++ {
			ss = append(ss, uuid.NewString())
		}

		for _, s := range ss {
			ev := fmt.Sprintf("v-%s", s)
			bp.Insert(s, ev)
			av, err := bp.Find(s)
			assert.Nil(t, err)
			assert.Equal(t, ev, av)
		}
	})
}
