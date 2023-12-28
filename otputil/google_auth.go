package otputil

import (
	"encoding/base32"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

type googleAuth struct {
	period    time.Duration // 有效周期时间
	digits    uint          // 密码字数
	algorithm string
}

func NewGoogleAuth() ITOTPAuth {
	return &googleAuth{
		period:    30 * time.Second,
		digits:    6,
		algorithm: "sha1",
	}
}

// GetQRCodeContent 获取绑定二维码
// @see https://github.com/google/google-authenticator/wiki/Key-Uri-Format
// 可使用 Google 身份验证器（[安卓](https://play.google.com/store/apps/details?id=com.google.android.apps.authenticator2)）、
// [腾讯身份验证器](https://sj.qq.com/appdetail/com.tencent.authenticator) 等 TOTP 的验证设备
func (ga googleAuth) GetQRCodeContent(issuer, account, secret string) string {
	q := url.Values{}
	q.Set("secret", secret)

	label := url.PathEscape(account)
	if issuer != "" {
		//label = url.PathEscape(issuer) + ":" + label
		q.Set("issuer", url.QueryEscape(issuer))
	}
	if ga.algorithm != "" && ga.algorithm != "sha1" {
		q.Set("algorithm", strings.ToUpper(ga.algorithm))
	}
	if ga.digits > 6 {
		q.Set("digits", fmt.Sprintf("%d", ga.digits))
	}
	q.Set("period", fmt.Sprintf("%d", int64(ga.period.Seconds())))
	u := url.URL{
		Scheme:   "otpauth",
		Host:     "totp",
		Path:     label,
		RawQuery: q.Encode(),
	}
	return u.String()
}

func (ga googleAuth) GenSecret(length int) string {
	if length <= 0 {
		length = 16
	}
	secret := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	rand.Read(secret)

	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(secret)
}

func (ga googleAuth) GetCode(secret string) string {
	return ga.GetCodeAt(secret, time.Now())
}

func (ga googleAuth) GetCodeAt(secret string, t time.Time) string {
	secretBs, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		log.Printf("[otputil] decode secret failed, err: %s", err)
		return ""
	}
	counter := uint64(t.Unix() / 30)
	return hotp(secretBs, counter, int(ga.digits))
}

func (ga googleAuth) Verify(secret, code string) bool {
	return ga.VerifyAt(secret, code, time.Now())
}

func (ga googleAuth) VerifyAt(secret, code string, t time.Time) bool {
	secretBs, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		log.Printf("[otputil] decode secret failed, err: %s", err)
		return false
	}
	counter := uint64(math.Floor(float64(t.Unix()) / ga.period.Seconds()))
	return hotp(secretBs, counter, int(ga.digits)) == code
}
