//给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。 
//
// 输入为三个整数：day、month 和 year，分别表示日、月、年。 
//
// 您返回的结果必须是这几个值中的一个 {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", 
//"Friday", "Saturday"}。 
//
// 
//
// 示例 1： 
//
// 输入：day = 31, month = 8, year = 2019
//输出："Saturday"
// 
//
// 示例 2： 
//
// 输入：day = 18, month = 7, year = 1999
//输出："Sunday"
// 
//
// 示例 3： 
//
// 输入：day = 15, month = 8, year = 1993
//输出："Sunday"
// 
//
// 
//
// 提示： 
//
// 
// 给出的日期一定是在 1971 到 2100 年之间的有效日期。 
// 
// Related Topics 数学 👍 97 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}

func dayOfTheWeek(day, month, year int) string {
days := 0
// 输入年份之前的年份的天数贡献
days += 365*(year-1971) + (year-1969)/4
// 输入年份中，输入月份之前的月份的天数贡献
for _, d := range monthDays[:month-1] {
days += d
}
if month >= 3 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
days++
}
// 输入月份中的天数贡献
days += day
return week[(days+3)%7]
}
//leetcode submit region end(Prohibit modification and deletion)
