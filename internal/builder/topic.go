package builder

// Topic .
type Topic struct {
    RegistrationId        string `json:"registrationId,omitempty"`
    Aliases               string `json:"aliases,omitempty"`
    UserAccount           string `json:"userAccount,omitempty"`
    Topic                 string `json:"topic"`
    Topics                string `json:"topics"`
    Category              string `json:"category"` //ios使用
    RestrictedPackageName string `json:"restrictedPackageName,omitempty"`
}
