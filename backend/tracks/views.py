from django.http import FileResponse
from django.db.models import F
from django_filters.rest_framework import DjangoFilterBackend
from rest_framework.generics import ListAPIView, RetrieveAPIView

from .models import Track
from .serializers import TrackSerializer
from .filters import TrackFilter


class TrackListAPIView(ListAPIView):
    queryset = Track.objects.annotate(artist_name=F("artist__name"))
    serializer_class = TrackSerializer
    filter_backends = (DjangoFilterBackend,)
    filterset_class = TrackFilter


class TrackGetAPIView(RetrieveAPIView):
    queryset = Track.objects.all()

    def retrieve(self, request, *args, **kwargs):
        track: Track = self.get_object()
        return FileResponse(
            open(track.path, "rb"),
            headers={
                "Access-Control-Allow-Origin": "*",
                "Access-Control-Expose-Headers": "Content-Length, Content-Range",
                "Accept-Ranges": "bytes",
            },
        )


class TrackCoverGetAPIView(RetrieveAPIView):
    queryset = Track.objects.all()

    def retrieve(self, request, *args, **kwargs):
        track: Track = self.get_object()
        return FileResponse(
            track.cover,
            headers={"Cache-Control": "public, max-age=604800, immutable"},
        )
