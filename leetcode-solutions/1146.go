type SnapshotArray struct {
    count int
    snapshot map[int]map[int]int
}


func Constructor(length int) SnapshotArray {
    snapshotArray := SnapshotArray{}
    snapshotArray.snapshot = make(map[int]map[int]int, 0)
    return snapshotArray
}


func (this *SnapshotArray) Set(index int, val int) {
    if this.snapshot[this.count] == nil {
        this.snapshot[this.count] = map[int]int{}
    }
    this.snapshot[this.count][index] = val
}


func (this *SnapshotArray) Snap() int {
    k := this.count
    this.count += 1
    return k
}


func (this *SnapshotArray) Get(index int, snap_id int) int {
    for snap_id >= 0 {
        if v, ok := this.snapshot[snap_id]; ok {
            if v2, ok := v[index]; ok {
                return v2
            }
        }
        snap_id--;
    }
    return 0
}


/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
