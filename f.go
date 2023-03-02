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

type dm_build_889 struct{}

var Dm_build_890 = &dm_build_889{}

func (Dm_build_892 *dm_build_889) Dm_build_891(dm_build_893 []byte, dm_build_894 int, dm_build_895 byte) int {
	dm_build_893[dm_build_894] = dm_build_895
	return 1
}

func (Dm_build_897 *dm_build_889) Dm_build_896(dm_build_898 []byte, dm_build_899 int, dm_build_900 int8) int {
	dm_build_898[dm_build_899] = byte(dm_build_900)
	return 1
}

func (Dm_build_902 *dm_build_889) Dm_build_901(dm_build_903 []byte, dm_build_904 int, dm_build_905 int16) int {
	dm_build_903[dm_build_904] = byte(dm_build_905)
	dm_build_904++
	dm_build_903[dm_build_904] = byte(dm_build_905 >> 8)
	return 2
}

func (Dm_build_907 *dm_build_889) Dm_build_906(dm_build_908 []byte, dm_build_909 int, dm_build_910 int32) int {
	dm_build_908[dm_build_909] = byte(dm_build_910)
	dm_build_909++
	dm_build_908[dm_build_909] = byte(dm_build_910 >> 8)
	dm_build_909++
	dm_build_908[dm_build_909] = byte(dm_build_910 >> 16)
	dm_build_909++
	dm_build_908[dm_build_909] = byte(dm_build_910 >> 24)
	dm_build_909++
	return 4
}

func (Dm_build_912 *dm_build_889) Dm_build_911(dm_build_913 []byte, dm_build_914 int, dm_build_915 int64) int {
	dm_build_913[dm_build_914] = byte(dm_build_915)
	dm_build_914++
	dm_build_913[dm_build_914] = byte(dm_build_915 >> 8)
	dm_build_914++
	dm_build_913[dm_build_914] = byte(dm_build_915 >> 16)
	dm_build_914++
	dm_build_913[dm_build_914] = byte(dm_build_915 >> 24)
	dm_build_914++
	dm_build_913[dm_build_914] = byte(dm_build_915 >> 32)
	dm_build_914++
	dm_build_913[dm_build_914] = byte(dm_build_915 >> 40)
	dm_build_914++
	dm_build_913[dm_build_914] = byte(dm_build_915 >> 48)
	dm_build_914++
	dm_build_913[dm_build_914] = byte(dm_build_915 >> 56)
	return 8
}

func (Dm_build_917 *dm_build_889) Dm_build_916(dm_build_918 []byte, dm_build_919 int, dm_build_920 float32) int {
	return Dm_build_917.Dm_build_936(dm_build_918, dm_build_919, math.Float32bits(dm_build_920))
}

func (Dm_build_922 *dm_build_889) Dm_build_921(dm_build_923 []byte, dm_build_924 int, dm_build_925 float64) int {
	return Dm_build_922.Dm_build_941(dm_build_923, dm_build_924, math.Float64bits(dm_build_925))
}

func (Dm_build_927 *dm_build_889) Dm_build_926(dm_build_928 []byte, dm_build_929 int, dm_build_930 uint8) int {
	dm_build_928[dm_build_929] = byte(dm_build_930)
	return 1
}

func (Dm_build_932 *dm_build_889) Dm_build_931(dm_build_933 []byte, dm_build_934 int, dm_build_935 uint16) int {
	dm_build_933[dm_build_934] = byte(dm_build_935)
	dm_build_934++
	dm_build_933[dm_build_934] = byte(dm_build_935 >> 8)
	return 2
}

func (Dm_build_937 *dm_build_889) Dm_build_936(dm_build_938 []byte, dm_build_939 int, dm_build_940 uint32) int {
	dm_build_938[dm_build_939] = byte(dm_build_940)
	dm_build_939++
	dm_build_938[dm_build_939] = byte(dm_build_940 >> 8)
	dm_build_939++
	dm_build_938[dm_build_939] = byte(dm_build_940 >> 16)
	dm_build_939++
	dm_build_938[dm_build_939] = byte(dm_build_940 >> 24)
	return 3
}

func (Dm_build_942 *dm_build_889) Dm_build_941(dm_build_943 []byte, dm_build_944 int, dm_build_945 uint64) int {
	dm_build_943[dm_build_944] = byte(dm_build_945)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 8)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 16)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 24)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 32)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 40)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 48)
	dm_build_944++
	dm_build_943[dm_build_944] = byte(dm_build_945 >> 56)
	return 3
}

