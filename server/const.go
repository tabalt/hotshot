package server

// magic const
const (
	MAGIC_REQ = "\x00REQ"
	MAGIC_RES = "\x00RES"
)

// type const
const (
	TYPE_CAN_DO             = 1
	TYPE_CANT_DO            = 2
	TYPE_RESET_ABILITIES    = 3
	TYPE_PRE_SLEEP          = 4
	TYPE_NOOP               = 6
	TYPE_SUBMIT_JOB         = 7
	TYPE_JOB_CREATED        = 8
	TYPE_GRAB_JOB           = 9
	TYPE_NO_JOB             = 10
	TYPE_JOB_ASSIGN         = 11
	TYPE_WORK_STATUS        = 12
	TYPE_WORK_COMPLETE      = 13
	TYPE_WORK_FAIL          = 14
	TYPE_GET_STATUS         = 15
	TYPE_ECHO_REQ           = 16
	TYPE_ECHO_RES           = 17
	TYPE_SUBMIT_JOB_BG      = 18
	TYPE_ERROR              = 19
	TYPE_STATUS_RES         = 20
	TYPE_SUBMIT_JOB_HIGH    = 21
	TYPE_SET_CLIENT_ID      = 22
	TYPE_CAN_DO_TIMEOUT     = 23
	TYPE_ALL_YOURS          = 24
	TYPE_WORK_EXCEPTION     = 25
	TYPE_OPTION_REQ         = 26
	TYPE_OPTION_RES         = 27
	TYPE_WORK_DATA          = 28
	TYPE_WORK_WARNING       = 29
	TYPE_GRAB_JOB_UNIQ      = 30
	TYPE_JOB_ASSIGN_UNIQ    = 31
	TYPE_SUBMIT_JOB_HIGH_BG = 32
	TYPE_SUBMIT_JOB_LOW     = 33
	TYPE_SUBMIT_JOB_LOW_BG  = 34
	TYPE_SUBMIT_JOB_SCHED   = 35
	TYPE_SUBMIT_JOB_EPOCH   = 36
)
