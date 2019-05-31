// Copyright (C) 2019  Patrick Wieschollek, <mail@patwie.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ianlancetaylor/demangle"
	flags "github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

var DemangledNamesMap = make(map[int64]string)

func GetRuntimeEvents(db *sqlx.DB) []*Event {
	activities := []CuptiActivityKindRuntime{}
	err := db.Select(&activities, "SELECT * FROM CUPTI_ACTIVITY_KIND_RUNTIME")
	if err != nil {
		panic(err)
	}
	events := []*Event{}

	for _, activity := range activities {
		event := NewEvent()
		if val, ok := CbidTable[activity.Cbid]; ok {
			event.Name = val
		} else {
			event.Name = fmt.Sprintf("<unkown %v>", activity.Cbid)
		}
		event.Type = "X"
		event.Category = "cuda"
		event.Timestamp = activity.Start
		event.Duration = activity.End - activity.Start
		event.TID = fmt.Sprintf("Thread %v: Runtime API", activity.ThreadID)
		event.PID = fmt.Sprintf("[%v] Process", activity.ProcessID)
		events = append(events, event)
	}

	return events
}

func GetMemcpyEvents(db *sqlx.DB) []*Event {
	activities := []CuptiActivityKindMemcpy{}
	err := db.Select(&activities, "SELECT * FROM CUPTI_ACTIVITY_KIND_MEMCPY")
	if err != nil {
		panic(err)
	}
	events := []*Event{}

	for _, activity := range activities {
		event := NewEvent()

		copyKind := FindInMap(activity.CopyKind, ActivityMemcpyKind)

		// CUPTI_ACTIVITY_FLAG_MEMCPY_ASYNC = 1<<0
		// Indicates the activity represents an asynchronous memcpy operation.
		// Valid for CUPTI_ACTIVITY_KIND_MEMCPY.
		flags := "sync"
		if activity.Flags&(1<<0) == 1 {
			flags = "async"
		}

		event.Name = fmt.Sprintf("Memcpy %v [%v]", copyKind, flags)
		event.Type = "X"
		event.Category = "cuda"
		event.Timestamp = activity.Start
		event.Duration = activity.End - activity.Start
		event.TID = fmt.Sprintf("MemCpy (%v)", copyKind)
		event.PID = fmt.Sprintf("[%v:%v] Overview", activity.DeviceID, activity.ContextID)
		event.Args["Src"] = FindInMap(activity.SrcKind, ActivityMemoryKind)
		event.Args["Dst"] = FindInMap(activity.DstKind, ActivityMemoryKind)
		events = append(events, event)

	}

	return events
}

func GetMemcpy2Events(db *sqlx.DB) []*Event {
	activities := []CuptiActivityKindMemcpy2{}
	err := db.Select(&activities, "SELECT * FROM CUPTI_ACTIVITY_KIND_MEMCPY2")
	if err != nil {
		panic(err)
	}
	events := []*Event{}

	for _, activity := range activities {
		event := NewEvent()

		copyKind := FindInMap(activity.CopyKind, ActivityMemcpyKind)

		// CUPTI_ACTIVITY_FLAG_MEMCPY_ASYNC = 1<<0
		// Indicates the activity represents an asynchronous memcpy operation.
		// Valid for CUPTI_ACTIVITY_KIND_MEMCPY.
		flags := "sync"
		if activity.Flags&(1<<0) == 1 {
			flags = "async"
		}

		event.Name = fmt.Sprintf("Memcpy %v [%v]", copyKind, flags)
		event.Type = "X"
		event.Category = "cuda"
		event.Timestamp = activity.Start
		event.Duration = activity.End - activity.Start
		event.TID = fmt.Sprintf("MemCpy (%v) %v -> %v", copyKind, activity.SrcDeviceID, activity.DstDeviceID)
		event.PID = fmt.Sprintf("[%v:%v] Overview", activity.DeviceID, activity.ContextID)
		events = append(events, event)
	}

	return events
}

func GetMemsetEvents(db *sqlx.DB) []*Event {
	activities := []CuptiActivityKindMemset{}
	err := db.Select(&activities, "SELECT * FROM CUPTI_ACTIVITY_KIND_MEMSET")
	if err != nil {
		panic(err)
	}
	events := []*Event{}

	for _, activity := range activities {
		event := NewEvent()

		event.TID = "Memset"
		event.Name = "Memset"
		event.Type = "X"
		event.Category = "cuda"
		event.Timestamp = activity.Start
		event.Duration = activity.End - activity.Start
		event.PID = fmt.Sprintf("[%v:%v] Overview", activity.DeviceID, activity.ContextID)
		event.Args["bytes"] = fmt.Sprintf("%v", activity.Bytes)
		events = append(events, event)
	}

	return events
}

