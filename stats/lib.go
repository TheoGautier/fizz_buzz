package stats

func GetHighestHits(hitsStats map[string]int) (string, int) {
	mostHits := 0
	mostUsedParams := ""
	if hitsStats == nil {
		return "", 0
	}
	for params, hits := range hitsStats {
		if hits > mostHits {
			mostHits = hits
			mostUsedParams = params
		}
	}
	return mostUsedParams, mostHits
}
