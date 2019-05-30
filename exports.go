package main

type Meta struct {
	Devices []CuptiActivityKindDevice `json:"devices"`
}

type Infomation struct {
	Events   []*Event `json:"traceEvents"`
	TimeUnit string   `json:"displayTimeUnit"`
	Meta     Meta     `json:"meta"`
}

func NewInfomation() *Infomation {
	info := &Infomation{TimeUnit: "ns"}
	return info
}

// https://docs.google.com/document/d/1CvAClvFfyA5R-PhYUmn5OOQtYMH4h6I0nSsKchNAySU/edit
type Event struct {
	Name      string            `json:"name"` // The name of the event, as displayed in Trace Viewer
	Type      string            `json:"ph"`   // The event type. This is a single character which changes depending on the type of event being output. The valid values are listed in the table below. We will discuss each phase type below.
	Category  string            `json:"cat"`  // The event categories. This is a comma separated list of categories for the event. The categories can be used to hide events in the Trace Viewer UI.
	Timestamp int64             `json:"ts"`   // The tracing clock timestamp of the event. The timestamps are provided at microsecond granularity.
	PID       string            `json:"pid"`  // The process ID for the process that output this event.
	TID       string            `json:"tid"`  //The thread ID for the thread that output this event.
	Duration  int64             `json:"dur"`  //
	Args      map[string]string `json:"args"` // Any arguments provided for the event. Some of the event types have required argument fields, otherwise, you can put any information you wish in here. The arguments are displayed in Trace Viewer when you view an event in the analysis section.
}

func NewEvent() *Event {
	rsl := &Event{}
	rsl.Args = make(map[string]string)
	return rsl
}

func (e *Event) Copy() *Event {
	rsl := &Event{
		Name:      e.Name,
		Type:      e.Type,
		Category:  e.Category,
		Timestamp: e.Timestamp,
		PID:       e.PID,
		TID:       e.TID,
		Duration:  e.Duration,
	}
	rsl.Args = make(map[string]string)
	for k, v := range e.Args {
		rsl.Args[k] = v
	}

	return rsl
}
