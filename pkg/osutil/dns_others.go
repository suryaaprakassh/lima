//go:build !darwin

// SPDX-FileCopyrightText: Copyright The Lima Authors
// SPDX-License-Identifier: Apache-2.0

package osutil

func DNSAddresses() ([]string, error) {
	// TODO: parse /etc/resolv.conf?
	return []string{}, nil
}

func ProxySettings() (map[string]string, error) {
	return make(map[string]string), nil
}
