package goldscore_test

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-16 21:10:42
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-27 23:20:28
 */

import (
	"testing"

	"github.com/monitor1379/golds/goldscore"
)

func TestBtoi64(t *testing.T) {
	testCases := []struct {
		Input       string
		GroundTruth int64
	}{
		{"0", 0},
		{"1", 1},
		{"123456", 123456},
		{"-1", -1},
		{"-123456", -123456},
	}

	for _, testCase := range testCases {
		output, err := goldscore.Btoi64([]byte(testCase.Input))
		if err != nil {
			t.Errorf("error: %s. testCase: %+v\n", err, testCase)
		}
		if output != testCase.GroundTruth {
			t.Errorf("expect: %+v, got %+v\n", testCase.GroundTruth, output)
		}
	}
}
