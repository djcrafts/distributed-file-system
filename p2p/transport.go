package p2p

// Peer is an interface that represents the remote node
type Peer interface {}

// Transport is anything that handles the communication between nodes in the network.
// this can be of form - TCP, UDP, Websockets...
type Transport interface {
	ListenAndAccept() error
}
