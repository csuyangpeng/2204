package idmgr

import (
	"fmt"
	"lite5gc/cmn/types"
	"testing"
)

func Test_RegisterIDMgr_01(t *testing.T) {
	//empty register
	idMgr := NewIDMgr()
	err := idMgr.RegisterIDMgr("SC", 1000)
	if err != nil {
		t.Errorf("Register failed, err: %s", err)
	}
}
func Test_RegisterIDMgr_02(t *testing.T) {
	idMgr := NewIDMgr()
	err := idMgr.RegisterIDMgr("SC", 1000)
	if err != nil {
		t.Errorf("Register failed, err: %s", err)
	}

	err = idMgr.RegisterIDMgr("SC", 1000)
	if err == nil {
		t.Errorf("Duplicated Register should return error.")
	}
	t.Logf("Duplciated Register. return: %s", err)
}

func Test_BorrowID_01(t *testing.T) {
	idMgr := NewIDMgr()
	idType := "SC"

	err := idMgr.RegisterIDMgr(idType, 1000)
	if err != nil {
		t.Errorf("Register failed, err: %s", err)
	}

	k, err := idMgr.BorrowID(idType)
	if err != nil {
		t.Errorf("Borror Id Failed: %s", err)
	}
	fmt.Println("1", k)

	err = idMgr.ReserveID(idType, 2)
	if err != nil {
		t.Errorf("reserve Id Failed: %s", err)
	}

	k, err = idMgr.BorrowID(idType)
	if err != nil {
		t.Errorf("Borror Id Failed: %s", err)
	}
	fmt.Println("2", k)

	err = idMgr.UnReserveID(idType, 1)
	if err != nil {
		t.Errorf("reserve Id Failed: %s", err)
	}

	k, err = idMgr.BorrowID(idType)
	if err != nil {
		t.Errorf("Borror Id Failed: %s", err)
	}
	fmt.Println("3", k)

	k, err = idMgr.BorrowID(idType)
	if err != nil {
		t.Errorf("Borror Id Failed: %s", err)
	}
	fmt.Println("4", k)

	err = idMgr.ReserveID(idType, 4)
	if err != nil {
		t.Errorf("reserve Id Failed: %s", err)
	}

	k, err = idMgr.BorrowID(idType)
	if err != nil {
		t.Errorf("Borror Id Failed: %s", err)
	}
	fmt.Println("5", k)

	k, err = idMgr.BorrowID(idType)
	if err != nil {
		t.Errorf("Borror Id Failed: %s", err)
	}
	fmt.Println("6", k)

	err = idMgr.UnReserveID(idType, 4)
	if err != nil {
		t.Errorf("reserve Id Failed: %s", err)
	}

	k, err = idMgr.BorrowID(idType)
	if err != nil {
		t.Errorf("Borror Id Failed: %s", err)
	}
	fmt.Println("7", k)

}

func Test_BorrowID_02(t *testing.T) {
	idMgr := NewIDMgr()
	idType := "SC"

	err := idMgr.RegisterIDMgr(idType, 1000)
	if err != nil {
		t.Errorf("Register failed, err: %s", err)
	}

	var id uint32
	for i := 0; i < 150; i++ {
		id, err = idMgr.BorrowID(idType)
		if err != nil {
			t.Errorf("The %d times, Borror Id Failed: %s", i, err)
			return
		}
	}

	var expectID uint32 = 149
	if id == expectID {
		t.Logf("Borror Succ: %d", id)
	} else {
		t.Errorf("Lastly Borrow id should be %d, return:%d", expectID, id)
	}
}

func Test_BorrowID_03(t *testing.T) {
	idMgr := NewIDMgr()
	idType := "SC"

	// err := idMgr.RegisterIDMgr(idType, 1000)
	// if err != nil {
	// 	t.Errorf("Register failed, err: %s", err)
	// }

	id, err := idMgr.BorrowID(idType)
	if err != nil {
		t.Logf("Not Registed, Borror Id Failed: %s", err)
	} else {
		t.Errorf("Not registered idType [%s], But Return id [%d]", idType, id)
	}
}

func Test_BorrowID_04(t *testing.T) {
	idMgr := NewIDMgr()
	idType := "SC"

	err := idMgr.RegisterIDMgr(idType, 1)
	if err != nil {
		t.Errorf("Register failed, err: %s", err)
	}

	id, err := idMgr.BorrowID(idType)
	if err != nil {
		t.Errorf("Not Registed, Borror Id Failed: %s", err)
	} else {
		t.Logf("Registered idType [%s], Return id [%d]", idType, id)
	}
}

