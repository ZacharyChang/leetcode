import time
from concurrent import futures

import requests
from requests_toolbelt import MultipartEncoder

from hack.readme.login.login_pb2 import *
from hack.readme.login.login_pb2_grpc import *

session = requests.Session()
user_agent = r'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 ' \
             r'Safari/537.36 '


def login(username, password):
    url = 'https://leetcode.com'
    cookies = session.get(url).cookies
    for cookie in cookies:
        if cookie.name == 'csrftoken':
            csrftoken = cookie.value

    url = "https://leetcode.com/accounts/login"

    params_data = {
        'csrfmiddlewaretoken': csrftoken,
        'login': username,
        'password': password,
        'next': 'problems'
    }
    headers = {'User-Agent': user_agent, 'Connection': 'keep-alive', 'Referer': 'https://leetcode.com/accounts/login/',
               "origin": "https://leetcode.com"}
    m = MultipartEncoder(params_data)

    headers['Content-Type'] = m.content_type
    session.post(url, headers=headers, data=m, timeout=10, allow_redirects=False)
    is_login = session.cookies.get('LEETCODE_SESSION') != None
    return is_login


def get_problems():
    url = "https://leetcode.com/api/problems/all/"

    headers = {'User-Agent': user_agent, 'Connection': 'keep-alive'}
    resp = session.get(url, headers=headers, timeout=10)

    return resp.content.decode('utf-8')


_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class LeetcodeServicer(LeetcodeServiceServicer):
    def Login(self, request, context):
        login(request.name, request.password)
        return LoginResponse(result=get_problems())


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_LeetcodeServiceServicer_to_server(LeetcodeServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print('server started')
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
