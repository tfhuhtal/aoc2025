#include <fstream>
#include <iostream>
#include <string>

using namespace std;

int main() {
  ifstream in("input.txt");
  if (!in) {
    cerr << "Failed to open input.txt" << endl;
  }

  string line;
  int dial = 50;
  int p1 = 0;
  int p2 = 0;

  while (in >> line) {
    char d = line[0];
    int amt = stoi(line.substr(1));

    for (int i = 0; i < amt; i++) {
      if (d == 'L') {
        dial = (dial - 1 + 100) % 100;
      } else {
        dial = (dial + 1) % 100;
      }
      if (dial == 0) {
        p2++;
      }
    }
    if (dial == 0) {
          p1++;
    }
  }

  cout << "part 1 password: " << p1 << endl;
  cout << "part 2 password: " << p2 << endl;

  return 0;
}
