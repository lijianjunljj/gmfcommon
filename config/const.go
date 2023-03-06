package config

const (
	//OauthServerURL 鉴权认证地址
	OauthServerURL = "http://127.0.0.1:8360/baas/bcs/bcs-oauth-provider"
	//ErrorToWechat 错误发送到微信
	ErrorToWechat = "WX"
	//ErrorToSMS 错误发送到短信
	ErrorToSMS = "SMS"
	//ErrorToMail 错误发送到邮箱
	ErrorToMail = "MAIL"
	//ErrorToLog 错误保存到日志
	ErrorToLog = "LOG"

	//GrpcIsTracing 是否开启GRPC请求链路追踪
	GrpcIsTracing = true

	//LogPath 日志存储目录
	LogPath = "./logs"
	//LogFileSuffix 日志文件后缀
	LogFileSuffix = ".log"
	//LogFileNum 日志文件保存份数
	LogFileNum = 30
	//LogSplitHours 日志切割时间间隔
	LogSplitHours = 24

	//MailHost 邮箱主机
	MailHost = "smtp.qq.com"
	//MailPort 邮箱端口
	MailPort = 465
	//MailUser 发件人
	MailUser = "qingmumao@qq.com"
	//MailPass 发件人密码
	MailPass = "qdntsguposnwbefe"
	//MailTo 收件人 多个用,分割
	MailTo = "634203427@qq.com"
	//MailNormalSubject 邮件告警主题
	MailNormalSubject = "【系统告警】项目运行报错"
	//MailRequestSubject 邮件请求告警主题
	MailRequestSubject = "【系统告警】%s请求报错"

	//RequestTimeout 连接超时时间，默认5秒超时时间
	RequestTimeout = 5
	//RequestIsTracing 是否开启HTTP请求链路追踪
	RequestIsTracing = true

	//CodeSuccess 请求成功响应码
	CodeSuccess = "0000"
	//CodeFail 请求失败响应码
	CodeFail = "1001"
	//CodeTokenExpired 登录失效响应码
	CodeTokenExpired = "1002"
	//CodeException 系统异常响应码
	CodeException = "1003"
	//CodeNotAdmin非管理员
	CodeNotAdmin = "1004"
	//ResponseCode 请求响应代码字段名
	ResponseCode = "code"
	//ResponseMsg 请求响应代码字段名
	ResponseMsg = "msg"
	//ResponseData 请求响应代码字段名
	ResponseData = "data"

	//SignKey 加签密钥
	SignKey = "IgkibX71IEf382PT"
	//SignRsaPublicFile RSA加签公钥
	SignRsaPublicFile = "rsa/public.pem"
	//SignRsaPrivateFile RSA加签私钥
	SignRsaPrivateFile = "rsa/private.pem"
	//SignExpiry 签名超时时间,单位s
	SignExpiry = "120"
	//SignTypeRsa RSA签名类型名称
	SignTypeRsa = "RSA"
	//SignTypeAes AES签名类型名称
	SignTypeAes = "AES"
	//SignTypeMd5 MD5签名类型名称
	SignTypeMd5 = "MD5"

	//DateTimeFormat 年月日时分秒时间格式
	DateTimeFormat = "2006/01/02 15:04:05"

	//FileUploadFormat 文件上传按时间划分目录的时间格式
	FileUploadFormat = "200601"
	//FileUploadRoot 文件上传根目录名
	FileUploadRoot = "upload"
	//FileUploadThumb 文件上传图片缩略图存储目录名
	FileUploadThumb = "thumbnails"
	//CDN加速地址 使用301重定向 2022-04-28
	FileUrlOfCDN = ""
	//TokenSecret Token加密密钥
	TokenSecret = "ABCDEFGHIJK"
	//TokenExp Token有效期
	TokenExp = "2h"
	//TokenMethod Token加密算法
	TokenMethod = "SHA1"

	//SmsServerURL 短信服务地址
	SmsServerURL = "http://139.159.195.164:8080/edeeserver/sendSMS.do"
	//SmsUserID 短信用户ID
	SmsUserID = "901422"
	//SmsPassword 短信用户密码
	SmsPassword = "888888"
	//SmsMethod 短信发送调用方法
	SmsMethod = "sendSms2"
	//SmsTemplate 短信内容模板
	SmsTemplate = "您注册的BaaS系统验证码是：%v，5分钟内有效！"
)
