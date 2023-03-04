package singleton

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"
)

// File ...
type File struct {
	Open time.Time
	Data map[string]string
}

var globalFile *File
var lock sync.Mutex

// ReadData ...
func ReadData() (*File, error) {
	lock.Lock()
	defer lock.Unlock()
	if globalFile == nil {
		data, err := ioutil.ReadFile("./data.json")
		if err != nil {
			return nil, err
		}

		var mapData map[string]string
		err = json.Unmarshal(data, &mapData)
		if err != nil {
			return nil, err
		}
		globalFile = &File{
			Open: time.Now(),
			Data: mapData,
		}
	}
	return globalFile, nil
}
