#include <iostream>
#include <vector>

using namespace std;

int main() {

  long long n;
  cin >> n;

  int i = 1;

  while (i <= n) {
    if (i + 1 <= n) {
      cout << i + 1 << "\n";
      cout << i << "\n";
      i += 2;
    } else {
      cout << i << "\n";
      i += 1;
    }
  }
}
