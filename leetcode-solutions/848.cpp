class Solution {
public:
    string shiftingLetters(string S, vector<int>& shifts) {
        if (shifts.size() == 0) {
            return S;
        }
        vector<int> v(shifts.size());
        v[shifts.size()-1] = shifts[shifts.size()-1];
        for (int i=shifts.size()-2; i>=0; i--) {
            v[i] = (v[i+1] + shifts[i]) % 26;
        }
        for (int i=S.length()-1; i>=0; i--) {
            S[i] = ((S[i] - 'a' + v[i]) % 26) + 'a';
        }
        return S;
    }
};
