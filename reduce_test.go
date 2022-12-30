package pleasure_go

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	res := m.Run()
	os.Exit(res)
}

func TestReduce(t *testing.T) {
	t.Run("Should be a valid result with primitive numeric value", func(t *testing.T) {
		t.Parallel()
		entity := []int{2, 11, 4}
		cb := func(acc, item int) int {
			acc += item
			return acc
		}
		baseValue := 0
		actualResult := Reduce(entity, cb, baseValue)
		expected := 17
		assert.Equal(t, expected, actualResult)
	})

	t.Run("Should be a valid result with primitive string value", func(t *testing.T) {
		t.Parallel()
		entity := []string{"lang"}
		cb := func(acc, item string) string {
			acc += item
			return acc
		}
		baseValue := "go"
		actualResult := Reduce(entity, cb, baseValue)
		expected := "golang"
		assert.Equal(t, expected, actualResult)
	})

	t.Run("Should be a valid result with primitive string value", func(t *testing.T) {
		t.Parallel()

		type User struct {
			id   int
			name string
		}

		entity := []User{{id: 1, name: "Jon"}, {id: 2, name: "Ben"}}
		cb := func(acc map[string]int, item User) map[string]int {
			acc[item.name] = item.id
			return acc
		}
		baseValue := map[string]int{}
		actualResult := Reduce(entity, cb, baseValue)
		expected := map[string]int{
			"Jon": 1,
			"Ben": 2,
		}
		assert.Equal(t, expected, actualResult)
	})
}
