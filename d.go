/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_1196 struct {
	dm_build_1197 *list.List
	dm_build_1198 *dm_build_1250
	dm_build_1199 int
}

func Dm_build_1200() *Dm_build_1196 {
	return &Dm_build_1196{
		dm_build_1197: list.New(),
		dm_build_1199: 0,
	}
}

func (dm_build_1202 *Dm_build_1196) Dm_build_1201() int {
	return dm_build_1202.dm_build_1199
}

func (dm_build_1204 *Dm_build_1196) Dm_build_1203(dm_build_1205 *Dm_build_1274, dm_build_1206 int) int {
	var dm_build_1207 = 0
	var dm_build_1208 = 0
	for dm_build_1207 < dm_build_1206 && dm_build_1204.dm_build_1198 != nil {
		dm_build_1208 = dm_build_1204.dm_build_1198.dm_build_1258(dm_build_1205, dm_build_1206-dm_build_1207)
		if dm_build_1204.dm_build_1198.dm_build_1253 == 0 {
			dm_build_1204.dm_build_1240()
		}
		dm_build_1207 += dm_build_1208
		dm_build_1204.dm_build_1199 -= dm_build_1208
	}
	return dm_build_1207
}

func (dm_build_1210 *Dm_build_1196) Dm_build_1209(dm_build_1211 []byte, dm_build_1212 int, dm_build_1213 int) int {
	var dm_build_1214 = 0
	var dm_build_1215 = 0
	for dm_build_1214 < dm_build_1213 && dm_build_1210.dm_build_1198 != nil {
		dm_build_1215 = dm_build_1210.dm_build_1198.dm_build_1262(dm_build_1211, dm_build_1212, dm_build_1213-dm_build_1214)
		if dm_build_1210.dm_build_1198.dm_build_1253 == 0 {
			dm_build_1210.dm_build_1240()
		}
		dm_build_1214 += dm_build_1215
		dm_build_1210.dm_build_1199 -= dm_build_1215
		dm_build_1212 += dm_build_1215
	}
	return dm_build_1214
}

func (dm_build_1217 *Dm_build_1196) Dm_build_1216(dm_build_1218 io.Writer, dm_build_1219 int) int {
	var dm_build_1220 = 0
	var dm_build_1221 = 0
	for dm_build_1220 < dm_build_1219 && dm_build_1217.dm_build_1198 != nil {
		dm_build_1221 = dm_build_1217.dm_build_1198.dm_build_1267(dm_build_1218, dm_build_1219-dm_build_1220)
		if dm_build_1217.dm_build_1198.dm_build_1253 == 0 {
			dm_build_1217.dm_build_1240()
		}
		dm_build_1220 += dm_build_1221
		dm_build_1217.dm_build_1199 -= dm_build_1221
	}
	return dm_build_1220
}

func (dm_build_1223 *Dm_build_1196) Dm_build_1222(dm_build_1224 []byte, dm_build_1225 int, dm_build_1226 int) {
	if dm_build_1226 == 0 {
		return
	}
	var dm_build_1227 = dm_build_1254(dm_build_1224, dm_build_1225, dm_build_1226)
	if dm_build_1223.dm_build_1198 == nil {
		dm_build_1223.dm_build_1198 = dm_build_1227
	} else {
		dm_build_1223.dm_build_1197.PushBack(dm_build_1227)
	}
	dm_build_1223.dm_build_1199 += dm_build_1226
}

func (dm_build_1229 *Dm_build_1196) dm_build_1228(dm_build_1230 int) byte {
	var dm_build_1231 = dm_build_1230
	var dm_build_1232 = dm_build_1229.dm_build_1198
	for dm_build_1231 > 0 && dm_build_1232 != nil {
		if dm_build_1232.dm_build_1253 == 0 {
			continue
		}
		if dm_build_1231 > dm_build_1232.dm_build_1253-1 {
			dm_build_1231 -= dm_build_1232.dm_build_1253
			dm_build_1232 = dm_build_1229.dm_build_1197.Front().Value.(*dm_build_1250)
		} else {
			break
		}
	}
	return dm_build_1232.dm_build_1271(dm_build_1231)
}
func (dm_build_1234 *Dm_build_1196) Dm_build_1233(dm_build_1235 *Dm_build_1196) {
	if dm_build_1235.dm_build_1199 == 0 {
		return
	}
	var dm_build_1236 = dm_build_1235.dm_build_1198
	for dm_build_1236 != nil {
		dm_build_1234.dm_build_1237(dm_build_1236)
		dm_build_1235.dm_build_1240()
		dm_build_1236 = dm_build_1235.dm_build_1198
	}
	dm_build_1235.dm_build_1199 = 0
}
func (dm_build_1238 *Dm_build_1196) dm_build_1237(dm_build_1239 *dm_build_1250) {
	if dm_build_1239.dm_build_1253 == 0 {
		return
	}
	if dm_build_1238.dm_build_1198 == nil {
		dm_build_1238.dm_build_1198 = dm_build_1239
	} else {
		dm_build_1238.dm_build_1197.PushBack(dm_build_1239)
	}
	dm_build_1238.dm_build_1199 += dm_build_1239.dm_build_1253
}

