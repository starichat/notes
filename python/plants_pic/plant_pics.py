
'''
	前提：维护一个词典数据csv文件
 1.自动化工具实现输入词条，跳转搜索页面
 2.分析搜索页面信息，拿到图片url
 3.根据url下载图片
'''

from multiprocessing.dummy import Pool

"""根据搜索词下载百度图片"""
import re
import urllib
import csv
import requests


# 根据植物条目csv文件获取关键词列表
def get_keywords():
	# 读取csv至列表
	csvFile = open("plants.csv",encoding='utf-8',mode="r")
	reader = csv.reader(csvFile)
	# 建立空列表
	result = []
	for item in reader:
		# 忽略第一行
		if reader.line_num == 1:
			continue
		result.append(item[0])
	
	csvFile.close()
	return result
	

# 获取keyword，第 page 页的 url
def getPage(keyword, page, n):
	page = page * n
	keyword = urllib.parse.quote(keyword, safe='/')
	url_begin = "http://image.baidu.com/search/flip?tn=baiduimage&ie=utf-8&word="
	url = url_begin + keyword + "&pn=" + str(page) + "&gsm=" + str(hex(page)) + "&ct=&ic=0&lm=-1&width=0&height=0"
	return url


def get_onepage_urls(onepageurl):
	try:
		html = requests.get(onepageurl).text
	except Exception as e:
		print(e)
		pic_urls = []
		return pic_urls
	pic_urls = re.findall('"objURL":"(.*?)",', html, re.S)
	return pic_urls


def down_pic(pic_urls,keyword):
	"""给出图片链接列表, 下载所有图片"""
	# for i, pic_url in enumerate(pic_urls):
	# 	try:
	# 		pic = requests.get(pic_url, timeout=15)
	# 		string = keyword + str(i+1)+ '.jpg'
	# 		with open(string, 'wb') as f:
	# 			f.write(pic.content)
	# 			print('成功下载第%s张图片: %s' % (str(i + 1), str(pic_url)))
	# 	except Exception as e:
	# 		print('下载第%s张图片时失败: %s' % (str(i + 1), str(pic_url)))
	# 		print(e)
	# 		continue
	#
	
	# 下载一张搜索结果第一张图片
	try:
		pic = requests.get(pic_urls[0], timeout=15)
		string = keyword + '.jpg'
		with open(string, 'wb') as f:
			f.write(pic.content)
			print('成功下载图片: %s' % str(pic_urls[0]))
	except Exception as e:
		print('下载图片时失败: %s'  % str(pic_urls[0]))
		print(e)
	
def main(keyword):
	page_begin = 0
	page_number = 1
	image_number = 1
	all_pic_urls = []
	while 1:
		if page_begin > image_number:
			break
		print("第%d次请求数据", [page_begin])
		url = getPage(keyword, page_begin, page_number)
		onepage_urls = get_onepage_urls(url)
		page_begin += 1
		
		all_pic_urls.extend(onepage_urls)
	
	down_pic(list(set(all_pic_urls)), keyword)


if __name__ == '__main__':
	# 利用多进程加快下载速度，由于自己电脑是 4 核，开 4 个进程
	keywords = get_keywords()
	pool = Pool()
	groups = ([x * 4 for x in keywords])
	pool.map(main, groups)
	pool.close()
	pool.join()
	
	
	

	
