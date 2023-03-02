/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"io"
	"math"
)

type Dm_build_1247 struct {
	dm_build_1248 []byte
	dm_build_1249 int
}

func Dm_build_1250(dm_build_1251 int) *Dm_build_1247 {
	return &Dm_build_1247{make([]byte, 0, dm_build_1251), 0}
}

func Dm_build_1252(dm_build_1253 []byte) *Dm_build_1247 {
	return &Dm_build_1247{dm_build_1253, 0}
}

func (dm_build_1255 *Dm_build_1247) dm_build_1254(dm_build_1256 int) *Dm_build_1247 {

	dm_build_1257 := len(dm_build_1255.dm_build_1248)
	dm_build_1258 := cap(dm_build_1255.dm_build_1248)

	if dm_build_1257+dm_build_1256 <= dm_build_1258 {
		dm_build_1255.dm_build_1248 = dm_build_1255.dm_build_1248[:dm_build_1257+dm_build_1256]
	} else {

		var calCap = int64(math.Max(float64(2*dm_build_1258), float64(dm_build_1256+dm_build_1257)))

		nbuf := make([]byte, dm_build_1256+dm_build_1257, calCap)
		copy(nbuf, dm_build_1255.dm_build_1248)
		dm_build_1255.dm_build_1248 = nbuf
	}

	return dm_build_1255
}

func (dm_build_1260 *Dm_build_1247) Dm_build_1259() int {
	return len(dm_build_1260.dm_build_1248)
}

func (dm_build_1262 *Dm_build_1247) Dm_build_1261(dm_build_1263 int) *Dm_build_1247 {
	for i := dm_build_1263; i < len(dm_build_1262.dm_build_1248); i++ {
		dm_build_1262.dm_build_1248[i] = 0
	}
	dm_build_1262.dm_build_1248 = dm_build_1262.dm_build_1248[:dm_build_1263]
	return dm_build_1262
}

func (dm_build_1265 *Dm_build_1247) Dm_build_1264(dm_build_1266 int) *Dm_build_1247 {
	dm_build_1265.dm_build_1249 = dm_build_1266
	return dm_build_1265
}

func (dm_build_1268 *Dm_build_1247) Dm_build_1267() int {
	return dm_build_1268.dm_build_1249
}

func (dm_build_1270 *Dm_build_1247) Dm_build_1269(dm_build_1271 bool) int {
	return len(dm_build_1270.dm_build_1248) - dm_build_1270.dm_build_1249
}

func (dm_build_1273 *Dm_build_1247) Dm_build_1272(dm_build_1274 int, dm_build_1275 bool, dm_build_1276 bool) *Dm_build_1247 {

	if dm_build_1275 {
		if dm_build_1276 {
			dm_build_1273.dm_build_1254(dm_build_1274)
		} else {
			dm_build_1273.dm_build_1248 = dm_build_1273.dm_build_1248[:len(dm_build_1273.dm_build_1248)-dm_build_1274]
		}
	} else {
		if dm_build_1276 {
			dm_build_1273.dm_build_1249 += dm_build_1274
		} else {
			dm_build_1273.dm_build_1249 -= dm_build_1274
		}
	}

	return dm_build_1273
}

func (dm_build_1278 *Dm_build_1247) Dm_build_1277(dm_build_1279 io.Reader, dm_build_1280 int) (int, error) {
	dm_build_1281 := len(dm_build_1278.dm_build_1248)
	dm_build_1278.dm_build_1254(dm_build_1280)
	dm_build_1282 := 0
	for dm_build_1280 > 0 {
		n, err := dm_build_1279.Read(dm_build_1278.dm_build_1248[dm_build_1281+dm_build_1282:])
		if n > 0 && err == io.EOF {
			dm_build_1282 += n
			dm_build_1278.dm_build_1248 = dm_build_1278.dm_build_1248[:dm_build_1281+dm_build_1282]
			return dm_build_1282, nil
		} else if n > 0 && err == nil {
			dm_build_1280 -= n
			dm_build_1282 += n
		} else if n == 0 && err != nil {
			return -1, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
		}
	}

	return dm_build_1282, nil
}

