/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_913 struct{}

var Dm_build_914 = &dm_build_913{}

func (Dm_build_916 *dm_build_913) Dm_build_915(dm_build_917 []byte, dm_build_918 int, dm_build_919 byte) int {
	dm_build_917[dm_build_918] = dm_build_919
	return 1
}

func (Dm_build_921 *dm_build_913) Dm_build_920(dm_build_922 []byte, dm_build_923 int, dm_build_924 int8) int {
	dm_build_922[dm_build_923] = byte(dm_build_924)
	return 1
}

func (Dm_build_926 *dm_build_913) Dm_build_925(dm_build_927 []byte, dm_build_928 int, dm_build_929 int16) int {
	dm_build_927[dm_build_928] = byte(dm_build_929)
	dm_build_928++
	dm_build_927[dm_build_928] = byte(dm_build_929 >> 8)
	return 2
}

func (Dm_build_931 *dm_build_913) Dm_build_930(dm_build_932 []byte, dm_build_933 int, dm_build_934 int32) int {
	dm_build_932[dm_build_933] = byte(dm_build_934)
	dm_build_933++
	dm_build_932[dm_build_933] = byte(dm_build_934 >> 8)
	dm_build_933++
	dm_build_932[dm_build_933] = byte(dm_build_934 >> 16)
	dm_build_933++
	dm_build_932[dm_build_933] = byte(dm_build_934 >> 24)
	dm_build_933++
	return 4
}

func (Dm_build_936 *dm_build_913) Dm_build_935(dm_build_937 []byte, dm_build_938 int, dm_build_939 int64) int {
	dm_build_937[dm_build_938] = byte(dm_build_939)
	dm_build_938++
	dm_build_937[dm_build_938] = byte(dm_build_939 >> 8)
	dm_build_938++
	dm_build_937[dm_build_938] = byte(dm_build_939 >> 16)
	dm_build_938++
	dm_build_937[dm_build_938] = byte(dm_build_939 >> 24)
	dm_build_938++
	dm_build_937[dm_build_938] = byte(dm_build_939 >> 32)
	dm_build_938++
	dm_build_937[dm_build_938] = byte(dm_build_939 >> 40)
	dm_build_938++
	dm_build_937[dm_build_938] = byte(dm_build_939 >> 48)
	dm_build_938++
	dm_build_937[dm_build_938] = byte(dm_build_939 >> 56)
	return 8
}

func (Dm_build_941 *dm_build_913) Dm_build_940(dm_build_942 []byte, dm_build_943 int, dm_build_944 float32) int {
	return Dm_build_941.Dm_build_960(dm_build_942, dm_build_943, math.Float32bits(dm_build_944))
}

func (Dm_build_946 *dm_build_913) Dm_build_945(dm_build_947 []byte, dm_build_948 int, dm_build_949 float64) int {
	return Dm_build_946.Dm_build_965(dm_build_947, dm_build_948, math.Float64bits(dm_build_949))
}

func (Dm_build_951 *dm_build_913) Dm_build_950(dm_build_952 []byte, dm_build_953 int, dm_build_954 uint8) int {
	dm_build_952[dm_build_953] = byte(dm_build_954)
	return 1
}

func (Dm_build_956 *dm_build_913) Dm_build_955(dm_build_957 []byte, dm_build_958 int, dm_build_959 uint16) int {
	dm_build_957[dm_build_958] = byte(dm_build_959)
	dm_build_958++
	dm_build_957[dm_build_958] = byte(dm_build_959 >> 8)
	return 2
}

func (Dm_build_961 *dm_build_913) Dm_build_960(dm_build_962 []byte, dm_build_963 int, dm_build_964 uint32) int {
	dm_build_962[dm_build_963] = byte(dm_build_964)
	dm_build_963++
	dm_build_962[dm_build_963] = byte(dm_build_964 >> 8)
	dm_build_963++
	dm_build_962[dm_build_963] = byte(dm_build_964 >> 16)
	dm_build_963++
	dm_build_962[dm_build_963] = byte(dm_build_964 >> 24)
	return 3
}

