# -*- coding: utf-8 -*-

import re
import requests

HEADERS = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36'
}

def test():
    base_url = 'http://news.cyol.com/gb/channels/vrGlAKDl/index.html'
    # base_url = 'http://m.cyol.com/gb/channels/vrGlAKDl/index.html'
    pattern = re.compile('<li>.*<a href="(http.*?\.html)"')
    resp = requests.get(base_url, headers=HEADERS)
    assert resp.status_code == 200
    res = re.findall(pattern, resp.content.decode('utf-8'))
    assert len(res) > 0
    dxx_img = res[0].replace('m.html', 'images/end.jpg').replace('index.html', 'images/end.jpg')
    pattern = re.compile('<title>(.*?)</title>')
    resp = requests.get(res[0], headers=HEADERS)
    assert resp.status_code == 200
    res = re.findall(pattern, resp.content.decode('utf-8'))
    assert len(res) > 0
    return {
        'title': res[0],
        'dxx_img': dxx_img
    }

if __name__ == '__main__':
    print(test())