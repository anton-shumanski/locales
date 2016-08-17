package or_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type or_IN struct {
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

// New returns a new instance of translator for the 'or_IN' locale
func New() locales.Translator {
	return &or_IN{
		locale:                 "or_IN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0xe0, 0xac, 0x9c, 0xe0, 0xac, 0xbe, 0xe0, 0xac, 0xa8, 0xe0, 0xad, 0x81, 0xe0, 0xac, 0x86, 0xe0, 0xac, 0xb0, 0xe0, 0xad, 0x80}, {0xe0, 0xac, 0xab, 0xe0, 0xad, 0x87, 0xe0, 0xac, 0xac, 0xe0, 0xad, 0x83, 0xe0, 0xac, 0x86, 0xe0, 0xac, 0xb0, 0xe0, 0xad, 0x80}, {0xe0, 0xac, 0xae, 0xe0, 0xac, 0xbe, 0xe0, 0xac, 0xb0, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0x9a, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0x9a}, {0xe0, 0xac, 0x85, 0xe0, 0xac, 0xaa, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0xb0, 0xe0, 0xad, 0x87, 0xe0, 0xac, 0xb2}, {0xe0, 0xac, 0xae, 0xe0, 0xac, 0x87}, {0xe0, 0xac, 0x9c, 0xe0, 0xad, 0x81, 0xe0, 0xac, 0xa8}, {0xe0, 0xac, 0x9c, 0xe0, 0xad, 0x81, 0xe0, 0xac, 0xb2, 0xe0, 0xac, 0xbe, 0xe0, 0xac, 0x87}, {0xe0, 0xac, 0x85, 0xe0, 0xac, 0x97, 0xe0, 0xac, 0xb7, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0x9f}, {0xe0, 0xac, 0xb8, 0xe0, 0xad, 0x87, 0xe0, 0xac, 0xaa, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0x9f, 0xe0, 0xad, 0x87, 0xe0, 0xac, 0xae, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0x85, 0xe0, 0xac, 0x95, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0x9f, 0xe0, 0xad, 0x8b, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0xa8, 0xe0, 0xac, 0xad, 0xe0, 0xad, 0x87, 0xe0, 0xac, 0xae, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0xa1, 0xe0, 0xac, 0xbf, 0xe0, 0xac, 0xb8, 0xe0, 0xad, 0x87, 0xe0, 0xac, 0xae, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xb0}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0xe0, 0xac, 0x9c, 0xe0, 0xac, 0xbe}, {0xe0, 0xac, 0xab, 0xe0, 0xad, 0x87}, {0xe0, 0xac, 0xae, 0xe0, 0xac, 0xbe}, {0xe0, 0xac, 0x85}, {0xe0, 0xac, 0xae, 0xe0, 0xac, 0x87}, {0xe0, 0xac, 0x9c, 0xe0, 0xad, 0x81}, {0xe0, 0xac, 0x9c, 0xe0, 0xad, 0x81}, {0xe0, 0xac, 0x85}, {0xe0, 0xac, 0xb8, 0xe0, 0xad, 0x87}, {0xe0, 0xac, 0x85}, {0xe0, 0xac, 0xa8}, {0xe0, 0xac, 0xa1, 0xe0, 0xac, 0xbf}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xe0, 0xac, 0x9c, 0xe0, 0xac, 0xbe, 0xe0, 0xac, 0xa8, 0xe0, 0xad, 0x81, 0xe0, 0xac, 0x86, 0xe0, 0xac, 0xb0, 0xe0, 0xad, 0x80}, {0xe0, 0xac, 0xab, 0xe0, 0xad, 0x87, 0xe0, 0xac, 0xac, 0xe0, 0xad, 0x83, 0xe0, 0xac, 0x86, 0xe0, 0xac, 0xb0, 0xe0, 0xad, 0x80}, {0xe0, 0xac, 0xae, 0xe0, 0xac, 0xbe, 0xe0, 0xac, 0xb0, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0x9a, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0x9a}, {0xe0, 0xac, 0x85, 0xe0, 0xac, 0xaa, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0xb0, 0xe0, 0xad, 0x87, 0xe0, 0xac, 0xb2}, {0xe0, 0xac, 0xae, 0xe0, 0xac, 0x87}, {0xe0, 0xac, 0x9c, 0xe0, 0xad, 0x81, 0xe0, 0xac, 0xa8}, {0xe0, 0xac, 0x9c, 0xe0, 0xad, 0x81, 0xe0, 0xac, 0xb2, 0xe0, 0xac, 0xbe, 0xe0, 0xac, 0x87}, {0xe0, 0xac, 0x85, 0xe0, 0xac, 0x97, 0xe0, 0xac, 0xb7, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0x9f}, {0xe0, 0xac, 0xb8, 0xe0, 0xad, 0x87, 0xe0, 0xac, 0xaa, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0x9f, 0xe0, 0xad, 0x87, 0xe0, 0xac, 0xae, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0x85, 0xe0, 0xac, 0x95, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0x9f, 0xe0, 0xad, 0x8b, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0xa8, 0xe0, 0xac, 0xad, 0xe0, 0xad, 0x87, 0xe0, 0xac, 0xae, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0xa1, 0xe0, 0xac, 0xbf, 0xe0, 0xac, 0xb8, 0xe0, 0xad, 0x87, 0xe0, 0xac, 0xae, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xb0}},
		daysAbbreviated:        [][]uint8{{0xe0, 0xac, 0xb0, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xbf}, {0xe0, 0xac, 0xb8, 0xe0, 0xad, 0x8b, 0xe0, 0xac, 0xae}, {0xe0, 0xac, 0xae, 0xe0, 0xac, 0x99, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0x97, 0xe0, 0xac, 0xb3}, {0xe0, 0xac, 0xac, 0xe0, 0xad, 0x81, 0xe0, 0xac, 0xa7}, {0xe0, 0xac, 0x97, 0xe0, 0xad, 0x81, 0xe0, 0xac, 0xb0, 0xe0, 0xad, 0x81}, {0xe0, 0xac, 0xb6, 0xe0, 0xad, 0x81, 0xe0, 0xac, 0x95, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0xb6, 0xe0, 0xac, 0xa8, 0xe0, 0xac, 0xbf}},
		daysNarrow:             [][]uint8{{0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0xb8, 0xe0, 0xad, 0x8b}, {0xe0, 0xac, 0xae}, {0xe0, 0xac, 0xac, 0xe0, 0xad, 0x81}, {0xe0, 0xac, 0x97, 0xe0, 0xad, 0x81}, {0xe0, 0xac, 0xb6, 0xe0, 0xad, 0x81}, {0xe0, 0xac, 0xb6}},
		daysWide:               [][]uint8{{0xe0, 0xac, 0xb0, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xbf, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xbe, 0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0xb8, 0xe0, 0xad, 0x8b, 0xe0, 0xac, 0xae, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xbe, 0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0xae, 0xe0, 0xac, 0x99, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0x97, 0xe0, 0xac, 0xb3, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xbe, 0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0xac, 0xe0, 0xad, 0x81, 0xe0, 0xac, 0xa7, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xbe, 0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0x97, 0xe0, 0xad, 0x81, 0xe0, 0xac, 0xb0, 0xe0, 0xad, 0x81, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xbe, 0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0xb6, 0xe0, 0xad, 0x81, 0xe0, 0xac, 0x95, 0xe0, 0xad, 0x8d, 0xe0, 0xac, 0xb0, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xbe, 0xe0, 0xac, 0xb0}, {0xe0, 0xac, 0xb6, 0xe0, 0xac, 0xa8, 0xe0, 0xac, 0xbf, 0xe0, 0xac, 0xac, 0xe0, 0xac, 0xbe, 0xe0, 0xac, 0xb0}},
		periodsAbbreviated:     [][]uint8{{0x61, 0x6d}, {0x70, 0x6d}},
		periodsNarrow:          [][]uint8{{0x61, 0x6d}, {0x70, 0x6d}},
		periodsWide:            [][]uint8{{0x61, 0x6d}, {0x70, 0x6d}},
	}
}

// Locale returns the current translators string locale
func (or *or_IN) Locale() string {
	return or.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'or_IN'
func (or *or_IN) PluralsCardinal() []locales.PluralRule {
	return or.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'or_IN'
func (or *or_IN) PluralsOrdinal() []locales.PluralRule {
	return or.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'or_IN'
func (or *or_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'or_IN'
func (or *or_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'or_IN'
func (or *or_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'or_IN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (or *or_IN) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(or.decimal) + len(or.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, or.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, or.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(or.minus) - 1; j >= 0; j-- {
			b = append(b, or.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'or_IN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (or *or_IN) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(or.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, or.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(or.minus) - 1; j >= 0; j-- {
			b = append(b, or.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, or.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'or_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (or *or_IN) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := or.currencies[currency]
	l := len(s) + len(or.decimal) + len(or.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, or.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, or.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(or.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, or.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(or.minus) - 1; j >= 0; j-- {
			b = append(b, or.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, or.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'or_IN'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (or *or_IN) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := or.currencies[currency]
	l := len(s) + len(or.decimal) + len(or.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, or.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, or.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
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

		for j := len(or.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, or.currencyNegativePrefix[j])
		}

		for j := len(or.minus) - 1; j >= 0; j-- {
			b = append(b, or.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(or.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, or.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, or.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'or_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (or *or_IN) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'or_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (or *or_IN) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, or.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'or_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (or *or_IN) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, or.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'or_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (or *or_IN) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, or.daysWide[t.Day()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, or.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'or_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (or *or_IN) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, or.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, or.periodsAbbreviated[0]...)
	} else {
		b = append(b, or.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'or_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (or *or_IN) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, or.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, or.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, or.periodsAbbreviated[0]...)
	} else {
		b = append(b, or.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'or_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (or *or_IN) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, or.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, or.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, or.periodsAbbreviated[0]...)
	} else {
		b = append(b, or.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'or_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (or *or_IN) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}
	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, or.timeSeparator...)
	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, or.timeSeparator...)
	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, or.periodsAbbreviated[0]...)
	} else {
		b = append(b, or.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}
