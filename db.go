package main

type StringTable struct {
	ID    int64  `db:"id"`
	Value string `db:"value"`
}

type CuptiActivityKindRuntime struct {
	ID            int64 `db:"_id_"`
	Cbid          int64 `db:"cbid"`
	Start         int64 `db:"start"`
	End           int64 `db:"end"`
	ProcessID     int64 `db:"processId"`
	ThreadID      int64 `db:"threadId"`
	CorrelationID int64 `db:"correlationId"`
	ReturnValue   int64 `db:"returnValue"`
}

// type CuptiActivityKindMarker struct {
//  Key        int64  `db:"_id_"`
//  Flags      int64  `db:"flags"`
//  Timestamp  int64  `db:"timestamp"`
//  ID         int64  `db:"id"`
//  ObjectKind int64  `db:"objectKind"`
//  ObjectID   []byte `db:"objectId"`
//  Name       int64  `db:"name"`
//  Domain     int64  `db:"domain"`
// }

type CuptiActivityKindMemcpy struct {
	Key                  int64 `db:"_id_"`
	CopyKind             int64 `db:"copyKind"`
	SrcKind              int64 `db:"srcKind"`
	DstKind              int64 `db:"dstKind"`
	Flags                int64 `db:"flags"`
	Bytes                int64 `db:"bytes"`
	Start                int64 `db:"start"`
	End                  int64 `db:"end"`
	DeviceID             int64 `db:"deviceId"`
	ContextID            int64 `db:"contextId"`
	StreamID             int64 `db:"streamId"`
	CorrelationID        int64 `db:"correlationId"`
	RuntimeCorrelationID int64 `db:"runtimeCorrelationId"`
}

// peer to peer
type CuptiActivityKindMemcpy2 struct {
	Key           int64 `db:"_id_"`
	CopyKind      int64 `db:"copyKind"`
	SrcKind       int64 `db:"srcKind"`
	DstKind       int64 `db:"dstKind"`
	Flags         int64 `db:"flags"`
	Bytes         int64 `db:"bytes"`
	Start         int64 `db:"start"`
	End           int64 `db:"end"`
	DeviceID      int64 `db:"deviceId"`
	ContextID     int64 `db:"contextId"`
	StreamID      int64 `db:"streamId"`
	SrcDeviceID   int64 `db:"srcDeviceId"`
	SrcContextID  int64 `db:"srcContextId"`
	DstDeviceID   int64 `db:"dstDeviceId"`
	DstContextID  int64 `db:"dstContextId"`
	CorrelationID int64 `db:"correlationId"`
}

type CuptiActivityKindMemset struct {
	Key           int64 `db:"_id_"`
	Value         int64 `db:"value"`
	Bytes         int64 `db:"bytes"`
	Start         int64 `db:"start"`
	End           int64 `db:"end"`
	DeviceID      int64 `db:"deviceId"`
	ContextID     int64 `db:"contextId"`
	StreamID      int64 `db:"streamId"`
	CorrelationID int64 `db:"correlationId"`
	Flags         int64 `db:"flags"`
	MemoryKind    int64 `db:"memoryKind"`
}

type CuptiActivityKindConcurrentKernel struct {
	Key                             int64  `db:"_id_"`
	CacheConfig                     []byte `db:"cacheConfig"`
	SharedMemoryConfig              int64  `db:"sharedMemoryConfig"`
	RegistersPerThread              int64  `db:"registersPerThread"`
	PartitionedGlobalCacheRequested int64  `db:"partitionedGlobalCacheRequested"`
	PartitionedGlobalCacheExecuted  int64  `db:"partitionedGlobalCacheExecuted"`
	Start                           int64  `db:"start"`
	End                             int64  `db:"end"`
	Completed                       int64  `db:"completed"`
	DeviceID                        int64  `db:"deviceId"`
	ContextID                       int64  `db:"contextId"`
	StreamID                        int64  `db:"streamId"`
	GridX                           int64  `db:"gridX"`
	GridY                           int64  `db:"gridY"`
	GridZ                           int64  `db:"gridZ"`
	BlockX                          int64  `db:"blockX"`
	BlockY                          int64  `db:"blockY"`
	BlockZ                          int64  `db:"blockZ"`
	StaticSharedMemory              int64  `db:"staticSharedMemory"`
	DynamicSharedMemory             int64  `db:"dynamicSharedMemory"`
	LocalMemoryPerThread            int64  `db:"localMemoryPerThread"`
	LocalMemoryTotal                int64  `db:"localMemoryTotal"`
	CorrelationID                   int64  `db:"correlationId"`
	GridID                          int64  `db:"gridId"`
	Name                            int64  `db:"name"`
	Queued                          int64  `db:"queued"`
	Submitted                       int64  `db:"submitted"`
	LaunchType                      int64  `db:"launchType"`
	IsSharedMemoryCarveoutRequested int64  `db:"isSharedMemoryCarveoutRequested"`
	SharedMemoryCarveoutRequested   int64  `db:"sharedMemoryCarveoutRequested"`
	SharedMemoryExecuted            int64  `db:"sharedMemoryExecuted"`
}