func (dm_build_1284 *Dm_build_1247) Dm_build_1283(dm_build_1285 io.Writer) (*Dm_build_1247, error) {
	if _, err := dm_build_1285.Write(dm_build_1284.dm_build_1248); err != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
	}
	return dm_build_1284, nil
}

func (dm_build_1287 *Dm_build_1247) Dm_build_1286(dm_build_1288 bool) int {
	dm_build_1289 := len(dm_build_1287.dm_build_1248)
	dm_build_1287.dm_build_1254(1)

	if dm_build_1288 {
		return copy(dm_build_1287.dm_build_1248[dm_build_1289:], []byte{1})
	} else {
		return copy(dm_build_1287.dm_build_1248[dm_build_1289:], []byte{0})
	}
}

func (dm_build_1291 *Dm_build_1247) Dm_build_1290(dm_build_1292 byte) int {
	dm_build_1293 := len(dm_build_1291.dm_build_1248)
	dm_build_1291.dm_build_1254(1)

	return copy(dm_build_1291.dm_build_1248[dm_build_1293:], Dm_build_890.Dm_build_1068(dm_build_1292))
}

func (dm_build_1295 *Dm_build_1247) Dm_build_1294(dm_build_1296 int16) int {
	dm_build_1297 := len(dm_build_1295.dm_build_1248)
	dm_build_1295.dm_build_1254(2)

	return copy(dm_build_1295.dm_build_1248[dm_build_1297:], Dm_build_890.Dm_build_1071(dm_build_1296))
}

func (dm_build_1299 *Dm_build_1247) Dm_build_1298(dm_build_1300 int32) int {
	dm_build_1301 := len(dm_build_1299.dm_build_1248)
	dm_build_1299.dm_build_1254(4)

	return copy(dm_build_1299.dm_build_1248[dm_build_1301:], Dm_build_890.Dm_build_1074(dm_build_1300))
}

func (dm_build_1303 *Dm_build_1247) Dm_build_1302(dm_build_1304 uint8) int {
	dm_build_1305 := len(dm_build_1303.dm_build_1248)
	dm_build_1303.dm_build_1254(1)

	return copy(dm_build_1303.dm_build_1248[dm_build_1305:], Dm_build_890.Dm_build_1086(dm_build_1304))
}

func (dm_build_1307 *Dm_build_1247) Dm_build_1306(dm_build_1308 uint16) int {
	dm_build_1309 := len(dm_build_1307.dm_build_1248)
	dm_build_1307.dm_build_1254(2)

	return copy(dm_build_1307.dm_build_1248[dm_build_1309:], Dm_build_890.Dm_build_1089(dm_build_1308))
}

func (dm_build_1311 *Dm_build_1247) Dm_build_1310(dm_build_1312 uint32) int {
	dm_build_1313 := len(dm_build_1311.dm_build_1248)
	dm_build_1311.dm_build_1254(4)

	return copy(dm_build_1311.dm_build_1248[dm_build_1313:], Dm_build_890.Dm_build_1092(dm_build_1312))
}

func (dm_build_1315 *Dm_build_1247) Dm_build_1314(dm_build_1316 uint64) int {
	dm_build_1317 := len(dm_build_1315.dm_build_1248)
	dm_build_1315.dm_build_1254(8)

	return copy(dm_build_1315.dm_build_1248[dm_build_1317:], Dm_build_890.Dm_build_1095(dm_build_1316))
}

func (dm_build_1319 *Dm_build_1247) Dm_build_1318(dm_build_1320 float32) int {
	dm_build_1321 := len(dm_build_1319.dm_build_1248)
	dm_build_1319.dm_build_1254(4)

	return copy(dm_build_1319.dm_build_1248[dm_build_1321:], Dm_build_890.Dm_build_1092(math.Float32bits(dm_build_1320)))
}

