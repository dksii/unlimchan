package unlimchan

// Create builds buffered unlimited channel. Actually creates two channels: in and out.
// First value is input channel for writing and closing the channel. Second is output (for reading).
func Create() (in chan interface{}, out chan interface{}) {
	in, out = make(chan interface{}), make(chan interface{})
	buffered := &Queue{}

	go func() {
		var outCh chan interface{}
		var val interface{}
		var readyToSend bool

		for in != nil || buffered.Size() > 0 || readyToSend {
			select {
			case val, ok := <-in:
				if !ok {
					in = nil
					break
				}
				buffered.Enqueue(val)
			case outCh <- val:
				outCh = nil
				readyToSend = false
			}

			if buffered.Size() > 0 && !readyToSend {
				val, _ = buffered.Dequeue()
				outCh = out
				readyToSend = true
			}
		}

		close(out)
	}()

	return
}
