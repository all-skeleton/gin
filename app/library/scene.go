package library

var UploadScenes = map[string]map[string]string{
	"1001": {
		"desc":  "头像",
		"value": "head_img",
	},
}

var scenes = make(map[string]map[string]string)

func init() {
	scenes["upload"] = parseScene(UploadScenes)
}

func parseScene(scene map[string]map[string]string) map[string]string {
	res := make(map[string]string, len(scene))
	for idx, v := range scene {
		res[v["value"]] = idx
	}
	return res
}

// key => value
func GetScenes(mark string) map[string]string {
	return scenes[mark]
}

// key => value
func GetSceneAll() map[string]map[string]string {
	return scenes
}

// key => value
func GetUploadSceneInfo(scene string) map[string]string {
	return UploadScenes[scene]
}