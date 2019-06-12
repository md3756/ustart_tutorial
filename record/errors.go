package record

import "errors"

var (
	// //errCustomerExists This user already has a customer
	errRecordExists = errors.New("This record already has a customer")

	//errProblemLoadingCustomer This user already has a customer
	errProblemLoadingRecord = errors.New("There was a problem loading one or more records in this list")
)
