#!/bin/sh
 
echo '<html><body>'
for i in $(ls -1 | grep -v gen.sh) 
do
echo  "<a href=$i>$i</a>\c"
echo '<br>'
done
echo '</body></html>'

