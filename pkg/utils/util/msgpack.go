/*
Copyright 2019 Cortex Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	s "github.com/cortexlabs/cortex/pkg/api/strings"
	"github.com/cortexlabs/cortex/pkg/utils/errors"
	"github.com/ugorji/go/codec"
)

var mh codec.MsgpackHandle

func init() {
	mh.RawToString = true
}

func MarshalMsgpack(obj interface{}) ([]byte, error) {
	var bytes []byte
	enc := codec.NewEncoderBytes(&bytes, &mh)
	err := enc.Encode(obj)
	if err != nil {
		return nil, errors.Wrap(err, s.ErrMarshalMsgpack)
	}
	return bytes, nil
}

func MustMarshalMsgpack(obj interface{}) []byte {
	msgpackBytes, err := MarshalMsgpack(obj)
	if err != nil {
		panic(err)
	}
	return msgpackBytes
}

func UnmarshalMsgpackToInterface(b []byte) (interface{}, error) {
	var obj interface{}
	err := UnmarshalMsgpack(b, &obj)
	if err != nil {
		return nil, errors.Wrap(err, s.ErrUnmarshalMsgpack)
	}
	return obj, nil
}

func UnmarshalMsgpack(b []byte, obj interface{}) error {
	dec := codec.NewDecoderBytes(b, &mh)
	return dec.Decode(&obj)
}
