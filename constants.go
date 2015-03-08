package drum

// This file defines some constants specific to the drum package

var SPLICE_FILE_SIZE int = 6
var FILE_SIZE int = 8
var VERSION_SIZE int = 32
var TEMPO_SIZE int = 4
var TRACK_ID_SIZE int = 1
var TRACK_NAME_SIZE int = 4
var STEP_SEQUENCE_SIZE = 16
var TRACK_SIZE = TRACK_ID_SIZE + TRACK_NAME_SIZE + STEP_SEQUENCE_SIZE
var EMPTY_BYTE = "\x00"
