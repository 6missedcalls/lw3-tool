package errors

// Here are constants that will be used for errors throughout the command line
const (
	FailedPathCreation = "Failed to create path, please rerun the command."
	FailedFileCreation = "Failed to create file, please rerun the command."
	FailedFileWrite    = "Failed to write to file, please rerun the command."

	FailedHomedirRetrieval = "Failed to retrieve home directory, please report this issue on GitHub."

	NilName   = "Name cannot be empty"
	NilTicker = "Ticker cannot be empty"
)
