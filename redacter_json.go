package redacter

import (
	"bytes"
	"encoding/json"
	"io"
)

// JSON redacts an object in json format
func (r *Redacter) JSON(v any) (string, error) {
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
	decoder := json.NewDecoder(&buff)

	if err := encoder.Encode(v); err != nil {
		return "", err
	}

	var data any
	if err := decoder.Decode(&data); err != nil {
		return "", err
	}

	data = r.redact(data)

	buff.Reset()

	if err := encoder.Encode(data); err != nil {
		return "", err
	}

	res := buff.Bytes()
	last := len(res) - 1

	if res[last] == '\n' {
		res = res[:last]
	}

	return string(res), nil
}

// JSONRaw redacts a byte slice in json format
func (r *Redacter) JSONRaw(raw []byte) (string, error) {
	var data any
	if err := json.Unmarshal(raw, &data); err != nil {
		return "", err
	}

	data = r.redact(data)

	res, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

// JSONRader redacts a reader in json format
func (r *Redacter) JSONRader(reader io.Reader) (string, error) {
	var data any
	if err := json.NewDecoder(reader).Decode(&data); err != nil {
		return "", err
	}

	data = r.redact(data)

	res, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