func GetConcurrentKernelEvents(db *sqlx.DB) []*Event {
	activities := []CuptiActivityKindConcurrentKernel{}
	err := db.Select(&activities, "SELECT * FROM CUPTI_ACTIVITY_KIND_CONCURRENT_KERNEL")
	if err != nil {
		panic(err)
	}
	events := []*Event{}

	for _, activity := range activities {
		event := NewEvent()
		event.Name = fmt.Sprintf("<unkown %v>", activity.Name)

		if val, ok := DemangledNamesMap[activity.Name]; ok {
			event.Name = val
		}

		event.TID = "Compute"
		event.Type = "X"
		event.Category = "cuda"
		event.Timestamp = activity.Start
		event.Duration = activity.End - activity.Start

		event.PID = fmt.Sprintf("[%v:%v] Overview", activity.DeviceID, activity.ContextID)
		event.Args["Grid"] = fmt.Sprintf("[%v,%v,%v]", activity.GridX, activity.GridY, activity.GridZ)
		event.Args["Block"] = fmt.Sprintf("[%v,%v,%v]", activity.BlockX, activity.BlockY, activity.BlockZ)
		event.Args["SharedMemoryConfig"] = fmt.Sprintf("%v", activity.SharedMemoryConfig)
		event.Args["StaticSharedMemory"] = fmt.Sprintf("%v", activity.StaticSharedMemory)
		event.Args["DynamicSharedMemory"] = fmt.Sprintf("%v", activity.DynamicSharedMemory)
		event.Args["StreamID"] = fmt.Sprintf("%v", activity.StreamID)
		event.Args["RegistersPerThread"] = fmt.Sprintf("%v", activity.RegistersPerThread)
		events = append(events, event)

		// There is an issue, but we keep this enabled.
		// https://github.com/catapult-project/catapult/issues/4636
		computeEvent := event.Copy()
		computeEvent.PID = "Compute"
		computeEvent.TID = event.Name
		events = append(events, computeEvent)

		streamEvent := event.Copy()
		streamEvent.PID = "Streams"
		streamEvent.TID = fmt.Sprintf("Stream %v", activity.StreamID)
		events = append(events, streamEvent)

	}

	return events
}

func GetSynchronizationEvents(db *sqlx.DB) []*Event {
	activities := []CuptiActivityKindSynchronization{}
	err := db.Select(&activities, "SELECT * FROM CUPTI_ACTIVITY_KIND_SYNCHRONIZATION")
	if err != nil {
		panic(err)
	}
	events := []*Event{}

	for _, activity := range activities {
		event := NewEvent()
		event.Name = FindInMap(activity.Type, ActivitySynchronizationType)
		event.Type = "X"
		event.Category = "cuda"
		event.Timestamp = activity.Start
		event.Duration = activity.End - activity.Start

		event.TID = "Synchronize"
		event.PID = "Synchronization"

		event.Args["CudaEventID"] = fmt.Sprintf("%v", activity.CudaEventID)
		events = append(events, event)
	}

	return events
}

var log = logrus.New()

func main() {

	var opts struct {
		OutputFile string `short:"o" long:"output" default:"[nvvpfile].json" description:"output file for Chrome tracing" required:"false" name:"output file"`
		Verbose    bool   `short:"v" long:"verbose"  description:"verbose logging" required:"false" name:"verbose"`
		Pretty     bool   `short:"p" long:"pretty"  description:"ident and prettify JSON output" required:"false" name:"pretty"`
		Override   bool   `short:"f"  description:"override output file if exists" required:"false" name:"override"`
		Args       struct {
			NVVPFile string `positional-arg-name:"file" description:"output from nvprof, e.g., 'nvprof -o [file] [your-app]'"`
		} `positional-args:"true" required:"1"`
	}

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: false,
	})
	log.SetOutput(os.Stdout)

	_, err := flags.Parse(&opts)
	if err == nil {

		if opts.Verbose {
			log.SetLevel(logrus.DebugLevel)
		}

		if _, err := os.Stat(opts.Args.NVVPFile); os.IsNotExist(err) {
			log.Fatalln(err)
		}

		if opts.OutputFile == "[nvvpfile].json" {
			opts.OutputFile = fmt.Sprintf("%s.json", strings.TrimSuffix(filepath.Base(opts.Args.NVVPFile), filepath.Ext(opts.Args.NVVPFile)))
			log.Debugln("Set output to", opts.OutputFile)
		}

		if _, err := os.Stat(opts.OutputFile); !os.IsNotExist(err) {
			if !opts.Override {
				log.Fatalln(fmt.Sprintf("file %s already exists use -f to override", opts.OutputFile))
			}
		}

		log.Debugln("Open sqlite3 database", opts.Args.NVVPFile)
		db, err := sqlx.Connect("sqlite3", opts.Args.NVVPFile)
		if err != nil {
			log.Fatalln(err)
		}

		stringTable := []StringTable{}
		err = db.Select(&stringTable, "SELECT _id_ as id, value FROM StringTable")
		if err != nil {
			log.Fatalln(err)
		}

		log.Debugln("Demangle Names")
		for _, p := range stringTable {
			demangledValue, err := demangle.ToString(p.Value)
			if err == nil {
				DemangledNamesMap[p.ID] = demangledValue
				log.Debugf("Demangle '%s' to '%s'\n", p.Value, demangledValue)
			} else {
				DemangledNamesMap[p.ID] = p.Value
				log.Debugf("Cannot demangle '%s'\n", p.Value)
			}
		}

		info := NewInfomation()

		err = db.Select(&info.Meta.Devices, "SELECT * FROM CUPTI_ACTIVITY_KIND_DEVICE")
		if err != nil {
			log.Fatalln(err)
		}

		// runtime events
		info.Events = append(info.Events, GetRuntimeEvents(db)...)
		info.Events = append(info.Events, GetMemcpyEvents(db)...)
		info.Events = append(info.Events, GetMemcpy2Events(db)...)
		info.Events = append(info.Events, GetMemsetEvents(db)...)
		info.Events = append(info.Events, GetConcurrentKernelEvents(db)...)
		info.Events = append(info.Events, GetSynchronizationEvents(db)...)

		var file []byte
		// dump for Google Chrome
		if opts.Pretty {
			file, _ = json.MarshalIndent(info, "", " ")
		} else {
			file, _ = json.Marshal(info)
		}
		_ = ioutil.WriteFile(opts.OutputFile, file, 0644)
	}
}
