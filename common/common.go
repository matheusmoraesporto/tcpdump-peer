package common

import sctp "github.com/thebagchi/sctp-go"

const (
	SCTPNetowrk = "sctp4"
)

func NewSCTPInitMessage() sctp.SCTPInitMsg {
	return sctp.SCTPInitMsg{
		NumOutStreams:  0xffff,
		MaxInStreams:   0,
		MaxAttempts:    0,
		MaxInitTimeout: 0,
	}
}
