func reverseWords(s string) string {
    out := trim(s)
    reverse(out, 0, len(out)-1)
    lo := 0
    for i := 0; i < len(out); i++ {
        if out[i] == ' ' {
            reverse(out, lo, i-1) 
            lo = i+1
        }
    }
    reverse(out, lo, len(out)-1)
    return string(out)
}

func trim(s string) []byte {
	out := make([]byte, 0)
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	
	for i < len(s) {
		if s[i] != ' ' {
			out = append(out, s[i])
			i++
			continue
		}
		
		for i < len(s) {
			if i+1 < len(s) && s[i+1] != ' ' {
				out = append(out, s[i])
				i++
				break
			}
			i++
		}
	}
	return out
}

func reverse(out []byte, lo, hi int) []byte {
    for lo < hi {
        out[lo], out[hi] = out[hi], out[lo]
        lo++
        hi--
    }
    return out
}

