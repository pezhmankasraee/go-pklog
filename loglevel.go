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

// Log level defines the severity of log report.
// FatalError prints out a message and also causes os.Exit(1).
// Error, Warning and Info prints out only a message.
const (
	FatalError  int = 0
	Error       int = 100
	Warning     int = 200
	Information int = 300
)
