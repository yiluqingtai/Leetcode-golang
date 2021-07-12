/**
 *
 * File:    leetcode6.go
 *          Z字形变换
 * 
 * Author:  yiluqingtai(1572236483@qq.com)
 *          Created on 21/7/12
 *          
 **/

 func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    l := 0
    if len(s) < numRows {
        l = len(s)
    } else {
        l = numRows
    }
    rows := make([]string, l)
    godown := false
    curRow := 0
    for _, c := range s {
        rows[curRow] += string(c)
        if curRow == 0 || curRow == numRows - 1 {
            godown = !godown
        }
        if godown {
            curRow++
        } else {
            curRow--
        }
    }
    res := ""
    for _, ss := range rows {
        res += ss
    }
    return res
}