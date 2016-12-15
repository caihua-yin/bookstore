from base import BaseTestCase


class VersionTestCase(BaseTestCase):
    def test_version(self):
        """
        Service should return current version
        """
        def check_version(result):
            self.assertEqual(result.keys(), ['version'])
            self.assert_(result['version'].startswith('v'))

        self.assert_request('get', '/_version', expected_json=check_version)
