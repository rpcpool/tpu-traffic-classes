package main

import "strconv"

type ValidatorPorts struct {
	TPU     uint16
	TPUfwd  uint16
	TPUvote uint16
}

func (vp *ValidatorPorts) TPUstr() string {
	return strconv.FormatUint(uint64(vp.TPU), 10)
}

func (vp *ValidatorPorts) Fwdstr() string {
	return strconv.FormatUint(uint64(vp.TPUfwd), 10)
}

func (vp *ValidatorPorts) Votestr() string {
	return strconv.FormatUint(uint64(vp.TPUvote), 10)
}

func NewValidatorPorts(tpu uint16) *ValidatorPorts {
	return &ValidatorPorts{
		TPU:     tpu,
		TPUfwd:  tpu + 1,
		TPUvote: tpu + 2,
	}
}