package bot

import (
	"errors"

	"github.com/freepkg-alpha/maths/bitfield"
)

const (
	Lex_SUDO = uint64(1) << iota,
	Lex_,
)

type Lex struct {
	Flags bitfield.Uint64
}

var ELex = errors.New("Lex failure")

func (com *Com) Lex() error {
	t := com.Tail
	f := uint64(0)
	if len(t) < 2 || t[0] != ',' {
		return
	}
	t = t[1:]
	if
	com.Tail = t
	com.Flags = bitfield.Uint64(f)
}
