from .base import BaseTestCase


class StoreTestCase(BaseTestCase):
    def test_post_get_item(self):
        """
        Post item then get it
        """
        item = {'brand': 'apple',
                'name': 'iPhone7',
                'description': 'The latest iphone'}

        def check_item(result):
            item['id'] = result['id']
            resp = self.make_request('get', '/store/items/{}'.format(item['id']))
            self.assertEqual(resp.json()['id'], item['id'])
            self.assertEqual(resp.json()['brand'], item['brand'])
            self.assertEqual(resp.json()['name'], item['name'])
            self.assertEqual(resp.json()['description'], item['description'])

        self.assert_request('post',
                            '/store/items',
                            data=item,
                            expected_code=201,
                            expected_json=check_item)
