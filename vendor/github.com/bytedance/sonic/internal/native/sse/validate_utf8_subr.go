// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__validate_utf8 = 0
)

const (
    _stack__validate_utf8 = 48
)

const (
    _size__validate_utf8 = 668
)

var (
    _pcsp__validate_utf8 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {11, 40},
        {623, 48},
        {627, 40},
        {628, 32},
        {630, 24},
        {632, 16},
        {634, 8},
        {635, 0},
        {668, 48},
    }
)

var _cfunc_validate_utf8 = []loader.CFunc{
    {"_validate_utf8_entry", 0,  _entry__validate_utf8, 0, nil},
    {"_validate_utf8", _entry__validate_utf8, _size__validate_utf8, _stack__validate_utf8, _pcsp__validate_utf8},
}