func (dm_build_1323 *Dm_build_1247) Dm_build_1322(dm_build_1324 float64) int {
	dm_build_1325 := len(dm_build_1323.dm_build_1248)
	dm_build_1323.dm_build_1254(8)

	return copy(dm_build_1323.dm_build_1248[dm_build_1325:], Dm_build_890.Dm_build_1095(math.Float64bits(dm_build_1324)))
}

func (dm_build_1327 *Dm_build_1247) Dm_build_1326(dm_build_1328 []byte) int {
	dm_build_1329 := len(dm_build_1327.dm_build_1248)
	dm_build_1327.dm_build_1254(len(dm_build_1328))
	return copy(dm_build_1327.dm_build_1248[dm_build_1329:], dm_build_1328)
}

func (dm_build_1331 *Dm_build_1247) Dm_build_1330(dm_build_1332 []byte) int {
	return dm_build_1331.Dm_build_1298(int32(len(dm_build_1332))) + dm_build_1331.Dm_build_1326(dm_build_1332)
}

func (dm_build_1334 *Dm_build_1247) Dm_build_1333(dm_build_1335 []byte) int {
	return dm_build_1334.Dm_build_1302(uint8(len(dm_build_1335))) + dm_build_1334.Dm_build_1326(dm_build_1335)
}

func (dm_build_1337 *Dm_build_1247) Dm_build_1336(dm_build_1338 []byte) int {
	return dm_build_1337.Dm_build_1306(uint16(len(dm_build_1338))) + dm_build_1337.Dm_build_1326(dm_build_1338)
}

func (dm_build_1340 *Dm_build_1247) Dm_build_1339(dm_build_1341 []byte) int {
	return dm_build_1340.Dm_build_1326(dm_build_1341) + dm_build_1340.Dm_build_1290(0)
}

func (dm_build_1343 *Dm_build_1247) Dm_build_1342(dm_build_1344 string, dm_build_1345 string, dm_build_1346 *DmConnection) int {
	dm_build_1347 := Dm_build_890.Dm_build_1103(dm_build_1344, dm_build_1345, dm_build_1346)
	return dm_build_1343.Dm_build_1330(dm_build_1347)
}

func (dm_build_1349 *Dm_build_1247) Dm_build_1348(dm_build_1350 string, dm_build_1351 string, dm_build_1352 *DmConnection) int {
	dm_build_1353 := Dm_build_890.Dm_build_1103(dm_build_1350, dm_build_1351, dm_build_1352)
	return dm_build_1349.Dm_build_1333(dm_build_1353)
}

func (dm_build_1355 *Dm_build_1247) Dm_build_1354(dm_build_1356 string, dm_build_1357 string, dm_build_1358 *DmConnection) int {
	dm_build_1359 := Dm_build_890.Dm_build_1103(dm_build_1356, dm_build_1357, dm_build_1358)
	return dm_build_1355.Dm_build_1336(dm_build_1359)
}

func (dm_build_1361 *Dm_build_1247) Dm_build_1360(dm_build_1362 string, dm_build_1363 string, dm_build_1364 *DmConnection) int {
	dm_build_1365 := Dm_build_890.Dm_build_1103(dm_build_1362, dm_build_1363, dm_build_1364)
	return dm_build_1361.Dm_build_1339(dm_build_1365)
}

func (dm_build_1367 *Dm_build_1247) Dm_build_1366() byte {
	dm_build_1368 := Dm_build_890.Dm_build_983(dm_build_1367.dm_build_1248, dm_build_1367.dm_build_1249)
	dm_build_1367.dm_build_1249++
	return dm_build_1368
}

func (dm_build_1370 *Dm_build_1247) Dm_build_1369() int16 {
	dm_build_1371 := Dm_build_890.Dm_build_987(dm_build_1370.dm_build_1248, dm_build_1370.dm_build_1249)
	dm_build_1370.dm_build_1249 += 2
	return dm_build_1371
}

