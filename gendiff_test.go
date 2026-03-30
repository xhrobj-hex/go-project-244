package code

import "testing"

func Test_GenDiff_OK(t *testing.T) {
	// 	file {
	//   "host": "hexlet.io",
	//   "timeout": 50,
	//   "proxy": "123.234.53.22",
	//   "follow": false
	// }

	//	{
	//	  "timeout": 20,
	//	  "verbose": true,
	//	  "host": "hexlet.io"
	//	}

	// {
	//   - follow: false
	//     host: hexlet.io
	//   - proxy: 123.234.53.22
	//   - timeout: 50
	//   + timeout: 20
	//   + verbose: true
	// }
}
