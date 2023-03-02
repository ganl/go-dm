/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_1169 struct {
	dm_build_1170 *list.List
	dm_build_1171 *dm_build_1223
	dm_build_1172 int
}

func Dm_build_1173() *Dm_build_1169 {
	return &Dm_build_1169{
		dm_build_1170: list.New(),
		dm_build_1172: 0,
	}
}

func (dm_build_1175 *Dm_build_1169) Dm_build_1174() int {
	return dm_build_1175.dm_build_1172
}

func (dm_build_1177 *Dm_build_1169) Dm_build_1176(dm_build_1178 *Dm_build_1247, dm_build_1179 int) int {
	var dm_build_1180 = 0
	var dm_build_1181 = 0
	for dm_build_1180 < dm_build_1179 && dm_build_1177.dm_build_1171 != nil {
		dm_build_1181 = dm_build_1177.dm_build_1171.dm_build_1231(dm_build_1178, dm_build_1179-dm_build_1180)
		if dm_build_1177.dm_build_1171.dm_build_1226 == 0 {
			dm_build_1177.dm_build_1213()
		}
		dm_build_1180 += dm_build_1181
		dm_build_1177.dm_build_1172 -= dm_build_1181
	}
	return dm_build_1180
}

func (dm_build_1183 *Dm_build_1169) Dm_build_1182(dm_build_1184 []byte, dm_build_1185 int, dm_build_1186 int) int {
	var dm_build_1187 = 0
	var dm_build_1188 = 0
	for dm_build_1187 < dm_build_1186 && dm_build_1183.dm_build_1171 != nil {
		dm_build_1188 = dm_build_1183.dm_build_1171.dm_build_1235(dm_build_1184, dm_build_1185, dm_build_1186-dm_build_1187)
		if dm_build_1183.dm_build_1171.dm_build_1226 == 0 {
			dm_build_1183.dm_build_1213()
		}
		dm_build_1187 += dm_build_1188
		dm_build_1183.dm_build_1172 -= dm_build_1188
		dm_build_1185 += dm_build_1188
	}
	return dm_build_1187
}

func (dm_build_1190 *Dm_build_1169) Dm_build_1189(dm_build_1191 io.Writer, dm_build_1192 int) int {
	var dm_build_1193 = 0
	var dm_build_1194 = 0
	for dm_build_1193 < dm_build_1192 && dm_build_1190.dm_build_1171 != nil {
		dm_build_1194 = dm_build_1190.dm_build_1171.dm_build_1240(dm_build_1191, dm_build_1192-dm_build_1193)
		if dm_build_1190.dm_build_1171.dm_build_1226 == 0 {
			dm_build_1190.dm_build_1213()
		}
		dm_build_1193 += dm_build_1194
		dm_build_1190.dm_build_1172 -= dm_build_1194
	}
	return dm_build_1193
}

func (dm_build_1196 *Dm_build_1169) Dm_build_1195(dm_build_1197 []byte, dm_build_1198 int, dm_build_1199 int) {
	if dm_build_1199 == 0 {
		return
	}
	var dm_build_1200 = dm_build_1227(dm_build_1197, dm_build_1198, dm_build_1199)
	if dm_build_1196.dm_build_1171 == nil {
		dm_build_1196.dm_build_1171 = dm_build_1200
	} else {
		dm_build_1196.dm_build_1170.PushBack(dm_build_1200)
	}
	dm_build_1196.dm_build_1172 += dm_build_1199
}

func (dm_build_1202 *Dm_build_1169) dm_build_1201(dm_build_1203 int) byte {
	var dm_build_1204 = dm_build_1203
	var dm_build_1205 = dm_build_1202.dm_build_1171
	for dm_build_1204 > 0 && dm_build_1205 != nil {
		if dm_build_1205.dm_build_1226 == 0 {
			continue
		}
		if dm_build_1204 > dm_build_1205.dm_build_1226-1 {
			dm_build_1204 -= dm_build_1205.dm_build_1226
			dm_build_1205 = dm_build_1202.dm_build_1170.Front().Value.(*dm_build_1223)
		} else {
			break
		}
	}
	return dm_build_1205.dm_build_1244(dm_build_1204)
}
func (dm_build_1207 *Dm_build_1169) Dm_build_1206(dm_build_1208 *Dm_build_1169) {
	if dm_build_1208.dm_build_1172 == 0 {
		return
	}
	var dm_build_1209 = dm_build_1208.dm_build_1171
	for dm_build_1209 != nil {
		dm_build_1207.dm_build_1210(dm_build_1209)
		dm_build_1208.dm_build_1213()
		dm_build_1209 = dm_build_1208.dm_build_1171
	}
	dm_build_1208.dm_build_1172 = 0
}
func (dm_build_1211 *Dm_build_1169) dm_build_1210(dm_build_1212 *dm_build_1223) {
	if dm_build_1212.dm_build_1226 == 0 {
		return
	}
	if dm_build_1211.dm_build_1171 == nil {
		dm_build_1211.dm_build_1171 = dm_build_1212
	} else {
		dm_build_1211.dm_build_1170.PushBack(dm_build_1212)
	}
	dm_build_1211.dm_build_1172 += dm_build_1212.dm_build_1226
}