func (dm_build_1373 *Dm_build_1247) Dm_build_1372() int32 {
	dm_build_1374 := Dm_build_890.Dm_build_992(dm_build_1373.dm_build_1248, dm_build_1373.dm_build_1249)
	dm_build_1373.dm_build_1249 += 4
	return dm_build_1374
}

func (dm_build_1376 *Dm_build_1247) Dm_build_1375() int64 {
	dm_build_1377 := Dm_build_890.Dm_build_997(dm_build_1376.dm_build_1248, dm_build_1376.dm_build_1249)
	dm_build_1376.dm_build_1249 += 8
	return dm_build_1377
}

func (dm_build_1379 *Dm_build_1247) Dm_build_1378() float32 {
	dm_build_1380 := Dm_build_890.Dm_build_1002(dm_build_1379.dm_build_1248, dm_build_1379.dm_build_1249)
	dm_build_1379.dm_build_1249 += 4
	return dm_build_1380
}

func (dm_build_1382 *Dm_build_1247) Dm_build_1381() float64 {
	dm_build_1383 := Dm_build_890.Dm_build_1006(dm_build_1382.dm_build_1248, dm_build_1382.dm_build_1249)
	dm_build_1382.dm_build_1249 += 8
	return dm_build_1383
}

func (dm_build_1385 *Dm_build_1247) Dm_build_1384() uint8 {
	dm_build_1386 := Dm_build_890.Dm_build_1010(dm_build_1385.dm_build_1248, dm_build_1385.dm_build_1249)
	dm_build_1385.dm_build_1249 += 1
	return dm_build_1386
}

func (dm_build_1388 *Dm_build_1247) Dm_build_1387() uint16 {
	dm_build_1389 := Dm_build_890.Dm_build_1014(dm_build_1388.dm_build_1248, dm_build_1388.dm_build_1249)
	dm_build_1388.dm_build_1249 += 2
	return dm_build_1389
}

func (dm_build_1391 *Dm_build_1247) Dm_build_1390() uint32 {
	dm_build_1392 := Dm_build_890.Dm_build_1019(dm_build_1391.dm_build_1248, dm_build_1391.dm_build_1249)
	dm_build_1391.dm_build_1249 += 4
	return dm_build_1392
}

func (dm_build_1394 *Dm_build_1247) Dm_build_1393(dm_build_1395 int) []byte {
	dm_build_1396 := Dm_build_890.Dm_build_1041(dm_build_1394.dm_build_1248, dm_build_1394.dm_build_1249, dm_build_1395)
	dm_build_1394.dm_build_1249 += dm_build_1395
	return dm_build_1396
}

func (dm_build_1398 *Dm_build_1247) Dm_build_1397() []byte {
	return dm_build_1398.Dm_build_1393(int(dm_build_1398.Dm_build_1372()))
}

func (dm_build_1400 *Dm_build_1247) Dm_build_1399() []byte {
	return dm_build_1400.Dm_build_1393(int(dm_build_1400.Dm_build_1366()))
}

func (dm_build_1402 *Dm_build_1247) Dm_build_1401() []byte {
	return dm_build_1402.Dm_build_1393(int(dm_build_1402.Dm_build_1369()))
}

func (dm_build_1404 *Dm_build_1247) Dm_build_1403(dm_build_1405 int) []byte {
	return dm_build_1404.Dm_build_1393(dm_build_1405)
}

func (dm_build_1407 *Dm_build_1247) Dm_build_1406() []byte {
	dm_build_1408 := 0
	for dm_build_1407.Dm_build_1366() != 0 {
		dm_build_1408++
	}
	dm_build_1407.Dm_build_1272(dm_build_1408, false, false)
	return dm_build_1407.Dm_build_1393(dm_build_1408)
}

func (dm_build_1410 *Dm_build_1247) Dm_build_1409(dm_build_1411 int, dm_build_1412 string, dm_build_1413 *DmConnection) string {
	return Dm_build_890.Dm_build_1140(dm_build_1410.Dm_build_1393(dm_build_1411), dm_build_1412, dm_build_1413)
}

