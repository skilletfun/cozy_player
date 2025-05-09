from django.test import TestCase
from django.urls import reverse
from rest_framework.status import HTTP_200_OK, HTTP_404_NOT_FOUND
from parameterized import parameterized

from artists.models import Artist
from tracks.models import Track


class TrackGetTestCase(TestCase):
    @classmethod
    def setUpTestData(cls) -> None:
        cls.artists = [Artist(pk=i, name=f"Name {i}") for i in range(1, 6)]
        Artist.objects.bulk_create(cls.artists)

        for artist in cls.artists:
            tracks = [
                Track(
                    title=f"Track {i}",
                    artist=artist,
                    path="",
                    duration=180,
                )
                for i in range(1, 51)
            ]
            Track.objects.bulk_create(tracks)

    def test_get_list(self):
        response = self.client.get(reverse("tracks-list"))
        self.assertEqual(response.status_code, HTTP_200_OK)
        self.assertEqual(len(response.data), Track.objects.count())
        self.assertEqual(len(response.data), 250)
        self.assertEqual(Track.objects.count(), 250)
        for prop in ["id", "title", "path", "artist", "artist_name", "duration", "play_count"]:
            for row in response.data:
                self.assertTrue(prop in row)

    @parameterized.expand(
        [
            ("title=Track", 250),
            ("title=Track 47", 5),
            ("title=Trackkk", 0),
            ("artist=1", 50),
            ("artist=1&artist=2", 100),
            ("artist=1&artist=2&title=41", 2),
        ]
    )
    def test_get_list_fitered(self, query: str, count: int):
        response = self.client.get(reverse("tracks-list") + f"?{query}")
        self.assertEqual(response.status_code, HTTP_200_OK)
        self.assertEqual(len(response.data), count)
        for prop in ["id", "title", "path", "artist", "artist_name", "duration", "play_count"]:
            for row in response.data:
                self.assertTrue(prop in row)

    @parameterized.expand([0, 999999])
    def test_invalid_get_cover(self, pk):
        response = self.client.get(reverse("artists-cover", args=(pk,)))
        self.assertEqual(response.status_code, HTTP_404_NOT_FOUND)