func (Dm_build_947 *dm_build_889) Dm_build_946(dm_build_948 []byte, dm_build_949 int, dm_build_950 []byte, dm_build_951 int, dm_build_952 int) int {
	copy(dm_build_948[dm_build_949:dm_build_949+dm_build_952], dm_build_950[dm_build_951:dm_build_951+dm_build_952])
	return dm_build_952
}

func (Dm_build_954 *dm_build_889) Dm_build_953(dm_build_955 []byte, dm_build_956 int, dm_build_957 []byte, dm_build_958 int, dm_build_959 int) int {
	dm_build_956 += Dm_build_954.Dm_build_936(dm_build_955, dm_build_956, uint32(dm_build_959))
	return 4 + Dm_build_954.Dm_build_946(dm_build_955, dm_build_956, dm_build_957, dm_build_958, dm_build_959)
}

func (Dm_build_961 *dm_build_889) Dm_build_960(dm_build_962 []byte, dm_build_963 int, dm_build_964 []byte, dm_build_965 int, dm_build_966 int) int {
	dm_build_963 += Dm_build_961.Dm_build_931(dm_build_962, dm_build_963, uint16(dm_build_966))
	return 2 + Dm_build_961.Dm_build_946(dm_build_962, dm_build_963, dm_build_964, dm_build_965, dm_build_966)
}

func (Dm_build_968 *dm_build_889) Dm_build_967(dm_build_969 []byte, dm_build_970 int, dm_build_971 string, dm_build_972 string, dm_build_973 *DmConnection) int {
	dm_build_974 := Dm_build_968.Dm_build_1103(dm_build_971, dm_build_972, dm_build_973)
	dm_build_970 += Dm_build_968.Dm_build_936(dm_build_969, dm_build_970, uint32(len(dm_build_974)))
	return 4 + Dm_build_968.Dm_build_946(dm_build_969, dm_build_970, dm_build_974, 0, len(dm_build_974))
}

func (Dm_build_976 *dm_build_889) Dm_build_975(dm_build_977 []byte, dm_build_978 int, dm_build_979 string, dm_build_980 string, dm_build_981 *DmConnection) int {
	dm_build_982 := Dm_build_976.Dm_build_1103(dm_build_979, dm_build_980, dm_build_981)

	dm_build_978 += Dm_build_976.Dm_build_931(dm_build_977, dm_build_978, uint16(len(dm_build_982)))
	return 2 + Dm_build_976.Dm_build_946(dm_build_977, dm_build_978, dm_build_982, 0, len(dm_build_982))
}

func (Dm_build_984 *dm_build_889) Dm_build_983(dm_build_985 []byte, dm_build_986 int) byte {
	return dm_build_985[dm_build_986]
}

func (Dm_build_988 *dm_build_889) Dm_build_987(dm_build_989 []byte, dm_build_990 int) int16 {
	var dm_build_991 int16
	dm_build_991 = int16(dm_build_989[dm_build_990] & 0xff)
	dm_build_990++
	dm_build_991 |= int16(dm_build_989[dm_build_990]&0xff) << 8
	return dm_build_991
}

func (Dm_build_993 *dm_build_889) Dm_build_992(dm_build_994 []byte, dm_build_995 int) int32 {
	var dm_build_996 int32
	dm_build_996 = int32(dm_build_994[dm_build_995] & 0xff)
	dm_build_995++
	dm_build_996 |= int32(dm_build_994[dm_build_995]&0xff) << 8
	dm_build_995++
	dm_build_996 |= int32(dm_build_994[dm_build_995]&0xff) << 16
	dm_build_995++
	dm_build_996 |= int32(dm_build_994[dm_build_995]&0xff) << 24
	return dm_build_996
}