func (Dm_build_966 *dm_build_913) Dm_build_965(dm_build_967 []byte, dm_build_968 int, dm_build_969 uint64) int {
	dm_build_967[dm_build_968] = byte(dm_build_969)
	dm_build_968++
	dm_build_967[dm_build_968] = byte(dm_build_969 >> 8)
	dm_build_968++
	dm_build_967[dm_build_968] = byte(dm_build_969 >> 16)
	dm_build_968++
	dm_build_967[dm_build_968] = byte(dm_build_969 >> 24)
	dm_build_968++
	dm_build_967[dm_build_968] = byte(dm_build_969 >> 32)
	dm_build_968++
	dm_build_967[dm_build_968] = byte(dm_build_969 >> 40)
	dm_build_968++
	dm_build_967[dm_build_968] = byte(dm_build_969 >> 48)
	dm_build_968++
	dm_build_967[dm_build_968] = byte(dm_build_969 >> 56)
	return 3
}

func (Dm_build_971 *dm_build_913) Dm_build_970(dm_build_972 []byte, dm_build_973 int, dm_build_974 []byte, dm_build_975 int, dm_build_976 int) int {
	copy(dm_build_972[dm_build_973:dm_build_973+dm_build_976], dm_build_974[dm_build_975:dm_build_975+dm_build_976])
	return dm_build_976
}

func (Dm_build_978 *dm_build_913) Dm_build_977(dm_build_979 []byte, dm_build_980 int, dm_build_981 []byte, dm_build_982 int, dm_build_983 int) int {
	dm_build_980 += Dm_build_978.Dm_build_960(dm_build_979, dm_build_980, uint32(dm_build_983))
	return 4 + Dm_build_978.Dm_build_970(dm_build_979, dm_build_980, dm_build_981, dm_build_982, dm_build_983)
}

func (Dm_build_985 *dm_build_913) Dm_build_984(dm_build_986 []byte, dm_build_987 int, dm_build_988 []byte, dm_build_989 int, dm_build_990 int) int {
	dm_build_987 += Dm_build_985.Dm_build_955(dm_build_986, dm_build_987, uint16(dm_build_990))
	return 2 + Dm_build_985.Dm_build_970(dm_build_986, dm_build_987, dm_build_988, dm_build_989, dm_build_990)
}

func (Dm_build_992 *dm_build_913) Dm_build_991(dm_build_993 []byte, dm_build_994 int, dm_build_995 string, dm_build_996 string, dm_build_997 *DmConnection) int {
	dm_build_998 := Dm_build_992.Dm_build_1130(dm_build_995, dm_build_996, dm_build_997)
	dm_build_994 += Dm_build_992.Dm_build_960(dm_build_993, dm_build_994, uint32(len(dm_build_998)))
	return 4 + Dm_build_992.Dm_build_970(dm_build_993, dm_build_994, dm_build_998, 0, len(dm_build_998))
}

func (Dm_build_1000 *dm_build_913) Dm_build_999(dm_build_1001 []byte, dm_build_1002 int, dm_build_1003 string, dm_build_1004 string, dm_build_1005 *DmConnection) int {
	dm_build_1006 := Dm_build_1000.Dm_build_1130(dm_build_1003, dm_build_1004, dm_build_1005)

	dm_build_1002 += Dm_build_1000.Dm_build_955(dm_build_1001, dm_build_1002, uint16(len(dm_build_1006)))
	return 2 + Dm_build_1000.Dm_build_970(dm_build_1001, dm_build_1002, dm_build_1006, 0, len(dm_build_1006))
}

func (Dm_build_1008 *dm_build_913) Dm_build_1007(dm_build_1009 []byte, dm_build_1010 int) byte {
	return dm_build_1009[dm_build_1010]
}

func (Dm_build_1012 *dm_build_913) Dm_build_1011(dm_build_1013 []byte, dm_build_1014 int) int16 {
	var dm_build_1015 int16
	dm_build_1015 = int16(dm_build_1013[dm_build_1014] & 0xff)
	dm_build_1014++
	dm_build_1015 |= int16(dm_build_1013[dm_build_1014]&0xff) << 8
	return dm_build_1015
}

func (Dm_build_1017 *dm_build_913) Dm_build_1016(dm_build_1018 []byte, dm_build_1019 int) int32 {
	var dm_build_1020 int32
	dm_build_1020 = int32(dm_build_1018[dm_build_1019] & 0xff)
	dm_build_1019++
	dm_build_1020 |= int32(dm_build_1018[dm_build_1019]&0xff) << 8
	dm_build_1019++
	dm_build_1020 |= int32(dm_build_1018[dm_build_1019]&0xff) << 16
	dm_build_1019++
	dm_build_1020 |= int32(dm_build_1018[dm_build_1019]&0xff) << 24
	return dm_build_1020
}

