#include <fstream>
#include <iostream>
#include <string>

using namespace std;


int main() {
  ifstream in("input.txt");
  if (!in) {
    cerr << "Failed to open input.txt\n";
  }

  constexpr int start = 0;
  constexpr int end = 99;

  string line;
  int dial = 50;
  int password = 0;
  int value = 0;
  int rotation = 0;

  while (in >> line) {
    int sign = int(line[0]);
    int len = (line.substr(1, line.length() - 1)).length();

    if (len == 3) {
      value = stoi(line.substr(2, line.length() - 1));
    } else {
      value = stoi(line.substr(1, line.length() - 1));
    }
    
    
    if (sign == 82) {
      cout << "Rotation dir to right: " << value << "\n"; 
      dial += value;
      if (dial >= 100) dial -= 100; 
    }

    if (sign == 76) {
      cout << "Rotation dir to left: " << value << "\n";
      dial -= value;
      if (dial < 0) dial += 100;
    }


    if (dial == 0) {
      password++;
    }

    cout << "Dial: " << dial << "\n";
  } 

  cout << "Password is: " << password << "\n";

  return 0;
}
