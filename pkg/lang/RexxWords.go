package lang

type rexxwords struct {
}

func rxwords() (rcvr *rexxwords) {
	rcvr = &rexxwords{}
	return
}

func abbrev(a []rune, b []rune, leng int32) bool {
	if int32(len(a)) < leng || int32(len(b)) < leng {
		return false
	}
	if int32(len(b)) > int32(len(a)) {
		return false
	}
	if int32(len(b)) == 0 && leng == 0 {
		return true
	}
	_1 := int32(len(b)) - 1
	var i int32 = 0
	for ; i <= _1; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func centre(s []rune, wid int32, pad rune) []rune {
	if wid < 0 {
		// Prevents Go Panic
		return make([]rune, 0)
	}
	chop := int32(len(s)) - wid
	if chop == 0 {
		return s
	}
	ret := make([]rune, wid)
	if chop > 0 {
		copy(ret, s[chop/2:])
	} else {
		gap := int(-chop) / 2
		_2 := gap - 1
		i := 0
		for ; i <= _2; i++ {
			ret[i] = pad
		}
		copy(ret[gap:], s)
		_3 := len(ret) - 1
		i = gap + len(s)
		for ; i <= _3; i++ {
			ret[i] = pad
		}
	}
	return ret
}

func changestr(needle []rune, haystack []rune, _new []rune) []rune {
	var p int32 = 0
	var i int32 = 0
	reps := countstr(needle, haystack)
	newlen := int32(len(haystack)) + reps*(int32(len(_new))-int32(len(needle)))
	res := make([]rune, newlen)
	var oin int32 = 0
	var out int32 = 0
	_4 := reps
	for ; _4 > 0; _4-- {
		p = pos(needle, haystack, oin+1)
		_5 := p - 2
		i = oin
		for ; i <= _5; i++ {
			res[out] = haystack[i]
			out++
		}
		_6 := int32(len(_new)) - 1
		var i int32 = 0
		for ; i <= _6; i++ {
			res[out] = _new[i]
			out++
		}
		oin = p + int32(len(needle)) - 1
	}
	_7 := int32(len(haystack)) - 1
	i = oin
	for ; i <= _7; i++ {
		res[out] = haystack[i]
		out++
	}
	return res
}

func compare(a []rune, b []rune, pad rune) int32 {
	maxlen := int32(len(a))
	if int32(len(b)) > maxlen {
		maxlen = int32(len(b))
	}
	_8 := maxlen
	var i int32 = 1
	for ; i <= _8; i++ {
		if i > int32(len(a)) {
			if b[i-1] != pad {
				return i
			}
		} else if i > int32(len(b)) {
			if a[i-1] != pad {
				return i
			}
		} else {
			if a[i-1] != b[i-1] {
				return i
			}
		}
	}
	return 0
}

func countstr(needle []rune, haystack []rune) int32 {
	var count int32 = 0
	p := pos(needle, haystack, 1)
	for {
		if !(p > 0) {
			break
		}
		count++
		p = pos(needle, haystack, p+int32(len(needle)))
	}
	return count
}

func delstr(s []rune, start int32, leng int32) []rune {
	if start <= 0 {
		return s
	}
	if leng < 0 {
		return s
	}
	if start > int32(len(s)) {
		return s
	}
	fin := start + leng
	if fin > int32(len(s)) {
		fin = int32(len(s)) + 1
	}
	need := start + int32(len(s)) - fin
	res := make([]rune, need)
	if start > 1 {
		copy(res, s)
	}
	if fin <= int32(len(s)) {
		copy(res[start-1:], s[fin-1:])
	}
	return res
}

func delword(s []rune, num int32, leng int32) []rune {
	if num <= 0 {
		num = 1
	}
	if leng < 0 {
		leng = 0
	}
	start := wordindex(s, num)
	if start == 0 {
		return s
	}
	fin := wordindex(s, num+leng)
	if fin == 0 {
		fin = int32(len(s)) + 1
	}
	need := start + int32(len(s)) - fin
	res := make([]rune, need)
	if start > 1 {
		copy(res, s)
	}
	if fin <= int32(len(s)) {
		copy(res[start-1:], s[fin-1:])
	}
	return res
}

func insert(chars []rune, newchars []rune, num int32, leng int32, padchar rune) []rune {
	if num < 0 {
		num = 0
	}
	if leng < 0 {
		leng = int32(len(newchars))
	}
	var i int32 = 0
	reslen := num + leng
	if num < int32(len(chars)) {
		reslen = reslen + int32(len(chars)) - num
	}
	res := make([]rune, reslen)
	if num > 0 {
		if num <= int32(len(chars)) {
			copy(res, chars[0:num])
		} else {
			copy(res, chars)
			_9 := num - 1
			i = int32(len(chars))
			for ; i <= _9; i++ {
				res[i] = padchar
			}
		}
	}
	if leng > 0 {
		if leng <= int32(len(newchars)) {
			copy(res[num:], newchars[0:leng])
		} else {
			copy(res[num:], newchars)
			_10 := leng - 1
			i = int32(len(newchars))
			for ; i <= _10; i++ {
				res[num+i] = padchar
			}
		}
	}
	if num < int32(len(chars)) {
		copy(res[num+leng:], chars[num:])
	}
	return res
}

func nextblank(s []rune, start int32) int32 {
	if len(s) != 0 {
		if start <= int32(len(s)) && start > 0 {
			_13 := int32(len(s))
			i := start
			for ; i <= _13; i++ {
				if s[i-1] == ' ' {
					return i
				}
			}
		}
	}
	return 0
}

func nextnonblank(s []rune, start int32) int32 {
	if len(s) != 0 {
		if start <= int32(len(s)) && start > 0 {
			_14 := int32(len(s))
			i := start
			for ; i <= _14; i++ {
				if s[i-1] != ' ' {
					return i
				}
			}
		}
	}
	return 0
}

func overlay(chars []rune, newchars []rune, num int32, leng int32, padchar rune) []rune {
	if num < 1 {
		num = 1
	}
	if leng < 0 {
		leng = int32(len(newchars))
	}
	var i int32 = 0
	reslen := num + leng - 1
	if reslen < int32(len(chars)) {
		reslen = int32(len(chars))
	}
	res := make([]rune, reslen)
	if num > 1 {
		if num-1 <= int32(len(chars)) {
			copy(res, chars[0:num-1])
		} else {
			copy(res, chars)
			_11 := num - 2
			i = int32(len(chars))
			for ; i <= _11; i++ {
				res[i] = padchar
			}
		}
	}
	if leng > 0 {
		if leng <= int32(len(newchars)) {
			copy(res[num-1:], newchars[0:leng])
		} else {
			copy(res[num-1:], newchars)
			_12 := leng - 1
			i = int32(len(newchars))
			for ; i <= _12; i++ {
				res[num-1+i] = padchar
			}
		}
	}
	if num+leng-1 < int32(len(chars)) {
		copy(res[num+leng-1:], chars[num+leng-1:])
	}
	return res
}

func pos(needle []rune, haystack []rune, start int32) int32 {
	if start < 1 {
		// Prevents Go Panic
		return 0
	}
	if needle == nil || haystack == nil {
		return 0
	}
	if len(needle) == 0 {
		return 0
	}
	_16 := int32(len(haystack)) - int32(len(needle)) + 1
	i := start
i:
	for ; i <= _16; i++ {
		_17 := int32(len(needle)) - 1
		var j int32 = 0
		for ; j <= _17; j++ {
			if needle[j] != haystack[i+j-1] {
				continue i
			}
		}
		return i
	}
	return 0
}

func space(s []rune, gap int32, pad rune) []rune {
	if gap < 0 {
		// Prevents Go Panic
		gap = 1
	}
	var start int32 = 1
	var count int32 = 0
	var nonspaces int32 = 0
	for {
		start = nextnonblank(s, start)
		if start == 0 {
			break
		}
		count++
		nextblank := nextblank(s, start+1)
		if nextblank == 0 {
			nonspaces = nonspaces + int32(len(s)) + 1 - start
			break
		}
		nonspaces = nonspaces + nextblank - start
		start = nextblank
	}
	if count == 0 {
		return make([]rune, 0)
	}
	leng := (count-1)*gap + nonspaces
	res := make([]rune, leng)
	start = 1
	var out int32 = 0
	_19 := count
	var c int32 = 1
c:
	for ; c <= _19; c++ {
		start = nextnonblank(s, start)
		if start == 0 {
			break c
		}
	_20:
		for {
			res[out] = s[start-1]
			start++
			if start > int32(len(s)) {
				break c
			}
			out++
			if s[start-1] == ' ' {
				break _20
			}
		}
		if c == count {
			break c
		}
		_21 := gap
		for ; _21 > 0; _21-- {
			res[out] = pad
			out++
		}
	}
	return res
}

func subword(s []rune, num int32, leng int32) []rune {
	if leng == 0 {
		return make([]rune, 0)
	}
	start := wordindex(s, num)
	if start == 0 {
		return make([]rune, 0)
	}
	fin := wordindex(s, num+leng)
	if fin == 0 {
		fin = int32(len(s)) + 1
	}
	_22 := start
	fin = fin - 1
	for ; fin >= _22; fin-- {
		if s[fin-1] != ' ' {
			break
		}
	}
	need := (fin - start) + 1
	res := make([]rune, need)
	copy(res, s[start-1:])
	return res
}

func verify(s []rune, v []rune, opt rune, start int32) int32 {
	if opt == 'N' {
		return verifyn(s, v, start)
	}
	return verifym(s, v, start)
}

func verifym(s []rune, v []rune, start int32) int32 {
	if start <= 0 {
		// Prevents Go Panic
		start = 1
	}
	_try := rune(0)
	var t int32 = 0
	last := int32(len(s))
	if start > last {
		return 0
	}
	if len(v) == 0 {
		return 0
	}
	_23 := last
	i := start
	for ; i <= _23; i++ {
		_try = s[i-1]
		if _try == v[0] {
			return i
		}
		t = pos(ToRunesFromRune(_try), v, 2)
		if t > 0 {
			return i
		}
	}
	return 0
}

func verifyn(s []rune, v []rune, start int32) int32 {
	var t int32 = 0
	last := int32(len(s))
	if start > last {
		return 0
	}
	if len(v) == 0 {
		return start
	}
	_24 := last
	i := start
	for ; i <= _24; i++ {
		t = pos(ToRunesFromRune(s[i-1]), v, 1)
		if t == 0 {
			return i
		}
	}
	return 0
}

func word(s []rune, num int32) []rune {
	return subword(s, num, 1)
}

func wordindex(s []rune, num int32) int32 {
	var start int32 = 1
	var count int32
	for count = 1; ; count++ {
		start = nextnonblank(s, start)
		if start == 0 {
			return 0
		}
		if count == num {
			break
		}
		start = nextblank(s, start+1)
		if start == 0 {
			return 0
		}
	}
	return start
}

func wordlength(s []rune, num int32) int32 {
	start := wordindex(s, num)
	if start == 0 {
		return 0
	}
	fin := nextblank(s, start+1)
	if fin == 0 {
		fin = int32(len(s)) + 1
	}
	return fin - start
}

func wordpos(needle []rune, haystack []rune, wpos int32) int32 {
	nlen := int32(len(needle))
	if nlen == 0 {
		return 0
	}
	nbeg := nextnonblank(needle, 1)
	if nbeg == 0 {
		return 0
	}
	hbeg := wordindex(haystack, wpos)
	if hbeg == 0 {
		return 0
	}
	hlen := int32(len(haystack))
_26:
	for {
		nb := nbeg
		hb := hbeg
	compare:
		for {
			nend := nextblank(needle, nb+1)
			if nend == 0 {
				nend = nlen + 1
			}
			hend := nextblank(haystack, hb+1)
			if hend == 0 {
				hend = hlen + 1
			}
			if hend-hb != nend-nb {
				break compare
			}
			h := hb - 1
			_25 := nend - 2
			n := nb - 1
			for ; n <= _25; n++ {
				if needle[n] != haystack[h] {
					break compare
				}
				h++
			}
			if nend > nlen {
				return wpos
			}
			nb = nextnonblank(needle, nend)
			if nb == 0 {
				return wpos
			}
			if hend > hlen {
				break compare
			}
			hb = nextnonblank(haystack, hend)
			if hb == 0 {
				break compare
			}
		}
		wpos++
		hbeg = nextblank(haystack, hbeg+1)
		if hbeg == 0 {
			break _26
		}
		hbeg = nextnonblank(haystack, hbeg+1)
		if hbeg == 0 {
			break _26
		}
	}
	return 0
}

func words(s []rune) int32 {
	var start int32 = 1
	var count int32 = 0
	for {
		start = nextnonblank(s, start)
		if start == 0 {
			break
		}
		count++
		start = nextblank(s, start+1)
		if start == 0 {
			break
		}
	}
	return count
}
