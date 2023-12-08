package evolution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringGeneIsNumeric(t *testing.T) {
	for i := byte(0x0); i < 0xA; i++ {
		g := NewStringGeneFromByte(i)
		assert.True(t, g.IsNumeric(), fmt.Sprintf("%04b (%d) IsNumeric() -> True", i, i))
	}
	for i := byte(0xA); i < 0xE; i++ {
		g := NewStringGeneFromByte(i)
		assert.False(t, g.IsNumeric(), fmt.Sprintf("%04b (%d) IsNumeric() -> False", i, i))
	}
}

func TestStringGeneIsOperator(t *testing.T) {
	for i := byte(0x0); i < 0xA; i++ {
		g := NewStringGeneFromByte(i)
		assert.False(t, g.IsOperator(), fmt.Sprintf("%04b (%d) IsOperator() -> False", i, i))
	}
	for i := byte(0xA); i < 0xE; i++ {
		g := NewStringGeneFromByte(i)
		assert.True(t, g.IsOperator(), fmt.Sprintf("%04b (%d) IsOperator() -> True", i, i))
	}
}

func TestStringGeneDecode(t *testing.T) {
	cases := []struct {
		v   byte
		e   string
		msg string
	}{
		{
			v:   0x0,
			e:   "0",
			msg: "0x0 Decode() -> 0",
		},
		{
			v:   0x1,
			e:   "1",
			msg: "0x1 Decode() -> 1",
		},
		{
			v:   0x2,
			e:   "2",
			msg: "0x2 Decode() -> 2",
		},
		{
			v:   0x3,
			e:   "3",
			msg: "0x3 Decode() -> 3",
		},
		{
			v:   0x4,
			e:   "4",
			msg: "0x4 Decode() -> 4",
		},
		{
			v:   0x5,
			e:   "5",
			msg: "0x5 Decode() -> 5",
		},
		{
			v:   0x6,
			e:   "6",
			msg: "0x6 Decode() -> 6",
		},
		{
			v:   0x7,
			e:   "7",
			msg: "0x7 Decode() -> 7",
		},
		{
			v:   0x8,
			e:   "8",
			msg: "0x8 Decode() -> 8",
		},
		{
			v:   0x9,
			e:   "9",
			msg: "0x9 Decode() -> 9",
		},
		{
			v:   0xA,
			e:   "+",
			msg: "0xA Decode() -> +",
		},
		{
			v:   0xB,
			e:   "-",
			msg: "0xB Decode() -> -",
		},
		{
			v:   0xC,
			e:   "*",
			msg: "0xC Decode() -> *",
		},
		{
			v:   0xD,
			e:   "/",
			msg: "0xD Decode() -> /",
		},
		{
			v:   0xE,
			e:   " ",
			msg: "0xE Decode() -> \" \"",
		},
	}
	for _, c := range cases {
		g := NewStringGeneFromByte(c.v)
		assert.Equal(t, c.e, g.Decode(), c.msg)
	}
}

func TestStringGeneMutate(t *testing.T) {
	g := NewStringGeneFromByte(0x0)
	g.Mutate(float64(1.0))
	assert.NotEqual(t, "0", g.Decode(), "0x0 Mutate(1.0) -> !0")
}

func TestByteGeneIsNumeric(t *testing.T) {
	for i := byte(0x0); i < 0xA; i++ {
		g := NewByteGene(i)
		assert.True(t, g.IsNumeric(), fmt.Sprintf("%04b (%d) IsNumeric() -> True", i, i))
	}
	for i := byte(0xA); i < 0xE; i++ {
		g := NewByteGene(i)
		assert.False(t, g.IsNumeric(), fmt.Sprintf("%04b (%d) IsNumeric() -> False", i, i))
	}
}

func TestByteGeneIsOperator(t *testing.T) {
	for i := byte(0x0); i < 0xA; i++ {
		g := NewByteGene(i)
		assert.False(t, g.IsOperator(), fmt.Sprintf("%04b (%d) IsOperator() -> False", i, i))
	}
	for i := byte(0xA); i < 0xE; i++ {
		g := NewByteGene(i)
		assert.True(t, g.IsOperator(), fmt.Sprintf("%04b (%d) IsOperator() -> True", i, i))
	}
}

func TestByteGeneDecode(t *testing.T) {
	cases := []struct {
		v   byte
		e   string
		msg string
	}{
		{
			v:   0x0,
			e:   "0",
			msg: "0x0 Decode() -> 0",
		},
		{
			v:   0x1,
			e:   "1",
			msg: "0x1 Decode() -> 1",
		},
		{
			v:   0x2,
			e:   "2",
			msg: "0x2 Decode() -> 2",
		},
		{
			v:   0x3,
			e:   "3",
			msg: "0x3 Decode() -> 3",
		},
		{
			v:   0x4,
			e:   "4",
			msg: "0x4 Decode() -> 4",
		},
		{
			v:   0x5,
			e:   "5",
			msg: "0x5 Decode() -> 5",
		},
		{
			v:   0x6,
			e:   "6",
			msg: "0x6 Decode() -> 6",
		},
		{
			v:   0x7,
			e:   "7",
			msg: "0x7 Decode() -> 7",
		},
		{
			v:   0x8,
			e:   "8",
			msg: "0x8 Decode() -> 8",
		},
		{
			v:   0x9,
			e:   "9",
			msg: "0x9 Decode() -> 9",
		},
		{
			v:   0xA,
			e:   "+",
			msg: "0xA Decode() -> +",
		},
		{
			v:   0xB,
			e:   "-",
			msg: "0xB Decode() -> -",
		},
		{
			v:   0xC,
			e:   "*",
			msg: "0xC Decode() -> *",
		},
		{
			v:   0xD,
			e:   "/",
			msg: "0xD Decode() -> /",
		},
		{
			v:   0xE,
			e:   " ",
			msg: "0xE Decode() -> \" \"",
		},
	}
	for _, c := range cases {
		g := NewByteGene(c.v)
		assert.Equal(t, c.e, g.Decode(), c.msg)
	}
}
