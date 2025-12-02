#include <bits/stdc++.h>
using namespace std;

bool invalid(long long x, bool p2) {
    string s = to_string(x);
    int n = (int)s.size();
    int kStart = 2;
    int kEnd = p2 ? (n + 1) : 3; // k in [2, len+1) when p2, else [2,3)

    for (int k = kStart; k < kEnd; ++k) {
        if (n % k == 0) {
            int sz = n / k;
            bool ok = true;
            for (int i = 0; i < n; i += sz) {
                if (s.compare(i, sz, s, 0, sz) != 0) {
                    ok = false;
                    break;
                }
            }
            if (ok) return true;
        }
    }
    return false;
}

int main() {
    string line;

    cin >> line;
    
    long long p1 = 0, p2 = 0;

    size_t start = 0;
    while (start < line.size()) {
        size_t comma = line.find(',', start);
        string range = (comma == string::npos) ? line.substr(start) : line.substr(start, comma - start);

        size_t dash = range.find('-');
        if (dash != string::npos) {
            long long first = stoll(range.substr(0, dash));
            long long last  = stoll(range.substr(dash + 1));
            if (first > last) swap(first, last);
            for (long long x = first; x <= last; ++x) {
                if (invalid(x, false)) p1 += x;
                if (invalid(x, true))  p2 += x;
            }
        }

        if (comma == string::npos) break;
        start = comma + 1;
    }

    cout << p1 << "\n" << p2 << "\n";
    return 0;
}
