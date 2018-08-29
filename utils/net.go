package utils

import (
	"net"
	"strconv"
	"strings"
)

// Convert uint to net.IP http://www.sharejs.com
func Inet_ntoa(ip uint) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ip & 0xFF)
	bytes[1] = byte((ip >> 8) & 0xFF)
	bytes[2] = byte((ip >> 16) & 0xFF)
	bytes[3] = byte((ip >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

// Convert net.IP to uint
func Inet_aton(ip net.IP) uint {
	bits := strings.Split(ip.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum uint

	sum += uint(b0) << 24
	sum += uint(b1) << 16
	sum += uint(b2) << 8
	sum += uint(b3)

	return sum
}

