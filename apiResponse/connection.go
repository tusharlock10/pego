package apiResponse

// Session stores data related to a current HiRezAPI Session.
type Session struct {
	RetMsg    string `json:"ret_msg"`
	SessionID string `json:"session_id"`
	Timestamp string `json:"timestamp"`
}

// HiRezServerStatus stores data related to a HiRez Service's current status.
type HiRezServerStatus struct {
	EntryDatetime string `json:"entry_datetime"`
	Environment   string `json:"environment"`
	LimitedAccess bool   `json:"limited_access"`
	Platform      string `json:"platform"`
	RetMsg        string `json:"ret_msg"`
	Status        string `json:"status"`
	Version       string `json:"version"`
}

// DataUsed stores data related to how many requests and sessions have been used of the daily cap.
type DataUsed struct {
	ActiveSessions     uint   `json:"Active_Sessions"`
	ConcurrentSessions uint   `json:"Concurrent_Sessions"`
	RequestDailyLimit  uint   `json:"Request_Limit_Daily"`
	SessionCap         uint   `json:"Session_Cap"`
	SessionTimeLimit   uint   `json:"Session_Time_Limit"`
	TotalRequestsToday uint   `json:"Total_Requests_Today"`
	TotalSessionsToday uint   `json:"Total_Sessions_Today"`
	RetMsg             string `json:"ret_msg"`
}
