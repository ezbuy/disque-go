package disque

import (
	"sort"
	"time"
)

const NO_INDEX int = 0

type fetchOption interface {
	Name() string
	Args() []interface{}
	Index() int
}

type timeoutOption struct {
	duration time.Duration
}

func (to timeoutOption) Name() string {
	return "TIMEOUT"
}

func (to timeoutOption) Args() []interface{} {
	return []interface{}{int64(to.duration.Seconds()) * 1000}
}

func (to timeoutOption) Index() int { return 1 }

type countOption struct {
	max int
}

func (c countOption) Name() string {
	return "COUNT"
}

func (c countOption) Args() []interface{} {
	return []interface{}{c.max}
}

func (c countOption) Index() int { return 2 }

type withCounterOption struct {
}

func (c withCounterOption) Name() string {
	return "WITHCOUNTERS"
}

func (c withCounterOption) Args() []interface{} {
	return []interface{}{}
}

func (c withCounterOption) Index() int { return 3 }

type fromOption struct {
	queue string
}

func (f fromOption) Name() string {
	return "FROM"
}

func (f fromOption) Args() []interface{} {
	return []interface{}{f.queue}
}

func (f fromOption) Index() int { return 4 }

// noHangOption asks the command to not block even if there are no jobs in all the specified queues. This way the caller can just check if there are available jobs without blocking at all.
// SCOPE: GETJOB
// See more at https://github.com/antirez/disque#getjob-nohang-timeout-ms-timeout-count-count-withcounters-from-queue1-queue2--queuen
type noHangOption struct{}

func (nh noHangOption) Name() string {
	return "NOHANG"
}

func (nh noHangOption) Args() []interface{} {
	return []interface{}{}
}

func (nh noHangOption) Index() int { return 0 }

// withFetchOption defines a set of options when fetch the JOB from disque.
// COMMAND: GETJOB [TIMEOUT <ms-timeout>] [COUNT <count>] [WITHCOUNTERS] FROM [queueName]
func withFetchOption(options ...fetchOption) []fetchOption {
	sort.SliceStable(options, func(i, j int) bool {
		return options[i].Index() < options[j].Index()
	})
	return options
}
