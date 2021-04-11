class Solution {
    public int lengthOfLongestSubstring(String s) {
        int answer = 0;
        int len = s.length();
        Map<Character, Integer> m = new HashMap<>();
        for (int i=0,j=0; j<len; j++) {
            char x = s.charAt(j);
            Integer k = m.get(x);
            if (k != null) {
                i = Math.max(k, i);
            }
            answer = Math.max(answer, j - i + 1);
            m.put(x, j+1);
        }
        return answer;
    }
}
