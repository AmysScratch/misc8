package bot

import (
)

const (
	Lex_SUDO = uint64(1) << iota,
	Lex_,
)

type Lex struct {
	Flags uint64
}

func (com *Com) Lex() error {
}
