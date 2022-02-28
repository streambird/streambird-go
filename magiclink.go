package streambird

type MagicLinkCreateReq struct {
	UserID            string                 `json:"user_id,omitempty"`
	ExpiresIn         int64                  `json:"expires_in,omitempty"`
	DeviceFingerprint map[string]interface{} `json:"DeviceFingerprint,omitempty"`
}

type MagicLinkCreateResp struct {
	Token string `json:"token"`
}
