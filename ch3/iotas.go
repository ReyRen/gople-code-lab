package main

const (
	_   = 1 << (10 * iota)
	KiB //1024
	MiB //1048576
	GiB //103741824
	TiB //1099511627776	(exceeds 1 << 32)
	PiB //1125899906842624
)
