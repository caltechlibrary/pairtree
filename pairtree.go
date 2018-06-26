//
// pairtree.go implements encoding/decoding of object identifiers and pairtree paths (ppaths) per
// https://confluence.ucop.edu/download/attachments/14254128/PairtreeSpec.pdf?version=2&modificationDate=1295552323000&api=v2
//
// Author R. S. Doiel, <rsdoiel@library.caltech.edu>
//
// Copyright (c) 2018, Caltech
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
	"strings"
)

var (
	stepOneEncoding = map[string]string{
		" ":  "^20",
		"\"": "^22",
		"<":  "^3c",
		"\\": "^5c",
		"*":  "^2a",
		"=":  "^3d",
		"^":  "^5e",
		"+":  "^2b",
		">":  "^3e",
		"|":  "^7c",
		",":  "^2c",
		"?":  "^3f",
	}
	stepTwoEncoding = map[string]string{
		"/": "=",
		":": "+",
		".": ",",
	}
)

func charEncode(s string) string {
	//NOTE: we need to replace ^ with ^5e and avoid collisions with other hex values
	// we split the string into an array of substrings then replace each one as as need to.
	p := strings.Split(s, "")
	for i, target := range p {
		if val, ok := stepOneEncoding[target]; ok == true {
			p[i] = val
		}
	}
	s = strings.Join(p, "")
	for target, replacement := range stepTwoEncoding {
		if strings.Contains(s, target) {
			s = strings.Replace(s, target, replacement, -1)
		}
	}
	return s
}

func charDecode(s string) string {
	for replacement, target := range stepTwoEncoding {
		if strings.Contains(s, target) {
			s = strings.Replace(s, target, replacement, -1)
		}
	}
	for replacement, target := range stepOneEncoding {
		if strings.Contains(s, target) {
			s = strings.Replace(s, target, replacement, -1)
		}
	}
	return s
}

func Encode(src string) string {
	s := charEncode(src)
	results := []string{}
	for i := 0; i < len(s); i += 2 {
		if (i + 2) < len(s) {
			t := s[i : i+2]
			results = append(results, t)
		} else {
			t := s[i:]
			results = append(results, t)
		}
	}
	results = append(results, "")
	return strings.Join(results, string(os.PathSeparator))
}

func Decode(src string) string {
	parts := strings.Split(src, string(os.PathSeparator))
	results := []string{}
	for _, segment := range parts {
		if segment == "obj" {
			break
		}
		if len(segment) > 2 {
			break
		}
		if len(segment) == 1 {
			results = append(results, segment)
			break
		}
		results = append(results, segment)
	}
	return charDecode(strings.Join(results, ""))
}
