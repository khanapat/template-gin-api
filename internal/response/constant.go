package response

const (
	SuccessCode                uint64 = 200
	ErrInvalidRequestCode      uint64 = 1000
	ErrRequestExpireCode       uint64 = 1001
	ErrBasicAuthenticationCode uint64 = 4007
	ErrDatabaseCode            uint64 = 5000
	ErrRedisCode               uint64 = 5001
	ErrOperationCode           uint64 = 5002
	ErrThirdPartyCode          uint64 = 5004
)

const (
	SuccessMessageEN           string = "Success."
	ErrInvalidRequestMessageEN string = "Invalid request."
	ErrInternalServerMessageEN string = "Internal server error."
	// BasicAuthen
	ErrBasicAuthenticationMessageEN string = "Authentication failed."
	// Desc
	ErrContactAdminDescEN   string = "Please contact administrator for more information."
	ErrThirdPartyDescEN     string = "Service is unavailable. Please try again later."
	ErrAuthenticationDescEN string = "Unable to access data. Please check user & password."
)

const (
	SuccessMessageTH           string = "สำเร็จ."
	ErrInternalServerMessageTH string = "มีข้อผิดพลาดภายในเซิร์ฟเวอร์."
	// BasicAuthen
	ErrBasicAuthenticationMessageTH string = "ยืนยันตัวตนล้มเหลว"
	// Desc
	ErrContactAdminDescTH   string = "กรุณาติดต่อเจ้าหน้าที่ดูแลระบบเพื่อรับข้อมูลเพิ่มเติม."
	ErrThirdPartyDescTH     string = "ไม่สามารถใช้บริการได้. กรุณาทำรายการใหม่อีกครั้งภายหลัง."
	ErrAuthenticationDescTH string = "ไม่สามารถเข้าถึงข้อมูลได้. กรุณาตรวจสอบรหัสผู้ใช้งานใหม่อีกครั้ง."
)

const (
	ValidateFieldError string = "Invalid Parameters"
	OperationError     string = "Invalid Operation"
)
