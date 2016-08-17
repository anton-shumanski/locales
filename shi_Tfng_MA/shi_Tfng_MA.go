package shi_Tfng_MA

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type shi_Tfng_MA struct {
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

// New returns a new instance of translator for the 'shi_Tfng_MA' locale
func New() locales.Translator {
	return &shi_Tfng_MA{
		locale:             "shi_Tfng_MA",
		pluralsCardinal:    []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:     nil,
		decimal:            []byte{0x2c},
		group:              []byte{0xc2, 0xa0},
		minus:              []byte{},
		percent:            []byte{},
		perMille:           []byte{},
		timeSeparator:      []byte{0x3a},
		currencies:         [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		monthsAbbreviated:  [][]uint8{[]uint8(nil), {0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x8f, 0xe2, 0xb5, 0x8f}, {0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x95, 0xe2, 0xb4, 0xb0}, {0xe2, 0xb5, 0x8e, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x95}, {0xe2, 0xb5, 0x89, 0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x94}, {0xe2, 0xb5, 0x8e, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa2}, {0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x8f}, {0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x8d}, {0xe2, 0xb5, 0x96, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x9b}, {0xe2, 0xb5, 0x9b, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x9c}, {0xe2, 0xb4, 0xbd, 0xe2, 0xb5, 0x9c, 0xe2, 0xb5, 0x93}, {0xe2, 0xb5, 0x8f, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0xa1}, {0xe2, 0xb4, 0xb7, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x8a}},
		monthsNarrow:       [][]uint8{[]uint8(nil), {0xe2, 0xb5, 0x89}, {0xe2, 0xb4, 0xb1}, {0xe2, 0xb5, 0x8e}, {0xe2, 0xb5, 0x89}, {0xe2, 0xb5, 0x8e}, {0xe2, 0xb5, 0xa2}, {0xe2, 0xb5, 0xa2}, {0xe2, 0xb5, 0x96}, {0xe2, 0xb5, 0x9b}, {0xe2, 0xb4, 0xbd}, {0xe2, 0xb5, 0x8f}, {0xe2, 0xb4, 0xb7}},
		monthsWide:         [][]uint8{[]uint8(nil), {0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x8f, 0xe2, 0xb5, 0x8f, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x94}, {0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x95, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x95}, {0xe2, 0xb5, 0x8e, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x95, 0xe2, 0xb5, 0x9a}, {0xe2, 0xb5, 0x89, 0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x94, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x94}, {0xe2, 0xb5, 0x8e, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93}, {0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x8f, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93}, {0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x8d, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0xa3}, {0xe2, 0xb5, 0x96, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x9b, 0xe2, 0xb5, 0x9c}, {0xe2, 0xb5, 0x9b, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x9c, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x8f, 0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x94}, {0xe2, 0xb4, 0xbd, 0xe2, 0xb5, 0x9c, 0xe2, 0xb5, 0x93, 0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x94}, {0xe2, 0xb5, 0x8f, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0xa1, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x8f, 0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x94}, {0xe2, 0xb4, 0xb7, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x8a, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x8f, 0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x94}},
		daysAbbreviated:    [][]uint8{{0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb4, 0xb0}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x8f}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb5, 0x89}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb4, 0xbd, 0xe2, 0xb5, 0x95}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb4, 0xbd, 0xe2, 0xb5, 0xa1}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x8e}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb5, 0x89, 0xe2, 0xb4, 0xb9}},
		daysWide:           [][]uint8{{0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x8e, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x8f, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x8f, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb4, 0xbd, 0xe2, 0xb5, 0x95, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb4, 0xbd, 0xe2, 0xb5, 0xa1, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}, {0xe2, 0xb5, 0x99, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x8e, 0xe2, 0xb5, 0xa1, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb5, 0x89, 0xe2, 0xb4, 0xb9, 0xe2, 0xb5, 0xa2, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}},
		periodsAbbreviated: [][]uint8{{0xe2, 0xb5, 0x9c, 0xe2, 0xb5, 0x89, 0xe2, 0xb4, 0xbc, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa1, 0xe2, 0xb5, 0x9c}, {0xe2, 0xb5, 0x9c, 0xe2, 0xb4, 0xb0, 0xe2, 0xb4, 0xb7, 0xe2, 0xb4, 0xb3, 0xe2, 0xb4, 0xb3, 0xe2, 0xb5, 0xaf, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x9c}},
		periodsWide:        [][]uint8{{0xe2, 0xb5, 0x9c, 0xe2, 0xb5, 0x89, 0xe2, 0xb4, 0xbc, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa1, 0xe2, 0xb5, 0x9c}, {0xe2, 0xb5, 0x9c, 0xe2, 0xb4, 0xb0, 0xe2, 0xb4, 0xb7, 0xe2, 0xb4, 0xb3, 0xe2, 0xb4, 0xb3, 0xe2, 0xb5, 0xaf, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x9c}},
		erasAbbreviated:    [][]uint8{{0xe2, 0xb4, 0xb7, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x84}, {0xe2, 0xb4, 0xb7, 0xe2, 0xb4, 0xbc, 0xe2, 0xb5, 0x84}},
		erasNarrow:         [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:           [][]uint8{{0xe2, 0xb4, 0xb7, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x9c, 0x20, 0xe2, 0xb5, 0x8f, 0x20, 0xe2, 0xb5, 0x84, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x99, 0xe2, 0xb4, 0xb0}, {0xe2, 0xb4, 0xb7, 0xe2, 0xb4, 0xbc, 0xe2, 0xb4, 0xbc, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x94, 0x20, 0xe2, 0xb5, 0x8f, 0x20, 0xe2, 0xb5, 0x84, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x99, 0xe2, 0xb4, 0xb0}},
	}
}

// Locale returns the current translators string locale
func (shi *shi_Tfng_MA) Locale() string {
	return shi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'shi_Tfng_MA'
func (shi *shi_Tfng_MA) PluralsCardinal() []locales.PluralRule {
	return shi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'shi_Tfng_MA'
func (shi *shi_Tfng_MA) PluralsOrdinal() []locales.PluralRule {
	return shi.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'shi_Tfng_MA'
func (shi *shi_Tfng_MA) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	} else if n >= 2 && n <= 10 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'shi_Tfng_MA'
func (shi *shi_Tfng_MA) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'shi_Tfng_MA'
func (shi *shi_Tfng_MA) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'shi_Tfng_MA' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'shi_Tfng_MA' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (shi *shi_Tfng_MA) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := shi.currencies[currency]
	l := len(s) + len(shi.decimal) + len(shi.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, shi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(shi.group) - 1; j >= 0; j-- {
					b = append(b, shi.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(shi.minus) - 1; j >= 0; j-- {
			b = append(b, shi.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, shi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'shi_Tfng_MA'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := shi.currencies[currency]
	l := len(s) + len(shi.decimal) + len(shi.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, shi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(shi.group) - 1; j >= 0; j-- {
					b = append(b, shi.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(shi.minus) - 1; j >= 0; j-- {
			b = append(b, shi.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, shi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, symbol...)
	} else {

		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, shi.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, shi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, shi.daysWide[t.Day()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, shi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}
