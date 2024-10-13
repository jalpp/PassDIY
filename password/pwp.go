package cmds

import (
	eng "github.com/gregoryv/english"
	config "github.com/jalpp/passdiy/config"
)

func GetPwp() string {

	var buffer string = ""

	for i := 0; i < config.PASSPHRASE_COUNT_NUM; i++ {
		buffer += eng.RandomWord() + " "
	}

	return buffer
}
