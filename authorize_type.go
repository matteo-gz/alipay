package alipay

const (
	kProductionPublicAppAuthorize = "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm"
	kSandboxPublicAppAuthorize    = "https://openauth.alipaydev.com/oauth2/publicAppAuthorize.htm"
)

const (
	kProductionAppToAppAuth = "https://openauth.alipay.com/oauth2/appToAppAuth.htm"
	kSandboxAppToAppAuth    = "https://openauth.alipaydev.com/oauth2/appToAppAuth.htm"
)

// SystemOauthToken 换取授权访问令牌接口请求参数 https://docs.open.alipay.com/api_9/alipay.system.oauth.token
type SystemOauthToken struct {
	AppAuthToken string `json:"-"` // 可选
	GrantType    string `json:"-"` // 值为 authorization_code 时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"-"`
	RefreshToken string `json:"-"`
}

func (this SystemOauthToken) APIName() string {
	return "alipay.system.oauth.token"
}

func (this SystemOauthToken) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["grant_type"] = this.GrantType
	if this.Code != "" {
		m["code"] = this.Code
	}
	if this.RefreshToken != "" {
		m["refresh_token"] = this.RefreshToken
	}
	return m
}

// SystemOauthTokenRsp 换取授权访问令牌接口请求参数
type SystemOauthTokenRsp struct {
	Content struct {
		Code         Code   `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		UserId       string `json:"user_id"`
		AccessToken  string `json:"access_token"`
		ExpiresIn    int64  `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		ReExpiresIn  int64  `json:"re_expires_in"`
	} `json:"alipay_system_oauth_token_response"`
	Error *struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"` // 不要访问此结构体
	Sign string `json:"sign"`
}

// UserInfoShare 支付宝会员授权信息查询接口请求参数 https://docs.open.alipay.com/api_2/alipay.user.info.share
type UserInfoShare struct {
	AppAuthToken string `json:"-"` // 可选
	AuthToken    string `json:"-"` // 是
}

func (this UserInfoShare) APIName() string {
	return "alipay.user.info.share"
}

func (this UserInfoShare) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["auth_token"] = this.AuthToken
	return m
}

// UserInfoShareRsp 支付宝会员授权信息查询接口响应参数
type UserInfoShareRsp struct {
	Content struct {
		Code                 Code             `json:"code"`
		Msg                  string           `json:"msg"`
		UserNation           string           `json:"user_nation"`
		IdentityCardAddress  string           `json:"identity_card_address"`
		IdentityCardProvince string           `json:"identity_card_province"`
		IdentityCardCity     string           `json:"identity_card_city"`
		UserId               string           `json:"user_id"`
		Avatar               string           `json:"avatar"`
		Phone                string           `json:"phone"`
		City                 string           `json:"city"`
		Address              string           `json:"address"`
		Email                string           `json:"email"`
		UserName             string           `json:"user_name"`
		Mobile               string           `json:"mobile"`
		IsCertified          string           `json:"is_certified"`
		CertNo               string           `json:"cert_no"`
		Gender               string           `json:"gender"`
		DeliverAddresses     DeliverAddresses `json:"deliver_addresses"`
		SubCode              string           `json:"sub_code"`
		SubMsg               string           `json:"sub_msg"`
		AuthNo               string           `json:"auth_no"`
		Province             string           `json:"province"`
		NickName             string           `json:"nick_name"`
		IsStudentCertified   string           `json:"is_student_certified"`
		UserType             string           `json:"user_type"`
		UserStatus           string           `json:"user_status"`
	} `json:"alipay_user_info_share_response"`
	Sign string `json:"sign"`
}

// DeliverAddresses 收货地址
type DeliverAddresses struct {
	DeliverMobile         string `json:"deliver_mobile"`
	DeliverPhone          string `json:"deliver_phone"`
	Address               string `json:"address"`
	Zip                   string `json:"zip"`
	DeliverProvince       string `json:"deliver_province"`
	DeliverCity           string `json:"deliver_city"`
	DeliverArea           string `json:"deliver_area"`
	AddressCode           string `json:"address_code"`
	DeliverFullname       string `json:"deliver_fullname"`
	DefaultDeliverAddress string `json:"default_deliver_address"`
}

// OpenAuthTokenApp 换取应用授权令牌请求参数 https://docs.open.alipay.com/api_9/alipay.open.auth.token.app
type OpenAuthTokenApp struct {
	GrantType    string `json:"grant_type"` // 值为 authorization_code 时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"code"`
	RefreshToken string `json:"refresh_token"`
}

func (this OpenAuthTokenApp) APIName() string {
	return "alipay.open.auth.token.app"
}

func (this OpenAuthTokenApp) Params() map[string]string {
	var m = make(map[string]string)
	m["grant_type"] = this.GrantType
	if this.Code != "" {
		m["code"] = this.Code
	}
	if this.RefreshToken != "" {
		m["refresh_token"] = this.RefreshToken
	}
	return m
}

// OpenAuthTokenAppRsp 换取应用授权令牌响应参数 新版返回值 参见 https://opendocs.alipay.com/open/20160728150111277227/intro
type OpenAuthTokenAppRsp struct {
	Content struct {
		Code    Code             `json:"code"`
		Msg     string           `json:"msg"`
		SubCode Code             `json:"sub_code"`
		SubMsg  string           `json:"sub_msg"`
		Tokens  []*OpenAuthToken `json:"tokens"`
	} `json:"alipay_open_auth_token_app_response"`
	Sign string `json:"sign"`
}

type OpenAuthToken struct {
	AppAuthToken    string `json:"app_auth_token"`    // 授权令牌信息
	AppRefreshToken string `json:"app_refresh_token"` // 令牌信息
	AuthAppId       string `json:"auth_app_id"`       // 授权方应用id
	ExpiresIn       int64  `json:"expires_in"`        // 令牌有效期
	ReExpiresIn     int64  `json:"re_expires_in"`     // 有效期
	UserId          string `json:"user_id"`           // 支付宝用户标识
}

// AccountAuth 支付宝登录时, 帮客户端做参数签名, 返回授权请求信息字串接口请求参数 https://docs.open.alipay.com/218/105327/
type AccountAuth struct {
	Pid      string `json:"pid"`
	TargetId string `json:"target_id"`
	AuthType string `json:"auth_type"`
}

func (this AccountAuth) APIName() string {
	return "alipay.open.auth.sdk.code.get"
}

func (this AccountAuth) Params() map[string]string {
	var m = make(map[string]string)
	m["apiname"] = "com.alipay.account.auth"
	m["app_name"] = "mc"
	m["biz_type"] = "openservice"
	m["pid"] = this.Pid
	m["product_id"] = "APP_FAST_LOGIN"
	m["scope"] = "kuaijie"
	m["target_id"] = this.TargetId
	m["auth_type"] = this.AuthType
	return m
}
