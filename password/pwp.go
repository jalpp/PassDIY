package cmds

import (
	eng "github.com/gregoryv/english"
)

func GetPwp() string {

	var buffer string = ""

	for i := 0; i < PASSPHRASE_COUNT_NUM; i++ {
		buffer += eng.RandomWord() + " "
	}

	return buffer
}
