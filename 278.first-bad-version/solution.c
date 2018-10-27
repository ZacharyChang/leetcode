// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

int firstBadVersion(int n) {
  int start = 1;
  int end = n;
  while (start < end) {
    int mid = start + (end - start) / 2;
    if (isBadVersion(mid)) {
      end = mid;
    } else {
      start = mid + 1;
    }
  }
  if (isBadVersion(start)) {
    return start;
  }
  return end;
}