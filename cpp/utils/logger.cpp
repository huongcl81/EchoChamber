#include<bits/stdc++.h>
using namespace std;
void log_event(string s) {
  ofstream out("/var/log/echod.log", ios::app);
  out << "[ECHOCHAMBER] " << s << endl;
  out.close();
}