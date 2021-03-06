// go-config - Library for reading cacophony config files.
// Copyright (C) 2018, The Cacophony Project
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package config

func init() {
	allSections[WindowsKey] = section{
		key:         WindowsKey,
		mapToStruct: windowsMapToStruct,
		validate:    noValidateFunc,
	}
}

const WindowsKey = "windows"

type Windows struct {
	StartRecording string `mapstructure:"start-recording"`
	StopRecording  string `mapstructure:"stop-recording"`
	PowerOn        string `mapstructure:"power-on"`
	PowerOff       string `mapstructure:"power-off"`
}

func DefaultWindows() Windows {
	return Windows{
		StartRecording: "-30m",
		StopRecording:  "+30m",
		PowerOn:        "12:00",
		PowerOff:       "12:00",
	}
}

func noValidateFunc(s interface{}) error {
	return nil
}

func windowsMapToStruct(m map[string]interface{}) (interface{}, error) {
	var s Windows
	if err := decodeStructFromMap(&s, m, nil); err != nil {
		return nil, err
	}
	return s, nil
}
