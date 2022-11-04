package sctp

import sctp "github.com/thebagchi/sctp-go"

const (
	maxAttempts = 30
	SCTPNetowrk = "sctp4"
)

type ConnectionSCTP struct{}

func NewSCTPInitMessage() sctp.SCTPInitMsg {
	return sctp.SCTPInitMsg{
		NumOutStreams:  0xffff,
		MaxInStreams:   0,
		MaxAttempts:    maxAttempts,
		MaxInitTimeout: 0,
	}
}
