//
// pairtree_test.go implements testing for encoding/decoding object identifiers and pairtree paths (ppaths) per
// https://confluence.ucop.edu/download/attachments/14254128/PairtreeSpec.pdf?version=2&modificationDate=1295552323000&api=v2
//
// Author R. S. Doiel, <rsdoiel@library.caltech.edu>
//
// Copyright (c) 2020, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package pairtree

import (
	"os"
	"path"
	"testing"
)

var (
	pathSep = string([]rune{os.PathSeparator})
)

func TestCharEncoding(t *testing.T) {
	//FIXME: test is posix-centric, need to handle other
	// paths delimiters.
	testCharEncoding := map[string]string{
		"ark:/13030/xt12t3":                     "ark+=13030=xt12t3",
		"http://n2t.info/urn:nbn:se:kb:repos-1": "http+==n2t,info=urn+nbn+se+kb+repos-1",
		"what-the-*@?#!^!?":                     "what-the-^2a@^3f#!^5e!^3f",
	}

	for src, expected := range testCharEncoding {
		result := string(charEncode([]rune(src)))
		if result != expected {
			t.Errorf("%q, expected %q, got %q", src, expected, result)
		}
	}
	for expected, src := range testCharEncoding {
		result := charDecode(src)
		if result != expected {
			t.Errorf("%q, expected %q, got %q", src, expected, result)
		}
	}
}

func TestBasic(t *testing.T) {
	// Test Basic encoding
	testEncodings := map[string]string{
		"abcd":       path.Join("ab", "cd") + pathSep,
		"abcdefg":    path.Join("ab", "cd", "ef", "g") + pathSep,
		"12-986xy4":  path.Join("12", "-9", "86", "xy", "4") + pathSep,
		"2018-06-01": path.Join("20", "18", "-0", "6-", "01") + pathSep,
		"a":          path.Join("a") + pathSep,
		"ab":         path.Join("ab") + pathSep,
		"abc":        path.Join("ab", "c") + pathSep,
		"abcde":      path.Join("ab", "cd", "e") + pathSep,
		"mnopqz":     path.Join("mn", "op", "qz") + pathSep,
	}
	for src, expected := range testEncodings {
		result := Encode(src)
		if result != expected {
			t.Errorf("encoding %q, expected %q, got %q", src, expected, result)
		}
	}

	testDecodings := map[string]string{}
	for val, key := range testEncodings {
		testDecodings[key] = val
	}

	// Test Basic decoding
	for src, expected := range testDecodings {
		result := Decode(src)
		if result != expected {
			t.Errorf("decoding %q, expected %q, got %q", src, expected, result)
		}
	}
}

func TestAdvanced(t *testing.T) {
	testData := map[string]string{
		"ark:/13030/xt12t3": path.Join("ar", "k+", "=1", "30", "30",
			"=x", "t1", "2t", "3") + pathSep,
		"http://n2t.info/urn:nbn:se:kb:repos-1": path.Join("ht", "tp",
			"+=", "=n", "2t", ",i", "nf", "o=", "ur", "n+", "nb", "n+",
			"se", "+k", "b+", "re", "po", "s-", "1") + pathSep,
		"what-the-*@?#!^!?": path.Join("wh", "at", "-t", "he", "-^",
			"2a", "@^", "3f", "#!", "^5", "e!", "^3",
			"f") + pathSep,
	}
	for src, expected := range testData {
		result := Encode(src)
		if result != expected {
			t.Errorf("encode %q, expected %q, got %q", src, expected, result)
		}
	}
	for expected, src := range testData {
		result := Decode(src)
		if result != expected {
			t.Errorf("decode %q, expected %q, got %q", src, expected, result)
		}
	}
}

func TestUTF8Names(t *testing.T) {
	testData := map[string]string{
		"Hänggi-P": path.Join("Hä", "ng", "gi", "-P") + pathSep,
	}
	for src, expected := range testData {
		result := Encode(src)
		if result != expected {
			t.Errorf("encode %q, expected %q, got %q", src, expected, result)
		}
	}
	for expected, src := range testData {
		result := Decode(src)
		if result != expected {
			t.Errorf("decode %q, expected %q, got %q", src, expected, result)
		}
	}
}
