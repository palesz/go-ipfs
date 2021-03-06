// Package peerstream is a peer-to-peer networking library that multiplexes
// connections to many hosts. It tried to simplify the complexity of:
//
//   * accepting incoming connections over **multiple** listeners
//   * dialing outgoing connections over **multiple** transports
//   * multiplexing **multiple** connections per-peer
//   * multiplexing **multiple** different servers or protocols
//   * handling backpressure correctly
//   * handling stream multiplexing (we use SPDY, but maybe QUIC some day)
//   * providing a **simple** interface to the user
//
package peerstream
