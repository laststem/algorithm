class Solution {
public:
    string removeOuterParentheses(string S) {
        int st = 0;
        string answer;
        int cnt = 0;
        for (int i=0; i<S.length(); i++) {
            char c = S[i];
            if (c == '(') {
                if (st == 0) {
                    st++;
                } else {
                    answer.push_back(c);
                    cnt++;
                }
            } else {
                if (cnt == 0) {
                    st--;
                } else {
                    answer.push_back(c);
                    cnt--;
                }
            }
        }
        return answer;
    }
};
