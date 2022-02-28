package streambird

type MagicLinkCreateReq struct {
	UserID            string                 `json:"user_id,omitempty"`
	ExpiresIn         int64                  `json:"expires_in,omitempty"`
	DeviceFingerprint map[string]interface{} `json:"device_fingerprint,omitempty"`
}

type MagicLinkCreateResp struct {
	UserID string `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type MagicLinkVerifyReq struct {
	Token            string `json:"token,omitempty"`
	SessionToken     string `json:"session_token,omitempty"`
	SessionExpiresIn int    `json:"session_expires_in,omitempty"`
}

type MagicLinkVerifyResp struct {
	UserID       string      `json:"user_id,omitempty"`
	MethodID     string      `json:"method_id,omitempty"`
	SessionToken string      `json:"session_token,omitempty"`
	Session      interface{} `json:"session,omitempty"`
}

type MagicLinkEmailSendReq struct {
	Email                   string `json:"email"`
	LoginRedirectURL        string `json:"login_redirect_url,omitempty"`
	RegistrationRedirectURL string `json:"registration_redirect_url,omitempty"`
	LoginExpiresIn          int    `json:"login_expires_in,omitempty"`
	RegistrationExpiresIn   int    `json:"registration_expires_in,omitempty"`
}

type MagicLinkEmailSendResp struct {
	StatusCode int    `json:"status_code,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	EmailID    string `json:"email_id,omitempty"`
}

type MagicLinkEmailLoginOrCreateReq struct {
	Email                   string `json:"email"`
	LoginRedirectURL        string `json:"login_redirect_url,omitempty"`
	RegistrationRedirectURL string `json:"registration_redirect_url,omitempty"`
	LoginExpiresIn          int    `json:"login_expires_in,omitempty"`
	RegistrationExpiresIn   int    `json:"registration_expires_in,omitempty"`
	RequiresVerification    bool   `json:"requires_verification,omitempty"`
}

type MagicLinkEmailLoginOrCreateResp struct {
	UserID      string `json:"user_id,omitempty"`
	EmailID     string `json:"email_id,omitempty"`
	UserCreated bool   `json:"user_created,omitempty"`
}
