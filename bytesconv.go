package golds

import "strconv"

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-16 21:08:08
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-16 21:09:01
 */

// Btoi64: https://github.com/CodisLabs/codis/blob/release3.2/pkg/proxy/redis/decoder.go
func Btoi64(b []byte) (int64, error) {
	if len(b) != 0 && len(b) < 10 {
		var neg, i = false, 0
		switch b[0] {
		case '-':
			neg = true
			fallthrough
		case '+':
			i++
		}
		if len(b) != i {
			var n int64
			for ; i < len(b) && b[i] >= '0' && b[i] <= '9'; i++ {
				n = int64(b[i]-'0') + n*10
			}
			if len(b) == i {
				if neg {
					n = -n
				}
				return n, nil
			}
		}
	}

	if n, err := strconv.ParseInt(string(b), 10, 64); err != nil {
		return 0, err
	} else {
		return n, nil
	}
}
