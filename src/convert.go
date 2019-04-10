package main

import (
	sc "strconv"
)

func convert(val string) (int, error) {
	return sc.Atoi(val)
}
