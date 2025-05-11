from django.test import TestCase
from django.urls import reverse
from rest_framework.status import HTTP_200_OK

from artists.models import Artist
from tracks.models import Track
from playlists.models import Playlist


class PlaylistGetTestCase(TestCase):
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

        cls.playlist_1 = Playlist.objects.create(name="Playlist 1")
        cls.playlist_1.artists.set(cls.artists[:2])
        cls.playlist_1.tracks.set(cls.artists[-1].tracks.all())

        cls.playlist_2 = Playlist.objects.create(name="Playlist 2")
        cls.playlist_2.artists.set(cls.artists[3:4])
        cls.playlist_2.tracks.set(cls.artists[2].tracks.all())

    def test_get_list(self):
        response = self.client.get(reverse("playlists-list"))
        self.assertEqual(response.status_code, HTTP_200_OK)
        self.assertEqual(len(response.data), Playlist.objects.count())
        self.assertEqual(len(response.data), 2)
        self.assertEqual(Playlist.objects.count(), 2)
        for prop in ["id", "name", "artists", "tracks"]:
            for row in response.data:
                self.assertTrue(prop in row)
