# ezlog
[![GoDoc](https://godoc.org/github.com/mohae/ezlog?status.svg)](https://godoc.org/github.com/mohae/ezlog)[![Build Status](https://travis-ci.org/mohae/ezlog.png)](https://travis-ci.org/mohae/ezlog)

ezlog is a leveled log library with an api similar to stdlib's log package. The main difference is that stdlib's log.Output method isn't exposed and there are additional log line output methods corresponding to the level of the log lines that they are writing.

Ezlog provides leveled log lines using the Error[f|ln], Info[f|ln], and Debug[f|ln] methods. Log lines will be prefixed with either the level name or the first character of the level name, depending on the logger's configuration. Log lines are only written for log levels less than or equal to the logger's configured level, all other lines are discarded. If the logger's level is set to LogNone, all output will be discarded.

Fatal[f|ln] and Panic[f|ln] methods that prefix the log lines with "FATAL" or "F" and "PANIC" or "P", respectively and then exits or panics are provided.

The log flags are the same as stdlib's constants for log flags.

## Severity Levels

* none
* error
* info
* debug
