

for jsonfile in `ls json`
do
	echo "======================================================="
	echo $jsonfile
	#image_name = echo $jsonfile | awk -F"." '{print $1}'
	python parse_jsonfile.py "json/$jsonfile"
done