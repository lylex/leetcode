#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
  void moveZeroes(vector<int>& nums) {
    int hollow = 0, current;
    while(nums[hollow]) {
        ++hollow;
    }

    current = hollow + 1;

    while(current < nums.size()) {
      if(!nums[current]) {
        ++current;
        continue;
      }
      nums[hollow++] = nums[current++];
    }

    while(hollow < nums.size()) {
      nums[hollow++] = 0;
    }
  }
};

int main() {
  Solution s;
  vector<int> a{0,1,0,3,12};
  s.moveZeroes(a);

  cout << "--zero-------" << endl;
  for (int i = 0; i < a.size(); ++i) {
    cout << a[i] << endl;
  }
}
