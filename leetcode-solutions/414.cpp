class Solution {
public:
    int thirdMax(vector<int>& nums) {
        int a[3] = {0, };
        
        int mn = nums[0];
        for (int i=0; i<nums.size(); i++) {
            if (mn > nums[i]) {
                mn = nums[i];
            }
        }
        
        int mx_idx = 0;
        a[0] = mn;
        for (int i=0; i<nums.size(); i++) {
            int k = nums[i];
            if (a[0] < k) {
                a[0] = k;
                mx_idx = i;
            }
        }
        
        int mx2_idx = 0;
        a[1] = mn;
        for (int i=0; i<nums.size(); i++) {
            if (i == mx_idx) {
                continue;
            }
            int k = nums[i];
            if (a[1] < k && a[0] > k) {
                a[1] = k;
                mx2_idx = i;
            }
        }
        
        int mx3 = mn;
        a[2] = mn;
        for (int i=0; i<nums.size(); i++) {
            if (i == mx_idx || i == mx2_idx) {
                continue;
            }
            int k = nums[i];
            if (a[2] < k && a[0] > k && a[1] > k) {
                a[2] = k;
            }
        }
        if (a[0] == a[1] || a[1] == a[2]) {
            return a[0];
        }
        return a[2];
    }
};
