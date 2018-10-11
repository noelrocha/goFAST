// Copyright 2018 Alexander Poltoratskiy. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package fast_test

/*func Test_newMsg(t *testing.T) {
	type Sequense struct {
		Test uint32 `fast:"33"`
	}
	type Msg struct {
		Test uint32 `fast:"11"`
		Seq []Sequense `fast:"22"`
	}

	var msg Msg
	m := newMsg(&msg)

	if v, ok := m.tagMap["11"]; !ok || len(v) != 1 {
		t.Fatal("invalid parsing tag '11'")
	}

	if v, ok := m.tagMap["22"]; !ok || len(v) != 1 {
		t.Fatal("invalid parsing tag '22'")
	}

	if v, ok := m.tagMap["33"]; !ok || len(v) != 2 {
		t.Fatal("invalid parsing tag '33'")
	}
}

func TestMessage_Assign(t *testing.T) {
	type Sequense struct {
		Test uint32 `fast:"33"`
	}
	type Msg struct {
		Test uint32 `fast:"11"`
		Seq []Sequense `fast:"22"`
	}

	var msg Msg
	m := newMsg(&msg)

	fld := &field{
		id: 11,
		value: uint32(1),
	}

	m.Assign(fld)

	if msg.Test != uint32(1) {
		t.Fatal("invalid assigning to 'msg.Test' field")
	}

	fld = &field{
		id: 33,
		value: uint32(2),
	}

	m.AssignSlice(fld, 0)

	if len(msg.Seq) !=1 || msg.Seq[0].Test != uint32(2) {
		t.Fatal("invalid assigning to 'msg.Seq[0].Test' field")
	}
}*/
