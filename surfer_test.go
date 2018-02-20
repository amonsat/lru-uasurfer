package lruuasurfer

import (
	"sync"
	"testing"

	"github.com/avct/uasurfer"
)

func BenchmarkParseUasurfer(b *testing.B) {
	b.ResetTimer()

	num := len(testUAVars)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uasurfer.Parse(testUAVars[i%num].UA)
	}
}

func BenchmarkParseLRU(b *testing.B) {
	s := New()

	b.ResetTimer()

	num := len(testUAVars)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Parse(testUAVars[i%num].UA)
	}
}

func BenchmarkParseMap(b *testing.B) {
	b.ResetTimer()

	num := len(testUAVars)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cacheParse(testUAVars[i%num].UA)
	}
}

var cache sync.Map

func cacheParse(ua string) *uasurfer.UserAgent {
	if val, ok := cache.Load(ua); ok {
		return val.(*uasurfer.UserAgent)
	}

	result := uasurfer.Parse(ua)
	cache.Store(ua, result)

	return result
}

var testUAVars = []struct {
	UA string
	uasurfer.UserAgent
}{
	// Empty
	{"",
		uasurfer.UserAgent{}},

	// Single char
	{"a",
		uasurfer.UserAgent{}},

	// Some random string
	{"some random string",
		uasurfer.UserAgent{}},

	// Potentially malformed ua
	{")(",
		uasurfer.UserAgent{}},
	// iPhone
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/546.10 (KHTML, like Gecko) Version/6.0 Mobile/7E18WD Safari/8536.25",
		uasurfer.UserAgent{
			uasurfer.Browser{uasurfer.BrowserSafari, uasurfer.Version{6, 0, 0}}, uasurfer.OS{uasurfer.PlatformiPhone, uasurfer.OSiOS, uasurfer.Version{7, 0, 0}}, uasurfer.DevicePhone}},
	// Chrome
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 Safari/537.36",
		uasurfer.UserAgent{
			uasurfer.Browser{uasurfer.BrowserChrome, uasurfer.Version{43, 0, 2357}}, uasurfer.OS{uasurfer.PlatformMac, uasurfer.OSMacOSX, uasurfer.Version{10, 10, 4}}, uasurfer.DeviceComputer}},
	// Safari
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12",
		uasurfer.UserAgent{
			uasurfer.Browser{uasurfer.BrowserSafari, uasurfer.Version{8, 0, 7}}, uasurfer.OS{uasurfer.PlatformMac, uasurfer.OSMacOSX, uasurfer.Version{10, 10, 4}}, uasurfer.DeviceComputer}},
}
