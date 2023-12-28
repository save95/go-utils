package otputil

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
)

func hotp(secret []byte, counter uint64, digits int) string {
	buf := make([]byte, 8)
	mac := hmac.New(sha1.New, secret)
	binary.BigEndian.PutUint64(buf, counter)

	mac.Write(buf)
	sum := mac.Sum(nil)

	// "Dynamic truncation" in RFC 4226
	// http://tools.ietf.org/html/rfc4226#section-5.4
	offset := sum[len(sum)-1] & 0xF
	value := int64(((int(sum[offset]) & 0x7F) << 24) |
		((int(sum[offset+1] & 0xFF)) << 16) |
		((int(sum[offset+2] & 0xFF)) << 8) |
		(int(sum[offset+3]) & 0xFF))

	code := int32(value % int64(math.Pow10(digits)))

	return fmt.Sprintf("%06d", code)
}
