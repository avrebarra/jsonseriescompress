package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

type Data struct {
	V1 string
	V4 int
	V5 bool
	V6 float64

	VX1  string
	VX2  string
	VX3  string
	VX4  string
	VX5  string
	VX6  string
	VX7  string
	VX8  string
	VX9  string
	VX10 string
	VX11 string
	VX12 string
	VX13 string
	VX14 string
	VX15 string
	VX16 string
	VX17 string
	VX18 string
	VX19 string
	VX20 string
	VX21 string
	VX22 string
	VX23 string
	VX24 string
	VX25 string
	VX26 string
	VX27 string
	VX28 string
	VX29 string
	VX30 string
}

func main() {
	enc := diffmatchpatch.New()
	timeline := []Data{
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, VX2: "almost static data", VX3: "almost static data", VX4: "almost static data", VX5: "almost static data", VX6: "almost static data", VX7: "almost static data", VX8: "almost static data", VX9: "almost static data", VX10: "almost static data", VX11: "almost static data", VX12: "almost static data", VX13: "almost static data", VX14: "almost static data", VX15: "almost static data", VX16: "almost static data", VX17: "almost static data", VX18: "almost static data", VX19: "almost static data", VX20: "almost static data", VX21: "almost static data", VX22: "almost static data", VX23: "almost static data", VX24: "almost static data", VX25: "almost static data", VX26: "almost static data", VX27: "almost static data", VX28: "almost static data", VX29: "almost static data", VX30: "almost static data"},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "123123", V6: 123.453},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: true},
		{V1: "123123", V6: 123.453, V4: 32, V5: false},
		{V1: "21", V6: 123.453, V4: 32, V5: true},
	}

	fmt.Println("NORMAL ENCODING ------")
	enc1, _ := json.Marshal(timeline)
	fmt.Println(string(enc1))
	fmt.Println()

	// build normal json encoding
	fmt.Println("SOURCE ------")
	fmt.Printf("%+v\n", timeline[len(timeline)-1])
	fmt.Println()

	// build mutation based json encoding
	fmt.Println("ENCODING ------")
	enc2 := ""
	lastentry := ""
	for i, entry := range timeline {
		encdiffs := []string{}

		jsonentry, _ := json.Marshal(entry)
		diffs := enc.DiffCleanupSemantic(enc.DiffMain(lastentry, string(jsonentry), false))
		for _, v := range diffs {
			encdiffs = append(encdiffs, Diff{v}.Encode(lastentry))
			// encdiffs = append(encdiffs, fmt.Sprintf("%+v", v))
		}
		fmt.Println(i, lastentry, "to", string(jsonentry), "---->", encdiffs)

		enc2 = strings.Join([]string{enc2, strings.Join(encdiffs, "$")}, "#")
		lastentry = string(jsonentry)
	}

	fmt.Println()
	fmt.Println("RESULTS ------")
	fmt.Println(string(enc2))
	fmt.Println()
	// return

	// recover timeline from encoded string
	fmt.Println("DECODING ------")
	curjson := ""
	for _, tlentry := range strings.Split(enc2, "#") {
		if tlentry == "" {
			continue
		}

		diffs := []diffmatchpatch.Diff{}
		for _, tlentrydiff := range strings.Split(tlentry, "$") {
			diffs = append(diffs, Diff{}.Decode(tlentrydiff, curjson).Diff)
		}

		newjson, _ := enc.PatchApply(enc.PatchMake(curjson, diffs), curjson)
		fmt.Println(diffs, "with current json", curjson, "results to", newjson)

		curjson = newjson
	}
	fmt.Println()
	fmt.Println("RESULTS ------")
	fmt.Println(curjson)
	fmt.Println()
	// fmt.Println(enc.PatchApply(enc.PatchMake("", diffs), ""))
}

// ***

type Diff struct {
	diffmatchpatch.Diff
}

// func (in Diff) Encode(prev string) (out string) {
// 	out = fmt.Sprintf("%s:%s", in.Type, in.Text)
// 	if in.Diff.Type == diffmatchpatch.DiffEqual {
// 		out = fmt.Sprintf("%s:%d@%d", in.Type, len(in.Text), strings.Index(prev, in.Text))
// 	}
// 	return
// }

func (in Diff) Encode(prev string) (out string) {
	out = fmt.Sprintf("%d:%s", int8(in.Type), in.Text)
	if in.Diff.Type == diffmatchpatch.DiffEqual {
		out = fmt.Sprintf("%d:%d@%d", int8(in.Type), len(in.Text), strings.Index(prev, in.Text))
	}
	return
}

func (Diff) Decode(in, prev string) (out Diff) {
	parsed := strings.SplitN(in, ":", 2)

	typeint := convtoint(parsed[0])
	text := parsed[1]

	if typeint == int64(diffmatchpatch.DiffEqual) && strings.Contains(text, "@") {
		parsed := strings.SplitN(text, "@", 2)
		text = prev[convtoint(parsed[1]) : convtoint(parsed[1])+convtoint(parsed[0])]
	}

	return Diff{
		Diff: diffmatchpatch.Diff{
			Type: diffmatchpatch.Operation(typeint),
			Text: text,
		},
	}
}

func convtoint(in string) (out int64) {
	out, _ = strconv.ParseInt(in, 10, 64)
	return
}