func (Dm_build_998 *dm_build_889) Dm_build_997(dm_build_999 []byte, dm_build_1000 int) int64 {
	var dm_build_1001 int64
	dm_build_1001 = int64(dm_build_999[dm_build_1000] & 0xff)
	dm_build_1000++
	dm_build_1001 |= int64(dm_build_999[dm_build_1000]&0xff) << 8
	dm_build_1000++
	dm_build_1001 |= int64(dm_build_999[dm_build_1000]&0xff) << 16
	dm_build_1000++
	dm_build_1001 |= int64(dm_build_999[dm_build_1000]&0xff) << 24
	dm_build_1000++
	dm_build_1001 |= int64(dm_build_999[dm_build_1000]&0xff) << 32
	dm_build_1000++
	dm_build_1001 |= int64(dm_build_999[dm_build_1000]&0xff) << 40
	dm_build_1000++
	dm_build_1001 |= int64(dm_build_999[dm_build_1000]&0xff) << 48
	dm_build_1000++
	dm_build_1001 |= int64(dm_build_999[dm_build_1000]&0xff) << 56
	return dm_build_1001
}

func (Dm_build_1003 *dm_build_889) Dm_build_1002(dm_build_1004 []byte, dm_build_1005 int) float32 {
	return math.Float32frombits(Dm_build_1003.Dm_build_1019(dm_build_1004, dm_build_1005))
}

func (Dm_build_1007 *dm_build_889) Dm_build_1006(dm_build_1008 []byte, dm_build_1009 int) float64 {
	return math.Float64frombits(Dm_build_1007.Dm_build_1024(dm_build_1008, dm_build_1009))
}

func (Dm_build_1011 *dm_build_889) Dm_build_1010(dm_build_1012 []byte, dm_build_1013 int) uint8 {
	return uint8(dm_build_1012[dm_build_1013] & 0xff)
}

func (Dm_build_1015 *dm_build_889) Dm_build_1014(dm_build_1016 []byte, dm_build_1017 int) uint16 {
	var dm_build_1018 uint16
	dm_build_1018 = uint16(dm_build_1016[dm_build_1017] & 0xff)
	dm_build_1017++
	dm_build_1018 |= uint16(dm_build_1016[dm_build_1017]&0xff) << 8
	return dm_build_1018
}

func (Dm_build_1020 *dm_build_889) Dm_build_1019(dm_build_1021 []byte, dm_build_1022 int) uint32 {
	var dm_build_1023 uint32
	dm_build_1023 = uint32(dm_build_1021[dm_build_1022] & 0xff)
	dm_build_1022++
	dm_build_1023 |= uint32(dm_build_1021[dm_build_1022]&0xff) << 8
	dm_build_1022++
	dm_build_1023 |= uint32(dm_build_1021[dm_build_1022]&0xff) << 16
	dm_build_1022++
	dm_build_1023 |= uint32(dm_build_1021[dm_build_1022]&0xff) << 24
	return dm_build_1023
}

func (Dm_build_1025 *dm_build_889) Dm_build_1024(dm_build_1026 []byte, dm_build_1027 int) uint64 {
	var dm_build_1028 uint64
	dm_build_1028 = uint64(dm_build_1026[dm_build_1027] & 0xff)
	dm_build_1027++
	dm_build_1028 |= uint64(dm_build_1026[dm_build_1027]&0xff) << 8
	dm_build_1027++
	dm_build_1028 |= uint64(dm_build_1026[dm_build_1027]&0xff) << 16
	dm_build_1027++
	dm_build_1028 |= uint64(dm_build_1026[dm_build_1027]&0xff) << 24
	dm_build_1027++
	dm_build_1028 |= uint64(dm_build_1026[dm_build_1027]&0xff) << 32
	dm_build_1027++
	dm_build_1028 |= uint64(dm_build_1026[dm_build_1027]&0xff) << 40
	dm_build_1027++
	dm_build_1028 |= uint64(dm_build_1026[dm_build_1027]&0xff) << 48
	dm_build_1027++
	dm_build_1028 |= uint64(dm_build_1026[dm_build_1027]&0xff) << 56
	return dm_build_1028
}

func (Dm_build_1030 *dm_build_889) Dm_build_1029(dm_build_1031 []byte, dm_build_1032 int) []byte {
	dm_build_1033 := Dm_build_1030.Dm_build_1019(dm_build_1031, dm_build_1032)

	dm_build_1034 := make([]byte, dm_build_1033)
	copy(dm_build_1034[:int(dm_build_1033)], dm_build_1031[dm_build_1032+4:dm_build_1032+4+int(dm_build_1033)])
	return dm_build_1034
}

