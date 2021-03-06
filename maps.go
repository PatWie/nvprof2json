package main

import "fmt"

// taken from cupti_driver_cbid.h and remove "CUPTI_DRIVER_TRACE_CBID_"
var CbidTable = map[int64]string{
	0:         "INVALID",
	1:         "cuInit",
	2:         "cuDriverGetVersion",
	3:         "cuDeviceGet",
	4:         "cuDeviceGetCount",
	5:         "cuDeviceGetName",
	6:         "cuDeviceComputeCapability",
	7:         "cuDeviceTotalMem",
	8:         "cuDeviceGetProperties",
	9:         "cuDeviceGetAttribute",
	10:        "cuCtxCreate",
	11:        "cuCtxDestroy",
	12:        "cuCtxAttach",
	13:        "cuCtxDetach",
	14:        "cuCtxPushCurrent",
	15:        "cuCtxPopCurrent",
	16:        "cuCtxGetDevice",
	17:        "cuCtxSynchronize",
	18:        "cuModuleLoad",
	19:        "cuModuleLoadData",
	20:        "cuModuleLoadDataEx",
	21:        "cuModuleLoadFatBinary",
	22:        "cuModuleUnload",
	23:        "cuModuleGetFunction",
	24:        "cuModuleGetGlobal",
	25:        "cu64ModuleGetGlobal",
	26:        "cuModuleGetTexRef",
	27:        "cuMemGetInfo",
	28:        "cu64MemGetInfo",
	29:        "cuMemAlloc",
	30:        "cu64MemAlloc",
	31:        "cuMemAllocPitch",
	32:        "cu64MemAllocPitch",
	33:        "cuMemFree",
	34:        "cu64MemFree",
	35:        "cuMemGetAddressRange",
	36:        "cu64MemGetAddressRange",
	37:        "cuMemAllocHost",
	38:        "cuMemFreeHost",
	39:        "cuMemHostAlloc",
	40:        "cuMemHostGetDevicePointer",
	41:        "cu64MemHostGetDevicePointer",
	42:        "cuMemHostGetFlags",
	43:        "cuMemcpyHtoD",
	44:        "cu64MemcpyHtoD",
	45:        "cuMemcpyDtoH",
	46:        "cu64MemcpyDtoH",
	47:        "cuMemcpyDtoD",
	48:        "cu64MemcpyDtoD",
	49:        "cuMemcpyDtoA",
	50:        "cu64MemcpyDtoA",
	51:        "cuMemcpyAtoD",
	52:        "cu64MemcpyAtoD",
	53:        "cuMemcpyHtoA",
	54:        "cuMemcpyAtoH",
	55:        "cuMemcpyAtoA",
	56:        "cuMemcpy2D",
	57:        "cuMemcpy2DUnaligned",
	58:        "cuMemcpy3D",
	59:        "cu64Memcpy3D",
	60:        "cuMemcpyHtoDAsync",
	61:        "cu64MemcpyHtoDAsync",
	62:        "cuMemcpyDtoHAsync",
	63:        "cu64MemcpyDtoHAsync",
	64:        "cuMemcpyDtoDAsync",
	65:        "cu64MemcpyDtoDAsync",
	66:        "cuMemcpyHtoAAsync",
	67:        "cuMemcpyAtoHAsync",
	68:        "cuMemcpy2DAsync",
	69:        "cuMemcpy3DAsync",
	70:        "cu64Memcpy3DAsync",
	71:        "cuMemsetD8",
	72:        "cu64MemsetD8",
	73:        "cuMemsetD16",
	74:        "cu64MemsetD16",
	75:        "cuMemsetD32",
	76:        "cu64MemsetD32",
	77:        "cuMemsetD2D8",
	78:        "cu64MemsetD2D8",
	79:        "cuMemsetD2D16",
	80:        "cu64MemsetD2D16",
	81:        "cuMemsetD2D32",
	82:        "cu64MemsetD2D32",
	83:        "cuFuncSetBlockShape",
	84:        "cuFuncSetSharedSize",
	85:        "cuFuncGetAttribute",
	86:        "cuFuncSetCacheConfig",
	87:        "cuArrayCreate",
	88:        "cuArrayGetDescriptor",
	89:        "cuArrayDestroy",
	90:        "cuArray3DCreate",
	91:        "cuArray3DGetDescriptor",
	92:        "cuTexRefCreate",
	93:        "cuTexRefDestroy",
	94:        "cuTexRefSetArray",
	95:        "cuTexRefSetAddress",
	96:        "cu64TexRefSetAddress",
	97:        "cuTexRefSetAddress2D",
	98:        "cu64TexRefSetAddress2D",
	99:        "cuTexRefSetFormat",
	100:       "cuTexRefSetAddressMode",
	101:       "cuTexRefSetFilterMode",
	102:       "cuTexRefSetFlags",
	103:       "cuTexRefGetAddress",
	104:       "cu64TexRefGetAddress",
	105:       "cuTexRefGetArray",
	106:       "cuTexRefGetAddressMode",
	107:       "cuTexRefGetFilterMode",
	108:       "cuTexRefGetFormat",
	109:       "cuTexRefGetFlags",
	110:       "cuParamSetSize",
	111:       "cuParamSeti",
	112:       "cuParamSetf",
	113:       "cuParamSetv",
	114:       "cuParamSetTexRef",
	115:       "cuLaunch",
	116:       "cuLaunchGrid",
	117:       "cuLaunchGridAsync",
	118:       "cuEventCreate",
	119:       "cuEventRecord",
	120:       "cuEventQuery",
	121:       "cuEventSynchronize",
	122:       "cuEventDestroy",
	123:       "cuEventElapsedTime",
	124:       "cuStreamCreate",
	125:       "cuStreamQuery",
	126:       "cuStreamSynchronize",
	127:       "cuStreamDestroy",
	128:       "cuGraphicsUnregisterResource",
	129:       "cuGraphicsSubResourceGetMappedArray",
	130:       "cuGraphicsResourceGetMappedPointer",
	131:       "cu64GraphicsResourceGetMappedPointer",
	132:       "cuGraphicsResourceSetMapFlags",
	133:       "cuGraphicsMapResources",
	134:       "cuGraphicsUnmapResources",
	135:       "cuGetExportTable",
	136:       "cuCtxSetLimit",
	137:       "cuCtxGetLimit",
	138:       "cuD3D10GetDevice",
	139:       "cuD3D10CtxCreate",
	140:       "cuGraphicsD3D10RegisterResource",
	141:       "cuD3D10RegisterResource",
	142:       "cuD3D10UnregisterResource",
	143:       "cuD3D10MapResources",
	144:       "cuD3D10UnmapResources",
	145:       "cuD3D10ResourceSetMapFlags",
	146:       "cuD3D10ResourceGetMappedArray",
	147:       "cuD3D10ResourceGetMappedPointer",
	148:       "cuD3D10ResourceGetMappedSize",
	149:       "cuD3D10ResourceGetMappedPitch",
	150:       "cuD3D10ResourceGetSurfaceDimensions",
	151:       "cuD3D11GetDevice",
	152:       "cuD3D11CtxCreate",
	153:       "cuGraphicsD3D11RegisterResource",
	154:       "cuD3D9GetDevice",
	155:       "cuD3D9CtxCreate",
	156:       "cuGraphicsD3D9RegisterResource",
	157:       "cuD3D9GetDirect3DDevice",
	158:       "cuD3D9RegisterResource",
	159:       "cuD3D9UnregisterResource",
	160:       "cuD3D9MapResources",
	161:       "cuD3D9UnmapResources",
	162:       "cuD3D9ResourceSetMapFlags",
	163:       "cuD3D9ResourceGetSurfaceDimensions",
	164:       "cuD3D9ResourceGetMappedArray",
	165:       "cuD3D9ResourceGetMappedPointer",
	166:       "cuD3D9ResourceGetMappedSize",
	167:       "cuD3D9ResourceGetMappedPitch",
	168:       "cuD3D9Begin",
	169:       "cuD3D9End",
	170:       "cuD3D9RegisterVertexBuffer",
	171:       "cuD3D9MapVertexBuffer",
	172:       "cuD3D9UnmapVertexBuffer",
	173:       "cuD3D9UnregisterVertexBuffer",
	174:       "cuGLCtxCreate",
	175:       "cuGraphicsGLRegisterBuffer",
	176:       "cuGraphicsGLRegisterImage",
	177:       "cuWGLGetDevice",
	178:       "cuGLInit",
	179:       "cuGLRegisterBufferObject",
	180:       "cuGLMapBufferObject",
	181:       "cuGLUnmapBufferObject",
	182:       "cuGLUnregisterBufferObject",
	183:       "cuGLSetBufferObjectMapFlags",
	184:       "cuGLMapBufferObjectAsync",
	185:       "cuGLUnmapBufferObjectAsync",
	186:       "cuVDPAUGetDevice",
	187:       "cuVDPAUCtxCreate",
	188:       "cuGraphicsVDPAURegisterVideoSurface",
	189:       "cuGraphicsVDPAURegisterOutputSurface",
	190:       "cuModuleGetSurfRef",
	191:       "cuSurfRefCreate",
	192:       "cuSurfRefDestroy",
	193:       "cuSurfRefSetFormat",
	194:       "cuSurfRefSetArray",
	195:       "cuSurfRefGetFormat",
	196:       "cuSurfRefGetArray",
	197:       "cu64DeviceTotalMem",
	198:       "cu64D3D10ResourceGetMappedPointer",
	199:       "cu64D3D10ResourceGetMappedSize",
	200:       "cu64D3D10ResourceGetMappedPitch",
	201:       "cu64D3D10ResourceGetSurfaceDimensions",
	202:       "cu64D3D9ResourceGetSurfaceDimensions",
	203:       "cu64D3D9ResourceGetMappedPointer",
	204:       "cu64D3D9ResourceGetMappedSize",
	205:       "cu64D3D9ResourceGetMappedPitch",
	206:       "cu64D3D9MapVertexBuffer",
	207:       "cu64GLMapBufferObject",
	208:       "cu64GLMapBufferObjectAsync",
	209:       "cuD3D11GetDevices",
	210:       "cuD3D11CtxCreateOnDevice",
	211:       "cuD3D10GetDevices",
	212:       "cuD3D10CtxCreateOnDevice",
	213:       "cuD3D9GetDevices",
	214:       "cuD3D9CtxCreateOnDevice",
	215:       "cu64MemHostAlloc",
	216:       "cuMemsetD8Async",
	217:       "cu64MemsetD8Async",
	218:       "cuMemsetD16Async",
	219:       "cu64MemsetD16Async",
	220:       "cuMemsetD32Async",
	221:       "cu64MemsetD32Async",
	222:       "cuMemsetD2D8Async",
	223:       "cu64MemsetD2D8Async",
	224:       "cuMemsetD2D16Async",
	225:       "cu64MemsetD2D16Async",
	226:       "cuMemsetD2D32Async",
	227:       "cu64MemsetD2D32Async",
	228:       "cu64ArrayCreate",
	229:       "cu64ArrayGetDescriptor",
	230:       "cu64Array3DCreate",
	231:       "cu64Array3DGetDescriptor",
	232:       "cu64Memcpy2D",
	233:       "cu64Memcpy2DUnaligned",
	234:       "cu64Memcpy2DAsync",
	235:       "cuCtxCreate_v2",
	236:       "cuD3D10CtxCreate_v2",
	237:       "cuD3D11CtxCreate_v2",
	238:       "cuD3D9CtxCreate_v2",
	239:       "cuGLCtxCreate_v2",
	240:       "cuVDPAUCtxCreate_v2",
	241:       "cuModuleGetGlobal_v2",
	242:       "cuMemGetInfo_v2",
	243:       "cuMemAlloc_v2",
	244:       "cuMemAllocPitch_v2",
	245:       "cuMemFree_v2",
	246:       "cuMemGetAddressRange_v2",
	247:       "cuMemHostGetDevicePointer_v2",
	248:       "cuMemcpy_v2",
	249:       "cuMemsetD8_v2",
	250:       "cuMemsetD16_v2",
	251:       "cuMemsetD32_v2",
	252:       "cuMemsetD2D8_v2",
	253:       "cuMemsetD2D16_v2",
	254:       "cuMemsetD2D32_v2",
	255:       "cuTexRefSetAddress_v2",
	256:       "cuTexRefSetAddress2D_v2",
	257:       "cuTexRefGetAddress_v2",
	258:       "cuGraphicsResourceGetMappedPointer_v2",
	259:       "cuDeviceTotalMem_v2",
	260:       "cuD3D10ResourceGetMappedPointer_v2",
	261:       "cuD3D10ResourceGetMappedSize_v2",
	262:       "cuD3D10ResourceGetMappedPitch_v2",
	263:       "cuD3D10ResourceGetSurfaceDimensions_v2",
	264:       "cuD3D9ResourceGetSurfaceDimensions_v2",
	265:       "cuD3D9ResourceGetMappedPointer_v2",
	266:       "cuD3D9ResourceGetMappedSize_v2",
	267:       "cuD3D9ResourceGetMappedPitch_v2",
	268:       "cuD3D9MapVertexBuffer_v2",
	269:       "cuGLMapBufferObject_v2",
	270:       "cuGLMapBufferObjectAsync_v2",
	271:       "cuMemHostAlloc_v2",
	272:       "cuArrayCreate_v2",
	273:       "cuArrayGetDescriptor_v2",
	274:       "cuArray3DCreate_v2",
	275:       "cuArray3DGetDescriptor_v2",
	276:       "cuMemcpyHtoD_v2",
	277:       "cuMemcpyHtoDAsync_v2",
	278:       "cuMemcpyDtoH_v2",
	279:       "cuMemcpyDtoHAsync_v2",
	280:       "cuMemcpyDtoD_v2",
	281:       "cuMemcpyDtoDAsync_v2",
	282:       "cuMemcpyAtoH_v2",
	283:       "cuMemcpyAtoHAsync_v2",
	284:       "cuMemcpyAtoD_v2",
	285:       "cuMemcpyDtoA_v2",
	286:       "cuMemcpyAtoA_v2",
	287:       "cuMemcpy2D_v2",
	288:       "cuMemcpy2DUnaligned_v2",
	289:       "cuMemcpy2DAsync_v2",
	290:       "cuMemcpy3D_v2",
	291:       "cuMemcpy3DAsync_v2",
	292:       "cuMemcpyHtoA_v2",
	293:       "cuMemcpyHtoAAsync_v2",
	294:       "cuMemAllocHost_v2",
	295:       "cuStreamWaitEvent",
	296:       "cuCtxGetApiVersion",
	297:       "cuD3D10GetDirect3DDevice",
	298:       "cuD3D11GetDirect3DDevice",
	299:       "cuCtxGetCacheConfig",
	300:       "cuCtxSetCacheConfig",
	301:       "cuMemHostRegister",
	302:       "cuMemHostUnregister",
	303:       "cuCtxSetCurrent",
	304:       "cuCtxGetCurrent",
	305:       "cuMemcpy",
	306:       "cuMemcpyAsync",
	307:       "cuLaunchKernel",
	308:       "cuProfilerStart",
	309:       "cuProfilerStop",
	310:       "cuPointerGetAttribute",
	311:       "cuProfilerInitialize",
	312:       "cuDeviceCanAccessPeer",
	313:       "cuCtxEnablePeerAccess",
	314:       "cuCtxDisablePeerAccess",
	315:       "cuMemPeerRegister",
	316:       "cuMemPeerUnregister",
	317:       "cuMemPeerGetDevicePointer",
	318:       "cuMemcpyPeer",
	319:       "cuMemcpyPeerAsync",
	320:       "cuMemcpy3DPeer",
	321:       "cuMemcpy3DPeerAsync",
	322:       "cuCtxDestroy_v2",
	323:       "cuCtxPushCurrent_v2",
	324:       "cuCtxPopCurrent_v2",
	325:       "cuEventDestroy_v2",
	326:       "cuStreamDestroy_v2",
	327:       "cuTexRefSetAddress2D_v3",
	328:       "cuIpcGetMemHandle",
	329:       "cuIpcOpenMemHandle",
	330:       "cuIpcCloseMemHandle",
	331:       "cuDeviceGetByPCIBusId",
	332:       "cuDeviceGetPCIBusId",
	333:       "cuGLGetDevices",
	334:       "cuIpcGetEventHandle",
	335:       "cuIpcOpenEventHandle",
	336:       "cuCtxSetSharedMemConfig",
	337:       "cuCtxGetSharedMemConfig",
	338:       "cuFuncSetSharedMemConfig",
	339:       "cuTexObjectCreate",
	340:       "cuTexObjectDestroy",
	341:       "cuTexObjectGetResourceDesc",
	342:       "cuTexObjectGetTextureDesc",
	343:       "cuSurfObjectCreate",
	344:       "cuSurfObjectDestroy",
	345:       "cuSurfObjectGetResourceDesc",
	346:       "cuStreamAddCallback",
	347:       "cuMipmappedArrayCreate",
	348:       "cuMipmappedArrayGetLevel",
	349:       "cuMipmappedArrayDestroy",
	350:       "cuTexRefSetMipmappedArray",
	351:       "cuTexRefSetMipmapFilterMode",
	352:       "cuTexRefSetMipmapLevelBias",
	353:       "cuTexRefSetMipmapLevelClamp",
	354:       "cuTexRefSetMaxAnisotropy",
	355:       "cuTexRefGetMipmappedArray",
	356:       "cuTexRefGetMipmapFilterMode",
	357:       "cuTexRefGetMipmapLevelBias",
	358:       "cuTexRefGetMipmapLevelClamp",
	359:       "cuTexRefGetMaxAnisotropy",
	360:       "cuGraphicsResourceGetMappedMipmappedArray",
	361:       "cuTexObjectGetResourceViewDesc",
	362:       "cuLinkCreate",
	363:       "cuLinkAddData",
	364:       "cuLinkAddFile",
	365:       "cuLinkComplete",
	366:       "cuLinkDestroy",
	367:       "cuStreamCreateWithPriority",
	368:       "cuStreamGetPriority",
	369:       "cuStreamGetFlags",
	370:       "cuCtxGetStreamPriorityRange",
	371:       "cuMemAllocManaged",
	372:       "cuGetErrorString",
	373:       "cuGetErrorName",
	374:       "cuOccupancyMaxActiveBlocksPerMultiprocessor",
	375:       "cuCompilePtx",
	376:       "cuBinaryFree",
	377:       "cuStreamAttachMemAsync",
	378:       "cuPointerSetAttribute",
	379:       "cuMemHostRegister_v2",
	380:       "cuGraphicsResourceSetMapFlags_v2",
	381:       "cuLinkCreate_v2",
	382:       "cuLinkAddData_v2",
	383:       "cuLinkAddFile_v2",
	384:       "cuOccupancyMaxPotentialBlockSize",
	385:       "cuGLGetDevices_v2",
	386:       "cuDevicePrimaryCtxRetain",
	387:       "cuDevicePrimaryCtxRelease",
	388:       "cuDevicePrimaryCtxSetFlags",
	389:       "cuDevicePrimaryCtxReset",
	390:       "cuGraphicsEGLRegisterImage",
	391:       "cuCtxGetFlags",
	392:       "cuDevicePrimaryCtxGetState",
	393:       "cuEGLStreamConsumerConnect",
	394:       "cuEGLStreamConsumerDisconnect",
	395:       "cuEGLStreamConsumerAcquireFrame",
	396:       "cuEGLStreamConsumerReleaseFrame",
	397:       "cuMemcpyHtoD_v2_ptds",
	398:       "cuMemcpyDtoH_v2_ptds",
	399:       "cuMemcpyDtoD_v2_ptds",
	400:       "cuMemcpyDtoA_v2_ptds",
	401:       "cuMemcpyAtoD_v2_ptds",
	402:       "cuMemcpyHtoA_v2_ptds",
	403:       "cuMemcpyAtoH_v2_ptds",
	404:       "cuMemcpyAtoA_v2_ptds",
	405:       "cuMemcpy2D_v2_ptds",
	406:       "cuMemcpy2DUnaligned_v2_ptds",
	407:       "cuMemcpy3D_v2_ptds",
	408:       "cuMemcpy_ptds",
	409:       "cuMemcpyPeer_ptds",
	410:       "cuMemcpy3DPeer_ptds",
	411:       "cuMemsetD8_v2_ptds",
	412:       "cuMemsetD16_v2_ptds",
	413:       "cuMemsetD32_v2_ptds",
	414:       "cuMemsetD2D8_v2_ptds",
	415:       "cuMemsetD2D16_v2_ptds",
	416:       "cuMemsetD2D32_v2_ptds",
	417:       "cuGLMapBufferObject_v2_ptds",
	418:       "cuMemcpyAsync_ptsz",
	419:       "cuMemcpyHtoAAsync_v2_ptsz",
	420:       "cuMemcpyAtoHAsync_v2_ptsz",
	421:       "cuMemcpyHtoDAsync_v2_ptsz",
	422:       "cuMemcpyDtoHAsync_v2_ptsz",
	423:       "cuMemcpyDtoDAsync_v2_ptsz",
	424:       "cuMemcpy2DAsync_v2_ptsz",
	425:       "cuMemcpy3DAsync_v2_ptsz",
	426:       "cuMemcpyPeerAsync_ptsz",
	427:       "cuMemcpy3DPeerAsync_ptsz",
	428:       "cuMemsetD8Async_ptsz",
	429:       "cuMemsetD16Async_ptsz",
	430:       "cuMemsetD32Async_ptsz",
	431:       "cuMemsetD2D8Async_ptsz",
	432:       "cuMemsetD2D16Async_ptsz",
	433:       "cuMemsetD2D32Async_ptsz",
	434:       "cuStreamGetPriority_ptsz",
	435:       "cuStreamGetFlags_ptsz",
	436:       "cuStreamWaitEvent_ptsz",
	437:       "cuStreamAddCallback_ptsz",
	438:       "cuStreamAttachMemAsync_ptsz",
	439:       "cuStreamQuery_ptsz",
	440:       "cuStreamSynchronize_ptsz",
	441:       "cuEventRecord_ptsz",
	442:       "cuLaunchKernel_ptsz",
	443:       "cuGraphicsMapResources_ptsz",
	444:       "cuGraphicsUnmapResources_ptsz",
	445:       "cuGLMapBufferObjectAsync_v2_ptsz",
	446:       "cuEGLStreamProducerConnect",
	447:       "cuEGLStreamProducerDisconnect",
	448:       "cuEGLStreamProducerPresentFrame",
	449:       "cuGraphicsResourceGetMappedEglFrame",
	450:       "cuPointerGetAttributes",
	451:       "cuOccupancyMaxActiveBlocksPerMultiprocessorWithFlags",
	452:       "cuOccupancyMaxPotentialBlockSizeWithFlags",
	453:       "cuEGLStreamProducerReturnFrame",
	454:       "cuDeviceGetP2PAttribute",
	455:       "cuTexRefSetBorderColor",
	456:       "cuTexRefGetBorderColor",
	457:       "cuMemAdvise",
	458:       "cuStreamWaitValue32",
	459:       "cuStreamWaitValue32_ptsz",
	460:       "cuStreamWriteValue32",
	461:       "cuStreamWriteValue32_ptsz",
	462:       "cuStreamBatchMemOp",
	463:       "cuStreamBatchMemOp_ptsz",
	464:       "cuNVNbufferGetPointer",
	465:       "cuNVNtextureGetArray",
	466:       "cuNNSetAllocator",
	467:       "cuMemPrefetchAsync",
	468:       "cuMemPrefetchAsync_ptsz",
	469:       "cuEventCreateFromNVNSync",
	470:       "cuEGLStreamConsumerConnectWithFlags",
	471:       "cuMemRangeGetAttribute",
	472:       "cuMemRangeGetAttributes",
	473:       "cuStreamWaitValue64",
	474:       "cuStreamWaitValue64_ptsz",
	475:       "cuStreamWriteValue64",
	476:       "cuStreamWriteValue64_ptsz",
	477:       "cuLaunchCooperativeKernel",
	478:       "cuLaunchCooperativeKernel_ptsz",
	479:       "cuEventCreateFromEGLSync",
	480:       "cuLaunchCooperativeKernelMultiDevice",
	481:       "cuFuncSetAttribute",
	482:       "cuDeviceGetUuid",
	483:       "cuStreamGetCtx",
	484:       "cuStreamGetCtx_ptsz",
	485:       "cuImportExternalMemory",
	486:       "cuExternalMemoryGetMappedBuffer",
	487:       "cuExternalMemoryGetMappedMipmappedArray",
	488:       "cuDestroyExternalMemory",
	489:       "cuImportExternalSemaphore",
	490:       "cuSignalExternalSemaphoresAsync",
	491:       "cuSignalExternalSemaphoresAsync_ptsz",
	492:       "cuWaitExternalSemaphoresAsync",
	493:       "cuWaitExternalSemaphoresAsync_ptsz",
	494:       "cuDestroyExternalSemaphore",
	495:       "cuStreamBeginCapture",
	496:       "cuStreamBeginCapture_ptsz",
	497:       "cuStreamEndCapture",
	498:       "cuStreamEndCapture_ptsz",
	499:       "cuStreamIsCapturing",
	500:       "cuStreamIsCapturing_ptsz",
	501:       "cuGraphCreate",
	502:       "cuGraphAddKernelNode",
	503:       "cuGraphKernelNodeGetParams",
	504:       "cuGraphAddMemcpyNode",
	505:       "cuGraphMemcpyNodeGetParams",
	506:       "cuGraphAddMemsetNode",
	507:       "cuGraphMemsetNodeGetParams",
	508:       "cuGraphMemsetNodeSetParams",
	509:       "cuGraphNodeGetType",
	510:       "cuGraphGetRootNodes",
	511:       "cuGraphNodeGetDependencies",
	512:       "cuGraphNodeGetDependentNodes",
	513:       "cuGraphInstantiate",
	514:       "cuGraphLaunch",
	515:       "cuGraphLaunch_ptsz",
	516:       "cuGraphExecDestroy",
	517:       "cuGraphDestroy",
	518:       "cuGraphAddDependencies",
	519:       "cuGraphRemoveDependencies",
	520:       "cuGraphMemcpyNodeSetParams",
	521:       "cuGraphKernelNodeSetParams",
	522:       "cuGraphDestroyNode",
	523:       "cuGraphClone",
	524:       "cuGraphNodeFindInClone",
	525:       "cuGraphAddChildGraphNode",
	526:       "cuGraphAddEmptyNode",
	527:       "cuLaunchHostFunc",
	528:       "cuLaunchHostFunc_ptsz",
	529:       "cuGraphChildGraphNodeGetGraph",
	530:       "cuGraphAddHostNode",
	531:       "cuGraphHostNodeGetParams",
	532:       "cuDeviceGetLuid",
	533:       "cuGraphHostNodeSetParams",
	534:       "cuGraphGetNodes",
	535:       "cuGraphGetEdges",
	536:       "SIZE",
	0x7ffffff: "FORCE_INT",
}

// CUpti_ActivitySynchronizationType
var ActivitySynchronizationType = map[int64]string{
	0:         "Unknown data",
	1:         "Event synchronize API",
	2:         "Stream wait event API",
	3:         "Stream synchronize API",
	4:         "Context synchronize API",
	0x7ffffff: "FORCE_INT",
}

// CUpti_ActivityMemcpyKind
var ActivityMemcpyKind = map[int64]string{
	0:  "unknown",
	1:  "host to device",
	2:  "device to host",
	3:  "host to device array",
	4:  "device array to host",
	5:  "device array to device array",
	6:  "device array to device",
	7:  "device to device array",
	8:  "device to device (same device)",
	9:  "host to host",
	10: "peer to peer (different devices)",
}

// CUpti_ActivityMemoryKind
var ActivityMemoryKind = map[int64]string{
	0: "unknown",
	1: "pageable",
	2: "pinned",
	3: "on the device",
	4: "an array",
	5: "managed",
	6: "device static",
	7: "managed static",
}

func FindInMap(needle int64, haystack map[int64]string) string {
	result := fmt.Sprintf("<not found %v>", needle)
	if val, ok := haystack[needle]; ok {
		result = val
	}
	return result
}
