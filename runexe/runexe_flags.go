package main

import (
	"strings"
	"strconv"
)

type TimeLimitFlag uint64

func (t *TimeLimitFlag) String() string {
	return strconv.Itoa(int(*t / 1000)) + "ms"
}

func (t *TimeLimitFlag) Set(v string) error {
	if strings.HasSuffix(v, "ms") {
		r, err := strconv.Atoi(v[:len(v) - 2])
		if err != nil {
			return err
		}
		*t = TimeLimitFlag(r * 1000)
		return nil
	}

	r, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return err
	}
	*t = TimeLimitFlag(r * 1000000)
	return nil
}

type MemoryLimitFlag uint64

func (t *MemoryLimitFlag) String() string {
	return strconv.Itoa(int(*t))
}

func (t *MemoryLimitFlag) Set(v string) error {
	if strings.HasSuffix(v, "M") {
		r, err := strconv.Atoi(v[:len(v) - 1])
		if err != nil {
			return err
		}
		*t = MemoryLimitFlag(r * 1024 * 1024)
		return nil
	}
	r, err := strconv.Atoi(v)
	if err != nil {
		return err
	}
	*t = MemoryLimitFlag(r)
	return nil
}