func (Dm_build_1036 *dm_build_889) Dm_build_1035(dm_build_1037 []byte, dm_build_1038 int) []byte {
	dm_build_1039 := Dm_build_1036.Dm_build_1014(dm_build_1037, dm_build_1038)

	dm_build_1040 := make([]byte, dm_build_1039)
	copy(dm_build_1040[:int(dm_build_1039)], dm_build_1037[dm_build_1038+2:dm_build_1038+2+int(dm_build_1039)])
	return dm_build_1040
}

func (Dm_build_1042 *dm_build_889) Dm_build_1041(dm_build_1043 []byte, dm_build_1044 int, dm_build_1045 int) []byte {

	dm_build_1046 := make([]byte, dm_build_1045)
	copy(dm_build_1046[:dm_build_1045], dm_build_1043[dm_build_1044:dm_build_1044+dm_build_1045])
	return dm_build_1046
}

func (Dm_build_1048 *dm_build_889) Dm_build_1047(dm_build_1049 []byte, dm_build_1050 int, dm_build_1051 int, dm_build_1052 string, dm_build_1053 *DmConnection) string {
	return Dm_build_1048.Dm_build_1140(dm_build_1049[dm_build_1050:dm_build_1050+dm_build_1051], dm_build_1052, dm_build_1053)
}

func (Dm_build_1055 *dm_build_889) Dm_build_1054(dm_build_1056 []byte, dm_build_1057 int, dm_build_1058 string, dm_build_1059 *DmConnection) string {
	dm_build_1060 := Dm_build_1055.Dm_build_1019(dm_build_1056, dm_build_1057)
	dm_build_1057 += 4
	return Dm_build_1055.Dm_build_1047(dm_build_1056, dm_build_1057, int(dm_build_1060), dm_build_1058, dm_build_1059)
}

func (Dm_build_1062 *dm_build_889) Dm_build_1061(dm_build_1063 []byte, dm_build_1064 int, dm_build_1065 string, dm_build_1066 *DmConnection) string {
	dm_build_1067 := Dm_build_1062.Dm_build_1014(dm_build_1063, dm_build_1064)
	dm_build_1064 += 2
	return Dm_build_1062.Dm_build_1047(dm_build_1063, dm_build_1064, int(dm_build_1067), dm_build_1065, dm_build_1066)
}

func (Dm_build_1069 *dm_build_889) Dm_build_1068(dm_build_1070 byte) []byte {
	return []byte{dm_build_1070}
}

func (Dm_build_1072 *dm_build_889) Dm_build_1071(dm_build_1073 int16) []byte {
	return []byte{byte(dm_build_1073), byte(dm_build_1073 >> 8)}
}

func (Dm_build_1075 *dm_build_889) Dm_build_1074(dm_build_1076 int32) []byte {
	return []byte{byte(dm_build_1076), byte(dm_build_1076 >> 8), byte(dm_build_1076 >> 16), byte(dm_build_1076 >> 24)}
}

func (Dm_build_1078 *dm_build_889) Dm_build_1077(dm_build_1079 int64) []byte {
	return []byte{byte(dm_build_1079), byte(dm_build_1079 >> 8), byte(dm_build_1079 >> 16), byte(dm_build_1079 >> 24), byte(dm_build_1079 >> 32),
		byte(dm_build_1079 >> 40), byte(dm_build_1079 >> 48), byte(dm_build_1079 >> 56)}
}

func (Dm_build_1081 *dm_build_889) Dm_build_1080(dm_build_1082 float32) []byte {
	return Dm_build_1081.Dm_build_1092(math.Float32bits(dm_build_1082))
}

func (Dm_build_1084 *dm_build_889) Dm_build_1083(dm_build_1085 float64) []byte {
	return Dm_build_1084.Dm_build_1095(math.Float64bits(dm_build_1085))
}

func (Dm_build_1087 *dm_build_889) Dm_build_1086(dm_build_1088 uint8) []byte {
	return []byte{byte(dm_build_1088)}
}

func (Dm_build_1090 *dm_build_889) Dm_build_1089(dm_build_1091 uint16) []byte {
	return []byte{byte(dm_build_1091), byte(dm_build_1091 >> 8)}
}

func (Dm_build_1093 *dm_build_889) Dm_build_1092(dm_build_1094 uint32) []byte {
	return []byte{byte(dm_build_1094), byte(dm_build_1094 >> 8), byte(dm_build_1094 >> 16), byte(dm_build_1094 >> 24)}
}

