package iputil

import (
	"math"
	"net"

	"github.com/save95/xerror"
)

// IsLocal 检测 IP 地址是否是内网地址
func IsLocal(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}

// ToLong 把 net.IP 转为整数
func ToLong(ip net.IP) (uint, error) {
	b := ip.To4()
	if b == nil {
		return 0, xerror.New("invalid ipv4 format")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

// Parse 将整数解析为 net.IP
func Parse(i uint) (net.IP, error) {
	if i > math.MaxUint32 {
		return nil, xerror.New("beyond the scope of ipv4")
	}

	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip, nil
}
