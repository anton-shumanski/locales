package nnh_CM

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type nnh_CM struct {
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

// New returns a new instance of translator for the 'nnh_CM' locale
func New() locales.Translator {
	return &nnh_CM{
		locale:                 "nnh_CM",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{},
		percent:                []byte{0x25},
		perMille:               []byte{},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0x45, 0x55, 0x52}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x74, 0x73, 0x65, 0x74, 0x73, 0xc9, 0x9b, 0xcc, 0x80, 0xc9, 0x9b, 0x20, 0x6c, 0xc3, 0xb9, 0x6d}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6b, 0xc3, 0xa0, 0x67, 0x20, 0x6e, 0x67, 0x77, 0xc3, 0xb3, 0xc5, 0x8b}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6c, 0x65, 0x70, 0x79, 0xc3, 0xa8, 0x20, 0x73, 0x68, 0xc3, 0xba, 0x6d}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x63, 0xc3, 0xbf, 0xc3, 0xb3}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x74, 0x73, 0xc9, 0x9b, 0xcc, 0x80, 0xc9, 0x9b, 0x20, 0x63, 0xc3, 0xbf, 0xc3, 0xb3}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6e, 0x6a, 0xc3, 0xbf, 0x6f, 0x6c, 0xc3, 0xa1, 0xca, 0xbc}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x74, 0x79, 0xc9, 0x9b, 0xcc, 0x80, 0x62, 0x20, 0x74, 0x79, 0xc9, 0x9b, 0xcc, 0x80, 0x62, 0x20, 0x6d, 0x62, 0xca, 0x89, 0xcc, 0x80, 0xc5, 0x8b}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6d, 0x62, 0xca, 0x89, 0xcc, 0x80, 0xc5, 0x8b}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6e, 0x67, 0x77, 0xc9, 0x94, 0xcc, 0x80, 0xca, 0xbc, 0x20, 0x6d, 0x62, 0xc3, 0xbf, 0xc9, 0x9b}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x74, 0xc3, 0xa0, 0xc5, 0x8b, 0x61, 0x20, 0x74, 0x73, 0x65, 0x74, 0x73, 0xc3, 0xa1, 0xca, 0xbc}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6d, 0x65, 0x6a, 0x77, 0x6f, 0xc5, 0x8b, 0xc3, 0xb3}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6c, 0xc3, 0xb9, 0x6d}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x74, 0x73, 0x65, 0x74, 0x73, 0xc9, 0x9b, 0xcc, 0x80, 0xc9, 0x9b, 0x20, 0x6c, 0xc3, 0xb9, 0x6d}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6b, 0xc3, 0xa0, 0x67, 0x20, 0x6e, 0x67, 0x77, 0xc3, 0xb3, 0xc5, 0x8b}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6c, 0x65, 0x70, 0x79, 0xc3, 0xa8, 0x20, 0x73, 0x68, 0xc3, 0xba, 0x6d}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x63, 0xc3, 0xbf, 0xc3, 0xb3}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x74, 0x73, 0xc9, 0x9b, 0xcc, 0x80, 0xc9, 0x9b, 0x20, 0x63, 0xc3, 0xbf, 0xc3, 0xb3}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6e, 0x6a, 0xc3, 0xbf, 0x6f, 0x6c, 0xc3, 0xa1, 0xca, 0xbc}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x74, 0x79, 0xc9, 0x9b, 0xcc, 0x80, 0x62, 0x20, 0x74, 0x79, 0xc9, 0x9b, 0xcc, 0x80, 0x62, 0x20, 0x6d, 0x62, 0xca, 0x89, 0xcc, 0x80, 0xc5, 0x8b}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6d, 0x62, 0xca, 0x89, 0xcc, 0x80, 0xc5, 0x8b}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6e, 0x67, 0x77, 0xc9, 0x94, 0xcc, 0x80, 0xca, 0xbc, 0x20, 0x6d, 0x62, 0xc3, 0xbf, 0xc9, 0x9b}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x74, 0xc3, 0xa0, 0xc5, 0x8b, 0x61, 0x20, 0x74, 0x73, 0x65, 0x74, 0x73, 0xc3, 0xa1, 0xca, 0xbc}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6d, 0x65, 0x6a, 0x77, 0x6f, 0xc5, 0x8b, 0xc3, 0xb3}, {0x73, 0x61, 0xc5, 0x8b, 0x20, 0x6c, 0xc3, 0xb9, 0x6d}},
		daysAbbreviated:        [][]uint8{{0x6c, 0x79, 0xc9, 0x9b, 0xca, 0xbc, 0xc9, 0x9b, 0xcc, 0x81, 0x20, 0x73, 0xe1, 0xba, 0x85, 0xc3, 0xad, 0xc5, 0x8b, 0x74, 0xc3, 0xa8}, {0x6d, 0x76, 0x66, 0xc3, 0xb2, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x6d, 0x62, 0xc9, 0x94, 0xcc, 0x81, 0xc9, 0x94, 0x6e, 0x74, 0xc3, 0xa8, 0x20, 0x6d, 0x76, 0x66, 0xc3, 0xb2, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x74, 0x73, 0xc3, 0xa8, 0x74, 0x73, 0xc9, 0x9b, 0xcc, 0x80, 0xc9, 0x9b, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x6d, 0x62, 0xc9, 0x94, 0xcc, 0x81, 0xc9, 0x94, 0x6e, 0x74, 0xc3, 0xa8, 0x20, 0x74, 0x73, 0x65, 0x74, 0x73, 0xc9, 0x9b, 0xcc, 0x80, 0xc9, 0x9b, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x6d, 0x76, 0x66, 0xc3, 0xb2, 0x20, 0x6d, 0xc3, 0xa0, 0x67, 0x61, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x6d, 0xc3, 0xa0, 0x67, 0x61, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}},
		daysShort:              [][]uint8{{0x6c, 0x79, 0xc9, 0x9b, 0xca, 0xbc, 0xc9, 0x9b, 0xcc, 0x81, 0x20, 0x73, 0xe1, 0xba, 0x85, 0xc3, 0xad, 0xc5, 0x8b, 0x74, 0xc3, 0xa8}, {0x6d, 0x76, 0x66, 0xc3, 0xb2, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x6d, 0x62, 0xc9, 0x94, 0xcc, 0x81, 0xc9, 0x94, 0x6e, 0x74, 0xc3, 0xa8, 0x20, 0x6d, 0x76, 0x66, 0xc3, 0xb2, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x74, 0x73, 0xc3, 0xa8, 0x74, 0x73, 0xc9, 0x9b, 0xcc, 0x80, 0xc9, 0x9b, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x6d, 0x62, 0xc9, 0x94, 0xcc, 0x81, 0xc9, 0x94, 0x6e, 0x74, 0xc3, 0xa8, 0x20, 0x74, 0x73, 0x65, 0x74, 0x73, 0xc9, 0x9b, 0xcc, 0x80, 0xc9, 0x9b, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x6d, 0x76, 0x66, 0xc3, 0xb2, 0x20, 0x6d, 0xc3, 0xa0, 0x67, 0x61, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x6d, 0xc3, 0xa0, 0x67, 0x61, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}},
		daysWide:               [][]uint8{{0x6c, 0x79, 0xc9, 0x9b, 0xca, 0xbc, 0xc9, 0x9b, 0xcc, 0x81, 0x20, 0x73, 0xe1, 0xba, 0x85, 0xc3, 0xad, 0xc5, 0x8b, 0x74, 0xc3, 0xa8}, {0x6d, 0x76, 0x66, 0xc3, 0xb2, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x6d, 0x62, 0xc9, 0x94, 0xcc, 0x81, 0xc9, 0x94, 0x6e, 0x74, 0xc3, 0xa8, 0x20, 0x6d, 0x76, 0x66, 0xc3, 0xb2, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x74, 0x73, 0xc3, 0xa8, 0x74, 0x73, 0xc9, 0x9b, 0xcc, 0x80, 0xc9, 0x9b, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x6d, 0x62, 0xc9, 0x94, 0xcc, 0x81, 0xc9, 0x94, 0x6e, 0x74, 0xc3, 0xa8, 0x20, 0x74, 0x73, 0x65, 0x74, 0x73, 0xc9, 0x9b, 0xcc, 0x80, 0xc9, 0x9b, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x6d, 0x76, 0x66, 0xc3, 0xb2, 0x20, 0x6d, 0xc3, 0xa0, 0x67, 0x61, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}, {0x6d, 0xc3, 0xa0, 0x67, 0x61, 0x20, 0x6c, 0x79, 0xc9, 0x9b, 0xcc, 0x8c, 0xca, 0xbc}},
		periodsAbbreviated:     [][]uint8{{0x6d, 0x62, 0x61, 0xca, 0xbc, 0xc3, 0xa1, 0x6d, 0x62, 0x61, 0xca, 0xbc}, {0x6e, 0x63, 0x77, 0xc3, 0xb2, 0x6e, 0x7a, 0xc3, 0xa9, 0x6d}},
		periodsWide:            [][]uint8{{0x6d, 0x62, 0x61, 0xca, 0xbc, 0xc3, 0xa1, 0x6d, 0x62, 0x61, 0xca, 0xbc}, {0x6e, 0x63, 0x77, 0xc3, 0xb2, 0x6e, 0x7a, 0xc3, 0xa9, 0x6d}},
		erasAbbreviated:        [][]uint8{{0x6d, 0x2e, 0x7a, 0x2e, 0x59, 0x2e}, {0x6d, 0x2e, 0x67, 0x2e, 0x6e, 0x2e, 0x59, 0x2e}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x6d, 0xc3, 0xa9, 0x20, 0x7a, 0x79, 0xc3, 0xa9, 0x20, 0x59, 0xc4, 0x9b, 0x73, 0xc3, 0xb4}, {0x6d, 0xc3, 0xa9, 0x20, 0x67, 0xc3, 0xbf, 0x6f, 0x20, 0xc5, 0x84, 0x7a, 0x79, 0xc3, 0xa9, 0x20, 0x59, 0xc4, 0x9b, 0x73, 0xc3, 0xb4}},
	}
}

