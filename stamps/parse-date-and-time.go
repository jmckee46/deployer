package stamps

import "github.com/jmckee46/deployer/flaw"

// ParseDateAndTime Convert ("2006-Jan-02", YYYYhMonhDD, "1504", HHMM) to UTC Timestamp
func ParseDateAndTime(dateString, dateFormat, timeString, timeFormat string) (*Timestamp, flaw.Flaw) {
	format := dateFormat + "T" + timeFormat
	dateAndTimeString := dateString + "T" + timeString

	return ParseStamp(dateAndTimeString, format)
}