func (Dm_build_1022 *dm_build_913) Dm_build_1021(dm_build_1023 []byte, dm_build_1024 int) int64 {
	var dm_build_1025 int64
	dm_build_1025 = int64(dm_build_1023[dm_build_1024] & 0xff)
	dm_build_1024++
	dm_build_1025 |= int64(dm_build_1023[dm_build_1024]&0xff) << 8
	dm_build_1024++
	dm_build_1025 |= int64(dm_build_1023[dm_build_1024]&0xff) << 16
	dm_build_1024++
	dm_build_1025 |= int64(dm_build_1023[dm_build_1024]&0xff) << 24
	dm_build_1024++
	dm_build_1025 |= int64(dm_build_1023[dm_build_1024]&0xff) << 32
	dm_build_1024++
	dm_build_1025 |= int64(dm_build_1023[dm_build_1024]&0xff) << 40
	dm_build_1024++
	dm_build_1025 |= int64(dm_build_1023[dm_build_1024]&0xff) << 48
	dm_build_1024++
	dm_build_1025 |= int64(dm_build_1023[dm_build_1024]&0xff) << 56
	return dm_build_1025
}

func (Dm_build_1027 *dm_build_913) Dm_build_1026(dm_build_1028 []byte, dm_build_1029 int) float32 {
	return math.Float32frombits(Dm_build_1027.Dm_build_1043(dm_build_1028, dm_build_1029))
}

func (Dm_build_1031 *dm_build_913) Dm_build_1030(dm_build_1032 []byte, dm_build_1033 int) float64 {
	return math.Float64frombits(Dm_build_1031.Dm_build_1048(dm_build_1032, dm_build_1033))
}

func (Dm_build_1035 *dm_build_913) Dm_build_1034(dm_build_1036 []byte, dm_build_1037 int) uint8 {
	return uint8(dm_build_1036[dm_build_1037] & 0xff)
}

func (Dm_build_1039 *dm_build_913) Dm_build_1038(dm_build_1040 []byte, dm_build_1041 int) uint16 {
	var dm_build_1042 uint16
	dm_build_1042 = uint16(dm_build_1040[dm_build_1041] & 0xff)
	dm_build_1041++
	dm_build_1042 |= uint16(dm_build_1040[dm_build_1041]&0xff) << 8
	return dm_build_1042
}

func (Dm_build_1044 *dm_build_913) Dm_build_1043(dm_build_1045 []byte, dm_build_1046 int) uint32 {
	var dm_build_1047 uint32
	dm_build_1047 = uint32(dm_build_1045[dm_build_1046] & 0xff)
	dm_build_1046++
	dm_build_1047 |= uint32(dm_build_1045[dm_build_1046]&0xff) << 8
	dm_build_1046++
	dm_build_1047 |= uint32(dm_build_1045[dm_build_1046]&0xff) << 16
	dm_build_1046++
	dm_build_1047 |= uint32(dm_build_1045[dm_build_1046]&0xff) << 24
	return dm_build_1047
}

func (Dm_build_1049 *dm_build_913) Dm_build_1048(dm_build_1050 []byte, dm_build_1051 int) uint64 {
	var dm_build_1052 uint64
	dm_build_1052 = uint64(dm_build_1050[dm_build_1051] & 0xff)
	dm_build_1051++
	dm_build_1052 |= uint64(dm_build_1050[dm_build_1051]&0xff) << 8
	dm_build_1051++
	dm_build_1052 |= uint64(dm_build_1050[dm_build_1051]&0xff) << 16
	dm_build_1051++
	dm_build_1052 |= uint64(dm_build_1050[dm_build_1051]&0xff) << 24
	dm_build_1051++
	dm_build_1052 |= uint64(dm_build_1050[dm_build_1051]&0xff) << 32
	dm_build_1051++
	dm_build_1052 |= uint64(dm_build_1050[dm_build_1051]&0xff) << 40
	dm_build_1051++
	dm_build_1052 |= uint64(dm_build_1050[dm_build_1051]&0xff) << 48
	dm_build_1051++
	dm_build_1052 |= uint64(dm_build_1050[dm_build_1051]&0xff) << 56
	return dm_build_1052
}

func (Dm_build_1054 *dm_build_913) Dm_build_1053(dm_build_1055 []byte, dm_build_1056 int) []byte {
	dm_build_1057 := Dm_build_1054.Dm_build_1043(dm_build_1055, dm_build_1056)

	dm_build_1058 := make([]byte, dm_build_1057)
	copy(dm_build_1058[:int(dm_build_1057)], dm_build_1055[dm_build_1056+4:dm_build_1056+4+int(dm_build_1057)])
	return dm_build_1058
}

