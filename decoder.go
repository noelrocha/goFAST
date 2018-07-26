package fast

import (
	"fmt"
)

const maxLoadBytes = (32 << (^uint(0) >> 63)) * 8 / 7 // max size of 7-th bits data

type Decoder struct {
	repo map[uint]*Template

	data []byte
	current *pmap

	Debug bool
}

func NewDecoder(tmps ...*Template) *Decoder {
	decoder := &Decoder{repo: make(map[uint]*Template)}
	for _, t := range tmps {
		decoder.repo[t.ID] = t
	}
	return decoder
}

func (d *Decoder) Decode(segment []byte, msg interface{}) {
	d.data = segment

	d.log("data: ", utoi(d.data))

	if !d.parsePmap() {
		d.skipTail()
		d.log("tail: ", utoi(d.data))
	}

	d.log("pmap: ", utoi(d.data), *d.current)

	if d.current.isNextBitSet() {
		templateID := d.parseTemplateID()
		d.log("template: ", utoi(d.data), templateID)


	}


	return
}

func (d *Decoder) parsePmap() bool {
	d.current = new(pmap)
	d.current.mask = 1;
	for i:=0; i < maxLoadBytes; i++ {
		d.current.bitmap <<= 7
		d.current.bitmap |= uint(d.data[i]) & '\x7F'
		d.current.mask <<= 7

		if ('\x80' == (d.data[i] & '\x80')) {
			d.data = d.data[i+1:]
			return true;
		}
	}

	return false
}

func (d *Decoder) skipTail() {
	for i, b := range d.data {
		if 0 == (b & 0x80) {
			d.data = d.data[i+1:]
			return
		}
	}
}

func (d *Decoder) parseTemplateID() uint32 {
	i := 0
	id := uint32(d.data[i] & 0x7F)

	for (d.data[i] & 0x80) == 0 {
		id <<= 7;
		i++
		id |= uint32(d.data[i] & 0x7F);
	}

	d.data = d.data[i+1:]
	return id
}

func (d *Decoder) log(a ...interface{}) {
	if d.Debug {
		fmt.Println(a...)
	}
}

// -----------

type pmap struct {
	bitmap uint
	mask uint
}

func (p *pmap) isNextBitSet() bool {
	p.mask >>= 1
	return (p.bitmap & p.mask) != 0;
}

// ------------

func utoi(d []byte) (r []int8) {
	for _, b := range d {
		r = append(r, int8(b))
	}
	return
}