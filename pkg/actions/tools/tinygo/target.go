package tinygo

import "github.com/carapace-sh/carapace"

// ActionTargets completes targets
//
//	arduino-mega1280
//	arduino-mega2560
func ActionTargets() carapace.Action {
	return carapace.ActionValues(
		"arduino-mega1280",
		"arduino-mega2560",
		"arduino-mkr1000",
		"arduino-mkrwifi1010",
		"arduino-nano-new",
		"arduino-nano",
		"arduino-nano33",
		"arduino-zero",
		"arduino",
		"atmega328p",
		"atmega1280",
		"atmega1284p",
		"atmega2560",
		"atsamd21e18a",
		"atsamd21g18a",
		"atsamd51g19a",
		"atsamd51j19a",
		"atsamd51j20a",
		"atsamd51p19a",
		"atsamd51p20a",
		"atsame51j19a",
		"atsame54-xpro",
		"atsame54p20a",
		"attiny85",
		"avr",
		"bluepill-clone",
		"bluepill",
		"circuitplay-bluefruit",
		"circuitplay-express",
		"clue-alpha",
		"clue",
		"cortex-m-qemu",
		"cortex-m",
		"cortex-m0",
		"cortex-m0plus",
		"cortex-m3",
		"cortex-m4",
		"cortex-m7",
		"cortex-m33",
		"d1mini",
		"digispark",
		"esp32-coreboard-v2",
		"esp32-mini32",
		"esp32",
		"esp32c3",
		"esp8266",
		"fe310",
		"feather-m0",
		"feather-m4-can",
		"feather-m4",
		"feather-nrf52840-sense",
		"feather-nrf52840",
		"feather-rp2040",
		"feather-stm32f405",
		"gameboy-advance",
		"grandcentral-m4",
		"hifive1-qemu",
		"hifive1b",
		"itsybitsy-m0",
		"itsybitsy-m4",
		"itsybitsy-nrf52840",
		"k210",
		"lgt92",
		"m5stack-core2",
		"maixbit",
		"matrixportal-m4",
		"mdbt50qrx-uf2",
		"metro-m4-airlift",
		"microbit-s110v8",
		"microbit-v2-s113v7",
		"microbit-v2",
		"microbit",
		"nano-33-ble-s140v7",
		"nano-33-ble",
		"nano-rp2040",
		"nicenano",
		"nintendoswitch",
		"nodemcu",
		"nrf51-s110v8",
		"nrf51",
		"nrf52-s132v6",
		"nrf52",
		"nrf52833-s113v7",
		"nrf52833",
		"nrf52840-mdk-usb-dongle",
		"nrf52840-mdk",
		"nrf52840-s140v7",
		"nrf52840",
		"nucleo-f103rb",
		"nucleo-f722ze",
		"nucleo-l031k6",
		"nucleo-l432kc",
		"nucleo-l552ze",
		"p1am-100",
		"particle-3rd-gen",
		"particle-argon",
		"particle-boron",
		"particle-xenon",
		"pca10031",
		"pca10040-s132v6",
		"pca10040",
		"pca10056-s140v7",
		"pca10056",
		"pca10059",
		"pico",
		"pinetime-devkit0",
		"pybadge",
		"pygamer",
		"pyportal",
		"qtpy",
		"reelboard-s140v7",
		"reelboard",
		"riscv-qemu",
		"riscv",
		"riscv32",
		"riscv64",
		"rp2040",
		"stm32f4disco-1",
		"stm32f4disco",
		"stm32l0x2",
		"teensy36",
		"teensy40",
		"trinket-m0",
		"wasi",
		"wasm",
		"wioterminal",
		"x9pro",
		"xiao",
		"xtensa",
	).Tag("targets")
}
