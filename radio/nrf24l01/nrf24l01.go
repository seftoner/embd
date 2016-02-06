package nrf24l01

import "github.com/kidoman/embd"

type NRF24L01 struct {
	CePin embd.DigitalPin

	CsnPin embd.DigitalPin

	Bus embd.SPIBus
}

func New(cePin embd.DigitalPin, csnPin embd.DigitalPin, bus embd.SPIBus) *NRF24L01 {
	return &NRF24L01{cePin, csnPin, bus}
}
