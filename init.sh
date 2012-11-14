#!/bin/bash
mtr --raw -n 8.8.8.8 -c 1 | egrep "^h" | awk ' { print $3 } ' > hops.log
#echo bill@cse.ucdavis.edu000 >> hops.log
#echo bill@cse.ucdavis.edu111 >> hops.log
#echo bill@cse.ucdavis.edu222 >> hops.log
#echo bill@cse.ucdavis.edu333 >> hops.log
#echo bill@cse.ucdavis.edu444 >> hops.log
#echo bill@cse.ucdavis.edu555 >> hops.log
#echo bill@cse.ucdavis.edu666 >> hops.log
#echo bill@cse.ucdavis.edu777 >> hops.log
#echo bill@cse.ucdavis.edu888 >> hops.log
#echo bill@cse.ucdavis.edu999 >> hops.log
#echo bill@cse.ucdavis.eduAAA >> hops.log
#echo bill@cse.ucdavis.eduBBB >> hops.log
#echo bill@cse.ucdavis.eduCCC >> hops.log
##echo bill@cse.ucdavis.eduDDD >> hops.log
#echo bill@cse.ucdavis.eduEEE >> hops.log
#echo bill@cse.ucdavis.eduFFF >> hops.log