func (Dm_build_1060 *dm_build_913) Dm_build_1059(dm_build_1061 []byte, dm_build_1062 int) []byte {
	dm_build_1063 := Dm_build_1060.Dm_build_1038(dm_build_1061, dm_build_1062)

	dm_build_1064 := make([]byte, dm_build_1063)
	copy(dm_build_1064[:int(dm_build_1063)], dm_build_1061[dm_build_1062+2:dm_build_1062+2+int(dm_build_1063)])
	return dm_build_1064
}

func (Dm_build_1066 *dm_build_913) Dm_build_1065(dm_build_1067 []byte, dm_build_1068 int, dm_build_1069 int) []byte {

	dm_build_1070 := make([]byte, dm_build_1069)
	copy(dm_build_1070[:dm_build_1069], dm_build_1067[dm_build_1068:dm_build_1068+dm_build_1069])
	return dm_build_1070
}

func (Dm_build_1072 *dm_build_913) Dm_build_1071(dm_build_1073 []byte, dm_build_1074 int, dm_build_1075 int, dm_build_1076 string, dm_build_1077 *DmConnection) string {
	return Dm_build_1072.Dm_build_1167(dm_build_1073[dm_build_1074:dm_build_1074+dm_build_1075], dm_build_1076, dm_build_1077)
}

func (Dm_build_1079 *dm_build_913) Dm_build_1078(dm_build_1080 []byte, dm_build_1081 int, dm_build_1082 string, dm_build_1083 *DmConnection) string {
	dm_build_1084 := Dm_build_1079.Dm_build_1043(dm_build_1080, dm_build_1081)
	dm_build_1081 += 4
	return Dm_build_1079.Dm_build_1071(dm_build_1080, dm_build_1081, int(dm_build_1084), dm_build_1082, dm_build_1083)
}

func (Dm_build_1086 *dm_build_913) Dm_build_1085(dm_build_1087 []byte, dm_build_1088 int, dm_build_1089 string, dm_build_1090 *DmConnection) string {
	dm_build_1091 := Dm_build_1086.Dm_build_1038(dm_build_1087, dm_build_1088)
	dm_build_1088 += 2
	return Dm_build_1086.Dm_build_1071(dm_build_1087, dm_build_1088, int(dm_build_1091), dm_build_1089, dm_build_1090)
}

func (Dm_build_1093 *dm_build_913) Dm_build_1092(dm_build_1094 byte) []byte {
	return []byte{dm_build_1094}
}

func (Dm_build_1096 *dm_build_913) Dm_build_1095(dm_build_1097 int8) []byte {
	return []byte{byte(dm_build_1097)}
}

func (Dm_build_1099 *dm_build_913) Dm_build_1098(dm_build_1100 int16) []byte {
	return []byte{byte(dm_build_1100), byte(dm_build_1100 >> 8)}
}

func (Dm_build_1102 *dm_build_913) Dm_build_1101(dm_build_1103 int32) []byte {
	return []byte{byte(dm_build_1103), byte(dm_build_1103 >> 8), byte(dm_build_1103 >> 16), byte(dm_build_1103 >> 24)}
}

func (Dm_build_1105 *dm_build_913) Dm_build_1104(dm_build_1106 int64) []byte {
	return []byte{byte(dm_build_1106), byte(dm_build_1106 >> 8), byte(dm_build_1106 >> 16), byte(dm_build_1106 >> 24), byte(dm_build_1106 >> 32),
		byte(dm_build_1106 >> 40), byte(dm_build_1106 >> 48), byte(dm_build_1106 >> 56)}
}

func (Dm_build_1108 *dm_build_913) Dm_build_1107(dm_build_1109 float32) []byte {
	return Dm_build_1108.Dm_build_1119(math.Float32bits(dm_build_1109))
}

func (Dm_build_1111 *dm_build_913) Dm_build_1110(dm_build_1112 float64) []byte {
	return Dm_build_1111.Dm_build_1122(math.Float64bits(dm_build_1112))
}

func (Dm_build_1114 *dm_build_913) Dm_build_1113(dm_build_1115 uint8) []byte {
	return []byte{byte(dm_build_1115)}
}

func (Dm_build_1117 *dm_build_913) Dm_build_1116(dm_build_1118 uint16) []byte {
	return []byte{byte(dm_build_1118), byte(dm_build_1118 >> 8)}
}

