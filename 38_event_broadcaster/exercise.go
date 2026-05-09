// Package broadcaster is a tiny fan-out broadcaster, the same shape as
// client-go's `watch.Broadcaster` and what informers use internally to
// fan a single event stream out to many event handlers.
//
// Design notes:
//
// A broadcaster has many subscribers and one publisher. Important
// engineering questions:
//
//  1. What if a subscriber is slow?
//     Real systems pick one of:
//       - block the publisher (back-pressure; bad for global pipelines)
//       - drop events for that subscriber (good for telemetry/metrics)
//       - queue unboundedly (memory bomb; never do this)
//     This exercise picks DROP: each subscriber has its own bounded
//     channel, and Publish performs a non-blocking send. If the buffer is
//     full, the event is dropped FOR THAT SUBSCRIBER ONLY. Other
//     subscribers must still receive it.
//
//  2. Who closes the subscriber channel?
//     The broadcaster owns its subscriber channels. When Close() is
//     called, the broadcaster closes every subscriber channel exactly
//     once. After Close(), Publish must be a no-op and Subscribe must
//     return a closed channel (so receivers immediately see EOF).
//
//  3. Unsubscribe must not race with Publish.
//     A safe implementation uses a mutex around the subscriber set and
//     does the non-blocking sends while holding (at least) a read lock,
//     OR copies the subscriber list under the lock and sends outside.
//     This exercise does NOT prescribe which; just make the semantics
//     correct.
package broadcaster

// Broadcaster fans out values of type T to many subscribers.
type Broadcaster[T any] struct {
	// TODO: store subscriber channels, a closed flag, and a lock.
}

// New returns an empty broadcaster.
func New[T any]() *Broadcaster[T] {
	panic("TODO: implement New")
}

// Subscribe returns a new channel that will receive future Publish'd
// values. The channel has the given buffer size. After Close(), Subscribe
// must return an already-closed channel.
func (b *Broadcaster[T]) Subscribe(buf int) <-chan T {
	panic("TODO: implement Subscribe")
}

// Unsubscribe removes a previously-returned channel and closes it.
// It is a no-op if ch was not registered (e.g. already removed).
func (b *Broadcaster[T]) Unsubscribe(ch <-chan T) {
	panic("TODO: implement Unsubscribe")
}

// Publish delivers v to every subscriber via a non-blocking send.
// If a subscriber's buffer is full, the event is dropped FOR THAT
// SUBSCRIBER ONLY; other subscribers must still receive it.
// Publish on a closed broadcaster is a silent no-op.
func (b *Broadcaster[T]) Publish(v T) {
	panic("TODO: implement Publish")
}

// Close stops the broadcaster and closes every subscriber channel.
// Close must be idempotent.
func (b *Broadcaster[T]) Close() {
	panic("TODO: implement Close")
}
