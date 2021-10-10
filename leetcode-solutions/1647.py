class Solution:
    def minDeletions(self, s: str) -> int:
        c = [0]*26
        b = ord('a')
        freq = set()
        answer = 0
        for i in range(26):
            c[i] = s.count(chr(i+b))
        for v in c:
            while v in freq:
                answer += 1
                v -= 1
            if v > 0:
                freq.add(v)
        return answer
