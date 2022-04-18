package streambird

type OTPVerifyRequest struct {
	MethodID         string `json:"method_id"`
	OTP              string `json:"otp"`
	SessionToken     string `json:"session_token,omitempty"`
	SessionExpiresIn int32  `json:"session_expires_in,omitempty"`
}

type OTPVerifyResponse struct {
	UserID       string  `json:"user_id,omitempty"`
	MethodID     string  `json:"method_id,omitempty"`
	SessionToken string  `json:"session_token,omitempty"`
	SessionJWT   string  `json:"session_jwt,omitempty"`
	Session      Session `json:"session,omitempty"`
}

type OTPSendBySMSRequest struct {
	PhoneNumber       string                 `json:"phone_number"`
	ExpiresIn         int                    `json:"expires_in"`
	DeviceFingerprint map[string]interface{} `json:"device_fingerprint,omitempty"`
}

type OTPSendBySMSResponse struct {
	PhoneNumberID string `json:"phone_number_id"`
	UserID        string `json:"user_id"`
}

type OTPSendByEmailRequest struct {
	Email             string                 `json:"email"`
	ExpiresIn         int                    `json:"expires_in"`
	DeviceFingerprint map[string]interface{} `json:"device_fingerprint,omitempty"`
}

type OTPSendByEmailResponse struct {
	EmailID string `json:"email_id"`
	UserID  string `json:"user_id"`
}

type OTPLoginOrCreateByEmailRequest struct {
	Email                string                 `json:"email"`
	ExpiresIn            int                    `json:"expires_in"`
	RequiresVerification bool                   `json:"requires_verification,omitempty"`
	DeviceFingerprint    map[string]interface{} `json:"device_fingerprint,omitempty"`
}

type OTPLoginOrCreateByEmailResponse struct {
	UserID      string `json:"user_id"`
	UserCreated bool   `json:"user_created"`
	Status      string `json:"status"`
	EmailID     string `json:"email_id"`
}

type OTPLoginOrCreateBySMSRequest struct {
	PhoneNumber          string                 `json:"phone_number"`
	ExpiresIn            int                    `json:"expires_in"`
	RequiresVerification bool                   `json:"requires_verification,omitempty"`
	DeviceFingerprint    map[string]interface{} `json:"device_fingerprint,omitempty"`
}

type OTPLoginOrCreateBySMSResponse struct {
	UserID        string `json:"user_id"`
	UserCreated   bool   `json:"user_created"`
	Status        string `json:"status"`
	PhoneNumberID string `json:"phone_number_id"`
}
