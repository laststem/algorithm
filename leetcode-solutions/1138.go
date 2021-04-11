type Position struct {
    x int
    y int
}

func (p *Position) Diff(other Position) string {
    var s string
    for p.x != other.x || p.y != other.y {
        if p.x > other.x {
            s += "U"
            p.x--
        } else if p.x < other.x {
            if p.x + 1 == 5 && p.y != 0 {
                s += "L"
                p.y--
                continue
            }
            s += "D"
            p.x++
        } else if p.y > other.y {
            s += "L"
            p.y--
        } else {
            s += "R"
            p.y++
        }
    }
    return s + "!"
}

func alphabetBoardPath(target string) string {
    s := "abcdefghijklmnopqrstuvwxyz"
    m := make(map[rune]Position, 0)
    for idx, c := range s {
        m[c] = Position{idx / 5, idx % 5}
    }
    var answer string
    current := &Position{0, 0}
    for _, c := range target {
        answer += current.Diff(m[c])
    }
    return answer
}