func (Dm_build_1096 *dm_build_889) Dm_build_1095(dm_build_1097 uint64) []byte {
	return []byte{byte(dm_build_1097), byte(dm_build_1097 >> 8), byte(dm_build_1097 >> 16), byte(dm_build_1097 >> 24), byte(dm_build_1097 >> 32), byte(dm_build_1097 >> 40), byte(dm_build_1097 >> 48), byte(dm_build_1097 >> 56)}
}

func (Dm_build_1099 *dm_build_889) Dm_build_1098(dm_build_1100 []byte, dm_build_1101 string, dm_build_1102 *DmConnection) []byte {
	if dm_build_1101 == "UTF-8" {
		return dm_build_1100
	}

	if dm_build_1102 == nil {
		if e := dm_build_1145(dm_build_1101); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1100), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1102.encodeBuffer == nil {
		dm_build_1102.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1102.encode = dm_build_1145(dm_build_1102.getServerEncoding())
		dm_build_1102.transformReaderDst = make([]byte, 4096)
		dm_build_1102.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1102.encode; e != nil {

		dm_build_1102.encodeBuffer.Reset()

		n, err := dm_build_1102.encodeBuffer.ReadFrom(
			Dm_build_1159(bytes.NewReader(dm_build_1100), e.NewEncoder(), dm_build_1102.transformReaderDst, dm_build_1102.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1102.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1104 *dm_build_889) Dm_build_1103(dm_build_1105 string, dm_build_1106 string, dm_build_1107 *DmConnection) []byte {
	return Dm_build_1104.Dm_build_1098([]byte(dm_build_1105), dm_build_1106, dm_build_1107)
}

func (Dm_build_1109 *dm_build_889) Dm_build_1108(dm_build_1110 []byte) byte {
	return Dm_build_1109.Dm_build_983(dm_build_1110, 0)
}

func (Dm_build_1112 *dm_build_889) Dm_build_1111(dm_build_1113 []byte) int16 {
	return Dm_build_1112.Dm_build_987(dm_build_1113, 0)
}

func (Dm_build_1115 *dm_build_889) Dm_build_1114(dm_build_1116 []byte) int32 {
	return Dm_build_1115.Dm_build_992(dm_build_1116, 0)
}

func (Dm_build_1118 *dm_build_889) Dm_build_1117(dm_build_1119 []byte) int64 {
	return Dm_build_1118.Dm_build_997(dm_build_1119, 0)
}

func (Dm_build_1121 *dm_build_889) Dm_build_1120(dm_build_1122 []byte) float32 {
	return Dm_build_1121.Dm_build_1002(dm_build_1122, 0)
}

func (Dm_build_1124 *dm_build_889) Dm_build_1123(dm_build_1125 []byte) float64 {
	return Dm_build_1124.Dm_build_1006(dm_build_1125, 0)
}

func (Dm_build_1127 *dm_build_889) Dm_build_1126(dm_build_1128 []byte) uint8 {
	return Dm_build_1127.Dm_build_1010(dm_build_1128, 0)
}

func (Dm_build_1130 *dm_build_889) Dm_build_1129(dm_build_1131 []byte) uint16 {
	return Dm_build_1130.Dm_build_1014(dm_build_1131, 0)
}

func (Dm_build_1133 *dm_build_889) Dm_build_1132(dm_build_1134 []byte) uint32 {
	return Dm_build_1133.Dm_build_1019(dm_build_1134, 0)
}

func (Dm_build_1136 *dm_build_889) Dm_build_1135(dm_build_1137 []byte, dm_build_1138 string, dm_build_1139 *DmConnection) []byte {
	if dm_build_1138 == "UTF-8" {
		return dm_build_1137
	}

	if dm_build_1139 == nil {
		if e := dm_build_1145(dm_build_1138); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1137), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1139.encodeBuffer == nil {
		dm_build_1139.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1139.encode = dm_build_1145(dm_build_1139.getServerEncoding())
		dm_build_1139.transformReaderDst = make([]byte, 4096)
		dm_build_1139.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1139.encode; e != nil {

		dm_build_1139.encodeBuffer.Reset()

		n, err := dm_build_1139.encodeBuffer.ReadFrom(
			Dm_build_1159(bytes.NewReader(dm_build_1137), e.NewDecoder(), dm_build_1139.transformReaderDst, dm_build_1139.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_1139.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1141 *dm_build_889) Dm_build_1140(dm_build_1142 []byte, dm_build_1143 string, dm_build_1144 *DmConnection) string {
	return string(Dm_build_1141.Dm_build_1135(dm_build_1142, dm_build_1143, dm_build_1144))
}

func dm_build_1145(dm_build_1146 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1146); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1147 struct {
	dm_build_1148 io.Reader
	dm_build_1149 transform.Transformer
	dm_build_1150 error

	dm_build_1151                []byte
	dm_build_1152, dm_build_1153 int

	dm_build_1154                []byte
	dm_build_1155, dm_build_1156 int

	dm_build_1157 bool
}

const dm_build_1158 = 4096

func Dm_build_1159(dm_build_1160 io.Reader, dm_build_1161 transform.Transformer, dm_build_1162 []byte, dm_build_1163 []byte) *Dm_build_1147 {
	dm_build_1161.Reset()
	return &Dm_build_1147{
		dm_build_1148: dm_build_1160,
		dm_build_1149: dm_build_1161,
		dm_build_1151: dm_build_1162,
		dm_build_1154: dm_build_1163,
	}
}

func (dm_build_1165 *Dm_build_1147) Read(dm_build_1166 []byte) (int, error) {
	dm_build_1167, dm_build_1168 := 0, error(nil)
	for {

		if dm_build_1165.dm_build_1152 != dm_build_1165.dm_build_1153 {
			dm_build_1167 = copy(dm_build_1166, dm_build_1165.dm_build_1151[dm_build_1165.dm_build_1152:dm_build_1165.dm_build_1153])
			dm_build_1165.dm_build_1152 += dm_build_1167
			if dm_build_1165.dm_build_1152 == dm_build_1165.dm_build_1153 && dm_build_1165.dm_build_1157 {
				return dm_build_1167, dm_build_1165.dm_build_1150
			}
			return dm_build_1167, nil
		} else if dm_build_1165.dm_build_1157 {
			return 0, dm_build_1165.dm_build_1150
		}

		if dm_build_1165.dm_build_1155 != dm_build_1165.dm_build_1156 || dm_build_1165.dm_build_1150 != nil {
			dm_build_1165.dm_build_1152 = 0
			dm_build_1165.dm_build_1153, dm_build_1167, dm_build_1168 = dm_build_1165.dm_build_1149.Transform(dm_build_1165.dm_build_1151, dm_build_1165.dm_build_1154[dm_build_1165.dm_build_1155:dm_build_1165.dm_build_1156], dm_build_1165.dm_build_1150 == io.EOF)
			dm_build_1165.dm_build_1155 += dm_build_1167

			switch {
			case dm_build_1168 == nil:
				if dm_build_1165.dm_build_1155 != dm_build_1165.dm_build_1156 {
					dm_build_1165.dm_build_1150 = nil
				}

				dm_build_1165.dm_build_1157 = dm_build_1165.dm_build_1150 != nil
				continue
			case dm_build_1168 == transform.ErrShortDst && (dm_build_1165.dm_build_1153 != 0 || dm_build_1167 != 0):

				continue
			case dm_build_1168 == transform.ErrShortSrc && dm_build_1165.dm_build_1156-dm_build_1165.dm_build_1155 != len(dm_build_1165.dm_build_1154) && dm_build_1165.dm_build_1150 == nil:

			default:
				dm_build_1165.dm_build_1157 = true

				if dm_build_1165.dm_build_1150 == nil || dm_build_1165.dm_build_1150 == io.EOF {
					dm_build_1165.dm_build_1150 = dm_build_1168
				}
				continue
			}
		}

		if dm_build_1165.dm_build_1155 != 0 {
			dm_build_1165.dm_build_1155, dm_build_1165.dm_build_1156 = 0, copy(dm_build_1165.dm_build_1154, dm_build_1165.dm_build_1154[dm_build_1165.dm_build_1155:dm_build_1165.dm_build_1156])
		}
		dm_build_1167, dm_build_1165.dm_build_1150 = dm_build_1165.dm_build_1148.Read(dm_build_1165.dm_build_1154[dm_build_1165.dm_build_1156:])
		dm_build_1165.dm_build_1156 += dm_build_1167
	}
}
