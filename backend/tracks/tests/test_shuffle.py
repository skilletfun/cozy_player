import random

from django.test import TestCase

from tracks.utils import shuffle_queue


class ShuffleQueueTestCase(TestCase):
    @classmethod
    def setUpTestData(cls):
        artists: int = 5
        total_items: int = 100
        artists: list[int] = [i + 1 for i in range(artists)]

        cls.data: list[dict] = []
        for _ in range(total_items):
            cls.data.append({"play_count": random.randint(0, 10), "artist": random.choice(artists)})

        random.shuffle(cls.data)

    def test_shuffle(self):
        result = shuffle_queue(self.data)

        counts = [row["play_count"] for row in result]
        self.assertEqual(counts, sorted(counts))
