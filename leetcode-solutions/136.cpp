class Solution {
public:
    int singleNumber(vector<int>& nums) {
        int answer = 0;
        for (auto& k : nums) {
            answer ^= k;
        }
        return answer;
    }
};
