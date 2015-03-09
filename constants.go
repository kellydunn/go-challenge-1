package drum

// This file defines some constants specific to the drum package

// The begining identifier string of a splice file
var SpliceFileHeader = "SPLICE"

// The length of the SPLICE file header at the begining of each file in bytes
var SpliceFileSize = len(SpliceFileHeader)

// The length of the body of the splice file in bytes
var FileSize = 8

// The length of the version string of the splice file in bytes
var VersionSize = 32
var TempoSize = 4
var TrackIDSize = 1
var TrackNameSize = 4
var StepSequenceSize = 16
var TrackSize = TrackIDSize + TrackNameSize + StepSequenceSize
var EmptyByte = "\x00"
