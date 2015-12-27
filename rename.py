




import os, sys



if __name__ == "__main__":
	path=r"C:\BaiduYunDownload\Node.js First Look"
	os.chdir(path)
	print os.getcwd()
	filelist = os.listdir(path)
	for fname in filelist:
		#print os.path.realpath(fname)
		
		#basename = os.path.basename(fname)
		basename, suffix = os.path.splitext(fname)
		#array = basename.strip().split("-")
		#assert len(array) == 2
		name, fileid = basename.strip().split("-")
		new_name = "%s-%s%s" % (fileid, name ,suffix)
		print new_name
		os.rename(fname, new_name)

		#print basename, suffix, "|", name, fileid

		pass