func (dm_build_1241 *Dm_build_1196) dm_build_1240() {
	var dm_build_1242 = dm_build_1241.dm_build_1197.Front()
	if dm_build_1242 == nil {
		dm_build_1241.dm_build_1198 = nil
	} else {
		dm_build_1241.dm_build_1198 = dm_build_1242.Value.(*dm_build_1250)
		dm_build_1241.dm_build_1197.Remove(dm_build_1242)
	}
}

func (dm_build_1244 *Dm_build_1196) Dm_build_1243() []byte {
	var dm_build_1245 = make([]byte, dm_build_1244.dm_build_1199)
	var dm_build_1246 = dm_build_1244.dm_build_1198
	var dm_build_1247 = 0
	var dm_build_1248 = len(dm_build_1245)
	var dm_build_1249 = 0
	for dm_build_1246 != nil {
		if dm_build_1246.dm_build_1253 > 0 {
			if dm_build_1248 > dm_build_1246.dm_build_1253 {
				dm_build_1249 = dm_build_1246.dm_build_1253
			} else {
				dm_build_1249 = dm_build_1248
			}
			copy(dm_build_1245[dm_build_1247:dm_build_1247+dm_build_1249], dm_build_1246.dm_build_1251[dm_build_1246.dm_build_1252:dm_build_1246.dm_build_1252+dm_build_1249])
			dm_build_1247 += dm_build_1249
			dm_build_1248 -= dm_build_1249
		}
		if dm_build_1244.dm_build_1197.Front() == nil {
			dm_build_1246 = nil
		} else {
			dm_build_1246 = dm_build_1244.dm_build_1197.Front().Value.(*dm_build_1250)
		}
	}
	return dm_build_1245
}

type dm_build_1250 struct {
	dm_build_1251 []byte
	dm_build_1252 int
	dm_build_1253 int
}

func dm_build_1254(dm_build_1255 []byte, dm_build_1256 int, dm_build_1257 int) *dm_build_1250 {
	return &dm_build_1250{
		dm_build_1255,
		dm_build_1256,
		dm_build_1257,
	}
}

func (dm_build_1259 *dm_build_1250) dm_build_1258(dm_build_1260 *Dm_build_1274, dm_build_1261 int) int {
	if dm_build_1259.dm_build_1253 <= dm_build_1261 {
		dm_build_1261 = dm_build_1259.dm_build_1253
	}
	dm_build_1260.Dm_build_1357(dm_build_1259.dm_build_1251[dm_build_1259.dm_build_1252 : dm_build_1259.dm_build_1252+dm_build_1261])
	dm_build_1259.dm_build_1252 += dm_build_1261
	dm_build_1259.dm_build_1253 -= dm_build_1261
	return dm_build_1261
}

func (dm_build_1263 *dm_build_1250) dm_build_1262(dm_build_1264 []byte, dm_build_1265 int, dm_build_1266 int) int {
	if dm_build_1263.dm_build_1253 <= dm_build_1266 {
		dm_build_1266 = dm_build_1263.dm_build_1253
	}
	copy(dm_build_1264[dm_build_1265:dm_build_1265+dm_build_1266], dm_build_1263.dm_build_1251[dm_build_1263.dm_build_1252:dm_build_1263.dm_build_1252+dm_build_1266])
	dm_build_1263.dm_build_1252 += dm_build_1266
	dm_build_1263.dm_build_1253 -= dm_build_1266
	return dm_build_1266
}

func (dm_build_1268 *dm_build_1250) dm_build_1267(dm_build_1269 io.Writer, dm_build_1270 int) int {
	if dm_build_1268.dm_build_1253 <= dm_build_1270 {
		dm_build_1270 = dm_build_1268.dm_build_1253
	}
	dm_build_1269.Write(dm_build_1268.dm_build_1251[dm_build_1268.dm_build_1252 : dm_build_1268.dm_build_1252+dm_build_1270])
	dm_build_1268.dm_build_1252 += dm_build_1270
	dm_build_1268.dm_build_1253 -= dm_build_1270
	return dm_build_1270
}
func (dm_build_1272 *dm_build_1250) dm_build_1271(dm_build_1273 int) byte {
	return dm_build_1272.dm_build_1251[dm_build_1272.dm_build_1252+dm_build_1273]
}
