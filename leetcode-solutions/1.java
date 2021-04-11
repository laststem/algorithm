class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> m = new HashMap<>();
        for(int i=0; i<nums.length; i++) {
            m.put(nums[i], i);
        }
        for(int i=0; i<nums.length; i++) {
            Integer k = m.get(target-nums[i]);
            if (k != null && k != i) {
                return new int[]{i, k};
            }
        }
        return null;
    }
}