func (dm_build_1214 *Dm_build_1169) dm_build_1213() {
	var dm_build_1215 = dm_build_1214.dm_build_1170.Front()
	if dm_build_1215 == nil {
		dm_build_1214.dm_build_1171 = nil
	} else {
		dm_build_1214.dm_build_1171 = dm_build_1215.Value.(*dm_build_1223)
		dm_build_1214.dm_build_1170.Remove(dm_build_1215)
	}
}

func (dm_build_1217 *Dm_build_1169) Dm_build_1216() []byte {
	var dm_build_1218 = make([]byte, dm_build_1217.dm_build_1172)
	var dm_build_1219 = dm_build_1217.dm_build_1171
	var dm_build_1220 = 0
	var dm_build_1221 = len(dm_build_1218)
	var dm_build_1222 = 0
	for dm_build_1219 != nil {
		if dm_build_1219.dm_build_1226 > 0 {
			if dm_build_1221 > dm_build_1219.dm_build_1226 {
				dm_build_1222 = dm_build_1219.dm_build_1226
			} else {
				dm_build_1222 = dm_build_1221
			}
			copy(dm_build_1218[dm_build_1220:dm_build_1220+dm_build_1222], dm_build_1219.dm_build_1224[dm_build_1219.dm_build_1225:dm_build_1219.dm_build_1225+dm_build_1222])
			dm_build_1220 += dm_build_1222
			dm_build_1221 -= dm_build_1222
		}
		if dm_build_1217.dm_build_1170.Front() == nil {
			dm_build_1219 = nil
		} else {
			dm_build_1219 = dm_build_1217.dm_build_1170.Front().Value.(*dm_build_1223)
		}
	}
	return dm_build_1218
}

type dm_build_1223 struct {
	dm_build_1224 []byte
	dm_build_1225 int
	dm_build_1226 int
}

func dm_build_1227(dm_build_1228 []byte, dm_build_1229 int, dm_build_1230 int) *dm_build_1223 {
	return &dm_build_1223{
		dm_build_1228,
		dm_build_1229,
		dm_build_1230,
	}
}

func (dm_build_1232 *dm_build_1223) dm_build_1231(dm_build_1233 *Dm_build_1247, dm_build_1234 int) int {
	if dm_build_1232.dm_build_1226 <= dm_build_1234 {
		dm_build_1234 = dm_build_1232.dm_build_1226
	}
	dm_build_1233.Dm_build_1326(dm_build_1232.dm_build_1224[dm_build_1232.dm_build_1225 : dm_build_1232.dm_build_1225+dm_build_1234])
	dm_build_1232.dm_build_1225 += dm_build_1234
	dm_build_1232.dm_build_1226 -= dm_build_1234
	return dm_build_1234
}

func (dm_build_1236 *dm_build_1223) dm_build_1235(dm_build_1237 []byte, dm_build_1238 int, dm_build_1239 int) int {
	if dm_build_1236.dm_build_1226 <= dm_build_1239 {
		dm_build_1239 = dm_build_1236.dm_build_1226
	}
	copy(dm_build_1237[dm_build_1238:dm_build_1238+dm_build_1239], dm_build_1236.dm_build_1224[dm_build_1236.dm_build_1225:dm_build_1236.dm_build_1225+dm_build_1239])
	dm_build_1236.dm_build_1225 += dm_build_1239
	dm_build_1236.dm_build_1226 -= dm_build_1239
	return dm_build_1239
}

func (dm_build_1241 *dm_build_1223) dm_build_1240(dm_build_1242 io.Writer, dm_build_1243 int) int {
	if dm_build_1241.dm_build_1226 <= dm_build_1243 {
		dm_build_1243 = dm_build_1241.dm_build_1226
	}
	dm_build_1242.Write(dm_build_1241.dm_build_1224[dm_build_1241.dm_build_1225 : dm_build_1241.dm_build_1225+dm_build_1243])
	dm_build_1241.dm_build_1225 += dm_build_1243
	dm_build_1241.dm_build_1226 -= dm_build_1243
	return dm_build_1243
}
func (dm_build_1245 *dm_build_1223) dm_build_1244(dm_build_1246 int) byte {
	return dm_build_1245.dm_build_1224[dm_build_1245.dm_build_1225+dm_build_1246]
}
