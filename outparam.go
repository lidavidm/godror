// Copyright 2026 The Godror Authors
//
//
// SPDX-License-Identifier: UPL-1.0 OR Apache-2.0

package godror

/*
#include <stdlib.h>
#include "dpiImpl.h"
*/
import "C"

// OracleType is the type used to represent Oracle types in ODPI-C.
type OracleType = C.dpiOracleTypeNum

// Exported Oracle type variables (cannot be constants since they reference C values).
var (
	OracleTypeNumber       OracleType = C.DPI_ORACLE_TYPE_NUMBER
	OracleTypeNativeFloat  OracleType = C.DPI_ORACLE_TYPE_NATIVE_FLOAT
	OracleTypeNativeDouble OracleType = C.DPI_ORACLE_TYPE_NATIVE_DOUBLE
	OracleTypeVarchar      OracleType = C.DPI_ORACLE_TYPE_VARCHAR
	OracleTypeRaw          OracleType = C.DPI_ORACLE_TYPE_RAW
	OracleTypeBoolean      OracleType = C.DPI_ORACLE_TYPE_BOOLEAN
	OracleTypeDate         OracleType = C.DPI_ORACLE_TYPE_DATE
	OracleTypeTimestamp    OracleType = C.DPI_ORACLE_TYPE_TIMESTAMP
	OracleTypeTimestampTZ  OracleType = C.DPI_ORACLE_TYPE_TIMESTAMP_TZ
	OracleTypeTimestampLTZ OracleType = C.DPI_ORACLE_TYPE_TIMESTAMP_LTZ
	OracleTypeBlob         OracleType = C.DPI_ORACLE_TYPE_BLOB
	OracleTypeClob         OracleType = C.DPI_ORACLE_TYPE_CLOB
	OracleTypeNclob        OracleType = C.DPI_ORACLE_TYPE_NCLOB
	OracleTypeLongVarchar  OracleType = C.DPI_ORACLE_TYPE_LONG_VARCHAR
	OracleTypeLongRaw      OracleType = C.DPI_ORACLE_TYPE_LONG_RAW
	OracleTypeJson         OracleType = C.DPI_ORACLE_TYPE_JSON
)

// OutParam represents an OUT or INOUT parameter for use with Oracle stored procedures.
// It allows callers to get raw DPI data from OUT parameters without going through
// sql.Out which forces type-specific conversions.
type OutParam struct {
	// In holds the input value for INOUT parameters. May be nil for pure OUT parameters.
	In any
	// OracleType specifies the Oracle type to bind this parameter as.
	OracleType C.dpiOracleTypeNum
	// Result holds the output value after statement execution.
	Result OutParamResult
}

// OutParamResult holds the result of an OUT parameter after execution.
type OutParamResult struct {
	IsNull    bool
	Bytes     []byte
	Timestamp *OutParamTimestamp
	Float     float32
	Double    float64
}

// OutParamTimestamp holds a timestamp value returned from an OUT parameter.
type OutParamTimestamp struct {
	Year                       int16
	Month, Day, Hour, Min, Sec uint8
	FSec                       uint32
	TZHourOffset, TZMinOffset  int8
}

// isTimestampOracleType returns true if the given Oracle type is a date or timestamp type.
func isTimestampOracleType(t C.dpiOracleTypeNum) bool {
	return t == C.DPI_ORACLE_TYPE_DATE ||
		t == C.DPI_ORACLE_TYPE_TIMESTAMP ||
		t == C.DPI_ORACLE_TYPE_TIMESTAMP_TZ ||
		t == C.DPI_ORACLE_TYPE_TIMESTAMP_LTZ
}

// outParamNativeType maps an Oracle type to the native type used for retrieving
// OUT parameter values.
func outParamNativeType(t C.dpiOracleTypeNum) C.dpiNativeTypeNum {
	switch {
	case isTimestampOracleType(t):
		return C.DPI_NATIVE_TYPE_TIMESTAMP
	case t == C.DPI_ORACLE_TYPE_NATIVE_FLOAT:
		return C.DPI_NATIVE_TYPE_FLOAT
	case t == C.DPI_ORACLE_TYPE_NATIVE_DOUBLE:
		return C.DPI_NATIVE_TYPE_DOUBLE
	}
	return C.DPI_NATIVE_TYPE_BYTES
}
