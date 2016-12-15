"""
Base for store-service functional testing.
"""

import unittest
import requests
import json

from server import service


class BaseTestCase(unittest.TestCase):
    """
    Base test case.
    """
    def setUp(self):
        service.start()

    def make_request(self, method, uri, headers={}, params={}, data=None):
        url = 'http://' + service.http_address + uri
        options = {}

        return getattr(requests, method)(url, data=data, headers=headers, params=params, **options)

    def assert_request(self, method, uri, headers={}, params={}, data=None,
                       expected_code=200, expected_body='', expected_json=None,
                       expected_headers=[]):
        if len(headers) == 0:
            headers = {}

        if data is not None and 'Content-Type' not in headers:
            data = json.dumps(data)
            headers['Content-Type'] = 'application/json'

        resp = self.make_request(method, uri, headers, params, data)

        # Validate status code first
        self.assertEqual(resp.status_code, expected_code, resp.text)

        for k, v in expected_headers:
            if callable(v):
                v(resp.headers[k])
            else:
                self.assertEqual(resp.headers[k], v)

        if expected_json is not None:
            self.assertEqual(resp.headers['Content-Type'], 'application/json')
            if callable(expected_json):
                expected_json(resp.json())
            else:
                self.assertEqual(resp.json(), expected_json)
        elif expected_body is not None:
            self.assertEqual(resp.text, expected_body)

        return resp
