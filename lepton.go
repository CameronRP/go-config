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

const LeptonKey = "lepton"

type Lepton struct {
	SPISpeed    int64  `mapstructure:"spi-speed"`
	FrameOutput string `mapstructure:"frame-output"`
}

func DefaultLepton() Lepton {
	return Lepton{
		SPISpeed:    2000000,
		FrameOutput: "/var/run/lepton-frames",
	}
}
