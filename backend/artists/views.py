from django.db.models import Count, Sum
from django.http import FileResponse
from rest_framework.generics import ListAPIView, RetrieveAPIView

from .models import Artist
from .serializers import ArtistSerializer, ArtistInfoSerializer


class ArtistListAPIView(ListAPIView):
    queryset = Artist.objects.all()
    serializer_class = ArtistSerializer


class ArtistGetAPIView(RetrieveAPIView):
    queryset = Artist.objects.annotate(
        tracks_count=Count("tracks"), duration=Sum("tracks__duration")
    )
    serializer_class = ArtistInfoSerializer


class ArtistCoverGetAPIView(RetrieveAPIView):
    queryset = Artist.objects.all()

    def retrieve(self, request, *args, **kwargs):
        artist: Artist = self.get_object()
        return FileResponse(
            open(artist.cover, "rb"),
            headers={"Cache-Control": "public, max-age=43200, immutable"},
        )
