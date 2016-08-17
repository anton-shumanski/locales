package ar

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ar struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
	percentPrefix          []byte
	percentSuffix          []byte
	perMille               []byte
	timeSeparator          []byte
	inifinity              []byte
	currencies             [][]byte // idx = enum of currency code
	currencyPositivePrefix []byte
	currencyPositiveSuffix []byte
	currencyNegativePrefix []byte
	currencyNegativeSuffix []byte
	monthsAbbreviated      [][]byte
	monthsNarrow           [][]byte
	monthsShort            [][]byte
	monthsWide             [][]byte
	daysAbbreviated        [][]byte
	daysNarrow             [][]byte
	daysShort              [][]byte
	daysWide               [][]byte
	periodsAbbreviated     [][]byte
	periodsNarrow          [][]byte
	periodsShort           [][]byte
	periodsWide            [][]byte
	erasAbbreviated        [][]byte
	erasNarrow             [][]byte
	erasWide               [][]byte
}

// New returns a new instance of translator for the 'ar' locale
func New() locales.Translator {
	return &ar{
		locale:                 "ar",
		pluralsCardinal:        []locales.PluralRule{1, 2, 3, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0xd9, 0xab},
		group:                  []byte{0xd9, 0xac},
		minus:                  []byte{0xe2, 0x80, 0x8f, 0x2d},
		percent:                []byte{0xd9, 0xaa},
		perMille:               []byte{0xd8, 0x89},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0xd8, 0xaf, 0x2e, 0xd8, 0xa5, 0x2e, 0xe2, 0x80, 0x8f}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x24}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0xd8, 0xaf, 0x2e, 0xd8, 0xa8, 0x2e, 0xe2, 0x80, 0x8f}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x52, 0x24}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x24}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0xc2, 0xa5}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0xd8, 0xaf, 0x2e, 0xd8, 0xac, 0x2e, 0xe2, 0x80, 0x8f}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0xd8, 0xac, 0x2e, 0xd9, 0x85, 0x2e, 0xe2, 0x80, 0x8f}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x24}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0xd8, 0xb1, 0x2e, 0xd8, 0xa5, 0xd9, 0x86, 0x2e}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0xe2, 0x82, 0xaa}, {0xe2, 0x82, 0xb9}, {0xd8, 0xaf, 0x2e, 0xd8, 0xb9, 0x2e, 0xe2, 0x80, 0x8f}, {0xd8, 0xb1, 0x2e, 0xd8, 0xa5, 0x2e}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0xd8, 0xaf, 0x2e, 0xd8, 0xa3, 0x2e, 0xe2, 0x80, 0x8f}, {0x4a, 0x50, 0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0xd9, 0x81, 0x2e, 0xd8, 0xac, 0x2e, 0xd9, 0x82, 0x2e, 0xe2, 0x80, 0x8f}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0xe2, 0x82, 0xa9}, {0xd8, 0xaf, 0x2e, 0xd9, 0x83, 0x2e, 0xe2, 0x80, 0x8f}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0xd9, 0x84, 0x2e, 0xd9, 0x84, 0x2e, 0xe2, 0x80, 0x8f}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0xd8, 0xaf, 0x2e, 0xd9, 0x84, 0x2e, 0xe2, 0x80, 0x8f}, {0xd8, 0xaf, 0x2e, 0xd9, 0x85, 0x2e, 0xe2, 0x80, 0x8f}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0xd8, 0xa3, 0x2e, 0xd9, 0x85, 0x2e, 0xe2, 0x80, 0x8f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x24}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x24}, {0xd8, 0xb1, 0x2e, 0xd8, 0xb9, 0x2e, 0xe2, 0x80, 0x8f}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0xd8, 0xb1, 0x2e, 0xd8, 0xa8, 0x2e}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0xd8, 0xb1, 0x2e, 0xd9, 0x82, 0x2e, 0xe2, 0x80, 0x8f}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0xd8, 0xb1, 0x2e, 0xd8, 0xb3, 0x2e, 0xe2, 0x80, 0x8f}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0xd8, 0xaf, 0x2e, 0xd8, 0xb3, 0x2e, 0xe2, 0x80, 0x8f}, {0xd8, 0xac, 0x2e, 0xd8, 0xb3, 0x2e}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0xd8, 0xac, 0x2e, 0xd8, 0xac, 0x2e, 0xd8, 0xb3, 0x2e}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0xd9, 0x84, 0x2e, 0xd8, 0xb3, 0x2e, 0xe2, 0x80, 0x8f}, {0x53, 0x5a, 0x4c}, {0xe0, 0xb8, 0xbf}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0xd8, 0xaf, 0x2e, 0xd8, 0xaa, 0x2e, 0xe2, 0x80, 0x8f}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0xd9, 0x84, 0x2e, 0xd8, 0xaa, 0x2e}, {0x54, 0x54, 0x44}, {0x4e, 0x54, 0x24}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x24}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x45, 0x43, 0x24}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x2a, 0x2a, 0x2a}, {0x59, 0x44, 0x44}, {0xd8, 0xb1, 0x2e, 0xd9, 0x8a, 0x2e, 0xe2, 0x80, 0x8f}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0xd9, 0x8a, 0xd9, 0x86, 0xd8, 0xa7, 0xd9, 0x8a, 0xd8, 0xb1}, {0xd9, 0x81, 0xd8, 0xa8, 0xd8, 0xb1, 0xd8, 0xa7, 0xd9, 0x8a, 0xd8, 0xb1}, {0xd9, 0x85, 0xd8, 0xa7, 0xd8, 0xb1, 0xd8, 0xb3}, {0xd8, 0xa3, 0xd8, 0xa8, 0xd8, 0xb1, 0xd9, 0x8a, 0xd9, 0x84}, {0xd9, 0x85, 0xd8, 0xa7, 0xd9, 0x8a, 0xd9, 0x88}, {0xd9, 0x8a, 0xd9, 0x88, 0xd9, 0x86, 0xd9, 0x8a, 0xd9, 0x88}, {0xd9, 0x8a, 0xd9, 0x88, 0xd9, 0x84, 0xd9, 0x8a, 0xd9, 0x88}, {0xd8, 0xa3, 0xd8, 0xba, 0xd8, 0xb3, 0xd8, 0xb7, 0xd8, 0xb3}, {0xd8, 0xb3, 0xd8, 0xa8, 0xd8, 0xaa, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd8, 0xa3, 0xd9, 0x83, 0xd8, 0xaa, 0xd9, 0x88, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd9, 0x86, 0xd9, 0x88, 0xd9, 0x81, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd8, 0xaf, 0xd9, 0x8a, 0xd8, 0xb3, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0xd9, 0x8a}, {0xd9, 0x81}, {0xd9, 0x85}, {0xd8, 0xa3}, {0xd9, 0x88}, {0xd9, 0x86}, {0xd9, 0x84}, {0xd8, 0xba}, {0xd8, 0xb3}, {0xd9, 0x83}, {0xd8, 0xa8}, {0xd8, 0xaf}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xd9, 0x8a, 0xd9, 0x86, 0xd8, 0xa7, 0xd9, 0x8a, 0xd8, 0xb1}, {0xd9, 0x81, 0xd8, 0xa8, 0xd8, 0xb1, 0xd8, 0xa7, 0xd9, 0x8a, 0xd8, 0xb1}, {0xd9, 0x85, 0xd8, 0xa7, 0xd8, 0xb1, 0xd8, 0xb3}, {0xd8, 0xa3, 0xd8, 0xa8, 0xd8, 0xb1, 0xd9, 0x8a, 0xd9, 0x84}, {0xd9, 0x85, 0xd8, 0xa7, 0xd9, 0x8a, 0xd9, 0x88}, {0xd9, 0x8a, 0xd9, 0x88, 0xd9, 0x86, 0xd9, 0x8a, 0xd9, 0x88}, {0xd9, 0x8a, 0xd9, 0x88, 0xd9, 0x84, 0xd9, 0x8a, 0xd9, 0x88}, {0xd8, 0xa3, 0xd8, 0xba, 0xd8, 0xb3, 0xd8, 0xb7, 0xd8, 0xb3}, {0xd8, 0xb3, 0xd8, 0xa8, 0xd8, 0xaa, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd8, 0xa3, 0xd9, 0x83, 0xd8, 0xaa, 0xd9, 0x88, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd9, 0x86, 0xd9, 0x88, 0xd9, 0x81, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd8, 0xaf, 0xd9, 0x8a, 0xd8, 0xb3, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}},
		daysAbbreviated:        [][]uint8{{0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xa3, 0xd8, 0xad, 0xd8, 0xaf}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xa7, 0xd8, 0xab, 0xd9, 0x86, 0xd9, 0x8a, 0xd9, 0x86}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xab, 0xd9, 0x84, 0xd8, 0xa7, 0xd8, 0xab, 0xd8, 0xa7, 0xd8, 0xa1}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xa3, 0xd8, 0xb1, 0xd8, 0xa8, 0xd8, 0xb9, 0xd8, 0xa7, 0xd8, 0xa1}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xae, 0xd9, 0x85, 0xd9, 0x8a, 0xd8, 0xb3}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xac, 0xd9, 0x85, 0xd8, 0xb9, 0xd8, 0xa9}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xb3, 0xd8, 0xa8, 0xd8, 0xaa}},
		daysNarrow:             [][]uint8{{0xd8, 0xad}, {0xd9, 0x86}, {0xd8, 0xab}, {0xd8, 0xb1}, {0xd8, 0xae}, {0xd8, 0xac}, {0xd8, 0xb3}},
		daysShort:              [][]uint8{{0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xa3, 0xd8, 0xad, 0xd8, 0xaf}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xa7, 0xd8, 0xab, 0xd9, 0x86, 0xd9, 0x8a, 0xd9, 0x86}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xab, 0xd9, 0x84, 0xd8, 0xa7, 0xd8, 0xab, 0xd8, 0xa7, 0xd8, 0xa1}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xa3, 0xd8, 0xb1, 0xd8, 0xa8, 0xd8, 0xb9, 0xd8, 0xa7, 0xd8, 0xa1}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xae, 0xd9, 0x85, 0xd9, 0x8a, 0xd8, 0xb3}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xac, 0xd9, 0x85, 0xd8, 0xb9, 0xd8, 0xa9}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xb3, 0xd8, 0xa8, 0xd8, 0xaa}},
		daysWide:               [][]uint8{{0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xa3, 0xd8, 0xad, 0xd8, 0xaf}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xa7, 0xd8, 0xab, 0xd9, 0x86, 0xd9, 0x8a, 0xd9, 0x86}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xab, 0xd9, 0x84, 0xd8, 0xa7, 0xd8, 0xab, 0xd8, 0xa7, 0xd8, 0xa1}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xa3, 0xd8, 0xb1, 0xd8, 0xa8, 0xd8, 0xb9, 0xd8, 0xa7, 0xd8, 0xa1}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xae, 0xd9, 0x85, 0xd9, 0x8a, 0xd8, 0xb3}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xac, 0xd9, 0x85, 0xd8, 0xb9, 0xd8, 0xa9}, {0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xb3, 0xd8, 0xa8, 0xd8, 0xaa}},
		periodsAbbreviated:     [][]uint8{{0xd8, 0xb5}, {0xd9, 0x85}},
		periodsNarrow:          [][]uint8{{0xd8, 0xb5}, {0xd9, 0x85}},
		periodsWide:            [][]uint8{{0xd8, 0xb5}, {0xd9, 0x85}},
		erasAbbreviated:        [][]uint8{[]uint8(nil), []uint8(nil)},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{[]uint8(nil), []uint8(nil)},
	}
}

