package keyboard

import (
	"github.com/eiannone/keyboard"
)

func KeyboardHandler(keyChannel chan uint32) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	var (
		PID    uint32 = 1
		exitId uint32 = 0
	)
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeySpace {
			keyChannel <- PID
			PID++
		} else if key == keyboard.KeyEsc {
			keyboard.Close()
			keyChannel <- exitId
			return
		}
	}
}