func (dm_build_1415 *Dm_build_1247) Dm_build_1414(dm_build_1416 string, dm_build_1417 *DmConnection) string {
	return Dm_build_890.Dm_build_1140(dm_build_1415.Dm_build_1397(), dm_build_1416, dm_build_1417)
}

func (dm_build_1419 *Dm_build_1247) Dm_build_1418(dm_build_1420 string, dm_build_1421 *DmConnection) string {
	return Dm_build_890.Dm_build_1140(dm_build_1419.Dm_build_1399(), dm_build_1420, dm_build_1421)
}

func (dm_build_1423 *Dm_build_1247) Dm_build_1422(dm_build_1424 string, dm_build_1425 *DmConnection) string {
	return Dm_build_890.Dm_build_1140(dm_build_1423.Dm_build_1401(), dm_build_1424, dm_build_1425)
}

func (dm_build_1427 *Dm_build_1247) Dm_build_1426(dm_build_1428 string, dm_build_1429 *DmConnection) string {
	return Dm_build_890.Dm_build_1140(dm_build_1427.Dm_build_1406(), dm_build_1428, dm_build_1429)
}

func (dm_build_1431 *Dm_build_1247) Dm_build_1430(dm_build_1432 int, dm_build_1433 byte) int {
	return dm_build_1431.Dm_build_1466(dm_build_1432, Dm_build_890.Dm_build_1068(dm_build_1433))
}

func (dm_build_1435 *Dm_build_1247) Dm_build_1434(dm_build_1436 int, dm_build_1437 int16) int {
	return dm_build_1435.Dm_build_1466(dm_build_1436, Dm_build_890.Dm_build_1071(dm_build_1437))
}

func (dm_build_1439 *Dm_build_1247) Dm_build_1438(dm_build_1440 int, dm_build_1441 int32) int {
	return dm_build_1439.Dm_build_1466(dm_build_1440, Dm_build_890.Dm_build_1074(dm_build_1441))
}

func (dm_build_1443 *Dm_build_1247) Dm_build_1442(dm_build_1444 int, dm_build_1445 int64) int {
	return dm_build_1443.Dm_build_1466(dm_build_1444, Dm_build_890.Dm_build_1077(dm_build_1445))
}

func (dm_build_1447 *Dm_build_1247) Dm_build_1446(dm_build_1448 int, dm_build_1449 float32) int {
	return dm_build_1447.Dm_build_1466(dm_build_1448, Dm_build_890.Dm_build_1080(dm_build_1449))
}

func (dm_build_1451 *Dm_build_1247) Dm_build_1450(dm_build_1452 int, dm_build_1453 float64) int {
	return dm_build_1451.Dm_build_1466(dm_build_1452, Dm_build_890.Dm_build_1083(dm_build_1453))
}

func (dm_build_1455 *Dm_build_1247) Dm_build_1454(dm_build_1456 int, dm_build_1457 uint8) int {
	return dm_build_1455.Dm_build_1466(dm_build_1456, Dm_build_890.Dm_build_1086(dm_build_1457))
}

func (dm_build_1459 *Dm_build_1247) Dm_build_1458(dm_build_1460 int, dm_build_1461 uint16) int {
	return dm_build_1459.Dm_build_1466(dm_build_1460, Dm_build_890.Dm_build_1089(dm_build_1461))
}

func (dm_build_1463 *Dm_build_1247) Dm_build_1462(dm_build_1464 int, dm_build_1465 uint32) int {
	return dm_build_1463.Dm_build_1466(dm_build_1464, Dm_build_890.Dm_build_1092(dm_build_1465))
}

func (dm_build_1467 *Dm_build_1247) Dm_build_1466(dm_build_1468 int, dm_build_1469 []byte) int {
	return copy(dm_build_1467.dm_build_1248[dm_build_1468:], dm_build_1469)
}

func (dm_build_1471 *Dm_build_1247) Dm_build_1470(dm_build_1472 int, dm_build_1473 []byte) int {
	return dm_build_1471.Dm_build_1438(dm_build_1472, int32(len(dm_build_1473))) + dm_build_1471.Dm_build_1466(dm_build_1472+4, dm_build_1473)
}

func (dm_build_1475 *Dm_build_1247) Dm_build_1474(dm_build_1476 int, dm_build_1477 []byte) int {
	return dm_build_1475.Dm_build_1430(dm_build_1476, byte(len(dm_build_1477))) + dm_build_1475.Dm_build_1466(dm_build_1476+1, dm_build_1477)
}

func (dm_build_1479 *Dm_build_1247) Dm_build_1478(dm_build_1480 int, dm_build_1481 []byte) int {
	return dm_build_1479.Dm_build_1434(dm_build_1480, int16(len(dm_build_1481))) + dm_build_1479.Dm_build_1466(dm_build_1480+2, dm_build_1481)
}

func (dm_build_1483 *Dm_build_1247) Dm_build_1482(dm_build_1484 int, dm_build_1485 []byte) int {
	return dm_build_1483.Dm_build_1466(dm_build_1484, dm_build_1485) + dm_build_1483.Dm_build_1430(dm_build_1484+len(dm_build_1485), 0)
}

func (dm_build_1487 *Dm_build_1247) Dm_build_1486(dm_build_1488 int, dm_build_1489 string, dm_build_1490 string, dm_build_1491 *DmConnection) int {
	return dm_build_1487.Dm_build_1470(dm_build_1488, Dm_build_890.Dm_build_1103(dm_build_1489, dm_build_1490, dm_build_1491))
}

func (dm_build_1493 *Dm_build_1247) Dm_build_1492(dm_build_1494 int, dm_build_1495 string, dm_build_1496 string, dm_build_1497 *DmConnection) int {
	return dm_build_1493.Dm_build_1474(dm_build_1494, Dm_build_890.Dm_build_1103(dm_build_1495, dm_build_1496, dm_build_1497))
}

func (dm_build_1499 *Dm_build_1247) Dm_build_1498(dm_build_1500 int, dm_build_1501 string, dm_build_1502 string, dm_build_1503 *DmConnection) int {
	return dm_build_1499.Dm_build_1478(dm_build_1500, Dm_build_890.Dm_build_1103(dm_build_1501, dm_build_1502, dm_build_1503))
}

func (dm_build_1505 *Dm_build_1247) Dm_build_1504(dm_build_1506 int, dm_build_1507 string, dm_build_1508 string, dm_build_1509 *DmConnection) int {
	return dm_build_1505.Dm_build_1482(dm_build_1506, Dm_build_890.Dm_build_1103(dm_build_1507, dm_build_1508, dm_build_1509))
}

func (dm_build_1511 *Dm_build_1247) Dm_build_1510(dm_build_1512 int) byte {
	return Dm_build_890.Dm_build_1108(dm_build_1511.Dm_build_1537(dm_build_1512, 1))
}

func (dm_build_1514 *Dm_build_1247) Dm_build_1513(dm_build_1515 int) int16 {
	return Dm_build_890.Dm_build_1111(dm_build_1514.Dm_build_1537(dm_build_1515, 2))
}

func (dm_build_1517 *Dm_build_1247) Dm_build_1516(dm_build_1518 int) int32 {
	return Dm_build_890.Dm_build_1114(dm_build_1517.Dm_build_1537(dm_build_1518, 4))
}

func (dm_build_1520 *Dm_build_1247) Dm_build_1519(dm_build_1521 int) int64 {
	return Dm_build_890.Dm_build_1117(dm_build_1520.Dm_build_1537(dm_build_1521, 8))
}

func (dm_build_1523 *Dm_build_1247) Dm_build_1522(dm_build_1524 int) float32 {
	return Dm_build_890.Dm_build_1120(dm_build_1523.Dm_build_1537(dm_build_1524, 4))
}

func (dm_build_1526 *Dm_build_1247) Dm_build_1525(dm_build_1527 int) float64 {
	return Dm_build_890.Dm_build_1123(dm_build_1526.Dm_build_1537(dm_build_1527, 8))
}

func (dm_build_1529 *Dm_build_1247) Dm_build_1528(dm_build_1530 int) uint8 {
	return Dm_build_890.Dm_build_1126(dm_build_1529.Dm_build_1537(dm_build_1530, 1))
}

func (dm_build_1532 *Dm_build_1247) Dm_build_1531(dm_build_1533 int) uint16 {
	return Dm_build_890.Dm_build_1129(dm_build_1532.Dm_build_1537(dm_build_1533, 2))
}

func (dm_build_1535 *Dm_build_1247) Dm_build_1534(dm_build_1536 int) uint32 {
	return Dm_build_890.Dm_build_1132(dm_build_1535.Dm_build_1537(dm_build_1536, 4))
}

func (dm_build_1538 *Dm_build_1247) Dm_build_1537(dm_build_1539 int, dm_build_1540 int) []byte {
	return dm_build_1538.dm_build_1248[dm_build_1539 : dm_build_1539+dm_build_1540]
}

func (dm_build_1542 *Dm_build_1247) Dm_build_1541(dm_build_1543 int) []byte {
	dm_build_1544 := dm_build_1542.Dm_build_1516(dm_build_1543)
	return dm_build_1542.Dm_build_1537(dm_build_1543+4, int(dm_build_1544))
}

func (dm_build_1546 *Dm_build_1247) Dm_build_1545(dm_build_1547 int) []byte {
	dm_build_1548 := dm_build_1546.Dm_build_1510(dm_build_1547)
	return dm_build_1546.Dm_build_1537(dm_build_1547+1, int(dm_build_1548))
}

func (dm_build_1550 *Dm_build_1247) Dm_build_1549(dm_build_1551 int) []byte {
	dm_build_1552 := dm_build_1550.Dm_build_1513(dm_build_1551)
	return dm_build_1550.Dm_build_1537(dm_build_1551+2, int(dm_build_1552))
}

func (dm_build_1554 *Dm_build_1247) Dm_build_1553(dm_build_1555 int) []byte {
	dm_build_1556 := 0
	for dm_build_1554.Dm_build_1510(dm_build_1555) != 0 {
		dm_build_1555++
		dm_build_1556++
	}

	return dm_build_1554.Dm_build_1537(dm_build_1555-dm_build_1556, int(dm_build_1556))
}

func (dm_build_1558 *Dm_build_1247) Dm_build_1557(dm_build_1559 int, dm_build_1560 string, dm_build_1561 *DmConnection) string {
	return Dm_build_890.Dm_build_1140(dm_build_1558.Dm_build_1541(dm_build_1559), dm_build_1560, dm_build_1561)
}

func (dm_build_1563 *Dm_build_1247) Dm_build_1562(dm_build_1564 int, dm_build_1565 string, dm_build_1566 *DmConnection) string {
	return Dm_build_890.Dm_build_1140(dm_build_1563.Dm_build_1545(dm_build_1564), dm_build_1565, dm_build_1566)
}

func (dm_build_1568 *Dm_build_1247) Dm_build_1567(dm_build_1569 int, dm_build_1570 string, dm_build_1571 *DmConnection) string {
	return Dm_build_890.Dm_build_1140(dm_build_1568.Dm_build_1549(dm_build_1569), dm_build_1570, dm_build_1571)
}

func (dm_build_1573 *Dm_build_1247) Dm_build_1572(dm_build_1574 int, dm_build_1575 string, dm_build_1576 *DmConnection) string {
	return Dm_build_890.Dm_build_1140(dm_build_1573.Dm_build_1553(dm_build_1574), dm_build_1575, dm_build_1576)
}