// Locale returns the current translators string locale
func (nnh *nnh_CM) Locale() string {
	return nnh.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'nnh_CM'
func (nnh *nnh_CM) PluralsCardinal() []locales.PluralRule {
	return nnh.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'nnh_CM'
func (nnh *nnh_CM) PluralsOrdinal() []locales.PluralRule {
	return nnh.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'nnh_CM'
func (nnh *nnh_CM) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'nnh_CM'
func (nnh *nnh_CM) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'nnh_CM'
func (nnh *nnh_CM) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'nnh_CM' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nnh *nnh_CM) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(nnh.decimal) + len(nnh.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nnh.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nnh.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(nnh.minus) - 1; j >= 0; j-- {
			b = append(b, nnh.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'nnh_CM' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (nnh *nnh_CM) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'nnh_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nnh *nnh_CM) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nnh.currencies[currency]
	l := len(s) + len(nnh.decimal) + len(nnh.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nnh.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nnh.group[0])
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

	for j := len(nnh.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, nnh.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(nnh.minus) - 1; j >= 0; j-- {
			b = append(b, nnh.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nnh.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'nnh_CM'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nnh *nnh_CM) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nnh.currencies[currency]
	l := len(s) + len(nnh.decimal) + len(nnh.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nnh.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nnh.group[0])
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

		for j := len(nnh.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, nnh.currencyNegativePrefix[j])
		}

		for j := len(nnh.minus) - 1; j >= 0; j-- {
			b = append(b, nnh.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(nnh.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, nnh.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nnh.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'nnh_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nnh *nnh_CM) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'nnh_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nnh *nnh_CM) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, nnh.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'nnh_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nnh *nnh_CM) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, []byte{0x27, 0x6c, 0x79, 0xc9, 0x9b, 0x27, 0xcc, 0x8c, 0xca, 0xbc, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x6e, 0x61, 0x27, 0x20}...)
	b = append(b, nnh.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'nnh_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nnh *nnh_CM) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, nnh.daysWide[t.Day()]...)
	b = append(b, []byte{0x20, 0x2c}...)
	b = append(b, []byte{0x27, 0x6c, 0x79, 0xc9, 0x9b, 0x27, 0xcc, 0x8c, 0xca, 0xbc, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x6e, 0x61, 0x27, 0x20}...)
	b = append(b, nnh.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'nnh_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nnh *nnh_CM) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'nnh_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nnh *nnh_CM) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'nnh_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nnh *nnh_CM) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'nnh_CM'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nnh *nnh_CM) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}
