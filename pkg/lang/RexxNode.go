package lang

type RexxNode struct {
	Leaf     *Rexx
	InitLeaf *Rexx
}

// NetRexx: method RexxNode(argleaf=Rexx)
//
// Differences: None
func RxNode(argleaf *Rexx) (rcvr *RexxNode) {
	rcvr = &RexxNode{}
	rcvr.Leaf = argleaf
	rcvr.InitLeaf = argleaf
	return
}

// NetRexx: method clear()
//
// Differences: None
func (rcvr *Rexx) Clear() {
	rcvr.coll = nil
	return
}

// NetRexx: method compareTo(i2=Object) returns int Signals ClassCastException
//
// Differences: This function only deals with type string.
func (rcvr *Rexx) CompareTo(i2 string) (int32, error) {
	j := RxFromString(i2)
	lt, err := rcvr.OpLt(nil, j)
	if err != nil {
		return 0, err
	} else {
		if lt {
			return -1, nil
		}
	}
	gt, err := rcvr.OpGt(nil, j)
	if err != nil {
		return 0, err
	} else {
		if gt {
			return +1, nil
		}
	}
	return 0, nil
}

// NetRexx: method containsKey(key=Object) returns boolean
//
// Differences: This function only deals with type Rexx
//
//	This function only deals with type Rexx
func (rcvr *Rexx) ContainsKey(key *Rexx) bool {
	return rcvr.TestNode(key)
}

// NetRexx: method containsValue(value=Object) returns boolean
//
// Differences: None
//
//	This function only deals with type Rexx
func (rcvr *Rexx) ContainsValue(value *Rexx) bool {
	var rv *Rexx
	var status bool // golang default bool value is false
	if rcvr.coll == nil {
		return status
	}
	for _, node := range rcvr.coll {
		rv = node.Leaf
		if rv != nil {
			if rv.OpEqS(nil, value) {
				status = true
				break
			}
		}
	}
	return status
}

// NetRexx: method copyindexed(r=Rexx) protect returns Rexx
//
// Differences: If r == nil, return rcvr.
func (rcvr *Rexx) CopyIndexed(r *Rexx) *Rexx {
	if r == nil || r.coll == nil {
		// Prevents Go Panic
		return rcvr
	}
	if rcvr.coll == nil {
		rcvr.mux.Lock()
		rcvr.coll = make(map[*Rexx]*RexxNode)
		rcvr.mux.Unlock()
	}
	r.mux.Lock()
	for rxkey, rxnode := range r.coll {
		if rxnode.Leaf == nil {
			continue
		}
		if rxnode.Leaf == rxnode.InitLeaf {
			continue
		}
		temp, err := rcvr.GetNode(rxkey)
		if err == nil {
			temp.Leaf = rxnode.Leaf
		}
	}
	r.mux.Unlock()
	return rcvr
}

// NetRexx: method equals(rhs=Object) returns boolean
//
// Differences: None
//
//	This function only deals with type Rexx
func (rcvr *Rexx) Equals(rhs *Rexx) bool {
	if rhs == nil {
		return false
	}
	return rcvr.docomparestrict(nil, rhs) == 0
}

// NetRexx: method exists(key=Rexx) returns Rexx
//
// Differences: None
func (rcvr *Rexx) Exists(key *Rexx) *Rexx {
	return RxFromBool(rcvr.TestNode(key))
}

// NetRexx: method get(key=Object) returns Object
//
// Differences: argument key cannot be nil
//
//	This function only deals with type Rexx
func (rcvr *Rexx) Get(key *Rexx) (*Rexx, error) {
	if key == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument key cannot be nil")
	}
	if rcvr.TestNode(key) {
		temp, _ := rcvr.GetNode(key)
		return temp.Leaf, nil
	}
	return nil, RxException(1, "key not found")
}

