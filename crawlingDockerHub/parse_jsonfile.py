

import os, sys, json



if __name__ == "__main__":
	jsonfile=sys.argv[1]
	print jsonfile

	image_info = json.loads(open(jsonfile,"r").read())
	print image_info["name"]
	print image_info["namespace"]
	print image_info["tags"]

	image_url = "{}/{}".format(image_info["namespace"], image_info["name"])
	for tag in image_info["tags"]:
		image_tag_url = "{}:{}".format(image_url, tag.strip())
		print image_tag_url
		#pull from docker hub
		os.system("docker pull {}".format(image_tag_url))
		#tag for 10.253.20.65
		os.system("docker tag {} {}/{}".format(image_tag_url, "10.253.20.65", image_tag_url))
		#push to 10.253.20.65
		os.system("docker push {}/{}".format("10.253.20.65", image_tag_url))