func Test_ReturnID_01(t *testing.T) {
	idMgr := NewIDMgr()
	idType := "SC"

	err := idMgr.RegisterIDMgr(idType, 1000)
	if err != nil {
		t.Errorf("Register failed, err: %s", err)
	}

	var id uint32
	for i := 0; i < 100; i++ {
		id, err = idMgr.BorrowID(idType)
		if err != nil {
			t.Errorf("The %d times, Borror Id Failed: %s", i, err)
		}
	}

	for i := 20; i < 40; i++ {
		err = idMgr.ReturnID(idType, uint32(i))
	}

	for i := 0; i < 920; i++ {
		id, err = idMgr.BorrowID(idType)
		if err != nil {
			t.Errorf("The %d times, Borror Id Failed: %s", i, err)
		}
	}
	if err != nil {
		t.Errorf("Return id (%d) failed.", id)
	}
	idMgr.DumpIDList(idType)
}

func Test_ReturnID_02(t *testing.T) {
	idMgr := NewIDMgr()
	idType := "SC"

	err := idMgr.RegisterIDMgr(idType, 1000)
	if err != nil {
		t.Errorf("Register failed, err: %s", err)
	}

	var id uint32
	for i := 0; i < 100; i++ {
		id, err = idMgr.BorrowID(idType)
		if err != nil {
			t.Errorf("The %d times, Borror Id Failed: %s", i, err)
		}
	}

	err = idMgr.ReturnID(idType, id+1)
	if err == nil {
		t.Errorf("Return id(%d) + 1 should be failed.", id)
	}
}

func Test_GetIdList_01(t *testing.T) {
	idMgr := NewIDMgr()
	idType := "SC"

	err := idMgr.RegisterIDMgr(idType, 1000)
	if err != nil {
		t.Errorf("Register failed, err: %s", err)
	}
	borrowTimes := 100
	for i := 0; i < borrowTimes; i++ {
		_, err = idMgr.BorrowID(idType)
		if err != nil {
			t.Errorf("The %d times, Borror Id Failed: %s", i, err)
		}
	}

	idList, err := idMgr.GetIDList(types.ModuleName(idType))
	if len(idList) != borrowTimes {
		t.Errorf("borror id for [%d] times, But Get len of idList [%d]", borrowTimes, len(idList))
	} else {
		t.Logf("borror id for [%d] times, And Get len of idList [%d]", borrowTimes, len(idList))
	}
}

func Test_GetIdList_02(t *testing.T) {
	idMgr := NewIDMgr()
	ScType := "SC"
	N2apType := "NGAP"

	err := idMgr.RegisterIDMgr(ScType, 1000)
	if err != nil {
		t.Errorf("Register failed, err: %s", err)
	}

	err = idMgr.RegisterIDMgr(N2apType, 1000)
	if err != nil {
		t.Errorf("Register failed, err: %s", err)
	}

	borrowTimes := 100

	for i := 0; i < borrowTimes; i++ {
		_, err = idMgr.BorrowID(ScType)
		if err != nil {
			t.Errorf("The %d times, Borror Id Failed: %s", i, err)
		}
	}

	for i := 0; i < borrowTimes*2; i++ {
		_, err = idMgr.BorrowID(N2apType)
		if err != nil {
			t.Errorf("The %d times, Borror Id Failed: %s", i, err)
		}
	}

	idList, err := idMgr.GetIDList(types.ModuleName(ScType))
	if len(idList) != borrowTimes {
		t.Errorf("borror id for [%d] times, But Get len of idList [%d]", borrowTimes, len(idList))
	} else {
		t.Logf("borror id for [%d] times, And Get len of idList [%d]", borrowTimes, len(idList))
	}

	idList, err = idMgr.GetIDList(types.ModuleName(N2apType))
	if len(idList) != borrowTimes*2 {
		t.Errorf("borror id for [%d] times, But Get len of idList [%d]", borrowTimes*2, len(idList))
	} else {
		t.Logf("borror id for [%d] times, And Get len of idList [%d]", borrowTimes*2, len(idList))
	}
}

func BenchmarkBorrowReturn_01(b *testing.B) {
	b.StopTimer()
	idMgr := NewIDMgr()
	idType := "SC"
	idMgr.RegisterIDMgr(idType, uint32(b.N))

	b.Logf("N (%d)", b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := idMgr.BorrowID(idType)
		if err != nil {
			b.Errorf("The %d times, Borror Id Failed: %s", i, err)
		}
	}

	idList, err := idMgr.GetIDList(types.ModuleName(idType))
	if b.N != len(idList) {
		b.Errorf("Failed to GetIDList %s", err)
	}

	for i := 0; i < b.N; i++ {
		err = idMgr.ReturnID(idType, uint32(i))
		if err != nil {
			b.Errorf("Failed to return id %d. err : %s", i, err)
		}
	}

	for i := 0; i < b.N; i++ {
		_, err := idMgr.BorrowID(idType)
		if err != nil {
			b.Errorf("The %d times, Borror Id Failed: %s", i, err)
		}
	}
}
