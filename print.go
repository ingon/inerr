package inerr

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Print(err error) {
	Fprint(os.Stderr, err)
}

func Sprint(err error) string {
	var buf strings.Builder
	Fprint(&buf, err)
	return buf.String()
}

func Fprint(out io.Writer, err error) {
	print0(out, err)
}

func print0(out io.Writer, err error) {
	switch x := err.(type) {
	case interface{ Message() string }:
		fmt.Fprintf(out, "%v", x.Message())
	default:
		fmt.Fprintf(out, "%v", err)
	}

	switch x := err.(type) {
	case interface{ Unwrap() error }:
		if err := x.Unwrap(); err != nil {
			print(out, err, 2)
		}
	case interface{ Unwrap() []error }:
		for _, err := range x.Unwrap() {
			if err != nil {
				print(out, err, 2)
			}
		}
	}
}

func print(out io.Writer, err error, depth int) {
	switch x := err.(type) {
	case interface{ Message() string }:
		fmt.Fprintf(out, "\n%s%s", strings.Repeat(" ", depth), x.Message())
	default:
		fmt.Fprintf(out, "\n%s%v", strings.Repeat(" ", depth), err)
	}

	switch x := err.(type) {
	case interface{ Unwrap() error }:
		if err := x.Unwrap(); err != nil {
			print(out, err, depth+2)
		}
	case interface{ Unwrap() []error }:
		for _, err := range x.Unwrap() {
			if err != nil {
				print(out, err, depth+2)
			}
		}
	}
}
