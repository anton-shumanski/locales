package ne_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ne_IN struct {
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

// New returns a new instance of translator for the 'ne_IN' locale
func New() locales.Translator {
	return &ne_IN{
		locale:                 "ne_IN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0xe0, 0xa4, 0x9c, 0xe0, 0xa4, 0xa8, 0xe0, 0xa4, 0xb5, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x80}, {0xe0, 0xa4, 0xab, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0xac, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0x85, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x80}, {0xe0, 0xa4, 0xae, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x9a}, {0xe0, 0xa4, 0x85, 0xe0, 0xa4, 0xaa, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb0, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xb2}, {0xe0, 0xa4, 0xae, 0xe0, 0xa5, 0x87}, {0xe0, 0xa4, 0x9c, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xa8}, {0xe0, 0xa4, 0x9c, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xb2, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0x88}, {0xe0, 0xa4, 0x85, 0xe0, 0xa4, 0x97, 0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x9f}, {0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0xaa, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x9f, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0xae, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0x85, 0xe0, 0xa4, 0x95, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x9f, 0xe0, 0xa5, 0x8b, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xa8, 0xe0, 0xa5, 0x8b, 0xe0, 0xa4, 0xad, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0xae, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xa1, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0xae, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xb0}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0xe0, 0xa5, 0xa7}, {0xe0, 0xa5, 0xa8}, {0xe0, 0xa5, 0xa9}, {0xe0, 0xa5, 0xaa}, {0xe0, 0xa5, 0xab}, {0xe0, 0xa5, 0xac}, {0xe0, 0xa5, 0xad}, {0xe0, 0xa5, 0xae}, {0xe0, 0xa5, 0xaf}, {0xe0, 0xa5, 0xa7, 0xe0, 0xa5, 0xa6}, {0xe0, 0xa5, 0xa7, 0xe0, 0xa5, 0xa7}, {0xe0, 0xa5, 0xa7, 0xe0, 0xa5, 0xa8}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xe0, 0xa4, 0x9c, 0xe0, 0xa4, 0xa8, 0xe0, 0xa4, 0xb5, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x80}, {0xe0, 0xa4, 0xab, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0xac, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0x85, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x80}, {0xe0, 0xa4, 0xae, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x9a}, {0xe0, 0xa4, 0x85, 0xe0, 0xa4, 0xaa, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb0, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xb2}, {0xe0, 0xa4, 0xae, 0xe0, 0xa4, 0x88}, {0xe0, 0xa4, 0x9c, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xa8}, {0xe0, 0xa4, 0x9c, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xb2, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0x88}, {0xe0, 0xa4, 0x85, 0xe0, 0xa4, 0x97, 0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x9f}, {0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0xaa, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x9f, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0xae, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0x85, 0xe0, 0xa4, 0x95, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x9f, 0xe0, 0xa5, 0x8b, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xa8, 0xe0, 0xa5, 0x8b, 0xe0, 0xa4, 0xad, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0xae, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xa1, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x87, 0xe0, 0xa4, 0xae, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xb0}},
		daysAbbreviated:        [][]uint8{{0xe0, 0xa4, 0x86, 0xe0, 0xa4, 0x87, 0xe0, 0xa4, 0xa4}, {0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x8b, 0xe0, 0xa4, 0xae}, {0xe0, 0xa4, 0xae, 0xe0, 0xa4, 0x99, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x97, 0xe0, 0xa4, 0xb2}, {0xe0, 0xa4, 0xac, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xa7}, {0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xb9, 0xe0, 0xa5, 0x80}, {0xe0, 0xa4, 0xb6, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0x95, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xb6, 0xe0, 0xa4, 0xa8, 0xe0, 0xa4, 0xbf}},
		daysNarrow:             [][]uint8{{0xe0, 0xa4, 0x86}, {0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x8b}, {0xe0, 0xa4, 0xae}, {0xe0, 0xa4, 0xac, 0xe0, 0xa5, 0x81}, {0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xbf}, {0xe0, 0xa4, 0xb6, 0xe0, 0xa5, 0x81}, {0xe0, 0xa4, 0xb6}},
		daysShort:              [][]uint8{{0xe0, 0xa4, 0x86, 0xe0, 0xa4, 0x87, 0xe0, 0xa4, 0xa4}, {0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x8b, 0xe0, 0xa4, 0xae}, {0xe0, 0xa4, 0xae, 0xe0, 0xa4, 0x99, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x97, 0xe0, 0xa4, 0xb2}, {0xe0, 0xa4, 0xac, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xa7}, {0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xb9, 0xe0, 0xa5, 0x80}, {0xe0, 0xa4, 0xb6, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0x95, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xb6, 0xe0, 0xa4, 0xa8, 0xe0, 0xa4, 0xbf}},
		daysWide:               [][]uint8{{0xe0, 0xa4, 0x86, 0xe0, 0xa4, 0x87, 0xe0, 0xa4, 0xa4, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xb8, 0xe0, 0xa5, 0x8b, 0xe0, 0xa4, 0xae, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xae, 0xe0, 0xa4, 0x99, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0x97, 0xe0, 0xa4, 0xb2, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xac, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0xa7, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xb9, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xb6, 0xe0, 0xa5, 0x81, 0xe0, 0xa4, 0x95, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb0, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}, {0xe0, 0xa4, 0xb6, 0xe0, 0xa4, 0xa8, 0xe0, 0xa4, 0xbf, 0xe0, 0xa4, 0xac, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb0}},
		periodsAbbreviated:     [][]uint8{{0xe0, 0xa4, 0xaa, 0xe0, 0xa5, 0x82, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb5, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb9, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xa8}, {0xe0, 0xa4, 0x85, 0xe0, 0xa4, 0xaa, 0xe0, 0xa4, 0xb0, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb9, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xa8}},
		periodsNarrow:          [][]uint8{{0x61}, {0x70}},
		periodsWide:            [][]uint8{{0xe0, 0xa4, 0xaa, 0xe0, 0xa5, 0x82, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb5, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb9, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xa8}, {0xe0, 0xa4, 0x85, 0xe0, 0xa4, 0xaa, 0xe0, 0xa4, 0xb0, 0xe0, 0xa4, 0xbe, 0xe0, 0xa4, 0xb9, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xa8}},
		erasAbbreviated:        [][]uint8{{0xe0, 0xa4, 0x88, 0xe0, 0xa4, 0xb8, 0xe0, 0xa4, 0xbe, 0x20, 0xe0, 0xa4, 0xaa, 0xe0, 0xa5, 0x82, 0xe0, 0xa4, 0xb0, 0xe0, 0xa5, 0x8d, 0xe0, 0xa4, 0xb5}, {0xe0, 0xa4, 0xb8, 0xe0, 0xa4, 0xa8, 0xe0, 0xa5, 0x8d}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{[]uint8(nil), []uint8(nil)},
	}
}

// Locale returns the current translators string locale
func (ne *ne_IN) Locale() string {
	return ne.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ne_IN'
func (ne *ne_IN) PluralsCardinal() []locales.PluralRule {
	return ne.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ne_IN'
func (ne *ne_IN) PluralsOrdinal() []locales.PluralRule {
	return ne.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ne_IN'
func (ne *ne_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ne_IN'
func (ne *ne_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 1 && n <= 4 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ne_IN'
func (ne *ne_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ne.CardinalPluralRule(num1, v1)
	end := ne.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ne_IN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ne *ne_IN) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ne.decimal) + len(ne.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ne.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ne.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ne.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ne_IN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ne *ne_IN) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ne.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ne.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ne.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ne.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ne_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ne *ne_IN) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ne.currencies[currency]
	l := len(s) + len(ne.decimal) + len(ne.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ne.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ne.group[0])
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

	for j := len(ne.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ne.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, ne.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ne.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ne_IN'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ne *ne_IN) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ne.currencies[currency]
	l := len(s) + len(ne.decimal) + len(ne.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ne.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ne.group[0])
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

		for j := len(ne.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ne.currencyNegativePrefix[j])
		}

		b = append(b, ne.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ne.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ne.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ne.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'ne_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ne *ne_IN) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'ne_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ne *ne_IN) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ne.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'ne_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ne *ne_IN) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ne.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'ne_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ne *ne_IN) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ne.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ne.daysWide[t.Day()]...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'ne_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ne *ne_IN) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ne.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ne.periodsAbbreviated[0]...)
	} else {
		b = append(b, ne.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'ne_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ne *ne_IN) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ne.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ne.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ne.periodsAbbreviated[0]...)
	} else {
		b = append(b, ne.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'ne_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ne *ne_IN) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ne.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ne.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ne.periodsAbbreviated[0]...)
	} else {
		b = append(b, ne.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'ne_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ne *ne_IN) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ne.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ne.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ne.periodsAbbreviated[0]...)
	} else {
		b = append(b, ne.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}
