package config

var AreaData map[uint]string
var DistrictData map[uint]string
var Relation map[int][]int

func init() {
	AreaData = map[uint]string{0: "不限", 1: "港島", 2: "九龍", 3: "新界", 4: "離島", 5: "海外"}
	DistrictData = map[uint]string{1: "鰂魚涌", 2: "太古/西灣河", 3: "筲箕灣", 4: "柴灣", 5: "壽臣山", 6: "淺水灣", 7: "赤柱/大潭", 8: "南區/石澳", 9: "香港仔", 10: "薄扶林/海怡", 11: "貝沙灣", 12: "西半山", 13: "西區", 14: "上環/中環", 15: "中半山/金鐘", 16: "灣仔", 17: "銅鑼灣", 18: "跑馬地", 19: "天后/大坑", 20: "北角", 21: "山頂", 22: "油塘/藍田", 23: "觀塘/秀茂坪", 24: "九龍灣", 25: "鑽石山/彩虹", 26: "康城/清水灣", 27: "將軍澳", 28: "調景嶺", 29: "九龍城", 30: "土瓜灣", 31: "佐敦/尖沙咀", 32: "九龍站", 33: "紅磡/黃埔", 34: "深水埗", 35: "長沙灣", 36: "長沙灣", 37: "荔枝角", 38: "美孚", 39: "大角咀", 40: "奧運", 41: "九龍塘", 42: "石硤尾", 43: "何文田/京士柏", 44: "黃大仙/新蒲崗", 45: "太子", 46: "旺角/油麻地", 47: "西貢", 48: "馬鞍山", 49: "大埔/太和", 50: "沙田/火炭", 51: "深井", 52: "荃灣", 53: "葵涌", 54: "青衣", 55: "元朗", 56: "天水圍", 57: "屯門", 58: "粉嶺", 59: "上水", 60: "愉景灣", 61: "東涌", 62: "馬灣", 63: "大嶼山", 64: "坪洲", 65: "南丫島", 66: "長洲", 67: "大窩口"}
	Relation = map[int][]int{
		1: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
		2: {22, 23, 24, 25, 26, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 90},
		3: {46, 47, 48, 27, 49, 50, 51, 52, 53, 54, 56, 57, 58, 59, 28, 67},
		4: {60, 61, 62, 63, 64, 65, 66},
		5: {68, 69, 70, 71, 72, 84, 86, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 85, 87, 88, 89, 93, 94, 95, 96},
	}
}
