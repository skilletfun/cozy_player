from django.test import TestCase
from django.urls import reverse
from rest_framework.status import HTTP_200_OK, HTTP_404_NOT_FOUND, HTTP_400_BAD_REQUEST
from parameterized import parameterized

from artists.models import Artist
from tracks.models import Track


class TrackUpdateTestCase(TestCase):
    @classmethod
    def setUpTestData(cls) -> None:
        cls.artist = Artist.objects.create(pk=1, name="Artist")
        cls.track = Track.objects.create(
            title="Track",
            artist=cls.artist,
            path="",
            duration=180,
        )

    def test_put(self):
        url = reverse("tracks-ru", args=(self.track.pk,))
        data = {"play_count": 11, "title": "Updated"}
        response = self.client.put(url, data=data, content_type="application/json")
        self.assertEqual(response.status_code, HTTP_200_OK)
        self.assertEqual(response.json()["play_count"], data["play_count"])
        self.assertEqual(response.json()["title"], data["title"])
        self.track.refresh_from_db()
        self.assertEqual(self.track.play_count, data["play_count"])
        self.assertEqual(self.track.title, data["title"])

    @parameterized.expand([0, 1, 10, 666, 999])
    def test_patch_play_count(self, value: int):
        url = reverse("tracks-ru", args=(self.track.pk,))
        data = {"play_count": value}
        response = self.client.patch(url, data=data, content_type="application/json")
        self.assertEqual(response.status_code, HTTP_200_OK)
        self.assertEqual(response.json()["play_count"], value)
        self.track.refresh_from_db()
        self.assertEqual(self.track.play_count, value)

    @parameterized.expand([0, 999999])
    def test_invalid_patch(self, pk):
        url = reverse("tracks-ru", args=(pk,))
        data = {"play_count": 999}
        response = self.client.patch(url, data=data, content_type="application/json")
        self.assertEqual(response.status_code, HTTP_404_NOT_FOUND)

    @parameterized.expand(["New", "UpDaTeD"])
    def test_patch_title(self, value: str):
        url = reverse("tracks-ru", args=(self.track.pk,))
        data = {"title": value}
        response = self.client.patch(url, data=data, content_type="application/json")
        self.assertEqual(response.status_code, HTTP_200_OK)
        self.track.refresh_from_db()
        self.assertEqual(self.track.title, value)

    @parameterized.expand(["", None])
    def test_invalid_patch_empty_title(self, value):
        url = reverse("tracks-ru", args=(self.track.pk,))
        data = {"title": value}
        response = self.client.patch(url, data=data, content_type="application/json")
        self.assertEqual(response.status_code, HTTP_400_BAD_REQUEST)
