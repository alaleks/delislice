package delislice

import "testing"

func TestDeleteItemSlice(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 4, 5}
	got, _ := DelItemInd(2, sl)
	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Errorf("got should be %v, but got = %v", want, got)
		}
	}
}

func TestDeleteItemSlice2(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5}
	_, err := DelItemInd(5, sl)
	if err == nil {
		t.Errorf("should be an error")
	}

}

func TestDeleteItemSliceVal(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 5, 5}
	want := []int{1, 2, 3, 4}
	got, _ := DelItemVal(5, sl)
	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Errorf("got should be %v, but got = %v", want, got)
		}
	}
}

func TestDeleteItemSliceVal2(t *testing.T) {
	sl := []string{"1", "2", "3", "4", "5", "5", "6", "5"}
	want := []string{"1", "2", "3", "4", "6"}
	got, _ := DelItemVal("5", sl)
	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Errorf("got should be %v, but got = %v", want, got)
		}
	}
}