type CuptiActivityKindSynchronization struct {
	Key           int64 `db:"_id_"`
	Type          int64 `db:"type"`
	Start         int64 `db:"start"`
	End           int64 `db:"end"`
	CorrelationID int64 `db:"correlationId"`
	ContextID     int64 `db:"contextId"`
	StreamID      int64 `db:"streamId"`
	CudaEventID   int64 `db:"cudaEventId"`
}

type CuptiActivityKindDevice struct {
	Key                              int64  `db:"_id_"`
	Flags                            int64  `db:"flags" json:"flags"`
	GlobalMemoryBandwidth            int64  `db:"globalMemoryBandwidth" json:"globalMemoryBandwidth"`
	GlobalMemorySize                 int64  `db:"globalMemorySize" json:"globalMemorySize"`
	ConstantMemorySize               int64  `db:"constantMemorySize" json:"constantMemorySize"`
	L2CacheSize                      int64  `db:"l2CacheSize" json:"l2CacheSize"`
	NumThreadsPerWarp                int64  `db:"numThreadsPerWarp" json:"numThreadsPerWarp"`
	CoreClockRate                    int64  `db:"coreClockRate" json:"coreClockRate"`
	NumMemcpyEngines                 int64  `db:"numMemcpyEngines" json:"numMemcpyEngines"`
	NumMultiprocessors               int64  `db:"numMultiprocessors" json:"numMultiprocessors"`
	MaxIPC                           int64  `db:"maxIPC" json:"maxIPC"`
	MaxWarpsPerMultiprocessor        int64  `db:"maxWarpsPerMultiprocessor" json:"maxWarpsPerMultiprocessor"`
	MaxBlocksPerMultiprocessor       int64  `db:"maxBlocksPerMultiprocessor" json:"maxBlocksPerMultiprocessor"`
	MaxSharedMemoryPerMultiprocessor int64  `db:"maxSharedMemoryPerMultiprocessor" json:"maxSharedMemoryPerMultiprocessor"`
	MaxRegistersPerMultiprocessor    int64  `db:"maxRegistersPerMultiprocessor" json:"maxRegistersPerMultiprocessor"`
	MaxRegistersPerBlock             int64  `db:"maxRegistersPerBlock" json:"maxRegistersPerBlock"`
	MaxSharedMemoryPerBlock          int64  `db:"maxSharedMemoryPerBlock" json:"maxSharedMemoryPerBlock"`
	MaxThreadsPerBlock               int64  `db:"maxThreadsPerBlock" json:"maxThreadsPerBlock"`
	MaxBlockDimX                     int64  `db:"maxBlockDimX" json:"maxBlockDimX"`
	MaxBlockDimY                     int64  `db:"maxBlockDimY" json:"maxBlockDimY"`
	MaxBlockDimZ                     int64  `db:"maxBlockDimZ" json:"maxBlockDimZ"`
	MaxGridDimX                      int64  `db:"maxGridDimX" json:"maxGridDimX"`
	MaxGridDimY                      int64  `db:"maxGridDimY" json:"maxGridDimY"`
	MaxGridDimZ                      int64  `db:"maxGridDimZ" json:"maxGridDimZ"`
	ComputeCapabilityMajor           int64  `db:"computeCapabilityMajor" json:"computeCapabilityMajor"`
	ComputeCapabilityMinor           int64  `db:"computeCapabilityMinor" json:"computeCapabilityMinor"`
	ID                               int64  `db:"id" json:"id"`
	EccEnabled                       int64  `db:"eccEnabled" json:"eccEnabled"`
	Uuid                             []byte `db:"uuid" json:"uuid"`
	Name                             int64  `db:"name" json:"name"`
}
