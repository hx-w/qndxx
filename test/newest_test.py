# -*- coding: utf-8 -*-

import re
import requests


def test():
    base_url = 'http://news.cyol.com/gb/channels/vrGlAKDl/index.html'
    pattern = re.compile('<li>.*<a href="(https.*?\.html)"')
    resp = requests.get(base_url)
    assert resp.status_code == 200
    res = re.findall(pattern, resp.content.decode('utf-8'))
    assert len(res) > 0
    dxx_img = res[0].replace('m.html', 'images/end.jpg').replace('index.html', 'images/end.jpg')
    pattern = re.compile('<title>(.*?)</title>')
    resp = requests.get(res[0])
    assert resp.status_code == 200
    res = re.findall(pattern, resp.content.decode('utf-8'))
    assert len(res) > 0
    return {
        'title': res[0],
        'dxx_img': dxx_img
    }

if __name__ == '__main__':
    print(test())