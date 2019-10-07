import json
import time
from concurrent import futures

import requests
from requests_toolbelt import MultipartEncoder

from hack.readme.leetcode.leetcode_pb2 import *
from hack.readme.leetcode.leetcode_pb2_grpc import *

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
    return session.cookies.get('LEETCODE_SESSION') is not None


def get_problems():
    url = "https://leetcode.com/api/problems/all/"

    headers = {'User-Agent': user_agent, 'Connection': 'keep-alive'}
    resp = session.get(url, headers=headers, timeout=10)

    return resp.content.decode('utf-8')


def get_problem_by_slug(slug):
    url = "https://leetcode.com/graphql"
    params = {
        'operationName': "getQuestionDetail",
        'variables': {'titleSlug': slug},
        'query': '''query getQuestionDetail($titleSlug: String!) {
            question(titleSlug: $titleSlug) {
                questionId
                questionFrontendId
                questionTitle
                questionTitleSlug
                content
                difficulty
                stats
                similarQuestions
                categoryTitle
                topicTags {
                    name
                    slug
                }
            }
        }'''
    }

    json_data = json.dumps(params).encode('utf8')

    headers = {
        'User-Agent': user_agent,
        'Connection': 'keep-alive',
        'Content-Type': 'application/json',
        'Referer': 'https://leetcode.com/problems/' + slug
    }
    resp = session.post(url, data=json_data, headers=headers, timeout=10)
    print(resp.content.decode("utf-8"))
    return resp.content.decode('utf-8')


_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class LeetcodeServicer(LeetcodeServiceServicer):

    def Login(self, request, context):
        login(request.name, request.password)
        return Response(result="ok")

    def ListAllProblems(self, request, context):
        res = get_problems()
        return Response(result=res)

    def QueryProblem(self, request, context):
        return Response(result=get_problem_by_slug(request.slug))


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
