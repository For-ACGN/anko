#!anko

r = -1
if (false) {
  r = 1
} else if (false) {
  r = 2
} else if (false) {
  r = 3
} else {
  r = 4
}
is(4, r, "switch/default")

# vim: set ft=anko:
