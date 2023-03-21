/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"io"
	"math"
)

type Dm_build_1274 struct {
	dm_build_1275 []byte
	dm_build_1276 int
}

func Dm_build_1277(dm_build_1278 int) *Dm_build_1274 {
	return &Dm_build_1274{make([]byte, 0, dm_build_1278), 0}
}

func Dm_build_1279(dm_build_1280 []byte) *Dm_build_1274 {
	return &Dm_build_1274{dm_build_1280, 0}
}

func (dm_build_1282 *Dm_build_1274) dm_build_1281(dm_build_1283 int) *Dm_build_1274 {

	dm_build_1284 := len(dm_build_1282.dm_build_1275)
	dm_build_1285 := cap(dm_build_1282.dm_build_1275)

	if dm_build_1284+dm_build_1283 <= dm_build_1285 {
		dm_build_1282.dm_build_1275 = dm_build_1282.dm_build_1275[:dm_build_1284+dm_build_1283]
	} else {

		var calCap = int64(math.Max(float64(2*dm_build_1285), float64(dm_build_1283+dm_build_1284)))

		nbuf := make([]byte, dm_build_1283+dm_build_1284, calCap)
		copy(nbuf, dm_build_1282.dm_build_1275)
		dm_build_1282.dm_build_1275 = nbuf
	}

	return dm_build_1282
}

func (dm_build_1287 *Dm_build_1274) Dm_build_1286() int {
	return len(dm_build_1287.dm_build_1275)
}

func (dm_build_1289 *Dm_build_1274) Dm_build_1288(dm_build_1290 int) *Dm_build_1274 {
	for i := dm_build_1290; i < len(dm_build_1289.dm_build_1275); i++ {
		dm_build_1289.dm_build_1275[i] = 0
	}
	dm_build_1289.dm_build_1275 = dm_build_1289.dm_build_1275[:dm_build_1290]
	return dm_build_1289
}

func (dm_build_1292 *Dm_build_1274) Dm_build_1291(dm_build_1293 int) *Dm_build_1274 {
	dm_build_1292.dm_build_1276 = dm_build_1293
	return dm_build_1292
}

func (dm_build_1295 *Dm_build_1274) Dm_build_1294() int {
	return dm_build_1295.dm_build_1276
}

func (dm_build_1297 *Dm_build_1274) Dm_build_1296(dm_build_1298 bool) int {
	return len(dm_build_1297.dm_build_1275) - dm_build_1297.dm_build_1276
}

func (dm_build_1300 *Dm_build_1274) Dm_build_1299(dm_build_1301 int, dm_build_1302 bool, dm_build_1303 bool) *Dm_build_1274 {

	if dm_build_1302 {
		if dm_build_1303 {
			dm_build_1300.dm_build_1281(dm_build_1301)
		} else {
			dm_build_1300.dm_build_1275 = dm_build_1300.dm_build_1275[:len(dm_build_1300.dm_build_1275)-dm_build_1301]
		}
	} else {
		if dm_build_1303 {
			dm_build_1300.dm_build_1276 += dm_build_1301
		} else {
			dm_build_1300.dm_build_1276 -= dm_build_1301
		}
	}

	return dm_build_1300
}

func (dm_build_1305 *Dm_build_1274) Dm_build_1304(dm_build_1306 io.Reader, dm_build_1307 int) (int, error) {
	dm_build_1308 := len(dm_build_1305.dm_build_1275)
	dm_build_1305.dm_build_1281(dm_build_1307)
	dm_build_1309 := 0
	for dm_build_1307 > 0 {
		n, err := dm_build_1306.Read(dm_build_1305.dm_build_1275[dm_build_1308+dm_build_1309:])
		if n > 0 && err == io.EOF {
			dm_build_1309 += n
			dm_build_1305.dm_build_1275 = dm_build_1305.dm_build_1275[:dm_build_1308+dm_build_1309]
			return dm_build_1309, nil
		} else if n > 0 && err == nil {
			dm_build_1307 -= n
			dm_build_1309 += n
		} else if n == 0 && err != nil {
			return -1, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
		}
	}

	return dm_build_1309, nil
}

func (dm_build_1311 *Dm_build_1274) Dm_build_1310(dm_build_1312 io.Writer) (*Dm_build_1274, error) {
	if _, err := dm_build_1312.Write(dm_build_1311.dm_build_1275); err != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
	}
	return dm_build_1311, nil
}

func (dm_build_1314 *Dm_build_1274) Dm_build_1313(dm_build_1315 bool) int {
	dm_build_1316 := len(dm_build_1314.dm_build_1275)
	dm_build_1314.dm_build_1281(1)

	if dm_build_1315 {
		return copy(dm_build_1314.dm_build_1275[dm_build_1316:], []byte{1})
	} else {
		return copy(dm_build_1314.dm_build_1275[dm_build_1316:], []byte{0})
	}
}

func (dm_build_1318 *Dm_build_1274) Dm_build_1317(dm_build_1319 byte) int {
	dm_build_1320 := len(dm_build_1318.dm_build_1275)
	dm_build_1318.dm_build_1281(1)

	return copy(dm_build_1318.dm_build_1275[dm_build_1320:], Dm_build_914.Dm_build_1092(dm_build_1319))
}

func (dm_build_1322 *Dm_build_1274) Dm_build_1321(dm_build_1323 int8) int {
	dm_build_1324 := len(dm_build_1322.dm_build_1275)
	dm_build_1322.dm_build_1281(1)

	return copy(dm_build_1322.dm_build_1275[dm_build_1324:], Dm_build_914.Dm_build_1095(dm_build_1323))
}

func (dm_build_1326 *Dm_build_1274) Dm_build_1325(dm_build_1327 int16) int {
	dm_build_1328 := len(dm_build_1326.dm_build_1275)
	dm_build_1326.dm_build_1281(2)

	return copy(dm_build_1326.dm_build_1275[dm_build_1328:], Dm_build_914.Dm_build_1098(dm_build_1327))
}

func (dm_build_1330 *Dm_build_1274) Dm_build_1329(dm_build_1331 int32) int {
	dm_build_1332 := len(dm_build_1330.dm_build_1275)
	dm_build_1330.dm_build_1281(4)

	return copy(dm_build_1330.dm_build_1275[dm_build_1332:], Dm_build_914.Dm_build_1101(dm_build_1331))
}

func (dm_build_1334 *Dm_build_1274) Dm_build_1333(dm_build_1335 uint8) int {
	dm_build_1336 := len(dm_build_1334.dm_build_1275)
	dm_build_1334.dm_build_1281(1)

	return copy(dm_build_1334.dm_build_1275[dm_build_1336:], Dm_build_914.Dm_build_1113(dm_build_1335))
}

func (dm_build_1338 *Dm_build_1274) Dm_build_1337(dm_build_1339 uint16) int {
	dm_build_1340 := len(dm_build_1338.dm_build_1275)
	dm_build_1338.dm_build_1281(2)

	return copy(dm_build_1338.dm_build_1275[dm_build_1340:], Dm_build_914.Dm_build_1116(dm_build_1339))
}

func (dm_build_1342 *Dm_build_1274) Dm_build_1341(dm_build_1343 uint32) int {
	dm_build_1344 := len(dm_build_1342.dm_build_1275)
	dm_build_1342.dm_build_1281(4)

	return copy(dm_build_1342.dm_build_1275[dm_build_1344:], Dm_build_914.Dm_build_1119(dm_build_1343))
}

func (dm_build_1346 *Dm_build_1274) Dm_build_1345(dm_build_1347 uint64) int {
	dm_build_1348 := len(dm_build_1346.dm_build_1275)
	dm_build_1346.dm_build_1281(8)

	return copy(dm_build_1346.dm_build_1275[dm_build_1348:], Dm_build_914.Dm_build_1122(dm_build_1347))
}

func (dm_build_1350 *Dm_build_1274) Dm_build_1349(dm_build_1351 float32) int {
	dm_build_1352 := len(dm_build_1350.dm_build_1275)
	dm_build_1350.dm_build_1281(4)

	return copy(dm_build_1350.dm_build_1275[dm_build_1352:], Dm_build_914.Dm_build_1119(math.Float32bits(dm_build_1351)))
}

func (dm_build_1354 *Dm_build_1274) Dm_build_1353(dm_build_1355 float64) int {
	dm_build_1356 := len(dm_build_1354.dm_build_1275)
	dm_build_1354.dm_build_1281(8)

	return copy(dm_build_1354.dm_build_1275[dm_build_1356:], Dm_build_914.Dm_build_1122(math.Float64bits(dm_build_1355)))
}

func (dm_build_1358 *Dm_build_1274) Dm_build_1357(dm_build_1359 []byte) int {
	dm_build_1360 := len(dm_build_1358.dm_build_1275)
	dm_build_1358.dm_build_1281(len(dm_build_1359))
	return copy(dm_build_1358.dm_build_1275[dm_build_1360:], dm_build_1359)
}

func (dm_build_1362 *Dm_build_1274) Dm_build_1361(dm_build_1363 []byte) int {
	return dm_build_1362.Dm_build_1329(int32(len(dm_build_1363))) + dm_build_1362.Dm_build_1357(dm_build_1363)
}

func (dm_build_1365 *Dm_build_1274) Dm_build_1364(dm_build_1366 []byte) int {
	return dm_build_1365.Dm_build_1333(uint8(len(dm_build_1366))) + dm_build_1365.Dm_build_1357(dm_build_1366)
}

func (dm_build_1368 *Dm_build_1274) Dm_build_1367(dm_build_1369 []byte) int {
	return dm_build_1368.Dm_build_1337(uint16(len(dm_build_1369))) + dm_build_1368.Dm_build_1357(dm_build_1369)
}

func (dm_build_1371 *Dm_build_1274) Dm_build_1370(dm_build_1372 []byte) int {
	return dm_build_1371.Dm_build_1357(dm_build_1372) + dm_build_1371.Dm_build_1317(0)
}

func (dm_build_1374 *Dm_build_1274) Dm_build_1373(dm_build_1375 string, dm_build_1376 string, dm_build_1377 *DmConnection) int {
	dm_build_1378 := Dm_build_914.Dm_build_1130(dm_build_1375, dm_build_1376, dm_build_1377)
	return dm_build_1374.Dm_build_1361(dm_build_1378)
}

func (dm_build_1380 *Dm_build_1274) Dm_build_1379(dm_build_1381 string, dm_build_1382 string, dm_build_1383 *DmConnection) int {
	dm_build_1384 := Dm_build_914.Dm_build_1130(dm_build_1381, dm_build_1382, dm_build_1383)
	return dm_build_1380.Dm_build_1364(dm_build_1384)
}

func (dm_build_1386 *Dm_build_1274) Dm_build_1385(dm_build_1387 string, dm_build_1388 string, dm_build_1389 *DmConnection) int {
	dm_build_1390 := Dm_build_914.Dm_build_1130(dm_build_1387, dm_build_1388, dm_build_1389)
	return dm_build_1386.Dm_build_1367(dm_build_1390)
}

func (dm_build_1392 *Dm_build_1274) Dm_build_1391(dm_build_1393 string, dm_build_1394 string, dm_build_1395 *DmConnection) int {
	dm_build_1396 := Dm_build_914.Dm_build_1130(dm_build_1393, dm_build_1394, dm_build_1395)
	return dm_build_1392.Dm_build_1370(dm_build_1396)
}

func (dm_build_1398 *Dm_build_1274) Dm_build_1397() byte {
	dm_build_1399 := Dm_build_914.Dm_build_1007(dm_build_1398.dm_build_1275, dm_build_1398.dm_build_1276)
	dm_build_1398.dm_build_1276++
	return dm_build_1399
}

func (dm_build_1401 *Dm_build_1274) Dm_build_1400() int16 {
	dm_build_1402 := Dm_build_914.Dm_build_1011(dm_build_1401.dm_build_1275, dm_build_1401.dm_build_1276)
	dm_build_1401.dm_build_1276 += 2
	return dm_build_1402
}

func (dm_build_1404 *Dm_build_1274) Dm_build_1403() int32 {
	dm_build_1405 := Dm_build_914.Dm_build_1016(dm_build_1404.dm_build_1275, dm_build_1404.dm_build_1276)
	dm_build_1404.dm_build_1276 += 4
	return dm_build_1405
}

func (dm_build_1407 *Dm_build_1274) Dm_build_1406() int64 {
	dm_build_1408 := Dm_build_914.Dm_build_1021(dm_build_1407.dm_build_1275, dm_build_1407.dm_build_1276)
	dm_build_1407.dm_build_1276 += 8
	return dm_build_1408
}

func (dm_build_1410 *Dm_build_1274) Dm_build_1409() float32 {
	dm_build_1411 := Dm_build_914.Dm_build_1026(dm_build_1410.dm_build_1275, dm_build_1410.dm_build_1276)
	dm_build_1410.dm_build_1276 += 4
	return dm_build_1411
}

func (dm_build_1413 *Dm_build_1274) Dm_build_1412() float64 {
	dm_build_1414 := Dm_build_914.Dm_build_1030(dm_build_1413.dm_build_1275, dm_build_1413.dm_build_1276)
	dm_build_1413.dm_build_1276 += 8
	return dm_build_1414
}

func (dm_build_1416 *Dm_build_1274) Dm_build_1415() uint8 {
	dm_build_1417 := Dm_build_914.Dm_build_1034(dm_build_1416.dm_build_1275, dm_build_1416.dm_build_1276)
	dm_build_1416.dm_build_1276 += 1
	return dm_build_1417
}

func (dm_build_1419 *Dm_build_1274) Dm_build_1418() uint16 {
	dm_build_1420 := Dm_build_914.Dm_build_1038(dm_build_1419.dm_build_1275, dm_build_1419.dm_build_1276)
	dm_build_1419.dm_build_1276 += 2
	return dm_build_1420
}

func (dm_build_1422 *Dm_build_1274) Dm_build_1421() uint32 {
	dm_build_1423 := Dm_build_914.Dm_build_1043(dm_build_1422.dm_build_1275, dm_build_1422.dm_build_1276)
	dm_build_1422.dm_build_1276 += 4
	return dm_build_1423
}

func (dm_build_1425 *Dm_build_1274) Dm_build_1424(dm_build_1426 int) []byte {
	dm_build_1427 := Dm_build_914.Dm_build_1065(dm_build_1425.dm_build_1275, dm_build_1425.dm_build_1276, dm_build_1426)
	dm_build_1425.dm_build_1276 += dm_build_1426
	return dm_build_1427
}

func (dm_build_1429 *Dm_build_1274) Dm_build_1428() []byte {
	return dm_build_1429.Dm_build_1424(int(dm_build_1429.Dm_build_1403()))
}

func (dm_build_1431 *Dm_build_1274) Dm_build_1430() []byte {
	return dm_build_1431.Dm_build_1424(int(dm_build_1431.Dm_build_1397()))
}

func (dm_build_1433 *Dm_build_1274) Dm_build_1432() []byte {
	return dm_build_1433.Dm_build_1424(int(dm_build_1433.Dm_build_1400()))
}

func (dm_build_1435 *Dm_build_1274) Dm_build_1434(dm_build_1436 int) []byte {
	return dm_build_1435.Dm_build_1424(dm_build_1436)
}

func (dm_build_1438 *Dm_build_1274) Dm_build_1437() []byte {
	dm_build_1439 := 0
	for dm_build_1438.Dm_build_1397() != 0 {
		dm_build_1439++
	}
	dm_build_1438.Dm_build_1299(dm_build_1439, false, false)
	return dm_build_1438.Dm_build_1424(dm_build_1439)
}

func (dm_build_1441 *Dm_build_1274) Dm_build_1440(dm_build_1442 int, dm_build_1443 string, dm_build_1444 *DmConnection) string {
	return Dm_build_914.Dm_build_1167(dm_build_1441.Dm_build_1424(dm_build_1442), dm_build_1443, dm_build_1444)
}

func (dm_build_1446 *Dm_build_1274) Dm_build_1445(dm_build_1447 string, dm_build_1448 *DmConnection) string {
	return Dm_build_914.Dm_build_1167(dm_build_1446.Dm_build_1428(), dm_build_1447, dm_build_1448)
}

func (dm_build_1450 *Dm_build_1274) Dm_build_1449(dm_build_1451 string, dm_build_1452 *DmConnection) string {
	return Dm_build_914.Dm_build_1167(dm_build_1450.Dm_build_1430(), dm_build_1451, dm_build_1452)
}

func (dm_build_1454 *Dm_build_1274) Dm_build_1453(dm_build_1455 string, dm_build_1456 *DmConnection) string {
	return Dm_build_914.Dm_build_1167(dm_build_1454.Dm_build_1432(), dm_build_1455, dm_build_1456)
}

func (dm_build_1458 *Dm_build_1274) Dm_build_1457(dm_build_1459 string, dm_build_1460 *DmConnection) string {
	return Dm_build_914.Dm_build_1167(dm_build_1458.Dm_build_1437(), dm_build_1459, dm_build_1460)
}

func (dm_build_1462 *Dm_build_1274) Dm_build_1461(dm_build_1463 int, dm_build_1464 byte) int {
	return dm_build_1462.Dm_build_1497(dm_build_1463, Dm_build_914.Dm_build_1092(dm_build_1464))
}

func (dm_build_1466 *Dm_build_1274) Dm_build_1465(dm_build_1467 int, dm_build_1468 int16) int {
	return dm_build_1466.Dm_build_1497(dm_build_1467, Dm_build_914.Dm_build_1098(dm_build_1468))
}

func (dm_build_1470 *Dm_build_1274) Dm_build_1469(dm_build_1471 int, dm_build_1472 int32) int {
	return dm_build_1470.Dm_build_1497(dm_build_1471, Dm_build_914.Dm_build_1101(dm_build_1472))
}

func (dm_build_1474 *Dm_build_1274) Dm_build_1473(dm_build_1475 int, dm_build_1476 int64) int {
	return dm_build_1474.Dm_build_1497(dm_build_1475, Dm_build_914.Dm_build_1104(dm_build_1476))
}

func (dm_build_1478 *Dm_build_1274) Dm_build_1477(dm_build_1479 int, dm_build_1480 float32) int {
	return dm_build_1478.Dm_build_1497(dm_build_1479, Dm_build_914.Dm_build_1107(dm_build_1480))
}

func (dm_build_1482 *Dm_build_1274) Dm_build_1481(dm_build_1483 int, dm_build_1484 float64) int {
	return dm_build_1482.Dm_build_1497(dm_build_1483, Dm_build_914.Dm_build_1110(dm_build_1484))
}

func (dm_build_1486 *Dm_build_1274) Dm_build_1485(dm_build_1487 int, dm_build_1488 uint8) int {
	return dm_build_1486.Dm_build_1497(dm_build_1487, Dm_build_914.Dm_build_1113(dm_build_1488))
}

func (dm_build_1490 *Dm_build_1274) Dm_build_1489(dm_build_1491 int, dm_build_1492 uint16) int {
	return dm_build_1490.Dm_build_1497(dm_build_1491, Dm_build_914.Dm_build_1116(dm_build_1492))
}

func (dm_build_1494 *Dm_build_1274) Dm_build_1493(dm_build_1495 int, dm_build_1496 uint32) int {
	return dm_build_1494.Dm_build_1497(dm_build_1495, Dm_build_914.Dm_build_1119(dm_build_1496))
}

func (dm_build_1498 *Dm_build_1274) Dm_build_1497(dm_build_1499 int, dm_build_1500 []byte) int {
	return copy(dm_build_1498.dm_build_1275[dm_build_1499:], dm_build_1500)
}

func (dm_build_1502 *Dm_build_1274) Dm_build_1501(dm_build_1503 int, dm_build_1504 []byte) int {
	return dm_build_1502.Dm_build_1469(dm_build_1503, int32(len(dm_build_1504))) + dm_build_1502.Dm_build_1497(dm_build_1503+4, dm_build_1504)
}

func (dm_build_1506 *Dm_build_1274) Dm_build_1505(dm_build_1507 int, dm_build_1508 []byte) int {
	return dm_build_1506.Dm_build_1461(dm_build_1507, byte(len(dm_build_1508))) + dm_build_1506.Dm_build_1497(dm_build_1507+1, dm_build_1508)
}

func (dm_build_1510 *Dm_build_1274) Dm_build_1509(dm_build_1511 int, dm_build_1512 []byte) int {
	return dm_build_1510.Dm_build_1465(dm_build_1511, int16(len(dm_build_1512))) + dm_build_1510.Dm_build_1497(dm_build_1511+2, dm_build_1512)
}

func (dm_build_1514 *Dm_build_1274) Dm_build_1513(dm_build_1515 int, dm_build_1516 []byte) int {
	return dm_build_1514.Dm_build_1497(dm_build_1515, dm_build_1516) + dm_build_1514.Dm_build_1461(dm_build_1515+len(dm_build_1516), 0)
}

func (dm_build_1518 *Dm_build_1274) Dm_build_1517(dm_build_1519 int, dm_build_1520 string, dm_build_1521 string, dm_build_1522 *DmConnection) int {
	return dm_build_1518.Dm_build_1501(dm_build_1519, Dm_build_914.Dm_build_1130(dm_build_1520, dm_build_1521, dm_build_1522))
}

func (dm_build_1524 *Dm_build_1274) Dm_build_1523(dm_build_1525 int, dm_build_1526 string, dm_build_1527 string, dm_build_1528 *DmConnection) int {
	return dm_build_1524.Dm_build_1505(dm_build_1525, Dm_build_914.Dm_build_1130(dm_build_1526, dm_build_1527, dm_build_1528))
}

func (dm_build_1530 *Dm_build_1274) Dm_build_1529(dm_build_1531 int, dm_build_1532 string, dm_build_1533 string, dm_build_1534 *DmConnection) int {
	return dm_build_1530.Dm_build_1509(dm_build_1531, Dm_build_914.Dm_build_1130(dm_build_1532, dm_build_1533, dm_build_1534))
}

func (dm_build_1536 *Dm_build_1274) Dm_build_1535(dm_build_1537 int, dm_build_1538 string, dm_build_1539 string, dm_build_1540 *DmConnection) int {
	return dm_build_1536.Dm_build_1513(dm_build_1537, Dm_build_914.Dm_build_1130(dm_build_1538, dm_build_1539, dm_build_1540))
}

func (dm_build_1542 *Dm_build_1274) Dm_build_1541(dm_build_1543 int) byte {
	return Dm_build_914.Dm_build_1135(dm_build_1542.Dm_build_1568(dm_build_1543, 1))
}

func (dm_build_1545 *Dm_build_1274) Dm_build_1544(dm_build_1546 int) int16 {
	return Dm_build_914.Dm_build_1138(dm_build_1545.Dm_build_1568(dm_build_1546, 2))
}

func (dm_build_1548 *Dm_build_1274) Dm_build_1547(dm_build_1549 int) int32 {
	return Dm_build_914.Dm_build_1141(dm_build_1548.Dm_build_1568(dm_build_1549, 4))
}

func (dm_build_1551 *Dm_build_1274) Dm_build_1550(dm_build_1552 int) int64 {
	return Dm_build_914.Dm_build_1144(dm_build_1551.Dm_build_1568(dm_build_1552, 8))
}

func (dm_build_1554 *Dm_build_1274) Dm_build_1553(dm_build_1555 int) float32 {
	return Dm_build_914.Dm_build_1147(dm_build_1554.Dm_build_1568(dm_build_1555, 4))
}

func (dm_build_1557 *Dm_build_1274) Dm_build_1556(dm_build_1558 int) float64 {
	return Dm_build_914.Dm_build_1150(dm_build_1557.Dm_build_1568(dm_build_1558, 8))
}

func (dm_build_1560 *Dm_build_1274) Dm_build_1559(dm_build_1561 int) uint8 {
	return Dm_build_914.Dm_build_1153(dm_build_1560.Dm_build_1568(dm_build_1561, 1))
}

func (dm_build_1563 *Dm_build_1274) Dm_build_1562(dm_build_1564 int) uint16 {
	return Dm_build_914.Dm_build_1156(dm_build_1563.Dm_build_1568(dm_build_1564, 2))
}

func (dm_build_1566 *Dm_build_1274) Dm_build_1565(dm_build_1567 int) uint32 {
	return Dm_build_914.Dm_build_1159(dm_build_1566.Dm_build_1568(dm_build_1567, 4))
}

func (dm_build_1569 *Dm_build_1274) Dm_build_1568(dm_build_1570 int, dm_build_1571 int) []byte {
	return dm_build_1569.dm_build_1275[dm_build_1570 : dm_build_1570+dm_build_1571]
}

func (dm_build_1573 *Dm_build_1274) Dm_build_1572(dm_build_1574 int) []byte {
	dm_build_1575 := dm_build_1573.Dm_build_1547(dm_build_1574)
	return dm_build_1573.Dm_build_1568(dm_build_1574+4, int(dm_build_1575))
}

func (dm_build_1577 *Dm_build_1274) Dm_build_1576(dm_build_1578 int) []byte {
	dm_build_1579 := dm_build_1577.Dm_build_1541(dm_build_1578)
	return dm_build_1577.Dm_build_1568(dm_build_1578+1, int(dm_build_1579))
}

func (dm_build_1581 *Dm_build_1274) Dm_build_1580(dm_build_1582 int) []byte {
	dm_build_1583 := dm_build_1581.Dm_build_1544(dm_build_1582)
	return dm_build_1581.Dm_build_1568(dm_build_1582+2, int(dm_build_1583))
}

func (dm_build_1585 *Dm_build_1274) Dm_build_1584(dm_build_1586 int) []byte {
	dm_build_1587 := 0
	for dm_build_1585.Dm_build_1541(dm_build_1586) != 0 {
		dm_build_1586++
		dm_build_1587++
	}

	return dm_build_1585.Dm_build_1568(dm_build_1586-dm_build_1587, int(dm_build_1587))
}

func (dm_build_1589 *Dm_build_1274) Dm_build_1588(dm_build_1590 int, dm_build_1591 string, dm_build_1592 *DmConnection) string {
	return Dm_build_914.Dm_build_1167(dm_build_1589.Dm_build_1572(dm_build_1590), dm_build_1591, dm_build_1592)
}

func (dm_build_1594 *Dm_build_1274) Dm_build_1593(dm_build_1595 int, dm_build_1596 string, dm_build_1597 *DmConnection) string {
	return Dm_build_914.Dm_build_1167(dm_build_1594.Dm_build_1576(dm_build_1595), dm_build_1596, dm_build_1597)
}

func (dm_build_1599 *Dm_build_1274) Dm_build_1598(dm_build_1600 int, dm_build_1601 string, dm_build_1602 *DmConnection) string {
	return Dm_build_914.Dm_build_1167(dm_build_1599.Dm_build_1580(dm_build_1600), dm_build_1601, dm_build_1602)
}

func (dm_build_1604 *Dm_build_1274) Dm_build_1603(dm_build_1605 int, dm_build_1606 string, dm_build_1607 *DmConnection) string {
	return Dm_build_914.Dm_build_1167(dm_build_1604.Dm_build_1584(dm_build_1605), dm_build_1606, dm_build_1607)
}
