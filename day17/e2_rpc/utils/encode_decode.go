package utils

import (
	"bytes"
	"encoding/gob"
)

// RPCData 定义RPC交互的数据结构
type RPCData struct {
	// 访问的函数
	Name string
	// 访问时的参数
	Args []interface{}
}

// Encode 编码
func Encode(data RPCData) ([]byte, error) {
	// 得到字节数组的编码器
	var buf bytes.Buffer
	bufEnc := gob.NewEncoder(&buf)
	// 编码器对数据编码
	if err := bufEnc.Encode(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode 解码
func Decode(b []byte) (RPCData, error) {
	buf := bytes.NewBuffer(b)
	// 得到字节数组解码器
	bufDec := gob.NewDecoder(buf)
	var data RPCData
	if err := bufDec.Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}
