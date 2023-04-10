package stats

import "testing"

func TestGetHighestHits(t *testing.T) {
	input := map[string]int{
		"int1=2&int2=3": 0,
		"int1=3&int2=5": 2,
		"int1=9&int2=1": 1,
	}
	wantParams := "int1=3&int2=5"
	wantHits := 2
	params, hits := GetHighestHits(input)
	if params != wantParams {
		t.Errorf("Error getting highest hits, params %s are not as wanted %s", params, wantParams)
		return
	} else if hits != wantHits {
		t.Errorf("Error getting highest hits, count %d is not as wanted %d", hits, wantHits)
		return
	}
}
