/**
 * Copyright 2015 @ z3q.net.
 * name : value_rep.go
 * author : jarryliu
 * date : 2016-05-27 15:28
 * description :
 * history :
 */
package valueobject

type (
	// 微信API设置
	WxApiConfig struct {

		/**===== 微信公众平台设置 =====**/

		//APP ID
		AppId string
		//APP 密钥
		AppSecret string
		//通信密钥
		MpToken string
		//通信AES KEY
		MpAesKey string
		//原始ID
		OriId string

		/**===== 用于微信支付 =====**/

		//商户编号
		MchId string
		//商户接口密钥
		MchApiKey string
		//微信支付的证书路径(上传)
		MchCertPath string
		//微信支付的证书公钥路径(上传)
		MchCertKeyPath string

		//MchPayNotifyPath string //微信支付异步通知的路径
	}

	// 注册权限设置
	RegisterPerm struct {
		// 注册模式,等于member.RegisterMode
		RegisterMode int
		// 是否允许匿名注册
		AnonymousRegistered bool
		// 注册提示
		Notice string
		// 用户条款内容
		Licence string
	}

	GlobSaleConf struct {
		// 是否启用分销模式
		FxSalesEnabled bool
		// 返现比例,0则不返现
		CashBackPercent float32
		// 一级比例
		CashBackTg1Percent float32
		// 二级比例
		CashBackTg2Percent float32
		// 会员比例
		CashBackMemberPercent float32
		// 每一元返多少积分
		IntegralBackNum int
		// 每单额外赠送
		IntegralBackExtra int
		// 自动设置订单
		AutoSetupOrder bool
		// 订单超时分钟数
		OrderTimeOutMinute int
		// 订单自动确认时间
		OrderConfirmAfterMinute int
		// 订单超时自动收货
		OrderTimeOutReceiveHour int
		// 提现手续费费率
		ApplyCsn float32
		// 转账手续费费率
		TransCsn float32
		// 活动账户转为赠送可提现奖金手续费费率
		FlowConvertCsn float32
		// 赠送账户转换手续费费率
		PresentConvertCsn float32
	}

	IValueRep interface {
		// 获取微信接口配置
		GetWxApiConfig() *WxApiConfig

		// 保存微信接口配置
		SaveWxApiConfig(v *WxApiConfig) error

		// 获取注册权限
		GetRegisterPerm() *RegisterPerm

		// 保存注册权限
		SaveRegisterPerm(v *RegisterPerm) error
	}
)