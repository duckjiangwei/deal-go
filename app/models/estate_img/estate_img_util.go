package estateimg

import (
	"hk591_go/pkg/config"
	"strings"
)

//图片格式化
func PathFormat(path string) map[string]interface{} {
	img := make(map[string]interface{})
	if path != "" {
		file := strings.Split(path, "/estate")
		file = strings.Split(file[1], "src_")
		url := config.Get("domain.p1_url")
		fileName := strings.Split(file[1], ".")
		img["list"] = url + "/estate/crop" + file[0] + fileName[0] + "_212x160.jpg"
		img["small"] = url + "/estate/crop" + file[0] + fileName[0] + "_118x88.jpg"
		img["middle"] = url + "/estate/crop" + file[0] + fileName[0] + "_730x460.jpg"
		img["large"] = url + "/estate/crop" + file[0] + fileName[0] + "_600x600.jpg"
	}
	return img
}
