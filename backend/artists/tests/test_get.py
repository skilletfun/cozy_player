from io import BytesIO

from django.test import TestCase
from django.urls import reverse
from django.conf import settings
from rest_framework.status import HTTP_200_OK, HTTP_404_NOT_FOUND
from parameterized import parameterized

from artists.models import Artist


class ArtistGetTestCase(TestCase):
    @classmethod
    def setUpTestData(cls) -> None:
        cls.artists = [Artist(pk=i, name=f"Name {i}") for i in range(1, 151)]
        Artist.objects.bulk_create(cls.artists)

    def test_get_list(self):
        response = self.client.get(reverse("artists-list"))
        self.assertEqual(response.status_code, HTTP_200_OK)
        self.assertEqual(len(response.data), Artist.objects.count())
        self.assertEqual(len(response.data), 150)
        self.assertEqual(Artist.objects.count(), 150)
        self.assertTrue(all(["id" in r and "name" in r for r in response.data]))

    @parameterized.expand([i + 1 for i in range(150)])
    def test_get_by_id(self, pk: int):
        response = self.client.get(reverse("artists-retrieve", args=(pk,)))
        self.assertEqual(response.status_code, HTTP_200_OK)
        data = response.json()
        self.assertEqual(data["id"], pk)
        self.assertEqual(data["name"], f"Name {pk}")

    @parameterized.expand([0, 999999])
    def test_invalid_get_by_id(self, pk):
        response = self.client.get(reverse("artists-retrieve", args=(pk,)))
        self.assertEqual(response.status_code, HTTP_404_NOT_FOUND)

    def test_get_cover_template(self):
        response = self.client.get(reverse("artists-cover", args=(1,)))
        self.assertEqual(response.status_code, HTTP_200_OK)
        r_image = BytesIO(b"".join(response.streaming_content))
        with open(settings.DEFAULT_ARTIST_COVER, "rb") as f:
            self.assertEqual(r_image.read(), f.read())

    @parameterized.expand([0, 999999])
    def test_invalid_get_cover(self, pk):
        response = self.client.get(reverse("artists-cover", args=(pk,)))
        self.assertEqual(response.status_code, HTTP_404_NOT_FOUND)
