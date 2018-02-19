package lruuasurfer

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"time"

	"github.com/avct/uasurfer"
	"github.com/vodilov/lru"
)

type LruUaSurfer struct {
	CacheLru lru.Cache
}

//New return new lru-uasurfer
func New() *LruUaSurfer {
	var s LruUaSurfer
	s.CacheLru.MaxEntries = 1000000
	return &s
}

func (d *LruUaSurfer) LoadDump(filename string) {
	fmt.Println("Load data from cache")
	t1 := time.Now()
	if filename == "" {
		filename = "cache.dump"
	}

	fileHandle, err := os.Open(filename)
	if err != nil {
		return
	}
	defer fileHandle.Close()

	dec := gob.NewDecoder(fileHandle)
	err = dec.Decode(&d)
	if err != nil {
		panic(err)
	}

	t2 := time.Now()
	fmt.Println("Loading time:", t2.Sub(t1))
}

func (d *LruUaSurfer) SaveDump(filename string) {
	fmt.Println("Dump cached data")

	t1 := time.Now()
	if filename == "" {
		filename = "cache.dump"
	}
	fileDump, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer fileDump.Close()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(d)
	if err != nil {
		panic(err)
	}

	n, err := fileDump.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}

	t2 := time.Now()
	fmt.Printf("Dumping %v bytes time %v, ", n, t2.Sub(t1))
}

func (d *LruUaSurfer) Parse(ua string) *uasurfer.UserAgent {
	if val, ok := d.CacheLru.Get(ua); ok {
		return val.(*uasurfer.UserAgent)
	}

	result := uasurfer.Parse(ua)
	d.CacheLru.Add(ua, result)

	return result
}
