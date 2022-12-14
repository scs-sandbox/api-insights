// Copyright 2022 Cisco Systems, Inc. and its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"io"
	"os"
)

// Different returns the difference between two slices
func Different(slice1 []string, slice2 []string) []string {
	diff := make([]string, 0, 10)

	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		// String not found. We add it to return slice
		if !found {
			diff = append(diff, s1)
		}
	}

	return diff
}

// Intersect returns the intersection of two slices
func Intersect(slice1 []string, slice2 []string) []string {
	set := make([]string, 0, 10)

	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		// String not found. We add it to return slice
		if found {
			set = append(set, s1)
		}
	}

	return set
}

// Hash the text using MD5
func Hash(text string) string {
	if text == "" {
		return ""
	}
	h := md5.New()
	_, _ = io.WriteString(h, text)
	hash := h.Sum(nil)
	encoded := base64.StdEncoding.EncodeToString([]byte(string(hash)))
	return encoded
}

// GetEnvironment return the current environment name
func GetEnvironment() string {
	if os.Getenv("AUTH_COOKIE_EXTENSION") == "_stage" {
		return "staging"
	}
	return "production"
}

// BoolPtr returns a pointer to a bool
func BoolPtr(b bool) *bool { return &b }

// StringPtr returns a pointer to the passed string.
func StringPtr(s string) *string { return &s }

// Float32Ptr returns a pointer to the passed float32.
func Float32Ptr(f float32) *float32 { return &f }

// UnmarshalMapInto (json) marshals m & (json) unmarshals it into v.
func UnmarshalMapInto(m map[string]interface{}, v interface{}) error {
	mapBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(mapBytes, v)
}

// GZIP compresses & outputs data in the GZIP format.
// Optional GZIP header metadata can be passed into the compressed data.
func GZIP(data []byte, header *gzip.Header) ([]byte, error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)

	if header != nil {
		w.Comment = header.Comment
		w.Extra = header.Extra
		w.ModTime = header.ModTime
		w.Name = header.Name
		w.OS = header.OS
	}

	if _, err := w.Write(data); err != nil {
		return nil, err
	}

	if err := w.Flush(); err != nil {
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// GUNZIP decompresses & outputs data from the GZIP format.
// Optional GZIP header metadata passed into the compressed data is also returned.
func GUNZIP(data []byte) ([]byte, *gzip.Header, error) {
	r, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, nil, err
	}

	var b bytes.Buffer
	_, err = b.ReadFrom(r)
	if err != nil {
		return nil, nil, err
	}

	return b.Bytes(), &gzip.Header{
		Comment: r.Comment,
		Extra:   r.Extra,
		ModTime: r.ModTime,
		Name:    r.Name,
		OS:      r.OS,
	}, nil
}
