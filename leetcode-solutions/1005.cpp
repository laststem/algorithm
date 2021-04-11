class Solution {
public:
    int largestSumAfterKNegations(vector<int>& A, int K) {
        priority_queue <int,vector<int>, greater<int>> pq;
        int answer = 0;
        for(auto w : A) {
            pq.push(w);
            answer += w;
        }
        while(K--) {
            int x = pq.top();
            answer -= x * 2;
            pq.pop();
            pq.push(-x);
        }
        return answer;
    }
};