// Locale returns the current translators string locale
func (ar *ar) Locale() string {
	return ar.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ar'
func (ar *ar) PluralsCardinal() []locales.PluralRule {
	return ar.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ar'
func (ar *ar) PluralsOrdinal() []locales.PluralRule {
	return ar.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ar'
func (ar *ar) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod100 := math.Mod(n, 100)

	if n == 0 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if nMod100 >= 3 && nMod100 <= 10 {
		return locales.PluralRuleFew
	} else if nMod100 >= 11 && nMod100 <= 99 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ar'
func (ar *ar) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ar'
func (ar *ar) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ar.CardinalPluralRule(num1, v1)
	end := ar.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleZero && end == locales.PluralRuleOne {
		return locales.PluralRuleZero
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleTwo {
		return locales.PluralRuleZero
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleTwo {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleTwo {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ar' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ar *ar) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ar.decimal) + len(ar.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ar.decimal) - 1; j >= 0; j-- {
				b = append(b, ar.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ar.group) - 1; j >= 0; j-- {
					b = append(b, ar.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ar.minus) - 1; j >= 0; j-- {
			b = append(b, ar.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ar' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ar *ar) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ar.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ar.decimal) - 1; j >= 0; j-- {
				b = append(b, ar.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ar.minus) - 1; j >= 0; j-- {
			b = append(b, ar.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ar.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ar'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ar *ar) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ar.currencies[currency]
	l := len(s) + len(ar.decimal) + len(ar.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ar.decimal) - 1; j >= 0; j-- {
				b = append(b, ar.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ar.group) - 1; j >= 0; j-- {
					b = append(b, ar.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(ar.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ar.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(ar.minus) - 1; j >= 0; j-- {
			b = append(b, ar.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ar.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ar'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ar *ar) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ar.currencies[currency]
	l := len(s) + len(ar.decimal) + len(ar.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ar.decimal) - 1; j >= 0; j-- {
				b = append(b, ar.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ar.group) - 1; j >= 0; j-- {
					b = append(b, ar.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ar.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ar.currencyNegativePrefix[j])
		}

		for j := len(ar.minus) - 1; j >= 0; j-- {
			b = append(b, ar.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ar.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ar.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ar.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'ar'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ar *ar) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe2, 0x80, 0x8f, 0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xe2, 0x80, 0x8f, 0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'ar'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ar *ar) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe2, 0x80, 0x8f, 0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0xe2, 0x80, 0x8f, 0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'ar'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ar *ar) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ar.monthsWide[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'ar'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ar *ar) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, ar.daysWide[t.Day()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ar.monthsWide[t.Month()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'ar'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ar *ar) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ar.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ar.periodsAbbreviated[0]...)
	} else {
		b = append(b, ar.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'ar'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ar *ar) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ar.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ar.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ar.periodsAbbreviated[0]...)
	} else {
		b = append(b, ar.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'ar'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ar *ar) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ar.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ar.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ar.periodsAbbreviated[0]...)
	} else {
		b = append(b, ar.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'ar'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ar *ar) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ar.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ar.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ar.periodsAbbreviated[0]...)
	} else {
		b = append(b, ar.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}
