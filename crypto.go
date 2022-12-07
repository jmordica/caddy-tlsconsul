package storageconsul

import (
	"encoding/json"

	"github.com/pteich/errors"
)

func (cs *ConsulStorage) EncryptStorageData(data *StorageData) ([]byte, error) {
	// JSON marshal, then encrypt if key is there
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal")
	}

	// Prefix with simple prefix and then encrypt
	bytes = append([]byte(cs.ValuePrefix), bytes...)
	return bytes, err
}

func (cs *ConsulStorage) DecryptStorageData(bytes []byte) (*StorageData, error) {

	// Simple sanity check of the beginning of the byte array just to check
	if len(bytes) < len(cs.ValuePrefix) || string(bytes[:len(cs.ValuePrefix)]) != cs.ValuePrefix {
		return nil, errors.New("invalid data format")
	}

	// Now just json unmarshal
	data := &StorageData{}
	if err := json.Unmarshal(bytes[len(cs.ValuePrefix):], data); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal result")
	}
	return data, nil
}
