package unlimchan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("first combined test", func(t *testing.T) {
		q := &Queue{}
		assert.Equal(t, 0, q.Size())
		q.Enqueue(1)
		assert.Equal(t, 1, q.Size())
		q.Enqueue(1)
		assert.Equal(t, 2, q.Size())
		_, err := q.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 1, q.Size())
	})

	t.Run("second combined test", func(t *testing.T) {
		q := &Queue{}
		q.Enqueue(123)
		assert.Equal(t, 123, q.head.val)
		assert.Equal(t, 123, q.tail.val)
		assert.Nil(t, q.head.next)
		assert.Nil(t, q.tail.next)
		val, err := q.Dequeue()
		assert.NoError(t, err)
		assert.Equal(t, 123, val)
		assert.Nil(t, q.head)
		assert.Nil(t, q.tail)
		_, err = q.Dequeue()
		assert.EqualError(t, err, "queue is empty")
		q.Enqueue(1)
		q.Enqueue(2)
		assert.Equal(t, 2, q.Size())
	})
}
