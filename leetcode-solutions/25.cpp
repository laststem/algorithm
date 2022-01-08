struct ListNode {
    int val;
    ListNode *next;

    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

ListNode* reverseKGroup(ListNode* head, int k) {
    if (k == 1) {
        return head;
    }
    k--;
    ListNode* answer = nullptr;
    ListNode* cur = head;
    ListNode* p = nullptr;
    while(true) {
        ListNode* first = cur;
        ListNode* last = cur;
        bool ok = true;
        for (int i=0; i<k && ok; i++) {
            if (last != nullptr) {
                last = last->next;
            } else {
                ok = false;
            }
        }
        if (!ok || !last) {
            break;
        }
        ListNode* f = last;
        last = last->next;
        for (int i=0; i<=k; i++) {
            ListNode* nextFirst = first->next;
            first->next = last;
            last = first;
            first = nextFirst;
        }
        if (p) {
            p->next = f;
        }
        p = cur;
        cur = first;
        if (!answer) {
            answer = last;
        }
    }
    return answer;
}
