// Code generated by "stringer -type ResumeTransferErrorReason"; DO NOT EDIT.

package rpc

import "fmt"

const _ResumeTransferErrorReason_name = "ResumeTransferErrorReasonNotImplementedResumeTransferErrorReasonDisabledResumeTransferErrorReasonZFSErrorPermanentResumeTransferErrorReasonZFSErrorMaybeTemporary"

var _ResumeTransferErrorReason_index = [...]uint8{0, 39, 72, 114, 161}

func (i ResumeTransferErrorReason) String() string {
	if i >= ResumeTransferErrorReason(len(_ResumeTransferErrorReason_index)-1) {
		return fmt.Sprintf("ResumeTransferErrorReason(%d)", i)
	}
	return _ResumeTransferErrorReason_name[_ResumeTransferErrorReason_index[i]:_ResumeTransferErrorReason_index[i+1]]
}