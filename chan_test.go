package unlimchan

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("test order", func(t *testing.T) {
		in, out := Create()

		for i := 0; i < 100; i++ {
			in <- i
		}

		for i := 0; i < 100; i++ {
			assert.Equal(t, i, <-out)
		}
	})

	t.Run("test close and receive count", func(t *testing.T) {
		in, out := Create()
		var count int
		writes := 10000

		go func() {
			for i := 0; i < writes; i++ {
				in <- i
			}
			close(in)
		}()

		for range out {
			count++
		}

		assert.Equal(t, writes, count)
	})

	t.Run("test close and receive order", func(t *testing.T) {
		in, out := Create()
		writes := 10000

		go func() {
			for i := 0; i < writes; i++ {
				in <- i
			}
			close(in)
		}()

		i := 0
		for val := range out {
			assert.Equal(t, i, val)
			i++
		}
	})

	t.Run("read after close", func(t *testing.T) {
		in, out := Create()
		writes := 10000

		for i := 0; i < writes; i++ {
			in <- i
		}
		close(in)

		i := 0
		for val := range out {
			assert.Equal(t, i, val)
			i++
		}
	})

	t.Run("blocks on read", func(t *testing.T) {
		_, out := Create()

		select {
		case <-out:
			t.Fail()
		case <-time.After(time.Second):
		}
	})
}