// NetRexx: method getnode(key=Rexx) returns RexxNode
//
// Differences: argument key cannot be nil
func (rcvr *Rexx) GetNode(key *Rexx) (*RexxNode, error) {
	if key == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument key cannot be nil")
	}
	if rcvr.coll == nil {
		rcvr.mux.Lock()
		rcvr.coll = make(map[*Rexx]*RexxNode)
		rcvr.mux.Unlock()
	}
	for rxkey, rxnode := range rcvr.coll {
		if rxkey.OpEqS(nil, key) {
			if rxnode.Leaf != nil {
				return rxnode, nil
			}
		}
	}
	node := RxNode(rxfromclone(rcvr))
	node.Leaf = rcvr
	rcvr.coll[key] = node
	return node, nil
}

// NetRexx: method isEmpty() returns boolean
//
// Differences: None
func (rcvr *Rexx) IsEmpty() bool {
	if rcvr.coll == nil {
		return true
	}
	return rcvr.Size(0) == 0
}

// NetRexx: method isindexed() protect returns Rexx
//
// Differences: None
func (rcvr *Rexx) IsIndexed() *Rexx {
	if rcvr.Size(1) == 0 {
		return RxFromInt8(int8(0))
	}
	return RxFromInt8(int8(1))
}

// NetRexx: method keys protect returns Enumeration
//
// Differences: None
func (rcvr *Rexx) Keys() []*Rexx {
	if rcvr.coll == nil {
		rcvr.mux.Lock()
		rcvr.coll = make(map[*Rexx]*RexxNode)
		rcvr.mux.Unlock()
	}
	start := 0
	rxkeys := make([]*Rexx, len(rcvr.coll))
	for rxkey := range rcvr.coll {
		rxkeys[start] = rxkey
		start++
	}
	return rxkeys
}

// NetRexx: method put(key=Object,value=Rexx) returns Object
//
// Differences: argument key cannot be nil
//
//	This function only deals with type Rexx
func (rcvr *Rexx) Put(key *Rexx, value *Rexx) (*Rexx, error) {
	if key == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument key cannot be nil")
	}
	returnval, err := rcvr.Get(key) // test for key
	if err != nil {
		return nil, RxException(1, "key not found")
	}
	temp, _ := rcvr.GetNode(key)
	temp.Leaf = value
	return returnval, nil
}

// NetRexx: method putAll(t=Rexx)
//
// Differences: None
func (rcvr *Rexx) PutAll(t *Rexx) (*Rexx, error) {
	if t == nil {
		return nil, RxException(1, "t cannot be nil")
	}
	return rcvr.CopyIndexed(t), nil
}

// NetRexx: method remove(key=Object) returns Object
//
// Differences: argument key cannot be nil
//
//	This function only deals with type Rexx
func (rcvr *Rexx) Remove(key *Rexx) (*Rexx, error) {
	if key == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument key cannot be nil")
	}
	returnval, err := rcvr.Get(key)
	if err != nil {
		return nil, err
	}
	temp, _ := rcvr.GetNode(key)
	temp.Leaf = nil
	return returnval, nil
}

// NetRexx: method size(check=0) protect returns int
//
// Differences: None
func (rcvr *Rexx) Size(check int32) int32 {
	var icount int32 = 0
	var index int32 = 0
	if rcvr.coll == nil {
		return 0
	}
	rxkeys := rcvr.Keys()
	rcvr.mux.Lock()
	for ; index < int32(len(rxkeys)); index++ {
		node := rcvr.coll[rxkeys[index]]
		if node.Leaf == nil {
			continue
		}
		if node.Leaf == node.InitLeaf {
			continue
		}
		icount++
		rx01, _ := RxFromInt32(check).ToBool()
		if rx01 {
			rcvr.mux.Unlock()
			return 1
		}
	}
	rcvr.mux.Unlock()
	return icount
}

// NetRexx: method testnode(key=Rexx) returns boolean
//
// Differences: None
func (rcvr *Rexx) TestNode(key *Rexx) bool {
	if rcvr.coll == nil {
		return false
	}
	for rxkey, rxnode := range rcvr.coll {
		if rxkey.OpEqS(nil, key) {
			if rxnode.Leaf == nil {
				return false
			}
			if rxnode.Leaf == rxnode.InitLeaf {
				return false
			}
			return true
		}
	}
	return false
}
