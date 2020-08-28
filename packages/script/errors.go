/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package script

import "errors"

const (
	eContractLoop    = `there is loop in %s contract`
	eSysVar          = `system variable $%s cannot be changed`
	eTypeParam       = `parameter %d has wrong type`
	eUndefinedParam  = `%s is not defined`
	eUnknownContract = `unknown contract %s`
	eWrongParams     = `function %s must have %d parameters`
	eArrIndex        = `index of array cannot be type %s`
	eMapIndex        = `index of map cannot be type %s`
	eUnknownIdent    = `unknown identifier %s`
	eWrongVar        = `wrong var %v`
	eDataType        = `expecting type of the data field [Ln:%d Col:%d]`
	eDataName        = `expecting name of the data field [Ln:%d Col:%d]`
	eDataTag         = `unexpected tag [Ln:%d Col:%d]`
)

var (
	errContractPars    = errors.New(`wrong contract parameters`)
	errWrongCountPars  = errors.New(`wrong count of parameters`)
	errDivZero         = errors.New(`divided by zero`)
	errOper            = errors.New(`unexpected operator; expecting operand`)
)