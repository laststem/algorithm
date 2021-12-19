class Solution {
public:
    int singleNumber(vector<int>& nums) {
        int answer = 0;
        int b = 0;
        for (auto& k : nums) {
            answer = ~b & (answer ^ k);
            b = ~answer & (b ^ k);
        }
        return answer;
    }
};
