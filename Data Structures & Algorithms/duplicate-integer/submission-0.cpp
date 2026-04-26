#include <set>
#include <vector>
using namespace std;

class Solution {
public:
    bool hasDuplicate(vector<int>& nums) {
    set<int> numset;
        for (int x : nums) {
            numset.insert(x);
        }
        return numset.size() != nums.size();
    }
};