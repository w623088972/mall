package errno

var (
	OK      = &Errno{Code: 10000, Message: map[string]string{"en": "OK", "chs": "操作成功!", "cht": "操作成功!"}}
	Loading = &Errno{Code: 11111, Message: map[string]string{"en": "程序员小哥哥正在加急开发哦", "chs": "程序员小哥哥正在加急开发哦!", "cht": "程序员小哥哥正在加急开发哦!"}}

	InternalServerError             = &Errno{Code: 10001, Message: map[string]string{"en": "Internal server error", "chs": "系统错误,请稍后重试或联系客服人员", "cht": "系統錯誤,請稍後重試或聯繫客服人員"}}
	ErrDatabase                     = &Errno{Code: 10002, Message: map[string]string{"en": "Database error", "chs": "数据库错误,请稍后重试或联系客服人员", "cht": "數據庫錯誤,請稍後重試或聯繫客服人員"}}
	ErrRedis                        = &Errno{Code: 10003, Message: map[string]string{"en": "Redis error", "chs": "Redis数据库错误,请稍后重试或联系客服人员", "cht": "Redis數據庫錯誤,請稍後重試或聯繫客服人員"}}
	ErrMissingParameter             = &Errno{Code: 10004, Message: map[string]string{"en": "Missing parameter", "chs": "缺少参数,请稍后重试或联系客服人员", "cht": "缺少參數,請稍後重試或聯繫客服人員"}}
	ErrBinding                      = &Errno{Code: 10005, Message: map[string]string{"en": "Request binding error", "chs": "绑定参数错误,请稍后重试或联系客服人员", "cht": "綁定參數錯誤,請稍後重試或聯繫客服人員"}}
	ErrParameterWrong               = &Errno{Code: 10006, Message: map[string]string{"en": "Parameter is wrong", "chs": "请求参数错误,请稍后重试或联系客服人员", "cht": "請求參數錯誤,請稍後重試或聯繫客服人員"}}
	ErrPermissionDenied             = &Errno{Code: 10007, Message: map[string]string{"en": "Permission denied", "chs": "无权限做此操作,请稍后重试或联系客服人员", "cht": "無權限做此操作,請稍後重試或聯繫客服人員"}}
	ErrItemStatusWrong              = &Errno{Code: 10008, Message: map[string]string{"en": "Item status is wrong", "chs": "编辑项状态错误,请稍后重试或联系客服人员", "cht": "編輯項狀態錯誤,請稍後重試或聯繫客服人員"}}
	ErrContentIsPorn                = &Errno{Code: 10009, Message: map[string]string{"en": "Content contains pornography, violence, advertising, infringing content.", "chs": "内容含有色情、暴力、广告、侵权内容", "cht": "內容含有色情、暴力、廣告、侵權內容"}}
	ErrTopicStatusIsErr             = &Errno{Code: 10009, Message: map[string]string{"en": "Topic has been banned.", "chs": "话题已被禁止", "cht": "話題已被禁止"}}
	ErrExpired                      = &Errno{Code: 10010, Message: map[string]string{"en": "expired.", "chs": "已失效", "cht": "已失效"}}
	ErrPostIsExist                  = &Errno{Code: 10011, Message: map[string]string{"en": "expired.", "chs": "已提交，不可重复提交", "cht": "已提交，不可重复提交"}}
	ErrPayStatusIsErr               = &Errno{Code: 10012, Message: map[string]string{"en": "Order timed out, please place an order again", "chs": "订单超时，请重新下单", "cht": "訂單超時，請重新下單"}}
	ErrDeviceAccountIsExist         = &Errno{Code: 10013, Message: map[string]string{"en": "", "chs": "设备账号已存在", "cht": ""}}
	ErrDeviceAccountSiteNameIsExist = &Errno{Code: 10014, Message: map[string]string{"en": "", "chs": "设备账号现场名称已存在", "cht": ""}}
	ErrEntranceNameIsExist          = &Errno{Code: 10015, Message: map[string]string{"en": "", "chs": "入口名称已存在", "cht": ""}}
	ErrEntranceDeviceIsExist        = &Errno{Code: 10016, Message: map[string]string{"en": "", "chs": "入口下已分配设备，无法删除", "cht": "入口下已分配設備，無法刪除"}}
	ErrDeviceAccountIsEmpty         = &Errno{Code: 10017, Message: map[string]string{"en": "", "chs": "设备账号不能为空", "cht": "設備賬號不能為空"}}
	ErrDeviceSiteNameIsExist        = &Errno{Code: 10018, Message: map[string]string{"en": "", "chs": "设备现场名称已存在", "cht": "設備現場名稱已存在"}}
	ErrQrCodeErr                    = &Errno{Code: 10019, Message: map[string]string{"en": "QR code error", "chs": "二维码错误", "cht": "二維碼錯誤"}}
	ErrIsUsed                       = &Errno{Code: 10020, Message: map[string]string{"en": "Used", "chs": "已使用", "cht": "已使用"}}
	ErrAssetsNumberIsExist          = &Errno{Code: 10024, Message: map[string]string{"en": "Asset number already exists", "chs": "资产编号已存在", "cht": "資產編號已存在"}}
	ErrBoothCategoryIsExist         = &Errno{Code: 10025, Message: map[string]string{"en": "Card seat category already exists", "chs": "卡座类别已存在", "cht": "卡座類別已存在"}}
	ErrBoothDeskIsExist             = &Errno{Code: 10026, Message: map[string]string{"en": "Card seat table already exists", "chs": "卡座桌位已存在", "cht": "卡座桌位已存在"}}
	ErrConsumableNumberIsExist      = &Errno{Code: 10027, Message: map[string]string{"en": "Low value consumables No. already exists", "chs": "低值易耗品编号已存在", "cht": "低值易耗品編號已存在"}}
	ErrConsumableNameUnitIsExist    = &Errno{Code: 10028, Message: map[string]string{"en": "Name or unit of low value consumables already exists", "chs": "低值易耗品名称或单位已存在", "cht": "低值易耗品名稱或單位已存在"}}
	ErrUserJurisdictionNot          = &Errno{Code: 10029, Message: map[string]string{"en": "User does not have permission, please reconfirm", "chs": "用户没有权限，请重新确认", "cht": "用戶滅有權限，請重新確認"}}
	ErrInvitationNotEmpty           = &Errno{Code: 10031, Message: map[string]string{"en": "Sorry, this invitation has been used", "chs": "抱歉，该邀请函已被使用", "cht": "抱歉，該邀請函已被使用"}}
	ErrInviteNum                    = &Errno{Code: 10032, Message: map[string]string{"en": "Invitation is full", "chs": "邀请人数已满", "cht": "邀請人數已滿"}}
	ErrInviteIsReceived             = &Errno{Code: 10033, Message: map[string]string{"en": "Invitation received", "chs": "邀请已领取", "cht": "邀請已領取"}}
	ErrInviteIsExpired              = &Errno{Code: 10034, Message: map[string]string{"en": "Invitation has expired", "chs": "邀请已失效", "cht": "邀請已失效"}}
	ErrInviteIsSelf                 = &Errno{Code: 10035, Message: map[string]string{"en": "Received, contact friends and attend on time~", "chs": "已领取，联系好友，准时参加哦~", "cht": "已領取，聯繫好友，準時參加哦~"}}
	ErrInviteIsMobile               = &Errno{Code: 10036, Message: map[string]string{"en": "Failed to collect, incorrect mobile number", "chs": "领取失败，填写手机号不正确", "cht": "領取失敗，填寫手機號不正確"}}
	ErrInviteIsLoginMobile          = &Errno{Code: 10037, Message: map[string]string{"en": "Failed to collect, login mobile number is incorrect", "chs": "领取失败，登录手机号不正确 ", "cht": "領取失敗，登錄手機號不正確"}}
	ErrUserUserIsInvite             = &Errno{Code: 10021, Message: map[string]string{"en": "You have accepted the invitation. You can gift it to others", "chs": "您已接受过邀请，可转赠邀请他人", "cht": "您已接受邀請，可轉贈邀請他人"}}
	ErrInviteIsExist                = &Errno{Code: 10022, Message: map[string]string{"en": "The invitation letter does not exist", "chs": "该邀请函不存在", "cht": "該邀請函不存在"}}
	ErrInviteIsInvalid              = &Errno{Code: 10023, Message: map[string]string{"en": "The invitation has expired", "chs": "该邀请函已失效", "cht": "該邀請函已失效"}}
	ErrTicketStoreIsOpening         = &Errno{Code: 10041, Message: map[string]string{"en": "The shop has been closed", "chs": "该店铺已歇业", "cht": "該店鋪已歇業"}}
	ErrDeliveryTypeIsValid          = &Errno{Code: 10042, Message: map[string]string{"en": "The delivery method is not within the specified time", "chs": "该配送方式不在规定时间内", "cht": "該配送方式不在規定時間內"}}

	//10105不能修改 小程序/h5根据此错误码重新登陆
	ErrTokenInvalid              = &Errno{Code: 10105, Message: map[string]string{"en": "The token was invalid", "chs": "登录凭证失效,请稍后重试或联系客服人员", "cht": "登錄憑證失效,請稍後重試或聯繫客服人員"}}
	ErrTokenCreateUpdateFailed   = &Errno{Code: 10101, Message: map[string]string{"en": "Create or update token failed", "chs": "更新登录凭证未成功,请稍后重试或联系客服人员", "cht": "更新登錄憑證未成功,請稍後重試或聯繫客服人員"}}
	ErrMiniPSignFailed           = &Errno{Code: 10102, Message: map[string]string{"en": "Minip sign in failed", "chs": "小程序授权未成功,请稍后重试或联系客服人员", "cht": "小程序授權未成功,請稍後重試或聯繫客服人員"}}
	ErrMiniPCreateUserFailed     = &Errno{Code: 10103, Message: map[string]string{"en": "Minip create user failed", "chs": "小程序登录未成功,请稍后重试或联系客服人员", "cht": "小程序登錄未成功,請稍後重試或聯繫客服人員"}}
	ErrMiniPCreateIdentityFailed = &Errno{Code: 10104, Message: map[string]string{"en": "Minip create identity failed", "chs": "小程序登录未成功,请稍后重试或联系客服人员", "cht": "小程序登錄未成功,請稍後重試或聯繫客服人員"}}
	ErrNeedNameorPortrait        = &Errno{Code: 10111, Message: map[string]string{"en": "No userName or no portrait", "chs": "需要昵称或头像", "cht": "需要暱稱或頭像"}}
	ErrSmsOneMinuteOneMessage    = &Errno{Code: 10112, Message: map[string]string{"en": "One minute only one message", "chs": "一分钟内只能请求一次验证码,请稍后重试", "cht": "一分鐘內只能請求一次驗證碼,請稍後重試"}}
	ErrSms24Hours5Messages       = &Errno{Code: 10113, Message: map[string]string{"en": "24 hours only 5 messages", "chs": "24小时内只能请求5次验证码,请稍后重试", "cht": "24小時內只能請求一次驗證碼,請稍後重試"}}
	ErrSmsSendFailed             = &Errno{Code: 10114, Message: map[string]string{"en": "Send verify code message failed", "chs": "发送验证码未成功,请稍后重试或联系客服人员", "cht": "發送驗證碼未成功,請稍後重試或聯繫客服人員"}}
	ErrSmsVerifyCodeWrong        = &Errno{Code: 10115, Message: map[string]string{"en": "Verify code is wrong", "chs": "验证码错误,请稍后重试或联系客服人员", "cht": "驗證碼錯誤,請稍後重試或聯繫客服人員"}}
	ErrWxSignFailed              = &Errno{Code: 10106, Message: map[string]string{"en": "Weixin sign in failed", "chs": "微信授权未成功,请稍后重试或联系客服人员", "cht": "微信授權未成功,請稍後重試或聯繫客服人員"}}
	ErrWxCreateUserFailed        = &Errno{Code: 10107, Message: map[string]string{"en": "Weixin create user failed", "chs": "微信登录未成功,请稍后重试或联系客服人员", "cht": "微信登錄未成功,請稍後重試或聯繫客服人員"}}
	ErrWxCreateIdentityFailed    = &Errno{Code: 10108, Message: map[string]string{"en": "Weixin create identity failed", "chs": "微信登录未成功,请稍后重试或联系客服人员", "cht": "微信登錄未成功,請稍後重試或聯繫客服人員"}}
	ErrTodayHadToComeSign        = &Errno{Code: 10109, Message: map[string]string{"en": "Signed in today", "chs": "今日已签到", "cht": "今日已簽到"}}
	ErrMobileIsErr               = &Errno{Code: 10110, Message: map[string]string{"en": "Please enter the correct cell phone number", "chs": "请输入正确手机号", "cht": "请输入正确手机号"}}
	ErrCodeIsEmpty               = &Errno{Code: 10117, Message: map[string]string{"en": "Please enter the verification code", "chs": "请输入验证码", "cht": "请输入验证码"}}
	ErrNationIsEmpty             = &Errno{Code: 10118, Message: map[string]string{"en": "Please enter the country code", "chs": "请输入国家编码", "cht": "请输入国家编码"}}
	ErrWxUserInfoErr             = &Errno{Code: 10119, Message: map[string]string{"en": "User wechat information acquisition failed", "chs": "用户微信信息获取失败", "cht": "用戶微信信息獲取失敗"}}
	ErrUserNotBindWx             = &Errno{Code: 10120, Message: map[string]string{"en": "User is not bound to wechat", "chs": "用户未绑定微信", "cht": "用戶未綁定微信"}}
	ErrMobileIsBindWx            = &Errno{Code: 10121, Message: map[string]string{"en": "This phone has been bound to other wechat, please change the phone number", "chs": "该手机已绑定其他微信，请更换手机号", "cht": "該手機已綁定其他微信，請更換手機號"}}
	ErrAccountIsBindMobile       = &Errno{Code: 10122, Message: map[string]string{"en": "This account has been bound with other mobile numbers", "chs": "该账号已绑定其他手机号", "cht": "該賬號已綁定其他手機號"}}
	ErrMobileIsBindApple         = &Errno{Code: 10123, Message: map[string]string{"en": "This phone has been bound to other apple, please change the phone number", "chs": "该手机已绑定其他苹果，请更换手机号", "cht": "該手機已綁定其他蘋果，請更換手機號"}}
	ErrUserNotBindApple          = &Errno{Code: 10124, Message: map[string]string{"en": "User is not bound to Apple", "chs": "用户未绑定苹果", "cht": "用戶未綁定蘋果"}}
	ErrAppleFailed               = &Errno{Code: 10125, Message: map[string]string{"en": "Apple authorization failed. Please try again later or contact customer service", "chs": "苹果授权未成功,请稍后重试或联系客服人员", "cht": "蘋果授權未成功,請稍後重試或聯繫客服人員"}}
	ErrAppleTokenInvalid         = &Errno{Code: 10126, Message: map[string]string{"en": "The token was invalid", "chs": "苹果凭证无效,请稍后重试或联系客服人员", "cht": "蘋果憑證無效,請稍後重試或聯繫客服人員"}}

	ErrUploadImgFailed = &Errno{Code: 10201, Message: map[string]string{"en": "Upload image Failed", "chs": "上传图片失败,请稍后重试或联系客服人员", "cht": "上傳圖片失敗,請稍後重試或聯繫客服人員"}}
	ErrStsFailed       = &Errno{Code: 10202, Message: map[string]string{"en": "AliyunSts Failed", "chs": "阿里云sts授权未成功,请稍后重试或联系客服人员", "cht": "阿里雲sts授權未成功,請稍後重試或聯繫客服人員"}}

	ErrProductNotExist         = &Errno{Code: 30001, Message: map[string]string{"en": "Product do not exist", "chs": "商品不存在或已下架,请稍后重试或联系客服人员", "cht": "商品不存在或已下架,請稍後重試或聯繫客服人員"}}
	ErrSpecNotExist            = &Errno{Code: 30002, Message: map[string]string{"en": "Spec do not exist", "chs": "规格不存在或已下架,请稍后重试或联系客服人员", "cht": "規格不存在或已下架,請稍後重試或聯繫客服人員"}}
	ErrOrderNotExist           = &Errno{Code: 30003, Message: map[string]string{"en": "Order do not exist", "chs": "订单不存在。", "cht": "訂單不存在。"}}
	ErrPayTypeWrong            = &Errno{Code: 30004, Message: map[string]string{"en": "PayType is wrong", "chs": "支付方式错误,请稍后重试或联系客服人员", "cht": "支付方式錯誤,請稍後重試或聯繫客服人員"}}
	ErrOrderPaymentTimeout     = &Errno{Code: 30005, Message: map[string]string{"en": "Order payment timeout", "chs": "订单支付超時,请稍后重试或联系客服人员", "cht": "訂單支付超時,請稍後重試或聯繫客服人員"}}
	ErrOrderNotPay             = &Errno{Code: 30006, Message: map[string]string{"en": "You have unpaid orders", "chs": "您有未支付订单，请先完成未支付订单或取消该订单", "cht": "您有訂單未支付,请先完成未支付訂單或取消該訂單"}}
	ErrQuantityRestriction     = &Errno{Code: 30007, Message: map[string]string{"en": "Beyond quantity restriction", "chs": "超出购买数量限制，请先完成未支付订单或取消该订单", "cht": "超出購買數量限制,请先完成未支付訂單或取消該訂單"}}
	ErrLowStocks               = &Errno{Code: 30008, Message: map[string]string{"en": "Low stocks", "chs": "库存不足，请先完成未支付订单或取消该订单", "cht": "庫存不足,请先完成未支付訂單或取消該訂單"}}
	ErrWeixinPayFailed         = &Errno{Code: 30009, Message: map[string]string{"en": "WeixinPay failed", "chs": "微信支付未成功,请稍后重试或联系客服人员", "cht": "微信支付未成功,請稍後重試或聯繫客服人員"}}
	ErrWeixinOrderCancelFailed = &Errno{Code: 30010, Message: map[string]string{"en": "Cancel WeixinOrder failed", "chs": "取消微信订单未成功,请稍后重试或联系客服人员", "cht": "取消微信訂單未成功,請稍後重試或聯繫客服人員"}}
	ErrAliPayFailed            = &Errno{Code: 30011, Message: map[string]string{"en": "AliPay failed", "chs": "阿里支付未成功,请稍后重试或联系客服人员", "cht": "阿里支付未成功,請稍後重試或聯繫客服人員"}}
	ErrPayPalFailed            = &Errno{Code: 30012, Message: map[string]string{"en": "PayPal failed", "chs": "PayPal支付未成功,请稍后重试或联系客服人员", "cht": "PayPal支付未成功,請稍後重試或聯繫客服人員"}}
	ErrCouponNotFound          = &Errno{Code: 30013, Message: map[string]string{"en": "Coupon does not exist", "chs": "优惠券不存在", "cht": "優惠券不存在"}}
	ErrCouponIsUsed            = &Errno{Code: 30014, Message: map[string]string{"en": "Coupon is already in use", "chs": "优惠券已使用", "cht": "優惠券已使用"}}
	ErrCouponIsNotUsed         = &Errno{Code: 30015, Message: map[string]string{"en": "Coupon not available.", "chs": "优惠券不可用", "cht": "優惠券不可用"}}
	ErrExpressIdIsErr          = &Errno{Code: 30016, Message: map[string]string{"en": "Express order number error.", "chs": "快递单号错误", "cht": "快遞單號錯誤"}}
	ErrExpressNotFound         = &Errno{Code: 30017, Message: map[string]string{"en": "Item not shipped.", "chs": "商品未发货", "cht": "商品未發貨"}}
	ErrOrderStatusIsErr        = &Errno{Code: 30018, Message: map[string]string{"en": "Order status is incorrect, not on-site ticket collection.", "chs": "订单状态错误，不可现场取票。", "cht": "訂單狀態錯誤，不可現場取票"}}
	ErrRefundFailed            = &Errno{Code: 30019, Message: map[string]string{"en": "The refund was unsuccessful. Please try again later or contact customer service.", "chs": "退款未成功,请稍后重试或联系客服人员", "cht": "退款未成功,請稍後重試或聯繫客服人員"}}
	ErrUnionPayFailed          = &Errno{Code: 30020, Message: map[string]string{"en": "UnionPay failed.", "chs": "银联支付失败", "cht": "銀聯支付失敗"}}
	ErrWeixinIsPayIng          = &Errno{Code: 30021, Message: map[string]string{"en": "Require user to enter payment password", "chs": "需要用户输入支付密码", "cht": "需要用戶輸入支付密碼"}}
	ErrWristbandIdIsErr        = &Errno{Code: 30022, Message: map[string]string{"en": "Bracelet error, please show this activity bracelet", "chs": "手环错误，请出示本活动手环", "cht": "手環錯誤，請出示本活動手環"}}
	ErrWristbandReissuePer     = &Errno{Code: 30023, Message: map[string]string{"en": "Please go to the service desk to reissue the bracelet.", "chs": "请前往服务台补发手环", "cht": "請前往服務台補發手環"}}
	ErrOrderChangeIsEnd        = &Errno{Code: 30024, Message: map[string]string{"en": "Can not be changed within one week from the departure time", "chs": "距发车时间一周内无法改签", "cht": "距發車時間一周內無法改簽"}}
	ErrUserCardIsErr           = &Errno{Code: 30025, Message: map[string]string{"en": "Please input the certified information correctly", "chs": "请正确输入认证用户本人储蓄银行卡信息", "cht": "請正確輸入認證用戶本人儲蓄銀行卡信息"}}
	ErrOrderChangeNumErr       = &Errno{Code: 30026, Message: map[string]string{"en": "You have only one chance to change your tickets.", "chs": "您已改签过一次，不可再次改签", "cht": "您已改簽過一次，不可再次改簽"}}
	ErrGotoAppPay              = &Errno{Code: 30027, Message: map[string]string{"en": "Please go to the APP to pay", "chs": "请前往APP支付", "cht": "請前往APP支付"}}
	ErrGotoMinipPay            = &Errno{Code: 30028, Message: map[string]string{"en": "Please go to the minip to pay", "chs": "请前往小程序支付", "cht": "請前往小程序支付"}}

	ErrNeedAddress           = &Errno{Code: 30101, Message: map[string]string{"en": "Please fill in the delivery address", "chs": "请填写收货地址", "cht": "請填寫收貨地址"}}
	ErrNeedContact           = &Errno{Code: 30102, Message: map[string]string{"en": "Please fill in the contact", "chs": "请填写购票联系人", "cht": "請填寫購票聯繫人"}}
	ErrContactNumLimit       = &Errno{Code: 30103, Message: map[string]string{"en": "ContactNumLimit", "chs": "联系人已达到上限", "cht": "聯繫人已達到上限"}}
	ErrContactNumWrong       = &Errno{Code: 30104, Message: map[string]string{"en": "The contact Num is wrong", "chs": "请选择与购票数相等的购票人数", "cht": "請選擇與票數相等的購票人數"}}
	ErrContactCanBuyOne      = &Errno{Code: 30105, Message: map[string]string{"en": "Each contact can only buy one time", "chs": "每个购买人同一场次只能购买一次", "cht": "每個購買人同一场次只能購買一次"}}
	ErrCorrectIdentity       = &Errno{Code: 30106, Message: map[string]string{"en": "Please fill in the correct identity information", "chs": "请填写正确身份信息", "cht": "請填寫正確身份信息"}}
	ErrIdentityIsExist       = &Errno{Code: 30106, Message: map[string]string{"en": "Please fill in the correct identity information", "chs": "您输入的身份证号或手机号已被绑定其他帐号，请重新输入", "cht": "請填寫正確身份信息"}}
	ErrRealAuthentication    = &Errno{Code: 30107, Message: map[string]string{"en": "Real person authentication failed", "chs": "实人认证失败,请稍后重试或联系客服人员", "cht": "實人認證失敗,請稍後重試或聯繫客服人員"}}
	ErrRealAuthenticationIng = &Errno{Code: 30108, Message: map[string]string{"en": "Real person authentication has not done, please wait", "chs": "实人认证中,请稍后查看或重试", "cht": "實人認證中,請稍後后查看或重試"}}
	ErrContactHasBind        = &Errno{Code: 30109, Message: map[string]string{"en": "The Contact or account has already been bound", "chs": "该证件或该账号已被绑定", "cht": "該證件已被他人綁定"}}
	ErrDelMeContact          = &Errno{Code: 30110, Message: map[string]string{"en": "The Contact is Yourself", "chs": "不能删除本人", "cht": "不能刪除本人"}}
	ErrSceneIsEndIng         = &Errno{Code: 30111, Message: map[string]string{"en": "The session has expired", "chs": "场次已过期", "cht": "場次已過期"}}
	ErrUserIsJuveniles       = &Errno{Code: 30112, Message: map[string]string{"en": "No tickets under 18", "chs": "18岁以下禁止购票", "cht": "18歲以下禁止購票"}}

	ErrHasActivated         = &Errno{Code: 30201, Message: map[string]string{"en": "Your account has already activated the ticket", "chs": "该账号已激活门票", "cht": "該賬號已激活門票"}}
	ErrActCodeWrong         = &Errno{Code: 30202, Message: map[string]string{"en": "Your code is wrong", "chs": "门票激活码错误", "cht": "門票激活碼錯誤"}}
	ErrActCodeUsed          = &Errno{Code: 30203, Message: map[string]string{"en": "Your code has already been used", "chs": "门票激活码已使用,请稍后重试或联系客服人员", "cht": "門票激活碼已使用,請稍後重試或聯繫客服人員"}}
	ErrNoMobile             = &Errno{Code: 30204, Message: map[string]string{"en": "You need to bind your mobile", "chs": "您的账号需要绑定手机号,请稍后重试或联系客服人员", "cht": "您的賬號需要綁定手機號,請稍後重試或聯繫客服人員"}}
	ErrMiniPGetMobile       = &Errno{Code: 30205, Message: map[string]string{"en": "MiniP get mobile failed", "chs": "小程序获取用户手机号授权未成功,请稍后重试或联系客服人员", "cht": "小程序獲取用戶手機號授權未成功,請稍後重試或聯繫客服人員"}}
	ErrNotActivated         = &Errno{Code: 30206, Message: map[string]string{"en": "Your have not activated the ticket", "chs": "您暂时没有票被激活,请稍后重试或联系客服人员", "cht": "您暫時還沒有票被激活,請稍後重試或聯繫客服人員"}}
	ErrMobileUsed           = &Errno{Code: 30207, Message: map[string]string{"en": "The mobile has already been used, consolidate the accounts？", "chs": "手机号已被其他账号使用,是否合并账号？", "cht": "手機號已被其他賬號使用,是否合併賬號？"}}
	ErrIdCardUsed           = &Errno{Code: 30208, Message: map[string]string{"en": "The idcard has already been used", "chs": "该身份证件已经被激活,请稍后重试或联系客服人员", "cht": "該身份證件已經被激活,請稍後重試或聯繫客服人員"}}
	ErrNoNeedConsolidate    = &Errno{Code: 30209, Message: map[string]string{"en": "Don't need to consolidate the accounts", "chs": "您不需要合并账号", "cht": "該身份證已經被激活,請稍後重試或聯繫客服人員"}}
	ErrWrongIdtype          = &Errno{Code: 30210, Message: map[string]string{"en": "The idcard type is wrong", "chs": "身份证件类型错误,请稍后重试或联系客服人员", "cht": "身份證件類型錯誤,請稍後重試或聯繫客服人員"}}
	ErrMobileUsedWeixin     = &Errno{Code: 30211, Message: map[string]string{"en": "The mobile has already been used.", "chs": "手机号已被其他账号使用,请更换手机账号", "cht": "手機號已被其他賬號使用,請更換手機賬號"}}
	ErrActCodeWrongTimes    = &Errno{Code: 30212, Message: map[string]string{"en": "You wrong too many times", "chs": "您激活错误次数过多,请稍后重试或联系客服人员", "cht": "您激活錯誤次數過多,請稍後重試或聯繫客服人員"}}
	ErrIdCardIsWrong        = &Errno{Code: 30213, Message: map[string]string{"en": "Identity card format error", "chs": "身份证格式错误。", "cht": "身份證格式錯誤。"}}
	ErrQingdaoWrong         = &Errno{Code: 30214, Message: map[string]string{"en": "Qingdao only support ChinaId", "chs": "青岛电子票支持身份证验证", "cht": "青島電子票支持身份證驗證"}}
	ErrYouareDianzi         = &Errno{Code: 30215, Message: map[string]string{"en": "You are an e-ticket, please go to e-ticket activation.", "chs": "您是电子票，请前往电子票激活", "cht": "您是電子票，請前往電子票激活。"}}
	ErrAlreadyHaveWristband = &Errno{Code: 30216, Message: map[string]string{"en": "You already have one wristband", "chs": "您已领取手环", "cht": "您已領取手環。"}}
	ErrBuyTypeIsErr         = &Errno{Code: 30217, Message: map[string]string{"en": "Free and paid items cannot be placed together", "chs": "免费商品和付费商品不可一起下单", "cht": "免費商品和付費商品不可一起下單"}}
	ErrWrongWristband       = &Errno{Code: 30218, Message: map[string]string{"en": "Your wristbandId is wrong", "chs": "您的手环ID不合法", "cht": "您的手環ID不合法"}}
	ErrNeedBuyTicketFirst   = &Errno{Code: 30219, Message: map[string]string{"en": "You need buy ticket first", "chs": "您需要先购买此商品的现场门票", "cht": "您需要先購買此商品的現場門票"}}
	ErrNotHaveSceneAuth     = &Errno{Code: 30220, Message: map[string]string{"en": "The store did not get permission for this session", "chs": "店铺未获得该场次权限", "cht": "店鋪未獲得該場次權限"}}
	ErrBraceletNotUser      = &Errno{Code: 30221, Message: map[string]string{"en": "The bracelet is not associated with the user", "chs": "该手环未关联用户", "cht": "該手環為關聯用戶"}}
	ErrWristbandIsHaveUser  = &Errno{Code: 30222, Message: map[string]string{"en": "The association failed and the account has been associated", "chs": "关联失败，该账号已被关联", "cht": "關聯失敗，該賬號已被關聯"}}

	ErrHomiezNoMobile = &Errno{Code: 30301, Message: map[string]string{"en": "The orderId or the mobile is wrong", "chs": "手机号或订单号错误,请稍后重试或联系客服人员", "cht": "手機號或訂單號錯誤,請稍後重試或聯繫客服人員"}}

	ErrStarCircleIsStar     = &Errno{Code: 30401, Message: map[string]string{"en": "You have already liked it.", "chs": "你已点过赞了", "cht": "你已點過讚了"}}
	ErrIntegralIsEmpty      = &Errno{Code: 30402, Message: map[string]string{"en": "Points have been exhausted.", "chs": "积分已用尽", "cht": "積分已用盡"}}
	ErrDreamIsEnd           = &Errno{Code: 30403, Message: map[string]string{"en": "Dream building has ended.", "chs": "筑梦活动已结束", "cht": "築夢活動已結束"}}
	ErrUserHasNoOpportunity = &Errno{Code: 30404, Message: map[string]string{"en": "Sorry, your lottery is running out. Please come back tomorrow!.", "chs": "Oops，你的抽奖次数用光了。请明天再来！", "cht": "Oops，你的抽獎次數用光了。請明天再來！"}}
	ErrNotGetPrize          = &Errno{Code: 30405, Message: map[string]string{"en": "Excuse me, this time it was only a little bit of luck from the realization of the dream.", "chs": "不好意思，这次离梦想实现只差了一点点运气。", "cht": "不好意思，這次離夢想實現只差了一點點運氣。"}}
	ErrCircleIsEmpty        = &Errno{Code: 30406, Message: map[string]string{"en": "Post does not exist or has been deleted.", "chs": "帖子不存在或者已删除。", "cht": "帖子不存在或者已刪除。"}}

	ErrSetRecordCheat        = &Errno{Code: 40001, Message: map[string]string{"en": "SetRecord is Cheat.", "chs": "作弊", "cht": "作弊"}}
	ErrNotEnoughDiamond      = &Errno{Code: 40002, Message: map[string]string{"en": "Diamond is not Enough.", "chs": "您的钻石不够", "cht": "您的鑽石不夠"}}
	ErrBuyLimit              = &Errno{Code: 40003, Message: map[string]string{"en": "The good is purchase limit.", "chs": "此商品限量购买,您已经购买过", "cht": "此商品限量購買,您已經購買過"}}
	ErrUserBan               = &Errno{Code: 40004, Message: map[string]string{"en": "The accounts were closed status.", "chs": "账号状态异常,请稍后重试或联系客服人员", "cht": "賬號狀態異常,請稍後重試或聯繫客服人員"}}
	ErrNotEnoughRaffleTicket = &Errno{Code: 40005, Message: map[string]string{"en": "raffle ticket  is not Enough.", "chs": "您的抽奖券不足", "cht": "您的抽獎券不足"}}
	ErrGameActivityEnd       = &Errno{Code: 40006, Message: map[string]string{"en": "game activity end.", "chs": "活动结束", "cht": "活動結束"}}
	ErrGameRoundEnd          = &Errno{Code: 40007, Message: map[string]string{"en": "game round end.", "chs": "本局游戏结束", "cht": "本局遊戲結束"}}
	ErrVersion               = &Errno{Code: 40008, Message: map[string]string{"en": "please upgrade your app.", "chs": "请升级您的应用", "cht": "請陞級您的應用程序"}}
	ErrInvalidActivityId     = &Errno{Code: 40009, Message: map[string]string{"en": "invalid activity.", "chs": "活动结束", "cht": "活動結束"}}
	ErrGoldNotEnough         = &Errno{Code: 40010, Message: map[string]string{"en": "Gold is not Enough.", "chs": "您的金币不够", "cht": "您的金幣不够"}}
	ErrPurchased             = &Errno{Code: 40011, Message: map[string]string{"en": "The song already purchased.", "chs": "您已经购买", "cht": "您已經購買"}}
	ErrCompetitionNotStart   = &Errno{Code: 40012, Message: map[string]string{"en": "The competition is not start yet.", "chs": "比赛尚未开始", "cht": "比賽尚未開始"}}
	ErrLowScore              = &Errno{Code: 40101, Message: map[string]string{"en": "The score is lower.", "chs": "新分数比旧分数低", "cht": "新分數比舊分數低"}}
	ErrNoBeatMapContent      = &Errno{Code: 40102, Message: map[string]string{"en": "No BeatMapContent.", "chs": "谱面内容不存在", "cht": "譜面內容不存在"}}
	ErrBeatMapContentStatus  = &Errno{Code: 40103, Message: map[string]string{"en": "BeatMapContent status is wrong.", "chs": "谱面内容状态不正确", "cht": "譜面內容不正確"}}
	ErrNoThisBeatMap         = &Errno{Code: 40104, Message: map[string]string{"en": "You don't have this beatmap.", "chs": "您未拥有这个谱面", "cht": "您未擁有這個譜面"}}
	ErrNotEnoughPoint        = &Errno{Code: 40105, Message: map[string]string{"en": "point is not enough.", "chs": "积分不足", "cht": "積分不足"}}
	ErrSongNotExist          = &Errno{Code: 40106, Message: map[string]string{"en": "song does not exist.", "chs": "歌曲不存在", "cht": "歌曲不存在"}}
	ErrUploadLimit           = &Errno{Code: 40107, Message: map[string]string{"en": "beyond upload limit", "chs": "超出上传个数限制", "cht": "超出上传个数限制"}}

	ErrCondomIsGeted     = &Errno{Code: 50001, Message: map[string]string{"en": "You have received.", "chs": "您已领取", "cht": "您已領取"}}
	ErrCondomIsNotStart  = &Errno{Code: 50002, Message: map[string]string{"en": "Please pick up at 2018.12.30.", "chs": "请于2018.12.30现场领取", "cht": "請於2018.12.30現場領取"}}
	ErrNotHaveOpp        = &Errno{Code: 50003, Message: map[string]string{"en": "No voting opportunities.", "chs": "没有投票机会了", "cht": "沒有投票機會了"}}
	ErrOnlyFivePlanet    = &Errno{Code: 50004, Message: map[string]string{"en": "Upload up to 5 works.", "chs": "最多上传5个作品。", "cht": "最多上傳5個作品。"}}
	ErrIsExistVotePlanet = &Errno{Code: 50005, Message: map[string]string{"en": "You have already voted, please come back tomorrow.", "chs": "您已经投过票了，明天再来吧。", "cht": "您已經投過票了，明天再來吧。"}}

	ErrGateNotEnter     = &Errno{Code: 60001, Message: map[string]string{"en": "Not purchased.", "chs": "未购票。", "cht": "未購票。"}}
	ErrGateIsETicket    = &Errno{Code: 60002, Message: map[string]string{"en": "Please take the Qingdao dedicated channel.", "chs": "请走青岛专用通道。", "cht": "請走青島專用通道。"}}
	ErrGateTimeIsWrong  = &Errno{Code: 60003, Message: map[string]string{"en": "Date does not match.", "chs": "日期不符。", "cht": "日期不符。"}}
	ErrGateIsExistEnter = &Errno{Code: 60004, Message: map[string]string{"en": "Admitted today.", "chs": "今日已入场", "cht": "今日已入場。"}}
	ErrGateIsGrepErr    = &Errno{Code: 60005, Message: map[string]string{"en": "Please enter from the corresponding channel.", "chs": "请从对应通道入场", "cht": "請從對應通道入場"}}
	ErrInIsNot          = &Errno{Code: 60006, Message: map[string]string{"en": "User has no access", "chs": "用户无入场权限", "cht": "用戶無入場權限"}}

	ErrEasemobAPI      = &Errno{Code: 70001, Message: map[string]string{"en": "chat api.", "chs": "聊天", "cht": "聊天"}}
	ErrMessageTooLarge = &Errno{Code: 70002, Message: map[string]string{"en": "message too large.", "chs": "消息字数超过上限", "cht": "消息字數超過上限"}}
	ErrGroupNotExist   = &Errno{Code: 70003, Message: map[string]string{"en": "group does not exist.", "chs": "群组不存在", "cht": "群組不存在"}}
	ErrNotGroupMember  = &Errno{Code: 70004, Message: map[string]string{"en": "member does not blong to this group.", "chs": "成员不属于该群组", "cht": "成員不屬於該群組"}}
	ErrKicedout        = &Errno{Code: 70005, Message: map[string]string{"en": "you have been kiced out.", "chs": "你已经被踢出该群", "cht": "你已經被踢出該群"}}
	//ErrMessageTooLarge = &Errno{Code: 70002, Message: map[string]string{"en": "message too large.", "chs": "消息字数超过上限", "cht": "消息字數超過上限"}}

	ErrPledgeNotFound   = &Errno{Code: 80001, Message: map[string]string{"en": "You do not have a deposit yet.", "chs": "您还没有押金。", "cht": "您還沒有押金。"}}
	ErrPledgeIsExist    = &Errno{Code: 80002, Message: map[string]string{"en": "You have paid the deposit.", "chs": "您已支付押金。", "cht": "您已支付押金。"}}
	ErrNotHaveBattery   = &Errno{Code: 80003, Message: map[string]string{"en": "There is no charging treasure in use.", "chs": "没有正在使用中的充电宝。", "cht": "沒有正在使用中的充電寶。"}}
	ErrBatteryBorrowing = &Errno{Code: 80004, Message: map[string]string{"en": "Being rented", "chs": "正在租借中", "cht": "正在租借中"}}
	ErrBatteryPledgeNot = &Errno{Code: 80005, Message: map[string]string{"en": "If you have an uncompleted order, you need to return the charging treasure in the rental before you can refund the deposit~", "chs": "您有未完成订单，需要先归还租借中的充电宝，才能退押金哦～", "cht": "您有未完成訂單，需要先歸還租借中的充電寶，才能退押金哦～"}}

	ErrVideoOrderIsExist = &Errno{Code: 90001, Message: map[string]string{"en": "You have already purchased this video.", "chs": "您已购买该视频。", "cht": "您已購買該視頻。"}}
	ErrNotHaveVideoOrder = &Errno{Code: 90002, Message: map[string]string{"en": "You do not have permission to view this video.", "chs": "您没有该视频的权限。", "cht": "您沒有該視頻的權限。"}}
	ErrBusNotBuy         = &Errno{Code: 90003, Message: map[string]string{"en": "Unpurchased trains", "chs": "未购本车次", "cht": "未購本車次"}}

	ErrUserWalletNotOpen   = &Errno{Code: 100001, Message: map[string]string{"en": "You have not opened a wallet yet.", "chs": "您还没有开通钱包", "cht": "您還沒有開通錢包。"}}
	ErrPassWordIsErr       = &Errno{Code: 100002, Message: map[string]string{"en": "wrong password.", "chs": "密码错误", "cht": "密碼錯誤"}}
	ErrWalletAmountErr     = &Errno{Code: 100003, Message: map[string]string{"en": "Amount error.", "chs": "金额错误", "cht": "金額錯誤"}}
	ErrPassWordFormatIsErr = &Errno{Code: 100004, Message: map[string]string{"en": "Please enter a 6-digit alphanumeric password.", "chs": "请输入6位纯数字密码", "cht": "請輸入6位純數字密碼"}}
	ErrWalletBalanceNotGou = &Errno{Code: 100005, Message: map[string]string{"en": "Insufficient wallet balance.", "chs": "钱包余额不足", "cht": "錢包餘額不足"}}
	ErrUserWalletisOpen    = &Errno{Code: 100006, Message: map[string]string{"en": "You have opened your wallet, please return to my page and try again!", "chs": "您已开通钱包，请返回我的页面重试！", "cht": "您已開通錢包，請返回我的頁面重試！"}}
	ErrPosUserNoExist      = &Errno{Code: 100007, Message: map[string]string{"en": "No account exists", "chs": "账号不存在", "cht": "账号不存在"}}
	ErrAccountInvalid      = &Errno{Code: 100008, Message: map[string]string{"en": "Account invalidation, contact customer service", "chs": "账号失效，联系客服", "cht": "賬號失效，聯繫客服"}}
	ErrInvalidOrderStatus  = &Errno{Code: 100009, Message: map[string]string{"en": "Invalid order status", "chs": "无效的订单状态", "cht": "无效的订单状态"}}
	ErrOrderCanled         = &Errno{Code: 100011, Message: map[string]string{"en": "Order closed", "chs": "订单已关闭", "cht": "訂單已關閉"}}

	ErrProjectNameExists = &Errno{Code: 200001, Message: map[string]string{"en": "Project name exists", "chs": "项目名称已经存在", "cht": "项目名称已经存在"}}
	ErrSceneNameExists   = &Errno{Code: 200002, Message: map[string]string{"en": "scene name exists", "chs": "场次名称已经存在", "cht": "场次名称已经存在"}}
	ErrNoProjectBind     = &Errno{Code: 200003, Message: map[string]string{"en": "No project bind", "chs": "未指定项目", "cht": "未指定项目"}}
	ErrNoHavePerDontIn   = &Errno{Code: 200004, Message: map[string]string{"en": "No permission, no access", "chs": "没有权限，不可进入", "cht": "沒有權限，不可進入"}}
	ErrAuthError         = &Errno{Code: 200005, Message: map[string]string{"en": "Identity error", "chs": "身份错误", "cht": "身份錯誤"}}
	ErrDonotEnter        = &Errno{Code: 200006, Message: map[string]string{"en": "Do not enter", "chs": "禁止通行", "cht": "禁止通行"}}
	ErrWristbandHasUser  = &Errno{Code: 200007, Message: map[string]string{"en": "The bracelet has been successfully associated with the user", "chs": "手环已成功关联用户", "cht": "手環已成功關聯用戶"}}

	ErrMobileInvalid                           = &Errno{Code: 300000, Message: map[string]string{"en": "cell phone number is invalid", "chs": "手机号无效。", "cht": "手機號無效。"}}
	ErrSummary                                 = &Errno{Code: 300001, Message: map[string]string{"en": "summary too long", "chs": "内容长度超过限制。", "cht": "。"}}
	ErrTitle                                   = &Errno{Code: 300002, Message: map[string]string{"en": "title too long", "chs": "标题长度超过限制 。", "cht": "。"}}
	ErrHealthCodeNotStart                      = &Errno{Code: 300003, Message: map[string]string{"en": "Not at registration time. Please try again later", "chs": "不在登记时间，请稍后重试", "cht": "不在登記時間，請稍後重"}}
	ErrHealthStatusEmptyNoGrantWristband       = &Errno{Code: 300004, Message: map[string]string{"en": "The health status is not filled in, and the bracelet cannot be sent", "chs": "健康状态未填写，不能发手环", "cht": "健康狀態未填寫，不能發手環"}}
	ErrHealthStatusInvalidNoGrantWristband     = &Errno{Code: 300005, Message: map[string]string{"en": "Failure of health status, no Bracelet", "chs": "健康状态失效，不能发手环", "cht": "健康狀態失效，不能發手環"}}
	ErrHealthStatusAbnormalNoGrantWristband    = &Errno{Code: 300006, Message: map[string]string{"en": "Abnormal health, no Bracelet", "chs": "健康状态异常，不能发手环", "cht": "健康狀態異常，不能發手環"}}
	ErrHealthStatusEmptyNoRelationWristband    = &Errno{Code: 300007, Message: map[string]string{"en": "Health status is not filled in, cannot associate with Bracelet", "chs": "健康状态未填写，不能关联手环", "cht": "健康狀態未填寫，不能關聯手環"}}
	ErrHealthStatusInvalidNoRelationWristband  = &Errno{Code: 300008, Message: map[string]string{"en": "The health status is invalid, and the bracelet cannot be associated", "chs": "健康状态失效，不能关联手环", "cht": "健康狀態失效，不能關聯手環"}}
	ErrHealthStatusAbnormalNoRelationWristband = &Errno{Code: 300009, Message: map[string]string{"en": "Abnormal health status, unable to associate Bracelet", "chs": "健康状态异常，不能关联手环", "cht": "健康狀態異常，不能關聯手環"}}

	ErrStorePosUserExist = &Errno{Code: 400000, Message: map[string]string{"en": "账号已经存在", "chs": "账号已经存在", "cht": ""}}
	ErrUserPwdAtypism    = &Errno{Code: 400001, Message: map[string]string{"en": "两次密码不一致", "chs": "两次密码不一致", "cht": ""}}
	ErrStorePosUserNum   = &Errno{Code: 400002, Message: map[string]string{"en": "Account has reached the upper limit", "chs": "账号已达上限", "cht": ""}}
	ErrAccountExist      = &Errno{Code: 400003, Message: map[string]string{"en": "Account number already exist", "chs": "账号已存在", "cht": ""}}
)

var (
	InfoStarTagAll       = map[string]string{"en": "DJ", "chs": "DJ巨星", "cht": "DJ巨星"}
	InfoStarTagFirst     = map[string]string{"en": "First", "chs": "第一期", "cht": "第一期"}
	InfoStarTagSecond    = map[string]string{"en": "Second", "chs": "第二期", "cht": "第貳期"}
	InfoStarTagThird     = map[string]string{"en": "Third", "chs": "第三期", "cht": "第叁期"}
	InfoStarTagFourth    = map[string]string{"en": "Fourth", "chs": "第四期", "cht": "第肆期"}
	InfoStarTagFiveth    = map[string]string{"en": "AFROBEATS", "chs": "AFROBEATS", "cht": "AFROBEATS"}
	InfoStarTagCanBuyOne = map[string]string{"en": " has bought one. ", "chs": " 已经购买过一次. ", "cht": " 已經購買過一次. "}
)
