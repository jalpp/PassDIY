package cmds

import (
	"math/rand"

	config "github.com/jalpp/passdiy/config"
)

func GetMul(picker string) string {
	var buffer string = "\n"

	for i := 0; i <= config.MULTIPLE_VALUE_COUNT; i++ {
		switch picker {
		case "pass":
			buffer += GetStrongPassword() + "\n"
		case "token":
			buffer += GetAPIToken() + "\n"
		case "pin":
			buffer += GetPin() + "\n"
		case "pwp":
			buffer += GetPwp() + "\n"
		}
	}

	return buffer
}

func GetHundPick(picker string) string {
	var lotteryWheel [config.LOTTERY_WHEEL_COUNT]string

	for i := range lotteryWheel {
		switch picker {
		case "pass":
			lotteryWheel[i] = GetStrongPassword()
		case "token":
			lotteryWheel[i] = GetAPIToken()
		case "pin":
			lotteryWheel[i] = GetPin()
		}
	}

	return lotteryWheel[rand.Intn(config.LOTTERY_WHEEL_COUNT)]
}

func GetTenKPick(picker string) string {

	var outterlotteryWheel [config.LOTTERY_WHEEL_COUNT]string

	for i := range outterlotteryWheel {
		switch picker {
		case "pass":
			outterlotteryWheel[i] = GetHundPick("pass")
		case "token":
			outterlotteryWheel[i] = GetHundPick("token")
		case "pin":
			outterlotteryWheel[i] = GetHundPick("pin")
		}
	}

	return outterlotteryWheel[rand.Intn(config.LOTTERY_WHEEL_COUNT)]
}