func (Dm_build_1120 *dm_build_913) Dm_build_1119(dm_build_1121 uint32) []byte {
	return []byte{byte(dm_build_1121), byte(dm_build_1121 >> 8), byte(dm_build_1121 >> 16), byte(dm_build_1121 >> 24)}
}

func (Dm_build_1123 *dm_build_913) Dm_build_1122(dm_build_1124 uint64) []byte {
	return []byte{byte(dm_build_1124), byte(dm_build_1124 >> 8), byte(dm_build_1124 >> 16), byte(dm_build_1124 >> 24), byte(dm_build_1124 >> 32), byte(dm_build_1124 >> 40), byte(dm_build_1124 >> 48), byte(dm_build_1124 >> 56)}
}

func (Dm_build_1126 *dm_build_913) Dm_build_1125(dm_build_1127 []byte, dm_build_1128 string, dm_build_1129 *DmConnection) []byte {
	if dm_build_1128 == "UTF-8" {
		return dm_build_1127
	}

	if dm_build_1129 == nil {
		if e := dm_build_1172(dm_build_1128); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1127), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1129.encodeBuffer == nil {
		dm_build_1129.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1129.encode = dm_build_1172(dm_build_1129.getServerEncoding())
		dm_build_1129.transformReaderDst = make([]byte, 4096)
		dm_build_1129.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1129.encode; e != nil {

		dm_build_1129.encodeBuffer.Reset()

		n, err := dm_build_1129.encodeBuffer.ReadFrom(
			Dm_build_1186(bytes.NewReader(dm_build_1127), e.NewEncoder(), dm_build_1129.transformReaderDst, dm_build_1129.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1129.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1131 *dm_build_913) Dm_build_1130(dm_build_1132 string, dm_build_1133 string, dm_build_1134 *DmConnection) []byte {
	return Dm_build_1131.Dm_build_1125([]byte(dm_build_1132), dm_build_1133, dm_build_1134)
}

func (Dm_build_1136 *dm_build_913) Dm_build_1135(dm_build_1137 []byte) byte {
	return Dm_build_1136.Dm_build_1007(dm_build_1137, 0)
}

func (Dm_build_1139 *dm_build_913) Dm_build_1138(dm_build_1140 []byte) int16 {
	return Dm_build_1139.Dm_build_1011(dm_build_1140, 0)
}

func (Dm_build_1142 *dm_build_913) Dm_build_1141(dm_build_1143 []byte) int32 {
	return Dm_build_1142.Dm_build_1016(dm_build_1143, 0)
}

func (Dm_build_1145 *dm_build_913) Dm_build_1144(dm_build_1146 []byte) int64 {
	return Dm_build_1145.Dm_build_1021(dm_build_1146, 0)
}

func (Dm_build_1148 *dm_build_913) Dm_build_1147(dm_build_1149 []byte) float32 {
	return Dm_build_1148.Dm_build_1026(dm_build_1149, 0)
}

func (Dm_build_1151 *dm_build_913) Dm_build_1150(dm_build_1152 []byte) float64 {
	return Dm_build_1151.Dm_build_1030(dm_build_1152, 0)
}

func (Dm_build_1154 *dm_build_913) Dm_build_1153(dm_build_1155 []byte) uint8 {
	return Dm_build_1154.Dm_build_1034(dm_build_1155, 0)
}

func (Dm_build_1157 *dm_build_913) Dm_build_1156(dm_build_1158 []byte) uint16 {
	return Dm_build_1157.Dm_build_1038(dm_build_1158, 0)
}

func (Dm_build_1160 *dm_build_913) Dm_build_1159(dm_build_1161 []byte) uint32 {
	return Dm_build_1160.Dm_build_1043(dm_build_1161, 0)
}

func (Dm_build_1163 *dm_build_913) Dm_build_1162(dm_build_1164 []byte, dm_build_1165 string, dm_build_1166 *DmConnection) []byte {
	if dm_build_1165 == "UTF-8" {
		return dm_build_1164
	}

	if dm_build_1166 == nil {
		if e := dm_build_1172(dm_build_1165); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1164), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1166.encodeBuffer == nil {
		dm_build_1166.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1166.encode = dm_build_1172(dm_build_1166.getServerEncoding())
		dm_build_1166.transformReaderDst = make([]byte, 4096)
		dm_build_1166.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1166.encode; e != nil {

		dm_build_1166.encodeBuffer.Reset()

		n, err := dm_build_1166.encodeBuffer.ReadFrom(
			Dm_build_1186(bytes.NewReader(dm_build_1164), e.NewDecoder(), dm_build_1166.transformReaderDst, dm_build_1166.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_1166.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1168 *dm_build_913) Dm_build_1167(dm_build_1169 []byte, dm_build_1170 string, dm_build_1171 *DmConnection) string {
	return string(Dm_build_1168.Dm_build_1162(dm_build_1169, dm_build_1170, dm_build_1171))
}

func dm_build_1172(dm_build_1173 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1173); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1174 struct {
	dm_build_1175 io.Reader
	dm_build_1176 transform.Transformer
	dm_build_1177 error

	dm_build_1178                []byte
	dm_build_1179, dm_build_1180 int

	dm_build_1181                []byte
	dm_build_1182, dm_build_1183 int

	dm_build_1184 bool
}

const dm_build_1185 = 4096

func Dm_build_1186(dm_build_1187 io.Reader, dm_build_1188 transform.Transformer, dm_build_1189 []byte, dm_build_1190 []byte) *Dm_build_1174 {
	dm_build_1188.Reset()
	return &Dm_build_1174{
		dm_build_1175: dm_build_1187,
		dm_build_1176: dm_build_1188,
		dm_build_1178: dm_build_1189,
		dm_build_1181: dm_build_1190,
	}
}

func (dm_build_1192 *Dm_build_1174) Read(dm_build_1193 []byte) (int, error) {
	dm_build_1194, dm_build_1195 := 0, error(nil)
	for {

		if dm_build_1192.dm_build_1179 != dm_build_1192.dm_build_1180 {
			dm_build_1194 = copy(dm_build_1193, dm_build_1192.dm_build_1178[dm_build_1192.dm_build_1179:dm_build_1192.dm_build_1180])
			dm_build_1192.dm_build_1179 += dm_build_1194
			if dm_build_1192.dm_build_1179 == dm_build_1192.dm_build_1180 && dm_build_1192.dm_build_1184 {
				return dm_build_1194, dm_build_1192.dm_build_1177
			}
			return dm_build_1194, nil
		} else if dm_build_1192.dm_build_1184 {
			return 0, dm_build_1192.dm_build_1177
		}

		if dm_build_1192.dm_build_1182 != dm_build_1192.dm_build_1183 || dm_build_1192.dm_build_1177 != nil {
			dm_build_1192.dm_build_1179 = 0
			dm_build_1192.dm_build_1180, dm_build_1194, dm_build_1195 = dm_build_1192.dm_build_1176.Transform(dm_build_1192.dm_build_1178, dm_build_1192.dm_build_1181[dm_build_1192.dm_build_1182:dm_build_1192.dm_build_1183], dm_build_1192.dm_build_1177 == io.EOF)
			dm_build_1192.dm_build_1182 += dm_build_1194

			switch {
			case dm_build_1195 == nil:
				if dm_build_1192.dm_build_1182 != dm_build_1192.dm_build_1183 {
					dm_build_1192.dm_build_1177 = nil
				}

				dm_build_1192.dm_build_1184 = dm_build_1192.dm_build_1177 != nil
				continue
			case dm_build_1195 == transform.ErrShortDst && (dm_build_1192.dm_build_1180 != 0 || dm_build_1194 != 0):

				continue
			case dm_build_1195 == transform.ErrShortSrc && dm_build_1192.dm_build_1183-dm_build_1192.dm_build_1182 != len(dm_build_1192.dm_build_1181) && dm_build_1192.dm_build_1177 == nil:

			default:
				dm_build_1192.dm_build_1184 = true

				if dm_build_1192.dm_build_1177 == nil || dm_build_1192.dm_build_1177 == io.EOF {
					dm_build_1192.dm_build_1177 = dm_build_1195
				}
				continue
			}
		}

		if dm_build_1192.dm_build_1182 != 0 {
			dm_build_1192.dm_build_1182, dm_build_1192.dm_build_1183 = 0, copy(dm_build_1192.dm_build_1181, dm_build_1192.dm_build_1181[dm_build_1192.dm_build_1182:dm_build_1192.dm_build_1183])
		}
		dm_build_1194, dm_build_1192.dm_build_1177 = dm_build_1192.dm_build_1175.Read(dm_build_1192.dm_build_1181[dm_build_1192.dm_build_1183:])
		dm_build_1192.dm_build_1183 += dm_build_1194
	}
}
