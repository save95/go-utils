package otputil

import "time"

type ITOTPAuth interface {
	// GetQRCodeContent 获取二维码内容
	GetQRCodeContent(account, issuer, secret string) string

	// GenSecret 生成密钥
	GenSecret(length int) string

	// GetCode 获取当前验证码
	GetCode(secret string) string

	// GetCodeAt 获取指定时间的验证码
	GetCodeAt(secret string, t time.Time) string

	// Verify 校验当前时间点验证码是否正确
	Verify(secret, code string) bool

	// VerifyAt 校验指定时间二维码是否正确
	VerifyAt(secret, code string, t time.Time) bool
}
