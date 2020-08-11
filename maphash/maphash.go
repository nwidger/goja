// +build go1.14

package maphash

import "hash/maphash" // standard library's maphash, as of Go 1.14

// A Hash computes a seeded hash of a byte sequence.
//
// The zero Hash is a valid Hash ready to use.
// A zero Hash chooses a random seed for itself during
// the first call to a Reset, Write, Seed, Sum64, or Seed method.
// For control over the seed, use SetSeed.
//
// The computed hash values depend only on the initial seed and
// the sequence of bytes provided to the Hash object, not on the way
// in which the bytes are provided. For example, the three sequences
//
//     h.Write([]byte{'f','o','o'})
//     h.WriteByte('f'); h.WriteByte('o'); h.WriteByte('o')
//     h.WriteString("foo")
//
// all have the same effect.
//
// Hashes are intended to be collision-resistant, even for situations
// where an adversary controls the byte sequences being hashed.
//
// A Hash is not safe for concurrent use by multiple goroutines, but a Seed is.
// If multiple goroutines must compute the same seeded hash,
// each can declare its own Hash and call SetSeed with a common Seed.
type Hash = maphash.Hash

// A Seed is a random value that selects the specific hash function
// computed by a Hash. If two Hashes use the same Seeds, they
// will compute the same hash values for any given input.
// If two Hashes use different Seeds, they are very likely to compute
// distinct hash values for any given input.
//
// A Seed must be initialized by calling MakeSeed.
// The zero seed is uninitialized and not valid for use with Hash's SetSeed method.
//
// Each Seed value is local to a single process and cannot be serialized
// or otherwise recreated in a different process.
type Seed = maphash.Seed

// MakeSeed returns a new random seed.
func MakeSeed() Seed {
	return maphash.MakeSeed()
}
