/**
Copyright 2019 Pezhman Kasraee

This file is part of go-pklog.

go-pklog is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

go-pklog is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package pklog

import (
	"github.com/pezhmankasraee/go-pklog/baseconfig"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

const productionLifeCycle string = "PRODUCTION"
const errorMessageEnding string = baseconfig.Reset + " ] "

// CreateLog prints out a log message depends on application life cycle.
// This module reads first the result of environment parameter REST_APP.
// If REST_APP is equal to "PRODUCTION", the logs will be generated in /var/log/syslog otherwise
// the logs will be displayed with color on standard display.
// logLevel accepts four types of log level: FatalError(0), Error (10), Warning (20) and Info (30).
// s is the log message which is of type string.
func CreateLog(logLevel int, s string) {

	softwareLifeCycle := os.Getenv("REST_APP")
	appName := filepath.Base(os.Args[0])

	logwriter, err := syslog.New(syslog.LOG_DEBUG, appName)
	if err != nil {
		panic(" [ panic ] : pklog cannot write on syslog")
	}

	switch logLevel {
	case FatalError:
		if softwareLifeCycle == productionLifeCycle {
			log.SetPrefix("[ Fatal Error ] ")
		} else {
			log.SetPrefix("[ " + baseconfig.FGMagneta + "Fatal Error" + errorMessageEnding)
		}
		break
	case Error:
		if softwareLifeCycle == productionLifeCycle {
			log.SetPrefix("[ Error ] ")
		} else {
			log.SetPrefix("[ " + baseconfig.FGRed + "Error" + errorMessageEnding)
		}
		break
	case Warning:
		if softwareLifeCycle == productionLifeCycle {
			log.SetPrefix("[ Warning ] ")
		} else {
			log.SetPrefix("[ " + baseconfig.FGYellow + "Warning" + errorMessageEnding)
		}
		break
	case Information:
		if softwareLifeCycle == productionLifeCycle {
			log.SetPrefix("[ Info ] ")
		} else {
			log.SetPrefix("[ " + baseconfig.FGTeal + "Info" + errorMessageEnding)
		}
		break
	default:
		panic("[ Panic ] Log Level is invalid")
	}

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.LUTC)
	if logLevel < 100 {
		log.Fatal(s)
	}

	if softwareLifeCycle == productionLifeCycle {
		log.SetOutput(logwriter)
	} else {
		log.SetOutput(os.Stdout)
	}

	log.Println(s)
}
