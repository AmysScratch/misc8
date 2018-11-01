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
	GuildRecord GuildRecord
	Flags       bitfield.Uint64
}

type GuildRecord struct {
	RowID  int64
	Prefix string
}

var ELex = errors.New("Lex failure")

func (com *Com) Lex() error {
	stmt, err := DB.Prepare(`SELECT "_ROWID_", "prefix" FROM "Guild" WHERE "gid" = ? LIMIT 1`)
	if err != nil {
		return err
	}
	rows, err := stmt.Query(com.GID64)
	if err != nil {
		return err
	}
	var p string
	if rows.Next() {
		err := rows.Scan(
			&com.Lex.GuildRecord.RowID,
			&p,
		)
		rows.Close()
		if err != nil {
			return err
		}
	} else {
		com.Lex.GuildRecord.RowID = -1
		p = ","
	}
	com.Lex.GuildRecord.Prefix = p
	t := com.Tail
	f := uint64(0)
	len_p := len(p)
	if len(t) < len_p {
		return ELex
	}
	if t[:len_p] != p {
		return ELex
	}
	t = t[len_p:]
	com.Tail = t
	com.Flags = bitfield.Uint64(f)
}
