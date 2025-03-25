package enums

import "anonsurvey/consts"

const (
	UserTypeAdmin    = consts.U_CREATE | consts.U_DELETE
	UserTypeSurveyor = consts.S_READ | consts.S_WRITE
)
