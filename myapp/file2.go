package main

import "github.com/vmihailenco/msgpack/v5"

func (t *MyType) EncodeMsgpack(enc *msgpack.Encoder) error {
	var err error

	if err = enc.EncodeArrayLen(2); err != nil {
		return err
	}

	if err = enc.EncodeInt(int64(t.a)); err != nil {
		return err
	}

	if err = enc.EncodeInt(int64(t.b)); err != nil {
		return err
	}

	return nil
}
