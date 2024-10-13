package config

const (
	PIN_DIGIT_LENGTH      int = 6   // number of ints in pin digit
	API_TOKEN_CHAR_LENGTH int = 60  // number of chars in a API token
	PASWORD_CHAR_LENGTH   int = 40  // number of chars in a password
	PASSPHRASE_COUNT_NUM  int = 5   // number of words in passphrase
	MULTIPLE_VALUE_COUNT  int = 5   // how many password/tokens you want to output
	LOTTERY_WHEEL_COUNT   int = 100 // how many times you want to generate token/password/pins to randomly pick one (pass100, pass10000)
	SALT_EXTRA_LENGTH     int = 10  // how many extra chars you want to add to a password/token
)
